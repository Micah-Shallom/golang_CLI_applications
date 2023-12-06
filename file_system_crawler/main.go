package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	f   = os.Stdout
	extensions []string
	err error
)

type config struct {
	//extension to filter out
	exts 	[]string
	//min file size
	size 	int64
	//list files
	list 	bool
	//delete files
	del 	bool
	//log destination writer
	wLog 	io.Writer
	//archive directory
	archive string
	//date and time before and after
	before	string
	after	string
}

type DateTime struct {
	dt		string
	info	string
}

func main() {
	//parsing command line flags
	root := flag.String("root", ".", "root directory to start")
	logFile := flag.String("log", "", "Log deletes to this file") //default value for flag is "", program will send to STDOUT if user provides nothing
	//action options
	list := flag.Bool("list", false, "list files only")
	del := flag.Bool("del", false, "Delete files")
	archive := flag.String("archive", "", "archive directory")
	//filter options
	exts := flag.String("ext", "", "file extension(s) to filter out")
	size := flag.Int64("size", 0, "minimum file size")
	before := flag.String("before", "", "return files created before specfied date and time")
	after := flag.String("after", "", "return files created after specified date and time")
	flag.Parse()

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if *exts != ""{
		extensions = strings.Split(*exts, ",")
	}

	c := config{
		exts:     extensions,
		size:    *size,
		list:    *list,
		del:     *del,
		wLog:    f,
		archive: *archive,
		before:	*before,
		after:  *after,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "Deleted File: ", log.LstdFlags)

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if cfg.before != "" && cfg.after != ""{
			log.Fatal("provide only one of before or after flags at a time")
			return nil
		}

		if cfg.before != ""{
			if filterOut(path, cfg.exts, cfg.size, DateTime{dt: cfg.before, info: "before"} ,info) {
				return nil
			}
		}

		if cfg.after != ""{
			if filterOut(path, cfg.exts, cfg.size, DateTime{dt: cfg.after, info: "after"} ,info) {
				return nil
			}
		}
		if filterOut(path, cfg.exts, cfg.size, DateTime{dt: "", info: ""} ,info){
			return nil
		}
		// If list was explicitly set, dont do anything else
		if cfg.list {
			return listFile(path, out)
		}

		//Archive files and continue if successful
		if cfg.archive != "" {
			if err := archiveFile(cfg.archive, root, path); err != nil {
				return err
			}
		}

		if cfg.del {
			return delFile(path, delLogger)
		}
		//list is the default option if nothing else is met
		return listFile(path, out)
	})
}
