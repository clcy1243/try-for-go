package main

import (
	"fmt"
	// "time"
	"database/sql"
	// "database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "yuanjingtt:4MIICfgs421p@tcp(yjclwdb1.mysql.rds.aliyuncs.com:3306)/carowner")
	if err != nil {
		fmt.Println("连接数据库失败")
		fmt.Println(err)
        return;
	}
	status := 0
	rows, err := db.Query("SELECT * FROM contract WHERE status=?", status)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	defer db.Close()
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for i, _ := range columns {
		valuePtrs[i] = &values[i]
	}

	htmlFile, err := os.Create("table.html")
	htmlFile.WriteString(
		fmt.Sprint(
			"<!DOCTYPE html>",
			"<html lang=\"en\">",
			"<head>",
			"<meta charset=\"UTF-8\" />",
			"<title>Document</title>",
			"<!-- 新 Bootstrap 核心 CSS 文件 -->",
			"<link rel=\"stylesheet\" href=\"http://cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css\">",
			"<!-- 可选的Bootstrap主题文件（一般不用引入） -->",
			"<link rel=\"stylesheet\" href=\"http://cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css\">",
			"<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->",
			"<script src=\"http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js\"></script>",
			"<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->",
			"<script src=\"http://cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js\"></script>",
			"</head><body><table class=\"table\"><thead><tr>"
			)
		)
	for _, column := range columns {
		htmlFile.WriteString(fmt.Sprintf("<th>%v</th>", column))
	}
	htmlFile.WriteString("</tr></thead><tbody>")
	// fmt.Println("contract: ", columns)
	for rows.Next() {
		htmlFile.WriteString("<tr>")
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		// fmt.Println(values)
		for i, _ := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			htmlFile.WriteString(fmt.Sprintf("<td>%v</td>", v))
		}
		htmlFile.WriteString("</tr>")
	}
	htmlFile.WriteString("</tbody></table></body></html>")
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}
