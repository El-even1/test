package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>简单的 Go Web 页面</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        }
        .container {
            text-align: center;
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.3);
        }
        h1 {
            color: #333;
            margin-bottom: 20px;
        }
        p {
            color: #666;
            font-size: 18px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Hello, Go Web!</h1>
        <p>这是一个用 Go 语言创建的最简单的 Web 页面</p>
    </div>
</body>
</html>
`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("服务器启动在 http://localhost:8080")
	fmt.Println("打开浏览器访问 http://localhost:8080 查看页面")
	http.ListenAndServe(":8080", nil)
}
