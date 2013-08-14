package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	//"strings"
	"time"
	"os"
	"runtime"
)

func main(){



	c:=make(chan string,20000)
	resp,err:=http.Get("http://hao.jianzhu800.com")
	if err!=nil{
		fmt.Println("http get error.")
	}

	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
        fmt.Println(time.Now(),"\n")
	if err!=nil{
		fmt.Println("http read error")
		return
	}

	src:=string(body)

	re,_:=regexp.Compile("\"(http:[\\S\\s]+?)\"")
	all:=re.FindAllStringSubmatch(src,-1)
	for  _,v:=range all{

		c<-v[1]
	}

         pa(c)
	


}

func pa(c chan string){
	fmt.Println(len(c))
	for i:=0;i<3;i++{
		//d:=<-c
		go geturl(string(i),c)
        }
runtime.Gosched()
}

func geturl(d string,c chan string){
	/*resp,err:=http.Get(d)
        if err!=nil{
                 fmt.Println("http get error.")
         }
        defer resp.Body.Close()
        body,err:=ioutil.ReadAll(resp.Body)
        if err!=nil{
        fmt.Println("http read error")
        }
        src:=string(body)
       re,_:=regexp.Compile("\"(http:[\\S\\s]+?)\"")
       all:=re.FindAllStringSubmatch(src,-1)
       for  _,v:=range all{
            c<-v[1]
       }*/
       userFile:=d
       fout,err:=os.Create(userFile)
       
       if err!=nil{
	       panic(err)
	       return
       }
       for i := 0; i < 10; i++ {
fout.WriteString("Just a test!\r\n")
fout.Write([]byte("Just a test!\r\n"))
      fout.Close()
}
runtime.Gosched()
       }



