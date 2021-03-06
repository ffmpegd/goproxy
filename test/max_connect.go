package main
import(
    "net"
    "log"
    "time"
    "runtime"
    "bytes"
)
func new_connect(end chan int,data []byte){
    defer func(){end<-1;}()
	addr,err:=net.ResolveTCPAddr("tcp","127.0.0.1:20002");
	if (err!=nil){
		log.Print(err);
		return;
	}
	newConn, err:= net.DialTCP("tcp",nil,addr)
	if (err!=nil){
		log.Print(err);
		return;
	}
	defer newConn.Close();
    newConn.Write(data);
    newConn.CloseWrite();
    readBuffer := make([]byte,1024);
    newConn.Read(readBuffer);
    
}
func dumpGoRoutineNum(){
    for {
        time.Sleep(1000*1000*1000);
        log.Print(runtime.NumGoroutine());
        runtime.GC();
    }
}
func main(){
    go dumpGoRoutineNum();
    connect_num:=1;
    data:=bytes.Repeat([]byte("1"),1000);
    end:=make(chan int,connect_num);
    for i:=0;i<connect_num;i++{
        go new_connect(end,data)
    }
    for i:=0;i<connect_num;i++{
        <-end;
    }
}