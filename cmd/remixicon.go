package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"html/template"
)

func ConvertRemixIcon() *gcli.Command {
	return &gcli.Command{
		Name:    "convert-r",
		UseFor:  "转换RemixIcon to flutter dart",
		Aliases: []string{"c-r"},
		Func: func(cmd *gcli.Command, args []string) error {
			e := convert()
			if e != nil {
				color.Info.Println(e)
				return e
			}
			color.Success.Println("生成成功")
			return nil
		},
	}
}

type Icon struct {
	V   []string
}

func convert() error {
	p := &Icon{}
	file,err := GetDirFiles("file")
	if err != nil {
		color.Info.Println("文件不能为空")
		return err
	}

	f,err := ioutil.ReadFile(file[0])
	if err != nil {
		color.Info.Println("文件读取失败")
		return err
	}
	reg := regexp.MustCompile(`.ri-(.*?):.*?"\\(.*?)";`)
	icon := reg.FindAllStringSubmatch(string(f),-1)

	for _,v := range icon {
		i := v[1]
		value := v[2]
		title := strings.Replace(strings.Title(strings.Replace(fmt.Sprintf("ri-%s",i),"-"," ",-1))," ","",-1)
		s := fmt.Sprintf("static const IconData %s = IconData(0x%s, fontFamily: _family);",title,value)
		p.V = append(p.V, s)
	}

	t1, err := template.ParseFiles(file[2])

	iconFile, error := os.Create("file/remixicon.dart")
	if error != nil {
		return error
	}

	var b1 bytes.Buffer
	e := t1.Execute(&b1, p)
	if e != nil {
		return errors.New("执行模板失败")
	}
	_, eee := iconFile.WriteString(b1.String())
	if eee != nil {
		return errors.New("生成失败")
	}
	defer iconFile.Close()
	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func GetDirFiles(dir string) ([]string, error) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesRet := make([]string, 0)

	for _, file := range dirList {
		if file.IsDir() {
			files, err := GetDirFiles(dir + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}

			filesRet = append(filesRet, files...)
		} else {
			filesRet = append(filesRet, dir+string(os.PathSeparator)+file.Name())
		}
	}

	return filesRet, nil
}