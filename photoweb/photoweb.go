/*
	支持图片上传
	查看上传的图片
	图片列表
	删除指定图片
*/

package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

// 上传图片的handler
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 提交表单页面
	if r.Method == "GET" {
		// io.WriteString(w, "<html><form method=\"POST\" action=\"/upload\" "+
		// 	" enctype=\"multipart/form-data\">"+
		// 	"Choose an image to upload: <input name=\"image\" type=\"file\" />"+"<input type=\"submit\" value=\"Upload\" />"+
		// 	"</form></html>")
		if err := renderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// 点击上传按钮
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if isExists := isExists(imagePath); !isExists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path) // 文件信息
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads") // 读取所有图片
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/* 注释的是html输出
	var listHtml string
	for _, fileInfo := range fileInfoArr {
		imgid := fileInfo.Name()
		listHtml += "<li><a href=\"/view?id=" + imgid + "\">" + imgid + "</a></li>"
	}
	io.WriteString(w, "<html><ol>"+listHtml+"</ol></html>")
	*/

	data := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	data["images"] = images
	if err := renderHtml(w, "list", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 抽离出的渲染模版的方法，使用模版渲染网页
func renderHtml(w http.ResponseWriter, tpl string, data map[string]interface{}) error {
	err := templates[tpl+".html"].Execute(w, data)
	return err
}

// 全局变量
var templates = make(map[string]*template.Template)

// templates = make(map[string]*template.Template)

// 初始化函数，在main函数前执行
func init() {

	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	for _, fileInfo := range fileInfoArr {

		templateName := fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			// 扩展名为html
			continue
		}
		templatePath := TEMPLATE_DIR + "/" + templateName
		t := template.Must(template.ParseFiles(templatePath)) // 强制有模版才可以向下进行，没有模版报错。
		templates[templateName] = t
	}
}

func main() {
	// 图片列表页
	http.HandleFunc("/", listHandler)
	// 查看图片
	http.HandleFunc("/view", viewHandler)
	// 上传图片
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
