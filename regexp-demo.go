package main
import(
    "fmt"
    iconv "github.com/djimenez/iconv-go"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
)
// embed regexp.Regexp in a new type so we can extend it
type myRegexp struct{
    *regexp.Regexp
}
// add a new method to our new regular expression type
func(r *myRegexp)FindStringSubmatchMap(s string)[](map[string]string){
    captures:=make([](map[string]string),0)
    matches:=r.FindAllStringSubmatch(s,-1)
    if matches==nil{
        return captures
    }
    names:=r.SubexpNames()
    for _,match:=range matches{
        cmap:=make(map[string]string)
        for pos,val:=range match{
            name:=names[pos]
            if name==""{
                continue
            }
            /*
                fmt.Println("+++++++++")
                fmt.Println(name)
                fmt.Println(val)
            */
            cmap[name]=val
        }
        captures=append(captures,cmap)
    }
    return captures
}
// 抓取限号信息的正则表达式
var myExp=myRegexp{regexp.MustCompile(`自(?P<byear>[\d]{4})年(?P<bmonth>[\d]{1,2})月(?P<bday>[\d]{1,2})日至(?P<eyear>[\d]{4})年(?P<emonth>[\d]{1,2})月(?P<eday>[\d]{1,2})日，星期一至星期五限行机动车车牌尾号分别为：(?P<n11>[\d])和(?P<n12>[\d])、(?P<n21>[\d])和(?P<n22>[\d])、(?P<n31>[\d])和(?P<n32>[\d])、(?P<n41>[\d])和(?P<n42>[\d])、(?P<n51>[\d])和(?P<n52>[\d])`)}
func ErrorAndExit(err error){
    fmt.Fprintln(os.Stderr,err)
    os.Exit(1)
}
func main(){
    response,err:=http.Get("http://www.bjjtgl.gov.cn/zhuanti/10weihao/index.html")
    defer response.Body.Close()
    if err!=nil{
        ErrorAndExit(err)
    }
    input,err:=ioutil.ReadAll(response.Body)
    if err!=nil{
        ErrorAndExit(err)
    }
    body :=make([]byte,len(input))
    iconv.Convert(input,body,"gb2312","utf-8")
    mmap:=myExp.FindStringSubmatchMap(string(body))
    fmt.Println(mmap)
}
