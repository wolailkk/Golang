package main

import (
	"fmt"
	"os/exec"
)

func main(){
	var(
		cmd *exec.Cmd
		output []byte
		err error
	)

	//cmd = exec.Command("G:\\cygwin64\\bin\\bash.exe","-c","sleep 5;ls")
	cmd = exec.Command("C:\\Users\\Administrator\\AppData\\Local\\Atlassian\\SourceTree\\git_local\\git-cmd.exe","-c","sleep 5;ls -l")
	if output,err = cmd.CombinedOutput(); err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(output))

}