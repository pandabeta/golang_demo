package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func main(){
	fmt.Println(time.Now(),"\n")

	resp,err:=http.Get("http://blog.csdn.net/duguteng/article/details/7412652")
	if err!=nil{
		fmt.Println("http get error.")
	}
	fmt.Println(time.Now(),"\n")
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
fmt.Println(time.Now(),"\n")
	if err!=nil{
		fmt.Println("http read error")
		return
	}

	src:=string(body)

	re,_:=regexp.Compile("\\<[\\S\\s]+?\\>")
	src=re.ReplaceAllStringFunc(src,strings.ToLower)
	
	re,_=regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src=re.ReplaceAllString(src,"")

	re,_=regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src=re.ReplaceAllString(src,"")

	re,_=regexp.Compile("\\<[\\S\\s]+?\\>")
	src=re.ReplaceAllString(src,"\n")

	re,_=regexp.Compile("\\s{2,}")
	src=re.ReplaceAllString(src,"\n")
fmt.Println(time.Now(),"\n")

	fmt.Println(strings.TrimSpace(src))
fmt.Println(time.Now(),"\n")

}
