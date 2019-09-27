package main
import (
	"github.com/xwb1989/sqlparser"
	"fmt"
	"strings"
	"io"
)
func main(){

//	sql := "SELECT * FROM table WHERE a = 'abc'"
    sql := "SELECT p.bid,p.meid,p.qrid,p.trade_type,(CASE WHEN q.checkstand_id IS NULL THEN 0 ELSE q.checkstand_id END ) AS checkstand_id,(CASE WHEN c.name IS NULL THEN '' ELSE c.name END ) AS checkstand_name,SUM(p.amount) AS amount,SUM(p.real_amount) AS receivable_amount,COUNT(*) AS receivable_count,SUM(p.merchant_coupon_fee) AS merchant_coupon_amount,SUM( CASE WHEN p.merchant_coupon_fee > 0 THEN 1 ELSE 0 END ) AS merchant_coupon_count,SUM( CASE WHEN p.type = \"refund\" THEN p.amount ELSE 0 END ) AS refund_amount,SUM( CASE WHEN p.type = \"refund\" THEN 1 ELSE 0 END ) AS refund_count,SUM( CASE WHEN p.coupon_off > 0 THEN p.coupon_off ELSE 0 END ) AS platform_discount_amount,SUM( CASE WHEN p.coupon_off > 0 THEN 1 ELSE 0 END ) AS platform_discount_count,SUM(p.commission) AS commission FROM payments p  FORCE INDEX(meid_status_updated_at) LEFT JOIN qrcodes q ON (p.qrid = q.qrid) LEFT JOIN checkstands c ON(q.checkstand_id = c.id) WHERE (p.meid IN ('460')) AND (p.status = 'success') AND (p.updated_at >= '2019-06-16 00:00:00') AND (p.updated_at <= '2019-06-16 23:59:59') AND (p.finished_at >= '2019-06-16 00:00:00') AND (p.finished_at <= '2019-06-16 23:59:59') AND (p.trade_type NOT LIKE '%deposit%') AND (p.vendor_terminal_id LIKE 'device:%') GROUP BY p.meid, p.qrid, p.trade_type, p.bid ";
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		// Do something with the err
		println(err)
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		_ = stmt
	case *sqlparser.Insert:
	}
	r := strings.NewReader("INSERT INTO table1 VALUES (1, 'a'); INSERT INTO table2 VALUES (3, 4);")

	tokens := sqlparser.NewTokenizer(r)
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		fmt.Printf("%v",stmt)
		if err == io.EOF {
			break
		}
		// Do something with stmt or err.
	}
}
