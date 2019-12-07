package main
//主要是context配合协程执行命令，自由控制开关
import (
	"context"
	"fmt"
	"os/exec"
	"time"
)
//定义结构体
type result struct{
	err error
	output []byte
}
//定义了一个main方法
func main(){
	//定义
	var(
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)
	resultChan = make(chan * result,1000)
	ctx,cancelFunc = context.WithCancel(context.TODO())

	go func(){
		//闭包内定义
		var(
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"C:\\Users\\Administrator\\AppData\\Local\\Atlassian\\SourceTree\\git_local\\bin\\bash.exe","-c","sleep 1;echo hello word!")
		output,err = cmd.CombinedOutput()
		resultChan <- &result{
			err:err,
			output:output,
		}
	}()

	time.Sleep(1 * time.Second)
	//取消上下文
	cancelFunc()
	res = <- resultChan
	fmt.Println(res.err,string(res.output))
}
