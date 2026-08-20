package main

import ("database/sql"; "net/http"; "os/exec")

// 演示用漏洞代码：提交到后 SCG 应报 SQL 注入 + 命令注入
func demo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	rows, _ := db.Query("SELECT * FROM users WHERE id = '" + id + "'")
	defer rows.Close()
	out, _ := exec.Command("sh", "-c", "ping "+id).Output()
	w.Write(out)
}
