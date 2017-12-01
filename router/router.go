package router

import (
	"fmt"
	"log"
	"mygo/dbhelper"
	"net/http"
	"time"
)

/***/
func Router() {
	
	//handler func(ResponseWriter, *Request)
	http.HandleFunc("/selan", index)
	http.HandleFunc("/aa", aa)
	s := &http.Server{
		Addr:           ":7088",
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("服务启动成功")
	log.Fatal(s.ListenAndServe())

}
func index(w http.ResponseWriter, r *http.Request) {
	sqlStr := "select * from slm_area"
	w.Write([]byte(dbhelper.SQLJSONData(sqlStr)))
}
func aa(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			fmt.Println("%s, %s", k, v[0])
		}
	}
	// w.Write([]byte("这是测试的内容"))
}
