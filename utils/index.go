package utils

import (
	"bytes"
	"fmt"
	"github.com/jellycheng/gosupport"
	"io/ioutil"
	"net/http"
)

func SaveGphParams(w http.ResponseWriter, r *http.Request, fileName string) {
	//curUrl := r.URL.Path                      // 当前地址
	//if r.URL.RawQuery != ""{
	//	curUrl += "?" + r.URL.RawQuery
	//}
	curUrl := r.RequestURI
	contentType := r.Header.Get("Content-Type")
	allHeaders := gosupport.ToJson(r.Header)  // 所有请求头，返回类型 map[string][]string
	getStr := gosupport.ToJson(r.URL.Query()) // 所有get参数，返回类型 map[string][]string
	//_ = r.ParseForm()
	_ = r.ParseMultipartForm(32 << 20)
	postStr := gosupport.ToJson(r.PostForm) //post参数，返回类型 map[string][]string
	postBodyStr := ""                       // post请求body内容
	if postByte, err := ioutil.ReadAll(r.Body); err == nil {
		_ = r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(postByte))
		postBodyStr = string(postByte)
	}

	str := fmt.Sprintf(`
====================%s====================
当前地址: %s 
请求方式: %s
Content-Type: %s
headers: %s
RawQuery: %s
get: %s
postform: %s 
postbody: %s
====================end====================
`, gosupport.TimeNow2Format(gosupport.TimeFormat), curUrl, r.Method, contentType, allHeaders, r.URL.RawQuery, getStr, postStr, postBodyStr)

	gosupport.CreateSuperiorDir(fileName)
	if _, err := gosupport.FilePutContents(fileName, str, 0666); err != nil {
		fmt.Println(err.Error())
	}

}

func DianBoResp(hRet int, cpOrderId string) string {
	str := `<cp_notify_resp>
    <h_ret>%d</h_ret>
    <cp_order_id>%s</cp_order_id>
</cp_notify_resp>
`
	return fmt.Sprintf(str, hRet, cpOrderId)
}
