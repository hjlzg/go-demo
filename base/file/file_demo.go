package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//使用ioutil读取文件
func ReadFile(fileName string){
	data,err:=ioutil.ReadFile(fileName)
	check(err)
	fmt.Println(string(data))
}

//读取文件夹
func ReadAllDir(path string){
	files,err:=ioutil.ReadDir(path)
	if err!=nil{
		log.Fatal(err)
	}
	for _,file:=range files{
		fmt.Println(file.Name())
	}
}

//这种会覆盖掉原先内容
func WriteFile(fileName,data string){
	err:=ioutil.WriteFile(fileName,[]byte(data),os.ModePerm)
	check(err)
}

//追加内容到文件末尾
func AppendToFile(fileName,data string){
	file,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_APPEND,os.ModePerm)
	defer file.Close()
	check(err)
	file.Write([]byte("data"))
}

//创建文件
func CreateFile(fileName string){
	file,err:=os.Create(fileName)
	defer file.Close()
	check(err)
}

//创建单个文件夹
func MkOneDir(dir string){
	err:=os.Mkdir(dir,os.ModePerm)
	check(err)
	os.RemoveAll(dir)
}

//创建多层文件夹
func MkAllDir(dirs string){
	if !IsExist(dirs){
		err:=os.MkdirAll(dirs,os.ModePerm)
		check(err)
		os.RemoveAll(strings.Split(dirs,"/")[0])
	}
}

//删除文件
func DeleteFile(fileName string){
	err:=os.Remove(fileName)
	check(err)
}

func FileAbs(path string) string{
	if absPath,err:=filepath.Abs(path);err==nil{
		return absPath
	}
	return ""
}

func IsExist(filePath string) bool{
	_,err:=os.Stat(filePath)
	if err!=nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}

func check(err error){
	if err!=nil{
		panic(err)
	}
}
