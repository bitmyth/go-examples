package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/astaxie/beego/logs"
)

// InitSignal register signals handler.
/*
    信号处理模块，可用于在线加载配置，配置动态加载的信号为SIGHUP。
*/
func InitSignal() {
    c := make(chan os.Signal, 1)
    /*
        func Notify(c chan<- os.Signal, sig ...os.Signal)
        Notify函数让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。

        signal包不会为了向c发送信息而阻塞（就是说如果发送时c阻塞了，signal包会直接放弃）：调用者应该保证c有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的通道，缓存为1就足够了。

        可以使用同一通道多次调用Notify：每一次都会扩展该通道接收的信号集。唯一从信号集去除信号的方法是调用Stop。可以使用同一信号和不同通道多次调用Notify：每一个通道都会独立接收到该信号的一个拷贝。
    */
    /*
        syscall    包含一个低级的操作系统原语的接口
    */
    /*
        SIGHUP        终端控制进程结束(终端连接断开)
        SIGQUIT        用户发送QUIT字符(Ctrl+/)触发
        SIGTERM        结束程序(可以被捕获、阻塞或忽略)
        SIGINT        用户发送INTR字符(Ctrl+C)触发
        SIGSTOP        停止进程(不能被捕获、阻塞或忽略)
    */
    signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
    for {
        s := <-c
        logs.Info("comet[%s] get a signal %s", s.String())
        switch s {
        case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
            return
        case syscall.SIGHUP:
            /*  
                触发条件   命令行
                $ kill -SIGHUP  + 进程的 PID

                Example：

                kill -SIGHUP 84899
            */
            reload()
        default:
            return
        }
    }
}

func reload() {
	Conf:=1;
    logs.Error("after Conf -->  %v", Conf)
    err := InitConf()
    if err != nil {
        logs.Error("InitConf() error(%v)", err)
        return
    }
    logs.Error("before Conf -->  %v", Conf)
}

func InitConf()error{
	return nil
}

func main() {
    InitSignal()

}
