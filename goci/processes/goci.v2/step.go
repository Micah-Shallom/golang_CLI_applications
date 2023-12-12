package main

import "os/exec"

type step struct {
	name 		string
	exe	 		string
	args		[]string
	message		string
	proj		string
}

//creating a constuctor that instantiates and returns a new step
func newStep(name,exe,message,proj string, args []string) step {
	return step{
		name: name,
		exe: exe,
		message: message,
		args: args,
		proj: proj,
	}
}

func (s step) execute() (string, error){
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj
	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step: s.name,
			msg: s.message,
			cause: err,
		}
	}
	return s.message, nil
}