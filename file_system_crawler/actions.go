package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func filterOut(path string, exts []string, minSize int64 , dateTime DateTime, info os.FileInfo) bool {
	if info.IsDir() || info.Size() < int64(minSize) {
		return true
	}

	if dateTime.dt != "" || dateTime.info != ""{
		layout := "2006-01-02 15:04:05"
	
		beforeDT, err1 := time.Parse(layout, dateTime.dt)
		afterDT, err2 := time.Parse(layout, dateTime.dt)
	
		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing dates: ", err1, err2)
			// return 
		}
	
		pathDT, _ := time.Parse(layout, info.ModTime().Local().Format(layout))
	
		switch dateTime.info{
		case "before":
			if !pathDT.Before(beforeDT) {
				return true
			}
		case "after":
			if !pathDT.After(afterDT) {
				return true
			}
		}
	}
	

	ext := filepath.Ext(path)
	if len(exts) != 0 {
		for _, extValue := range exts{
			if extValue != "" && ext != extValue{
				continue
			}
			return false
		}
		return true
	}
	return false
}

func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}

func delFile(path string, delLogger *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	delLogger.Println(path)
	return nil
}

//function to archive path before deletion
func archiveFile(destDir, root, path string) error {
	//checking if the destDir is a directory
	info, err := os.Stat(destDir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", destDir)
	}

	//determine the relative directory of the file to be archived in relation to its source root directory
	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		return err
	}
	//create new file name by adding the .gz suffix to the original filename obtained by calling the filepath.Base() function
	dest := fmt.Sprintf("%s.gz", filepath.Base(path))
	targetPath := filepath.Join(destDir, relDir, dest)

	//with targetPath defined, create the target directory tree
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	out, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	zw := gzip.NewWriter(out)
	zw.Name = filepath.Base(path)

	if _, err = io.Copy(zw, in); err != nil {
		return err
	}
	if err := zw.Close(); err != nil {
		return err
	}
	return out.Close()
}
