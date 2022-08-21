package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
 * 打开和关闭文件
 * os.Open() 函数可以打开一个文件，返回一个 *File 和一个 err
 * close() 可以关闭一个打开的文件实例
 */

/*
 * 读取文件
 * os.Read() 它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时，会返回 0 和 io.EOF
 */

/*
 * bufio读取文件
 * bufio 是在 file 的基础上封装了一层 api ，支持更多的功能
 */

/*
 * ioutil 读取整个文件
 * io/ioutil 包的readFile方法能够读取完整的文件，只需要将文件名作为参数传入
 */

/*
 * os.OpenFile()函数可以指定模式打开文件，从而实现文件写入的功能
 *
 * func OpenFile(name string, flag int, perm FileMode) (*File, error)
 *
 * os.O_WRONLY 只写
 * os.O_CREATE 创建文件
 * os.O_RDONLY 只读
 * os.O_RDWR 读写
 * os.O_TRUNC 清空
 * os.O_APPEND 追加
 *
 * perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
 */

/*
 * bufio.NewWriter
 * ioutil.WriteFile
 * copyFile 借助 io.Copy实现一个拷贝文件函数
 */
func main() {

	//打开和关闭文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed err：", err)
		return
	}
	defer file.Close()

	//读取文件
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed err:", err)
		return
	}
	fmt.Printf("读取到%d字节数据", n)
	fmt.Println(tmp[:n])
	fmt.Println(string(tmp[:n]))
	fmt.Println("---------------------------")

	//循环读取到文件中
	var content []byte
	tmp = make([]byte, 128)
	for {
		n, err = file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
	fmt.Println("---------------------------")

	//bufio读取文件
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //这里是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed err:", err)
			return
		}
		fmt.Print(line)
	}

	fmt.Println("---------------------------")

	//ioutil.ReadFile读取整个文件
	content, err = ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
	fmt.Println("---------------------------")

	file, err = os.OpenFile("./aa.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	str := "这个是一个字符串"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("哈哈哈，我是大帅哥") //直接希尔字符串数据
	fmt.Println("---------------------------")

	//bufio.NewWriter
	file, err = os.OpenFile("bb.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	writer := bufio.NewWriter(file)
	//现将数据写入缓存
	for i := 0; i < 3; i++ {
		writer.WriteString("我是大帅哥\n")
	}
	//将缓存数据内容写入文件
	writer.Flush()
	fmt.Println("---------------------------")

	//ioutil.WriteFile
	str = "我是大美女"
	err = ioutil.WriteFile("bb.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
	fmt.Println("---------------------------")

	_, err = CopyFile("aa.txt", "bb.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")

}

func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}
