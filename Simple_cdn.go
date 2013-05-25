package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

var new_config = make(map[string]string) //配置文件
var staticHandler http.Handler
var File_cdn_addr, Addr_cdn_url, Addr, Cdn_Host, port string

const (
	CONFIG_FILE string = "Simple_cdn.conf"
)

// 初始化参数
func init() {
	//读取配置文件
	read_config()
	//绑定路径
	cdn_path := http.Dir(File_cdn_addr)
	staticHandler = http.FileServer(cdn_path)

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", StaticServer)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {

	if strings.Index(Cdn_Host, req.Host) != -1 {
		if req.URL.Path == "/" {
			// 跳转
			w.Header().Add("Location", Addr)
			w.WriteHeader(302)
			return
		}
		status := FileExist(File_cdn_addr + req.URL.Path)
		if status == false {
			// 跳转
			w.Header().Add("Location", Addr_cdn_url+req.URL.Path)
			w.WriteHeader(302)
			go down_file(req.URL.Path)
			return
		}
		staticHandler.ServeHTTP(w, req)
		return
	}
	w.WriteHeader(404)
}

//下载文件
func down_file(down_url string) {
	status := FileExist(File_cdn_addr + down_url)
	if status == true {
		return
	}
	response, err := http.Get(Addr_cdn_url + down_url)
	if err != nil {
		fmt.Println("http error!")
		return
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		//创建多级目录
		if strings.Index(File_cdn_addr+down_url, "/") != -1 {
			sp := strings.Split(File_cdn_addr+down_url, "/")
			//创建多层目录
			dir := ""
			for i := 0; i < len(sp)-1; i++ {
				dir += sp[i] + "/"
				os.Mkdir(dir, 0777)
			}
		}
		//写入文件
		file, err := os.Create(File_cdn_addr + down_url)
		if err == nil {
			defer file.Close()
			io.Copy(file, response.Body)
		}

	}
}

//读取配置文件
func read_config() {
	var config = make(map[string]string)
	dir, _ := path.Split(os.Args[0])
	os.Chdir(dir)
	path, _ := os.Getwd()
	config_file, err := os.Open(path + "/" + CONFIG_FILE) //打开文件
	if err != nil {
		fmt.Print("Can not read configuration file. now exit\n")
		os.Exit(0)
	}
	defer config_file.Close()
	buff := bufio.NewReader(config_file) //读入缓存
	//读取配置文件
	for {
		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil {
			break
		}
		rs := []rune(line)
		if string(rs[0:1]) == `#` || len(line) < 3 {
			continue
		}

		str_type := string(rs[0:strings.Index(line, " ")])
		detail := string(rs[strings.Index(line, " ")+1 : len(rs)-1])
		config[str_type] = detail
	}
	//再次过滤 (防止没有配置文件)
	verify(config)

	return
}

// 检查文件或目录是否存在--指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//验证配置文件
func verify(config map[string]string) {
	//配置文件--端口
	if len(config["port"]) >= 1 {
		port = config["port"]
	} else {
		port = "80"
	}

	//配置文件--cdn文件存放目录
	if len(config["File_cdn_addr"]) > 3 {
		File_cdn_addr = config["File_cdn_addr"]
	} else {
		fmt.Print("Can not read File_cdn_addr parameter. now exit\n")
		os.Exit(1)
	}

	//配置文件--加速的域名地址
	if len(config["Addr_cdn_url"]) > 3 {
		Addr_cdn_url = config["Addr_cdn_url"]
	} else {
		fmt.Print("Can not read Addr_cdn_url parameter. now exit\n")
		os.Exit(1)
	}

	//配置文件--加速的域名
	if len(config["Cdn_Host"]) > 3 {
		Cdn_Host = config["Cdn_Host"]
	} else {
		fmt.Print("Can not read Cdn_Host parameter. now exit\n")
		os.Exit(1)
	}

	//配置文件--默认域名
	if len(config["Addr"]) > 3 {
		Addr = config["Addr"]
	} else {
		fmt.Print("Can not read Addr parameter. now exit\n")
	}
	return
}
