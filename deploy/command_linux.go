package deploy

import (
	"os/exec"
	"fmt"
	"bytes"
)

// 执行 git reset 命令
func GitResetCommand(dir string, remoteBranch string, localBranch string) (error)  {
	cmd := exec.Command("/bin/sh","-c","git reset --hard "+ localBranch +"/" + remoteBranch);
	cmd.Dir = dir;
	fmt.Println(cmd.Args)

	output,err := cmd.Output();
	if err != nil {
		return  err;
	}
	fmt.Println(bytes.NewBuffer(output).String());

	cmd.Wait();
	return  nil;
}

//执行 git pull 命令
func GitPullCommand(dir string) (error) {
	cmd := exec.Command("/bin/sh","-c","git pull");
	cmd.Dir = dir;

	output,err := cmd.Output();

	if err != nil {
		return  err;
	}

	fmt.Println(bytes.NewBuffer(output).String());

	cmd.Wait();
	return  nil;
}
//执行 git checkout 命令
func GitCheckoutCommand(dir string, remoteBranch string) (error) {
	cmd := exec.Command("/bin/sh","-c","git checkout " + remoteBranch);
	cmd.Dir = dir;

	output,err := cmd.Output();

	if err != nil {
		return  err;
	}
	fmt.Println(bytes.NewBuffer(output).String());

	cmd.Wait();
	return  nil;
}

func GitCommand(dir string,remoteBranch string,localBranch string) (error) {

	err := GitResetCommand(dir,remoteBranch,localBranch);
	if err != nil {
		fmt.Print("git reset error:",err.Error());
		return  err;
	}

	err = GitPullCommand(dir);

	if err != nil {
		fmt.Print("git pull error:",err.Error());
		return  err;
	}

	err = GitCheckoutCommand(dir,remoteBranch);

	if err != nil {
		fmt.Print("git checkout error:",err.Error());
		return  err;
	}

	return nil;
}

