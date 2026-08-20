package main

import ("database/sql"; "net/http")

// 演示用修复代码：参数化查询、移除 exec，复测应通过
func demo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	rows, _ := db.Query("SELECT * FROM users WHERE id = ?", id)
	defer rows.Close()
	w.Write([]byte("ok"))
}
