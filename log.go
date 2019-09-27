// https://www.flysnow.org/2017/05/06/go-in-action-go-log.html

package main
import (
	"log"
)
func init(){
	// log.SetFlags(log.Ldate|log.Lshortfile)
	log.SetPrefix("【UserCenter】")

	log.SetFlags(log.Ldate|log.Ltime |log.LUTC)

}

func main() {
	log.Println("飞雪无情的博客:","http://www.flysnow.org")
	log.Printf("飞雪无情的微信公众号：%s\n","flysnow_org")
}
