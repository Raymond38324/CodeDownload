package main

import (
	"fmt"
	"os/exec"
)

func Exec(code string){
	cmd := exec.Command("bash","-c", fmt.Sprintf("echo %s | pbcopy", code))
	if err := cmd.Run(); err != nil{
		panic(err)
	}
}
