/*
Author: FreeKnight
服务器配置信息
*/
//------------------------------------------------------------
package server

//------------------------------------------------------------
import (
	"FKTrojan/common"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	MySQLUser string `json:"mysql_user"`
	MySQLPass string `json:"mysql_pass"`
	MySQLHost string `json:"mysql_host"`
	MySQLName string `json:"mysql_name"`
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 在有config.json的情况下使用配置，否则使用默认值
func loadConfig(configPath string) {

	//fmt.Printf("use %s \n", configPath)
	if !exist(configPath) {
		//fmt.Printf("do not use %s \n", configPath)
		return
	}
	recordByte, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}
	var config Config
	err = json.Unmarshal(recordByte, &config)
	if err != nil {
		return
	}
	//fmt.Printf("use %s %+v\n", configPath, config)
	if config.MySQLHost != "" {
		MySQLHost = config.MySQLHost
	}
	if config.MySQLName != "" {
		MySQLName = config.MySQLName
	}
	if config.MySQLPass != "" {
		MySQLPass = config.MySQLPass
	}
	if config.MySQLUser != "" {
		MySQLUser = config.MySQLUser
	}
}
func init() {
	currentPath := common.CurrentBinaryDir()
	configPath := filepath.Join(currentPath, "config.json")
	loadConfig(configPath)
}

