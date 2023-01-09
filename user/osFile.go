package main

/*
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name())
	}
}

// 创建目录
func createDir() {
	err := os.Mkdir("test", os.ModePerm) //创建一个目录
	if err != nil {
		fmt.Println(err)
	}
	err = os.MkdirAll("test/test", os.ModePerm) //创建多级目录、
	if err != nil {
		fmt.Println(err)
	}

}

//删除文件和目录

func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("test")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	createFile()
	createDir()
	remove()
	readops()
	readDir()
	write()

	f, _ := os.Open("a.txt")
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

func env() {
	s := os.Environ()
	fmt.Println(s)
	s1 := os.Getenv("GOPATH")
	fmt.Println(s1)
	s2 := os.Setenv("env1", "env1")
	fmt.Println(s2)
	s3, b := os.LookupEnv("env1")
	fmt.Println(s3)
	fmt.Println(b)
}

func write() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0777)
	f.Write([]byte("这是一个内容"))
	f.WriteString("sadfajslkdfjqweo爱上的发就算了打飞机去了为")
	f.Close()

}

func readDir() {
	dir, _ := os.ReadDir("user")
	for _, v := range dir {
		fmt.Println(v.IsDir())
		fmt.Println(v.Name())
	}
}

func readops() {

	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	for {
		buf := make([]byte, 10)
		n, err1 := f.Read(buf)
		if err1 == io.EOF {
			return
		}
		fmt.Println(n)           //查询结果自己长度
		fmt.Println(string(buf)) //查询内容
	}

}
*/