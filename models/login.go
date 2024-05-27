package models

import (
	"fmt"
	"github.com/dchest/captcha"
	"html/template"
	"net/http"
)

func capInit(w http.ResponseWriter, r *http.Request) {
	captcha.Server(200, 80)
}

func LoginInit() {
	// 设置captcha包的全局配置（可选）
	//captcha.DefaultLen = 6
	//h := captcha.Server(200, 80)

	http.HandleFunc("/captcha", capInit)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", homeHandler)

	// 解析并加载模板（如果有的话）
	tmpl := template.Must(template.ParseFiles("html/template/login.html"))
	http.HandleFunc("/login-form", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 假设你有一个处理登录请求的函数
		// loginRequest(r.FormValue("username"), r.FormValue("password"), r.FormValue("captcha_id"), r.FormValue("captcha_value"))

		// 验证captcha
		captchaID := r.FormValue("captcha_id")
		userInput := r.FormValue("captcha_value")

		if captcha.VerifyString(captchaID, userInput) {
			fmt.Fprint(w, "Captcha verified! Username and password to be checked...")
		} else {
			fmt.Fprint(w, "Incorrect captcha!")
		}

		return
	}
	// 如果是GET请求或其他不支持的方法，重定向到登录表单
	http.Redirect(w, r, "/login-form", http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 渲染主页，可能有登录链接等
	fmt.Fprint(w, "Welcome to the login page!")
}