//------------------------------------------------------------
var (
	ControlUser string = "root"                             // 管理员账号
	ControlPass string = "082DFC1514DFBF08AE4F6014A2FEBFEA" // 管理员密码 (MD5: bardman)

	UseSSL       bool   = true                               // 是否使用SSL
	MyIP         string = "127.0.0.1"                        // 服务器IP
	MyPort       string = "7777"                             // 服务器Port，如果使用SSL则此项巫妖
	UserAgentKey string = "0AB9008394AA329280DB3FCD6A328EDC" // 服务器用来鉴别这个应用是不是客户端的标示。当前是 MD5("FreeKnight")。

	MySQLUser string = "root"                // SQL数据库用户名
	MySQLPass string = "root"                // SQL数据库密码
	MySQLHost string = "tcp(127.0.0.1:3306)" // SQL数据库Host
	MySQLName string = "panel"               // SQL数据库表名

	ProfileDir string = "./Profiles" // 文件存储目录

	Banner         string = common.Base64Decode(`IyMjIyMjIyMjIyMjIyMjIyMjIyMjI35+Li4nfC4jIyMjIyMjIyMjIyMjIy58YC4ufn4jIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwojIyMjIyMjIyMjIyMjI34uL2Aufn4uLycgLi8gIyMjIyMjIyMjIyMjIyMjIyBcLiBgXC4gfn4uYFwufiMjIyMjIyMjIyMjIyMjCiMjIyMjIyMjIyMjI34uJyBgLmAtJyAgIC8gICB+IyMjIyMjIyMjIyMjI34gLiAgXCAgIGAtJy4nICBgLn4jIyMjIyMjIyMjIyMKIyMjIyMjIyMjI34uJyAgICB8ICAgICB8ICAuJ1wgfiMjIyMjIyMjIyN+IC9gLiAgfCAgICAgfCAgICAgYC5+IyMjIyMjIyMjIwojIyMjIyMjI34uJyAgICAgIHwgICAgIHwgIHxgLmAuXyB+IyMjI34gXy4nLid8ICB8ICAgICB8ICAgICAgIGAufiMjIyMjIyMjCiMjIyMjI34uJyAgICAgICAgYC4gICAgfCAgYC4uYC5ffFwuLS0uL3xfLicuLicgIHwgICAgLicgICAgICAgICBgLn4jIyMjIyMKIyMjI34uJyAgICAgICAgICAgIFwgICB8ICMuYC5gLl9gLictLWAuJ18uJy4nLiMgfCAgIC8gICAgICAgICAgICAgYC5+IyMjIwojI34uJyAgICAgICAuLi4uLi4gIFwgIHwgIyMjLmB+JyhvXHx8L28pYH4nLiMjIyB8ICAvICAuLi4uLi4gICAgICAgIGAufiMjCn4uYC4uLi4uLi4nfiAgICAgIGAuIFwgIFx+IyMjIyAgfFxfICBfL3wgICMjIyN+LyAgLyAuJyAgICAgIH5gLi4uLi4uLi4nLn4KOy4nICAgICAgICAgICAgICAgICBcIC4tLS0tLl9fLidgKG58fG4pJ2AuX18uLS0tLS4gLyAgICAgICAgICAgICAgICAgIGA7IApgLiAgICAgICAgICAgICAgICAgIC4nICAgIC4nICAgYC4gXGAnLyAuJyAgIGAuICAgIGAuICAgICAgICAgICAgICAgICAgLicgCiM6ICAgICAgICAgICAgICAgLi4nOiAgICAgIDogICAgJy4gfn4gLmAgICAgOiAgICAgIDpgLi4gICAgICAgICAgICAgICA6IyAKIzogICAgICAgICAgICAgLicgICA6ICAgICAuJyAgICAgLicgIGAuICAgICBgLiAgICAgOiAgIGAuICAgICAgICAgICAgIDojIAojOiAgICAgICAgICAgLicgICAgLmAgICAuJyAgICAgICAuIHx8IC4gICAgICAgYC4gICAnLiAgICBgLiAgICAgICAgICAgOiMgCiM6ICAgICAgICAgLicgICAgLicgOiAgICAgICAuLi4uJyAgICAgIGAuLi4uICAgICAgIDogYC4gICAgYC4gICAgICAgICA6IyAKIzogICAgICAgLicgICAgLicgKSApICAgICAgICAoICAgICAgKSAgICAgKCAgICAgICggICAgKWAuICAgIGAuICAgICAgIDojIAojOiAgICAgLi4nICAgIC4gICggKCggICApICApICkpICggICgoICAoICApKSAgKSAgKSkgICgoICBgLiAgIGAuLiAgICAgOiMgCiM6ICAuLicgICAgICAuJyMgKSApICkgKCggKCAoICggICkgKSApICkpKCAoICgoICggKCAgKSApICNgLiAgICAgYC4uICA6IyAKIzsuJyAgICAgICAgLicjI3woKCAgKCApICkgKSApICkoICggICgoICggKSApKSApICkgKSggKHx8IyNgLiAgICAgICBgLjojIAojYC4gICAgICAgIC4nIyMjfFxfXyAgKSggKCggKCAoICkgICkgICkpICkgICggICgoICggKV8pL3wjIyNgLiAgICAgICAuJyMgCiMjLmAgICAgICAgJyMjIyMjXF9ffn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+X18vIyMjIyNgICAgICAgJy4jIyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiMjIyAgICAgICAgIyMjIyMjIyAgfn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+fn5+ICAjIyMjIyMjICAgICAgICMjIyA=`)
	LoginHTML      string = common.Base64Decode(`PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MS4wIj4KCiAgICA8dGl0bGU+TG9naW4gUG9ydGFsPC90aXRsZT4KCiAgICA8bGluayByZWw9InN0eWxlc2hlZXQiIGhyZWY9Imh0dHA6Ly95dWkueWFob29hcGlzLmNvbS9wdXJlLzAuNi4wL3B1cmUtbWluLmNzcyI+CgkKPHN0eWxlPgpib2R5IHsKICBiYWNrZ3JvdW5kOiAjNzZiODUyOyAvKiBmYWxsYmFjayBmb3Igb2xkIGJyb3dzZXJzICovCiAgYmFja2dyb3VuZDogLXdlYmtpdC1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IC1tb3otbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiAtby1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IGxpbmVhci1ncmFkaWVudCh0byBsZWZ0LCAjNzZiODUyLCAjOERDMjZGKTsKICAtd2Via2l0LWZvbnQtc21vb3RoaW5nOiBhbnRpYWxpYXNlZDsKICAtbW96LW9zeC1mb250LXNtb290aGluZzogZ3JheXNjYWxlOyAgICAgIAp9CiAgICAgICAgLmJ1dHRvbi1zdWNjZXNzLAogICAgICAgIC5idXR0b24tZXJyb3IsCiAgICAgICAgLmJ1dHRvbi13YXJuaW5nLAogICAgICAgIC5idXR0b24tc2Vjb25kYXJ5IHsKICAgICAgICAgICAgY29sb3I6IHdoaXRlOwogICAgICAgICAgICBib3JkZXItcmFkaXVzOiA0cHg7CiAgICAgICAgICAgIHRleHQtc2hhZG93OiAwIDFweCAxcHggcmdiYSgwLCAwLCAwLCAwLjIpOwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLXN1Y2Nlc3MgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoMjgsIDE4NCwgNjUpOyAvKiB0aGlzIGlzIGEgZ3JlZW4gKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi1lcnJvciB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYigyMDIsIDYwLCA2MCk7IC8qIHRoaXMgaXMgYSBtYXJvb24gKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi13YXJuaW5nIHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDIyMywgMTE3LCAyMCk7IC8qIHRoaXMgaXMgYW4gb3JhbmdlICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24tc2Vjb25kYXJ5IHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDY2LCAxODQsIDIyMSk7IC8qIHRoaXMgaXMgYSBsaWdodCBibHVlICovCiAgICAgICAgfQo8L3N0eWxlPgo8L2hlYWQ+Cgo8Ym9keT4KCjxkaXYgY2xhc3M9InB1cmUtbWVudSBwdXJlLW1lbnUtaG9yaXpvbnRhbCIgYWxpZ249ImNlbnRlciI+CiAgICA8dWwgY2xhc3M9InB1cmUtbWVudS1saXN0Ij4KICAgICAgICA8c3BhbiBjbGFzcz0icHVyZS1tZW51LWhlYWRpbmciIGFsaWduPSJsZWZ0Ij5QbGVhc2UgTG9naW48L3NwYW4+CiAgICAgICAgPC9saT4KICAgIDwvdWw+CjwvZGl2Pgo8ZGl2IGFsaWduPSJjZW50ZXIiPgoJPGZvcm0gY2xhc3M9InB1cmUtZm9ybSIgbWV0aG9kPSJQT1NUIiBhY3Rpb249Ii9sb2dpbiI+CiAgICA8ZmllbGRzZXQgY2xhc3M9InB1cmUtZ3JvdXAiPgogICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBuYW1lPSJ1c2VybmFtZSIgY2xhc3M9InB1cmUtaW5wdXQtMS0yIiBwbGFjZWhvbGRlcj0iVXNlcm5hbWUiPgogICAgICAgIDxpbnB1dCB0eXBlPSJwYXNzd29yZCIgbmFtZT0icGFzc3dvcmQiIGNsYXNzPSJwdXJlLWlucHV0LTEtMiIgcGxhY2Vob2xkZXI9IlBhc3N3b3JkIj4KICAgIDwvZmllbGRzZXQ+CiAgICA8YnV0dG9uIHR5cGU9InN1Ym1pdCIgY2xhc3M9InB1cmUtYnV0dG9uIHB1cmUtaW5wdXQtMS0yIHB1cmUtYnV0dG9uLXByaW1hcnkiPlNpZ24gaW48L2J1dHRvbj4KPC9mb3JtPgo8L2Rpdj4KPC9ib2R5Pgo8L2h0bWw+`)
	FilebrowseHTML string = common.Base64Decode(`PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MS4wIj4KCiAgICA8dGl0bGU+Qm90bmV0IEMmQzogRmlsZSBCcm93c2VyPC90aXRsZT4KCiAgICA8bGluayByZWw9InN0eWxlc2hlZXQiIGhyZWY9Imh0dHA6Ly95dWkueWFob29hcGlzLmNvbS9wdXJlLzAuNi4wL3B1cmUtbWluLmNzcyI+CgkKCTxzdHlsZT4KYm9keSB7CiAgYmFja2dyb3VuZDogIzc2Yjg1MjsgLyogZmFsbGJhY2sgZm9yIG9sZCBicm93c2VycyAqLwogIGJhY2tncm91bmQ6IC13ZWJraXQtbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiAtbW96LWxpbmVhci1ncmFkaWVudChyaWdodCwgIzc2Yjg1MiwgIzhEQzI2Rik7CiAgYmFja2dyb3VuZDogLW8tbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiBsaW5lYXItZ3JhZGllbnQodG8gbGVmdCwgIzc2Yjg1MiwgIzhEQzI2Rik7CiAgLXdlYmtpdC1mb250LXNtb290aGluZzogYW50aWFsaWFzZWQ7CiAgLW1vei1vc3gtZm9udC1zbW9vdGhpbmc6IGdyYXlzY2FsZTsgICAgICAKfQogICAgICAgIC5idXR0b24tc3VjY2VzcywKICAgICAgICAuYnV0dG9uLWVycm9yLAogICAgICAgIC5idXR0b24td2FybmluZywKICAgICAgICAuYnV0dG9uLXNlY29uZGFyeSB7CiAgICAgICAgICAgIGNvbG9yOiB3aGl0ZTsKICAgICAgICAgICAgYm9yZGVyLXJhZGl1czogNHB4OwogICAgICAgICAgICB0ZXh0LXNoYWRvdzogMCAxcHggMXB4IHJnYmEoMCwgMCwgMCwgMC4yKTsKICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi1zdWNjZXNzIHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDI4LCAxODQsIDY1KTsgLyogdGhpcyBpcyBhIGdyZWVuICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24tZXJyb3IgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoMjAyLCA2MCwgNjApOyAvKiB0aGlzIGlzIGEgbWFyb29uICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24td2FybmluZyB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYigyMjMsIDExNywgMjApOyAvKiB0aGlzIGlzIGFuIG9yYW5nZSAqLwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLXNlY29uZGFyeSB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYig2NiwgMTg0LCAyMjEpOyAvKiB0aGlzIGlzIGEgbGlnaHQgYmx1ZSAqLwogICAgICAgIH0KCQkgICogeyBtYXJnaW46IDA7IHBhZGRpbmc6IDA7IH0KCSAgIHAgeyBwYWRkaW5nOiAxMHB4OyB9CgkgICAjbGVmdCB7IHBvc2l0aW9uOiBhYnNvbHV0ZTsgbGVmdDogMDsgIHdpZHRoOiA1MCU7IH0KCSAgICNyaWdodCB7IHBvc2l0aW9uOiBhYnNvbHV0ZTsgcmlnaHQ6IDA7ICB3aWR0aDogNTAlOyB9Cjwvc3R5bGU+CjwvaGVhZD4KCjxib2R5Pgo8ZGl2IGNsYXNzPSJwdXJlLW1lbnUgcHVyZS1tZW51LWhvcml6b250YWwiIGFsaWduPSJjZW50ZXIiPgogICAgPHVsIGNsYXNzPSJwdXJlLW1lbnUtbGlzdCI+Cgk8c3BhbiBjbGFzcz0icHVyZS1tZW51LWhlYWRpbmciIGFsaWduPSJsZWZ0Ij5Cb3RuZXQgQyZDPC9zcGFuPgogICAgICAgIDxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0gcHVyZS1tZW51LXNlbGVjdGVkIj48YSBocmVmPSJ7SE9TVH0vcGFuZWwiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+Q29udHJvbCBQYW5lbDwvYT48L2xpPgoJCTxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0gcHVyZS1tZW51LXNlbGVjdGVkIj48YSBocmVmPSJ7SE9TVH0vZmlsZXMiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+RmlsZSBCcm93c2VyPC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtc2VsZWN0ZWQiPjxhIGhyZWY9IntIT1NUfS90YXNrbWdyIiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPlRhc2sgTWFuYWdtZW50PC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtc2VsZWN0ZWQiPjxhIGhyZWY9IntIT1NUfS9zeXNsb2dzIiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPkMmQyBMb2dzPC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtaGFzLWNoaWxkcmVuIHB1cmUtbWVudS1hbGxvdy1ob3ZlciI+CiAgICAgICAgICAgIDxhICBpZD0ibWVudUxpbmsxIiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPlRvb2xzPC9hPgogICAgICAgICAgICA8dWwgY2xhc3M9InB1cmUtbWVudS1jaGlsZHJlbiI+CiAgICAgICAgICAgICAgICA8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIj48YSBocmVmPSJ7SE9TVH0vYnVpbGQiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+Qm90IEJ1aWxkZXI8L2E+PC9saT4KICAgICAgICAgICAgICAgIDxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0iPjxhIGhyZWY9IntIT1NUfS9kYnVpbGQiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+RG93bmxvYWRlciBCdWlsZGVyPC9hPjwvbGk+CiAgICAgICAgICAgICAgICA8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIj48YSBocmVmPSJ7SE9TVH0vdG9vbHMiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+Q3J5cHRvZ3JhcGh5IFRvb2xzPC9hPjwvbGk+CiAgICAgICAgICAgIDwvdWw+CiAgICAgICAgPC9saT4KCQk8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIHB1cmUtbWVudS1oYXMtY2hpbGRyZW4gcHVyZS1tZW51LWFsbG93LWhvdmVyIj4KICAgICAgICAgICAgPGEgaWQ9Im1lbnVMaW5rMSIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5Vc2VyIE1hbmFnbWVudDwvYT4KICAgICAgICAgICAgPHVsIGNsYXNzPSJwdXJlLW1lbnUtY2hpbGRyZW4iPgogICAgICAgICAgICAgICAgPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSI+PGEgaHJlZj0ie0hPU1R9L2FjY291bnQiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+QWNjb3VudCBTZXR0aW5nczwvYT48L2xpPgogICAgICAgICAgICAgICAgPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSI+PGEgaHJlZj0ie0hPU1R9L2xvZ291dCIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5Mb2dvdXQ8L2E+PC9saT4KICAgICAgICAgICAgPC91bD4KICAgICAgICA8L2xpPgogICAgPC91bD4KPC9kaXY+Cjxicj4KPGJyPgo8ZGl2IGFsaWduPSJjZW50ZXIiPjxoMj5GaWxlIEJyb3dzZXI8L2gyPjwvZGl2Pgo8YnI+CntTVEFUU30KPGJyPgoJPGRpdiBhbGlnbj0iY2VudGVyIj4KCTx1bD4KCXtGSUxFU30KCTwvdWw+Cgk8L2Rpdj4KPC9ib2R5Pgo8L2h0bWw+`)
	InfoHTML       string = common.Base64Decode(`PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MS4wIj4KCiAgICA8dGl0bGU+Qm90bmV0IEMmQzogSW5mb3JtYXRpb248L3RpdGxlPgoKICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cDovL3l1aS55YWhvb2FwaXMuY29tL3B1cmUvMC42LjAvcHVyZS1taW4uY3NzIj4KCTxzY3JpcHQgc3JjPSJodHRwczovL3VzZS5mb250YXdlc29tZS5jb20vZTNlNmEzOThkZi5qcyI+PC9zY3JpcHQ+CgkKCTxzdHlsZT4KYm9keSB7CiAgYmFja2dyb3VuZDogIzc2Yjg1MjsgLyogZmFsbGJhY2sgZm9yIG9sZCBicm93c2VycyAqLwogIGJhY2tncm91bmQ6IC13ZWJraXQtbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiAtbW96LWxpbmVhci1ncmFkaWVudChyaWdodCwgIzc2Yjg1MiwgIzhEQzI2Rik7CiAgYmFja2dyb3VuZDogLW8tbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiBsaW5lYXItZ3JhZGllbnQodG8gbGVmdCwgIzc2Yjg1MiwgIzhEQzI2Rik7CiAgLXdlYmtpdC1mb250LXNtb290aGluZzogYW50aWFsaWFzZWQ7CiAgLW1vei1vc3gtZm9udC1zbW9vdGhpbmc6IGdyYXlzY2FsZTsgICAgICAKfQogICAgICAgIC5idXR0b24tc3VjY2VzcywKICAgICAgICAuYnV0dG9uLWVycm9yLAogICAgICAgIC5idXR0b24td2FybmluZywKICAgICAgICAuYnV0dG9uLXNlY29uZGFyeSB7CiAgICAgICAgICAgIGNvbG9yOiB3aGl0ZTsKICAgICAgICAgICAgYm9yZGVyLXJhZGl1czogNHB4OwogICAgICAgICAgICB0ZXh0LXNoYWRvdzogMCAxcHggMXB4IHJnYmEoMCwgMCwgMCwgMC4yKTsKICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi1zdWNjZXNzIHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDI4LCAxODQsIDY1KTsgLyogdGhpcyBpcyBhIGdyZWVuICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24tZXJyb3IgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoMjAyLCA2MCwgNjApOyAvKiB0aGlzIGlzIGEgbWFyb29uICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24td2FybmluZyB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYigyMjMsIDExNywgMjApOyAvKiB0aGlzIGlzIGFuIG9yYW5nZSAqLwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLXNlY29uZGFyeSB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYig2NiwgMTg0LCAyMjEpOyAvKiB0aGlzIGlzIGEgbGlnaHQgYmx1ZSAqLwogICAgICAgIH0KCQkgICogeyBtYXJnaW46IDA7IHBhZGRpbmc6IDA7IH0KCSAgIHAgeyBwYWRkaW5nOiAxMHB4OyB9CgkgICAjbGVmdCB7IHBvc2l0aW9uOiBhYnNvbHV0ZTsgbGVmdDogMDsgIHdpZHRoOiA1MCU7IH0KCSAgICNyaWdodCB7IHBvc2l0aW9uOiBhYnNvbHV0ZTsgcmlnaHQ6IDA7ICB3aWR0aDogNTAlOyB9Cjwvc3R5bGU+Cgo8L2hlYWQ+Cgo8Ym9keT4KPGRpdiBjbGFzcz0icHVyZS1tZW51IHB1cmUtbWVudS1ob3Jpem9udGFsIiBhbGlnbj0iY2VudGVyIj4KICAgIDx1bCBjbGFzcz0icHVyZS1tZW51LWxpc3QiPgoJPHNwYW4gY2xhc3M9InB1cmUtbWVudS1oZWFkaW5nIiBhbGlnbj0ibGVmdCI+Qm90bmV0IEMmQzwvc3Bhbj4KICAgICAgICA8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIHB1cmUtbWVudS1zZWxlY3RlZCI+PGEgaHJlZj0iLi9wYW5lbCIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5Db250cm9sIFBhbmVsPC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtc2VsZWN0ZWQiPjxhIGhyZWY9Ii4vZmlsZXMiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+RmlsZSBCcm93c2VyPC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtaGFzLWNoaWxkcmVuIHB1cmUtbWVudS1hbGxvdy1ob3ZlciI+CiAgICAgICAgICAgIDxhIGlkPSJtZW51TGluazEiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+VXNlciBNYW5hZ21lbnQ8L2E+CiAgICAgICAgICAgIDx1bCBjbGFzcz0icHVyZS1tZW51LWNoaWxkcmVuIj4KICAgICAgICAgICAgICAgIDxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0iPjxhIGhyZWY9Ii4vbG9nb3V0IiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPkxvZ291dDwvYT48L2xpPgogICAgICAgICAgICA8L3VsPgogICAgICAgIDwvbGk+CiAgICA8L3VsPgo8L2Rpdj4KPGJyPgo8YnI+CjxkaXYgYWxpZ249ImNlbnRlciI+PGgyPkJvdCBJbmZvcm1hdGlvbjwvaDI+PC9kaXY+Cjxicj4Ke1NUQVRTfQo8YnI+CjxkaXYgYWxpZ249ImNlbnRlciI+CjxpbWcgYWx0PSIiIHNyYz0iLi9maWxlcy97R1VJRH0vU2NyZWVuc2hvdHMvRGVmYXVsdC5wbmciIHN0eWxlPSJ3aWR0aDogNTAwcHg7IGhlaWdodDogMjgxcHg7IiAvPgo8YnI+CjxhPkxhc3QgQ2hlY2stSW46IHtMQVNEQVRFfSA8L2E+Cjxicj4KPGE+Q2xpZW50IFZlcnNpb246IHtWRVJTSU9OfSA8L2E+Cjxicj4KCTxhIGNsYXNzPSJidXR0b24td2FybmluZyBwdXJlLWJ1dHRvbiIgaHJlZj0iLi9wdXJnZT9ndWlkPXtHVUlEfSI+PGkgY2xhc3M9ImZhIGZhLXRpbWVzIGZhLWxnIj48L2k+ICBEZWxldGUgZnJvbSBEYXRhYmFzZTwvYT4KPGJyPgo8ZGl2IGlkPSJsZWZ0IiBhbGlnbj0iY2VudGVyIj4KPGZvcm0gY2xhc3M9InB1cmUtZm9ybSBwdXJlLWZvcm0tc3RhY2tlZCI+Cgk8bGFiZWwgZm9yPSIxIj5HVUlEPC9sYWJlbD4KICAgIDxpbnB1dCBuYW1lPSIxIiB0eXBlPSJ0ZXh0IiB2YWx1ZT0ie0dVSUR9IiByZWFkb25seT4KCTxsYWJlbCBmb3I9IjIiPklQIEFkZHJlc3M8L2xhYmVsPgogICAgPGlucHV0IG5hbWU9IjIiIHR5cGU9InRleHQiIHZhbHVlPSJ7SVB9IiByZWFkb25seT4KCTxsYWJlbCBmb3I9IjIiPldobyBhbSBJPC9sYWJlbD4KICAgIDxpbnB1dCBuYW1lPSIyIiB0eXBlPSJ0ZXh0IiB2YWx1ZT0ie1dIT0FNSX0iIHJlYWRvbmx5PgoJPGxhYmVsIGZvcj0iMiI+T3BlcmF0aW5nIFN5c3RlbTwvbGFiZWw+CiAgICA8aW5wdXQgbmFtZT0iMiIgdHlwZT0idGV4dCIgdmFsdWU9IntPU30iIHJlYWRvbmx5PgoJPGxhYmVsIGZvcj0iMiI+RGF0ZSBJbnN0YWxsZWQ8L2xhYmVsPgogICAgPGlucHV0IG5hbWU9IjIiIHR5cGU9InRleHQiIHZhbHVlPSJ7SU5TREFURX0iIHJlYWRvbmx5PgoJPGxhYmVsIGZvcj0iMiI+SGF2ZSBBZG1pbiBSaWdodHM/PC9sYWJlbD4KICAgIDxpbnB1dCBuYW1lPSIyIiB0eXBlPSJ0ZXh0IiB2YWx1ZT0ie0FETUlOfSIgcmVhZG9ubHk+Cgk8bGFiZWwgZm9yPSIyIj5BbnRpLVZpcnVzPC9sYWJlbD4KICAgIDxpbnB1dCBuYW1lPSIyIiB0eXBlPSJ0ZXh0IiB2YWx1ZT0ie0FWfSIgcmVhZG9ubHk+Cgk8bGFiZWwgZm9yPSIyIj5DUFUgSW5mb3JtYXRpb248L2xhYmVsPgogICAgPGlucHV0IG5hbWU9IjIiIHR5cGU9InRleHQiIHZhbHVlPSJ7Q1BVfSIgcmVhZG9ubHk+Cgk8bGFiZWwgZm9yPSIyIj5HUFUgSW5mb3JtYXRpb248L2xhYmVsPgogICAgPGlucHV0IG5hbWU9IjIiIHR5cGU9InRleHQiIHZhbHVlPSJ7R1BVfSIgcmVhZG9ubHk+Cgk8L2Rpdj4KCTxkaXYgaWQ9InJpZ2h0IiBhbGlnbj0iY2VudGVyIj4KCTxhIGhyZWY9Ii4vZmlsZXMve0dVSUR9L0ZpbGVzLyIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5GaWxlcyBGb2xkZXI8L2E+Cgk8YSBocmVmPSIuL2ZpbGVzL3tHVUlEfS9LZXlsb2dzLyIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5LZXlsb2dzIEZvbGRlcjwvYT4KCTxhIGhyZWY9Ii4vZmlsZXMve0dVSUR9L1NjcmVlbnNob3RzLyIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5TY3JlZW5zaG90cyBGb2xkZXI8L2E+Cgk8YnI+Cgk8YSBocmVmPSIuL2ZpbGVzL3tHVUlEfS9GaWxlcy9JbnN0YWxsZWQgU29mdHdhcmUudHh0IiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPkluc3RhbGxlZCBTb2Z0d2FyZTwvYT4KCTxhIGhyZWY9Ii4vZmlsZXMve0dVSUR9L0ZpbGVzL0lQIENvbmZpZy50eHQiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+SVAgQ29uZmlnPC9hPgoJPGEgaHJlZj0iLi9maWxlcy97R1VJRH0vRmlsZXMvU3lzdGVtIEluZm9ybWF0aW9uLnR4dCIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5TeXN0ZW0gSW5mb3JtYXRpb248L2E+Cgk8YSBocmVmPSIuL2ZpbGVzL3tHVUlEfS9GaWxlcy9XaUZpIEluZm9ybWF0aW9uLnR4dCIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5XaUZpIEluZm9ybWF0aW9uPC9hPgoJPC9kaXY+CjwvZm9ybT4KPC9kaXY+CjwvYm9keT4=`)
	BotTableHTML   string = common.Base64Decode(`IDx0Ym9keT4KICAgICAgICA8dHIgY2xhc3M9InB1cmUtdGFibGUtb2RkIj4KCQkJPHRkPjxpbnB1dCBuYW1lPSJzZWxlY3RlZGJvdCIgdHlwZT0iY2hlY2tib3giIHZhbHVlPSJ7R1VJRH0iPjwvdGQ+CiAgICAgICAgICAgIDx0ZD48YSBocmVmPSIuL2ZpbGVzL3tHVUlEfSI+e0dVSUR9PC9hPjwvdGQ+CiAgICAgICAgICAgIDx0ZD48YSBocmVmPSJodHRwczovL25vcmVmZXJyYWwubmV0Lz9odHRwOi8ve0lQfS5pcGFkZHJlc3MuY29tIj57SVB9PC9hPjwvdGQ+CiAgICAgICAgICAgIDx0ZD57V0hPQU1JfTwvdGQ+CiAgICAgICAgICAgIDx0ZD57T1N9PC90ZD4KCQkJPHRkPntBRE1JTn08L3RkPgoJCQk8dGQ+e0xBU0RBVEV9PC90ZD4KCQkJPHRkPjxhIGhyZWY9Ii9pbmZvP2d1aWQ9e0dVSUR9Ij48aSBjbGFzcz0iZmEgZmEtaWQtY2FyZC1vIGZhLTJ4Ij48L2k+PC9hPjwvdGQ+CiAgICAgICAgPC90cj4KICAgIDwvdGJvZHk+`)
	PanelHTML      string = common.Base64Decode(`PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MS4wIj4KCiAgICA8dGl0bGU+Qm90bmV0IEMmQzogQ29udHJvbCBQYW5lbDwvdGl0bGU+CgogICAgPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiBocmVmPSJodHRwOi8veXVpLnlhaG9vYXBpcy5jb20vcHVyZS8wLjYuMC9wdXJlLW1pbi5jc3MiPgoJPHNjcmlwdCBzcmM9Imh0dHBzOi8vdXNlLmZvbnRhd2Vzb21lLmNvbS9lM2U2YTM5OGRmLmpzIj48L3NjcmlwdD4KCQoJPHN0eWxlPgpib2R5IHsKICBiYWNrZ3JvdW5kOiAjNzZiODUyOyAvKiBmYWxsYmFjayBmb3Igb2xkIGJyb3dzZXJzICovCiAgYmFja2dyb3VuZDogLXdlYmtpdC1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IC1tb3otbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiAtby1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IGxpbmVhci1ncmFkaWVudCh0byBsZWZ0LCAjNzZiODUyLCAjOERDMjZGKTsKICAtd2Via2l0LWZvbnQtc21vb3RoaW5nOiBhbnRpYWxpYXNlZDsKICAtbW96LW9zeC1mb250LXNtb290aGluZzogZ3JheXNjYWxlOyAgICAgIAp9CiAgICAgICAgLmJ1dHRvbi1zdWNjZXNzLAogICAgICAgIC5idXR0b24tZXJyb3IsCiAgICAgICAgLmJ1dHRvbi13YXJuaW5nLAogICAgICAgIC5idXR0b24tc2Vjb25kYXJ5IHsKICAgICAgICAgICAgY29sb3I6IHdoaXRlOwogICAgICAgICAgICBib3JkZXItcmFkaXVzOiA0cHg7CiAgICAgICAgICAgIHRleHQtc2hhZG93OiAwIDFweCAxcHggcmdiYSgwLCAwLCAwLCAwLjIpOwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLXN1Y2Nlc3MgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoMjgsIDE4NCwgNjUpOyAvKiB0aGlzIGlzIGEgZ3JlZW4gKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi1lcnJvciB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYigyMDIsIDYwLCA2MCk7IC8qIHRoaXMgaXMgYSBtYXJvb24gKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi13YXJuaW5nIHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDIyMywgMTE3LCAyMCk7IC8qIHRoaXMgaXMgYW4gb3JhbmdlICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24tc2Vjb25kYXJ5IHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDY2LCAxODQsIDIyMSk7IC8qIHRoaXMgaXMgYSBsaWdodCBibHVlICovCiAgICAgICAgfQoJCSAgKiB7IG1hcmdpbjogMDsgcGFkZGluZzogMDsgfQoJICAgcCB7IHBhZGRpbmc6IDEwcHg7IH0KCSAgICNsZWZ0IHsgcG9zaXRpb246IGFic29sdXRlOyBsZWZ0OiAwOyAgd2lkdGg6IDUwJTsgfQoJICAgI3JpZ2h0IHsgcG9zaXRpb246IGFic29sdXRlOyByaWdodDogMDsgIHdpZHRoOiA1MCU7IH0KPC9zdHlsZT4KPC9oZWFkPgo8Ym9keT4KPGRpdiBjbGFzcz0icHVyZS1tZW51IHB1cmUtbWVudS1ob3Jpem9udGFsIiBhbGlnbj0iY2VudGVyIj4KICAgIDx1bCBjbGFzcz0icHVyZS1tZW51LWxpc3QiPgoJPHNwYW4gY2xhc3M9InB1cmUtbWVudS1oZWFkaW5nIiBhbGlnbj0ibGVmdCI+Qm90bmV0IEMmQzwvc3Bhbj4KICAgICAgICA8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIHB1cmUtbWVudS1zZWxlY3RlZCI+PGEgaHJlZj0iLi9wYW5lbCIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5Db250cm9sIFBhbmVsPC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtc2VsZWN0ZWQiPjxhIGhyZWY9Ii4vZmlsZXMiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+RmlsZSBCcm93c2VyPC9hPjwvbGk+CgkJPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtaGFzLWNoaWxkcmVuIHB1cmUtbWVudS1hbGxvdy1ob3ZlciI+CiAgICAgICAgICAgIDxhIGlkPSJtZW51TGluazEiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+VXNlciBNYW5hZ21lbnQ8L2E+CiAgICAgICAgICAgIDx1bCBjbGFzcz0icHVyZS1tZW51LWNoaWxkcmVuIj4KICAgICAgICAgICAgICAgIDxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0iPjxhIGhyZWY9Ii4vbG9nb3V0IiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPkxvZ291dDwvYT48L2xpPgogICAgICAgICAgICA8L3VsPgogICAgICAgIDwvbGk+CiAgICA8L3VsPgo8L2Rpdj4KPGJyPgo8YnI+CjxkaXYgYWxpZ249ImNlbnRlciI+PGgyPkNvbnRyb2wgUGFuZWw8L2gyPjwvZGl2Pgo8YnI+CntTVEFUU30KPGJyPgo8Zm9ybSBjbGFzcz0icHVyZS1mb3JtIiBtZXRob2Q9IlBPU1QiIGFjdGlvbj0iLi9zZW5kY21kIj4KPGRpdiBhbGlnbj0iY2VudGVyIj4KPHRhYmxlIGNsYXNzPSJwdXJlLXRhYmxlIiBhbGlnbj0iY2VudGVyIj4KICAgIDx0aGVhZD4KICAgICAgICA8dHI+CgkJCTx0aD48L3RoPgogICAgICAgICAgICA8dGg+R1VJRDwvdGg+CiAgICAgICAgICAgIDx0aD5JUDwvdGg+CiAgICAgICAgICAgIDx0aD5XSE8gQU0gSTwvdGg+CiAgICAgICAgICAgIDx0aD5PUzwvdGg+CgkJCTx0aD5BRE1JTiBSSUdIVFM8L3RoPgoJCQk8dGg+TEFTVCBDSEVDSy1JTjwvdGg+CgkJCTx0aD48L3RoPgogICAgICAgIDwvdHI+CiAgICA8L3RoZWFkPgogICB7UkFXVEFCTEV9CjwvdGFibGU+Cjxicj4KPGRpdiBpZD0ibGVmdCIgYWxpZ249ImNlbnRlciI+Cgk8YSBjbGFzcz0iYnV0dG9uLXN1Y2Nlc3MgcHVyZS1idXR0b24iIGhyZWY9Ij9wYWdlPXtCQUNLfSI+PGkgY2xhc3M9ImZhIGZhLWFycm93LWxlZnQgZmEtbGciPjwvaT4gIFByZXZpb3VzPC9hPgo8L2Rpdj4KPGRpdiBpZD0icmlnaHQiIGFsaWduPSJjZW50ZXIiPgoJPGEgY2xhc3M9ImJ1dHRvbi1zdWNjZXNzIHB1cmUtYnV0dG9uIiBocmVmPSI/cGFnZT17TkVYVH0iPk5leHQgIDxpIGNsYXNzPSJmYSBmYS1hcnJvdy1yaWdodCBmYS1sZyI+PC9pPjwvYT4KPC9kaXY+CjwvZGl2Pgo8YnI+Cjxicj4KPGJyPgo8ZGl2IGlkPSJyaWdodCIgYWxpZ249ImNlbnRlciI+CiAgICA8ZmllbGRzZXQgY2xhc3M9InB1cmUtZ3JvdXAiPgoJPGxlZ2VuZD5Db21tYW5kPC9sZWdlbmQ+CiAgICAgICAgIDxzZWxlY3QgaWQ9ImJvdHNzZWxlY3Rpb24iIG5hbWU9ImJvdHNzZWxlY3Rpb24iIGNsYXNzPSJwdXJlLWlucHV0LTEtMiI+CiAgICAgICAgICAgICAgICAgICAgPG9wdGlvbiB2YWx1ZT0iMDAwIj5BbGw8L29wdGlvbj4KICAgICAgICAgICAgICAgICAgICA8b3B0aW9uIHZhbHVlPSJzZWxlY3RlZCI+U2VsZWN0ZWQ8L29wdGlvbj4KICAgICAgICAgICAgICAgIDwvc2VsZWN0PgogICAgICAgIDxzZWxlY3QgaWQ9ImNvbW1hbmR0eXBlIiBuYW1lPSJjb21tYW5kdHlwZSIgY2xhc3M9InB1cmUtaW5wdXQtMS0yIj4KICAgICAgICAgICAgICAgICAgICA8b3B0aW9uIHZhbHVlPSIweDEiPk9wZW4gVVJMPC9vcHRpb24+CiAgICAgICAgICAgICAgICAgICAgPG9wdGlvbiB2YWx1ZT0iMHgyIj5SdW4gQXBwbGljYXRpb248L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIweDUiPkRvd25sb2FkIGFuZCBSdW48L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIyeDQiPkVuYWJsZS9EaXNhYmxlIEtleWxvZ2dlcjwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjB4MCI+S2lsbCBNZTwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjB4NiI+UnVuIFBvd2Vyc2hlbGw8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIweDciPlNwcmVhZGVyPC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iMHg4Ij5TdGFydCBXZWJzZXJ2ZXI8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIweDkiPkVkaXQgV2Vic2VydmVyIFBhZ2U8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIxeDAiPkhpZGUgUHJvY2VzcyBXaW5kb3c8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIxeDEiPlNlZWQgVG9ycmVudDwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjF4MiI+U3lzdGVtIFBvd2VyIE9wdGlvbnM8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIxeDMiPlNldCBIb21lcGFnZTwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjF4NCI+U2V0IFdhbGxwYXBlcjwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjF4NSI+RWRpdCBIT1NUIEZpbGU8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIxeDYiPlVuaW5zdGFsbCBNZTwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjF4NyI+T3BlbiBQb3J0PC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iMXg4Ij5SdW4gU2NyaXB0PC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iMXg5Ij5SdW4gQ29tbWFuZDwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjJ4MCI+RW5hYmxlIFJldmVyc2UgUHJveHkgU2VydmVyPC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iMngxIj5QdXNoIEZpbGUgdG8gQ2xpZW50PC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iMngyIj5LaWxsIFByb2Nlc3M8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIyeDMiPlVwZGF0ZSBDbGllbnQ8L29wdGlvbj4KICAgICAgICAgICAgICAgIDwvc2VsZWN0PgogICAgPC9maWVsZHNldD4KCiAgICA8ZmllbGRzZXQgY2xhc3M9InB1cmUtZ3JvdXAiPgogICAgICAgIDx0ZXh0YXJlYSBpZD0iYXJnMSIgbmFtZT0iYXJnMSIgY2xhc3M9InB1cmUtaW5wdXQtMS0yIiBwbGFjZWhvbGRlcj0iQXJndW1lbnRzIj48L3RleHRhcmVhPgoJCTxicj4KCQkgVXNlICd8JyB0byBzcGxpdCBBcmd1bWVudHMKICAgIDwvZmllbGRzZXQ+CgogICAgPGJ1dHRvbiB0eXBlPSJzdWJtaXQiIGNsYXNzPSJwdXJlLWJ1dHRvbiBwdXJlLWlucHV0LTEtMiBwdXJlLWJ1dHRvbi1wcmltYXJ5Ij48aSBjbGFzcz0iZmEgZmEtcG9kY2FzdCBmYS1sZyI+PC9pPiAgU2VuZCBDb21tYW5kPC9idXR0b24+IDxhIGNsYXNzPSJidXR0b24tc2Vjb25kYXJ5IHB1cmUtYnV0dG9uIiBocmVmPSIuL3JlZnJlc2g/Z3VpZD17R1VJRH0iPjxpIGNsYXNzPSJmYSBmYS1yZWZyZXNoIGZhLWxnIj48L2k+ICBSZWZyZXNoIEluZm88L2E+CjwvZm9ybT4KPC9kaXY+CjxkaXYgaWQ9ImxlZnQiIGFsaWduPSJjZW50ZXIiPgo8Zm9ybSBjbGFzcz0icHVyZS1mb3JtIiBtZXRob2Q9IlBPU1QiIGFjdGlvbj0iLi9jbWRkZG9zIj4KPGZpZWxkc2V0IGNsYXNzPSJwdXJlLWdyb3VwIj4KPGxlZ2VuZD5ERG9TIEF0dGFjazwvbGVnZW5kPgogICAgICAgICA8c2VsZWN0IGlkPSJkZG9zbW9kZSIgbmFtZT0iZGRvc21vZGUiIGNsYXNzPSJwdXJlLWlucHV0LTEtMiI+CiAgICAgICAgICAgICAgICAgICAgPG9wdGlvbiB2YWx1ZT0iMCI+SFRUUCBHZXQ8L29wdGlvbj4KICAgICAgICAgICAgICAgICAgICA8b3B0aW9uIHZhbHVlPSI1Ij5UQ1AgRmxvb2Q8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSI0Ij5VRFAgRmxvb2Q8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIxIj5TbG93bG9yaXM8L29wdGlvbj4KCQkJCQk8b3B0aW9uIHZhbHVlPSIyIj5IVUxLPC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iMyI+VExTIEZsb29kPC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iNyI+QmFuZHdpZHRoIERyYWluPC9vcHRpb24+CgkJCQkJPG9wdGlvbiB2YWx1ZT0iNiI+R29sZGVuIEV5ZTwvb3B0aW9uPgoJCQkJCTxvcHRpb24gdmFsdWU9IjgiPkFjZTwvb3B0aW9uPgogICAgICAgICAgICAgICAgPC9zZWxlY3Q+CgkJCQk8L2ZpZWxkc2V0PgogICAgICAgIDxpbnB1dCB0eXBlPSJ0ZXh0IiBuYW1lPSJpcCIgY2xhc3M9InB1cmUtaW5wdXQtMS0yIiBwbGFjZWhvbGRlcj0iMTI3LjAuMC4xIj4KICAgICAgICA8aW5wdXQgdHlwZT0idGV4dCIgbmFtZT0icG9ydCIgY2xhc3M9InB1cmUtaW5wdXQtMS0yIiBwbGFjZWhvbGRlcj0iODAiPgoJCTxmaWVsZHNldCBjbGFzcz0icHVyZS1ncm91cCI+CiAgICAgICAgPGlucHV0IHR5cGU9InRleHQiIG5hbWU9InRocmVhZHMiIGNsYXNzPSJwdXJlLWlucHV0LTEtMiIgcGxhY2Vob2xkZXI9IlRocmVhZHMgMTAwMCI+CgkJPGlucHV0IHR5cGU9InRleHQiIG5hbWU9ImludGVydmFsIiBjbGFzcz0icHVyZS1pbnB1dC0xLTIiIHBsYWNlaG9sZGVyPSJJbnRlcnZhbCA1MDAgKiBNaWxsaXNlY29uZHMiPgoJCTxicj4KCQkgRERvUyBBdHRhY2tzIHV0aWxpemUgYWxsIGJvdHMKCQk8L2ZpZWxkc2V0PgogICAgPGJ1dHRvbiB0eXBlPSJzdWJtaXQiIGNsYXNzPSJwdXJlLWJ1dHRvbiBwdXJlLWlucHV0LTEtMiBwdXJlLWJ1dHRvbi1wcmltYXJ5Ij48aSBjbGFzcz0iZmEgZmEtcG9kY2FzdCBmYS1sZyI+PC9pPiAgU2VuZCBDb21tYW5kPC9idXR0b24+IDxhIGNsYXNzPSJidXR0b24tZXJyb3IgcHVyZS1idXR0b24iIGhyZWY9Ii4vc3RvcGRkb3MiPjxpIGNsYXNzPSJmYSBmYS1zdG9wIGZhLWxnIj48L2k+ICBTdG9wIEF0dGFjazwvYT4KPC9kaXY+CjwvZm9ybT4KCjwvYm9keT4=`)
	ErrorHTML      string = common.Base64Decode(`PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MS4wIj4KCiAgICA8dGl0bGU+Qm90bmV0IEMmQzogRXJyb3I8L3RpdGxlPgoKICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cDovL3l1aS55YWhvb2FwaXMuY29tL3B1cmUvMC42LjAvcHVyZS1taW4uY3NzIj4KCQoJPHN0eWxlPgpib2R5IHsKICBiYWNrZ3JvdW5kOiAjNzZiODUyOyAvKiBmYWxsYmFjayBmb3Igb2xkIGJyb3dzZXJzICovCiAgYmFja2dyb3VuZDogLXdlYmtpdC1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IC1tb3otbGluZWFyLWdyYWRpZW50KHJpZ2h0LCAjNzZiODUyLCAjOERDMjZGKTsKICBiYWNrZ3JvdW5kOiAtby1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IGxpbmVhci1ncmFkaWVudCh0byBsZWZ0LCAjNzZiODUyLCAjOERDMjZGKTsKICAtd2Via2l0LWZvbnQtc21vb3RoaW5nOiBhbnRpYWxpYXNlZDsKICAtbW96LW9zeC1mb250LXNtb290aGluZzogZ3JheXNjYWxlOyAgICAgIAp9CiAgICAgICAgLmJ1dHRvbi1zdWNjZXNzLAogICAgICAgIC5idXR0b24tZXJyb3IsCiAgICAgICAgLmJ1dHRvbi13YXJuaW5nLAogICAgICAgIC5idXR0b24tc2Vjb25kYXJ5IHsKICAgICAgICAgICAgY29sb3I6IHdoaXRlOwogICAgICAgICAgICBib3JkZXItcmFkaXVzOiA0cHg7CiAgICAgICAgICAgIHRleHQtc2hhZG93OiAwIDFweCAxcHggcmdiYSgwLCAwLCAwLCAwLjIpOwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLXN1Y2Nlc3MgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoMjgsIDE4NCwgNjUpOyAvKiB0aGlzIGlzIGEgZ3JlZW4gKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi1lcnJvciB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYigyMDIsIDYwLCA2MCk7IC8qIHRoaXMgaXMgYSBtYXJvb24gKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi13YXJuaW5nIHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDIyMywgMTE3LCAyMCk7IC8qIHRoaXMgaXMgYW4gb3JhbmdlICovCiAgICAgICAgfQogICAgICAgIC5idXR0b24tc2Vjb25kYXJ5IHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDY2LCAxODQsIDIyMSk7IC8qIHRoaXMgaXMgYSBsaWdodCBibHVlICovCiAgICAgICAgfQoJCSAgKiB7IG1hcmdpbjogMDsgcGFkZGluZzogMDsgfQoJICAgcCB7IHBhZGRpbmc6IDEwcHg7IH0KCSAgICNsZWZ0IHsgcG9zaXRpb246IGFic29sdXRlOyBsZWZ0OiAwOyAgd2lkdGg6IDUwJTsgfQoJICAgI3JpZ2h0IHsgcG9zaXRpb246IGFic29sdXRlOyByaWdodDogMDsgIHdpZHRoOiA1MCU7IH0KPC9zdHlsZT4KCjwvaGVhZD4KCjxib2R5Pgo8ZGl2IGNsYXNzPSJwdXJlLW1lbnUgcHVyZS1tZW51LWhvcml6b250YWwiIGFsaWduPSJjZW50ZXIiPgogICAgPHVsIGNsYXNzPSJwdXJlLW1lbnUtbGlzdCI+Cgk8c3BhbiBjbGFzcz0icHVyZS1tZW51LWhlYWRpbmciIGFsaWduPSJsZWZ0Ij5Cb3RuZXQgQyZDPC9zcGFuPgogICAgICAgIDxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0gcHVyZS1tZW51LXNlbGVjdGVkIj48YSBocmVmPSIuL3BhbmVsIiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPkNvbnRyb2wgUGFuZWw8L2E+PC9saT4KCQk8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIHB1cmUtbWVudS1zZWxlY3RlZCI+PGEgaHJlZj0iLi9maWxlcyIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5GaWxlIEJyb3dzZXI8L2E+PC9saT4KCQk8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIHB1cmUtbWVudS1oYXMtY2hpbGRyZW4gcHVyZS1tZW51LWFsbG93LWhvdmVyIj4KICAgICAgICAgICAgPGEgaWQ9Im1lbnVMaW5rMSIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5Vc2VyIE1hbmFnbWVudDwvYT4KICAgICAgICAgICAgPHVsIGNsYXNzPSJwdXJlLW1lbnUtY2hpbGRyZW4iPgogICAgICAgICAgICAgICAgPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSI+PGEgaHJlZj0iLi9sb2dvdXQiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+TG9nb3V0PC9hPjwvbGk+CiAgICAgICAgICAgIDwvdWw+CiAgICAgICAgPC9saT4KICAgIDwvdWw+CjwvZGl2Pgo8YnI+Cjxicj4KPGRpdiBhbGlnbj0iY2VudGVyIj48aDI+RXJyb3I6IHtFUlJPUn08L2gyPjwvZGl2Pgo8YnI+CntTVEFUU30KPGJyPgo8ZGl2IGFsaWduPSJjZW50ZXIiPgo8YSBjbGFzcz0iYnV0dG9uLWVycm9yIHB1cmUtYnV0dG9uIiBocmVmPSIuL3BhbmVsIj57RVJST1J9IFJldHVybiB0byBDb250cm9sIFBhbmVsPC9hPgo8L2Rpdj4KPC9ib2R5Pgo8L2h0bWw+`)
	SuccessHTML    string = common.Base64Decode(`PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MS4wIj4KCiAgICA8dGl0bGU+Qm90bmV0IEMmQzogU3VjY2VzczwvdGl0bGU+CgogICAgPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiBocmVmPSJodHRwOi8veXVpLnlhaG9vYXBpcy5jb20vcHVyZS8wLjYuMC9wdXJlLW1pbi5jc3MiPgoJCgk8c3R5bGU+CmJvZHkgewogIGJhY2tncm91bmQ6ICM3NmI4NTI7IC8qIGZhbGxiYWNrIGZvciBvbGQgYnJvd3NlcnMgKi8KICBiYWNrZ3JvdW5kOiAtd2Via2l0LWxpbmVhci1ncmFkaWVudChyaWdodCwgIzc2Yjg1MiwgIzhEQzI2Rik7CiAgYmFja2dyb3VuZDogLW1vei1saW5lYXItZ3JhZGllbnQocmlnaHQsICM3NmI4NTIsICM4REMyNkYpOwogIGJhY2tncm91bmQ6IC1vLWxpbmVhci1ncmFkaWVudChyaWdodCwgIzc2Yjg1MiwgIzhEQzI2Rik7CiAgYmFja2dyb3VuZDogbGluZWFyLWdyYWRpZW50KHRvIGxlZnQsICM3NmI4NTIsICM4REMyNkYpOwogIC13ZWJraXQtZm9udC1zbW9vdGhpbmc6IGFudGlhbGlhc2VkOwogIC1tb3otb3N4LWZvbnQtc21vb3RoaW5nOiBncmF5c2NhbGU7ICAgICAgCn0KICAgICAgICAuYnV0dG9uLXN1Y2Nlc3MsCiAgICAgICAgLmJ1dHRvbi1lcnJvciwKICAgICAgICAuYnV0dG9uLXdhcm5pbmcsCiAgICAgICAgLmJ1dHRvbi1zZWNvbmRhcnkgewogICAgICAgICAgICBjb2xvcjogd2hpdGU7CiAgICAgICAgICAgIGJvcmRlci1yYWRpdXM6IDRweDsKICAgICAgICAgICAgdGV4dC1zaGFkb3c6IDAgMXB4IDFweCByZ2JhKDAsIDAsIDAsIDAuMik7CiAgICAgICAgfQogICAgICAgIC5idXR0b24tc3VjY2VzcyB7CiAgICAgICAgICAgIGJhY2tncm91bmQ6IHJnYigyOCwgMTg0LCA2NSk7IC8qIHRoaXMgaXMgYSBncmVlbiAqLwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLWVycm9yIHsKICAgICAgICAgICAgYmFja2dyb3VuZDogcmdiKDIwMiwgNjAsIDYwKTsgLyogdGhpcyBpcyBhIG1hcm9vbiAqLwogICAgICAgIH0KICAgICAgICAuYnV0dG9uLXdhcm5pbmcgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoMjIzLCAxMTcsIDIwKTsgLyogdGhpcyBpcyBhbiBvcmFuZ2UgKi8KICAgICAgICB9CiAgICAgICAgLmJ1dHRvbi1zZWNvbmRhcnkgewogICAgICAgICAgICBiYWNrZ3JvdW5kOiByZ2IoNjYsIDE4NCwgMjIxKTsgLyogdGhpcyBpcyBhIGxpZ2h0IGJsdWUgKi8KICAgICAgICB9CgkJICAqIHsgbWFyZ2luOiAwOyBwYWRkaW5nOiAwOyB9CgkgICBwIHsgcGFkZGluZzogMTBweDsgfQoJICAgI2xlZnQgeyBwb3NpdGlvbjogYWJzb2x1dGU7IGxlZnQ6IDA7ICB3aWR0aDogNTAlOyB9CgkgICAjcmlnaHQgeyBwb3NpdGlvbjogYWJzb2x1dGU7IHJpZ2h0OiAwOyAgd2lkdGg6IDUwJTsgfQo8L3N0eWxlPgoKPC9oZWFkPgoKPGJvZHk+CjxkaXYgY2xhc3M9InB1cmUtbWVudSBwdXJlLW1lbnUtaG9yaXpvbnRhbCIgYWxpZ249ImNlbnRlciI+CiAgICA8dWwgY2xhc3M9InB1cmUtbWVudS1saXN0Ij4KCTxzcGFuIGNsYXNzPSJwdXJlLW1lbnUtaGVhZGluZyIgYWxpZ249ImxlZnQiPkJvdG5ldCBDJkM8L3NwYW4+CiAgICAgICAgPGxpIGNsYXNzPSJwdXJlLW1lbnUtaXRlbSBwdXJlLW1lbnUtc2VsZWN0ZWQiPjxhIGhyZWY9Ii4vcGFuZWwiIGNsYXNzPSJwdXJlLW1lbnUtbGluayI+Q29udHJvbCBQYW5lbDwvYT48L2xpPgoJCTxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0gcHVyZS1tZW51LXNlbGVjdGVkIj48YSBocmVmPSIuL2ZpbGVzIiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPkZpbGUgQnJvd3NlcjwvYT48L2xpPgoJCTxsaSBjbGFzcz0icHVyZS1tZW51LWl0ZW0gcHVyZS1tZW51LWhhcy1jaGlsZHJlbiBwdXJlLW1lbnUtYWxsb3ctaG92ZXIiPgogICAgICAgICAgICA8YSBpZD0ibWVudUxpbmsxIiBjbGFzcz0icHVyZS1tZW51LWxpbmsiPlVzZXIgTWFuYWdtZW50PC9hPgogICAgICAgICAgICA8dWwgY2xhc3M9InB1cmUtbWVudS1jaGlsZHJlbiI+CiAgICAgICAgICAgICAgICA8bGkgY2xhc3M9InB1cmUtbWVudS1pdGVtIj48YSBocmVmPSIuL2xvZ291dCIgY2xhc3M9InB1cmUtbWVudS1saW5rIj5Mb2dvdXQ8L2E+PC9saT4KICAgICAgICAgICAgPC91bD4KICAgICAgICA8L2xpPgogICAgPC91bD4KPC9kaXY+Cjxicj4KPGJyPgo8ZGl2IGFsaWduPSJjZW50ZXIiPjxoMj5TdWNjZXNzOiB7TUVTU0FHRX08L2gyPjwvZGl2Pgo8YnI+CntTVEFUU30KPGJyPgo8ZGl2IGFsaWduPSJjZW50ZXIiPgo8YSBjbGFzcz0iYnV0dG9uLXN1Y2Nlc3MgcHVyZS1idXR0b24iIGhyZWY9Ii4vcGFuZWwiPntNRVNTQUdFfSBSZXR1cm4gdG8gQ29udHJvbCBQYW5lbDwvYT4KPC9kaXY+CjwvYm9keT4KPC9odG1sPg==`)
)

var (
	DBPointer  *sql.DB
	isPanel    bool = true
	isNew      bool = true
	isEnabled  bool = true
	maxBotList int  = 2
)

//------------------------------------------------------------
