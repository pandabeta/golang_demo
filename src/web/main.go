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
//	"rand"
//"crypto/md5"
	"strconv"
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
	for i:=0;i<100;i++{
		d:=<-c
		go geturl(strconv.Itoa(i),c,d)
        }
runtime.Gosched()
}

func geturl(d string,c chan string,f string){
	resp,err:=http.Get(f)
        if err!=nil{
                 fmt.Println("http get error.")
         }
        defer resp.Body.Close()
        body,err:=ioutil.ReadAll(resp.Body)
        if err!=nil{
        fmt.Println("http read error")
        }
        src:=string(body)
       //re,_:=regexp.Compile("\"(http:[\\S\\s]+?)\"")
      // all:=re.FindAllStringSubmatch(src,-1)
       
	//  a:=[4]string{"a1","a2","a3","a4"}
       userFile:=d
       fout,err:=os.Create(userFile)
       
       if err!=nil{
	       panic(err)
       }
fout.WriteString(src)
      
fmt.Println(f,len(c),"\n")
fout.Close()
runtime.Gosched()
       }



