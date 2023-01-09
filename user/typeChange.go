package main

/*
import (
	"bytes"
	"fmt"
)

func main() {
	var i int = 1
	var j byte = 2
	//类型转换
	j = byte(i)
	fmt.Printf("j: %v\n", j)
	//查询是否包含
	b := []byte("duoke368.com") //字符串强转为byte切片
	sublice1 := []byte("duoke36a")
	sublice2 := []byte("DuoKe368")
	fmt.Println(bytes.Contains(b, sublice1)) //true
	fmt.Println(bytes.Contains(b, sublice2)) //false
	//查询出现的次数
	s := []byte("hellooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("1")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1)) //1
	fmt.Println(bytes.Count(s, sep2)) //2
	fmt.Println(bytes.Count(s, sep3)) //9

	//重复出现
	b = []byte("hi")
	fmt.Println(string(bytes.Repeat(b, 1))) //hifmt.Println(string(bytes .Repeat(b,3))) //hihihi
	//替换字符串
	s = []byte("hello,world")
	old := []byte("o")
	news := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  //hello, world
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  //hellee, world
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  //hellee, weerldnews ,
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee, weerld
	//字符集长度
	s = []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("转换前字符串的长度:", len(s)) //12"
	fmt.Println("转换后字符串的长度:", len(r)) //4
	//字节链接
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4))) //你好,世界
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5))) //你好#世界
}
*/
/*
import (
	"encoding/json"
	"errors"

	"fmt"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}
func main() {
	s, err := check("hello")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	} else {
		fmt.Printf("s: %v\n", s)
	}
	Marshal()
}

// 结构体转换为json
type Person struct {
	Name  string
	Age   int
	Email string
}

func Marshal() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))

	b1 := []byte(` {"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var m Person
	json.Unmarshal(b1, &m)
	fmt.Println(m)

}

//解析嵌套类型
func test3() {
	b := []byte(`{"Name":"tom","Age":20 ,"Email":"tom@gmail.com","parents":["big tom","kite"]}`)
	var f map[string]interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f)
	for k, v := range f {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}*/
