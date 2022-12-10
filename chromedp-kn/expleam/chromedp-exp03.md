```golang
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "strconv"
    "time"

    "github.com/chromedp/cdproto/cdp"
    "github.com/chromedp/chromedp"
    _ "github.com/go-sql-driver/mysql"
)

type Registrar struct {
    r_id   int64
    name   string
    url    string
    contry string
    number string
    email  string
}

func addToDb(l Registrar, db *sql.DB) {
    stmt, err := db.Prepare("INSERT registrar SET r_id=?,name=?,url=?,contry=?,number=?,email=?")
    checkErr(err)
    res, err := stmt.Exec(l.r_id, l.name, l.url, l.contry, l.number, l.email)
    checkErr(err)
    insert_id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(insert_id)
}

func initDb() *sql.DB {
    connectionString := fmt.Sprintf("root:%s@tcp(ip:3306)/test?charset=utf8&parseTime=True&loc=Local",
        "password")
    db, err := sql.Open("mysql", connectionString)
    checkErr(err)
    err = db.Ping()
    checkErr(err)
    return db
}

func main() {
    //连接数据库
    db := initDb()
    // 禁用chrome headless
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", false),
    )
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()
    ctx, cancel := chromedp.NewContext(
        allocCtx,
        chromedp.WithLogf(log.Printf),
    )
    defer cancel()

    var t1, t2, t3, t4, t6, t7 []*cdp.Node
    //开启一个chrome
    err := chromedp.Run(ctx,
        chromedp.Navigate("https://www.ggcx.com/main/globalRegistrar"), //打开首页
        chromedp.WaitVisible(`div[class="tabb"]`),                      //等着页面异步加载完成
        //获取需要的数据
        chromedp.Nodes(`.//td[@class="t1"]`, &t1),
        chromedp.Nodes(`.//td[@class="t2"]`, &t2),
        chromedp.Nodes(`.//td[@class="t3"]`, &t3),
        chromedp.Nodes(`.//td[@class="t4"]`, &t4),
        chromedp.Nodes(`.//td[@class="t6"]`, &t6),
        chromedp.Nodes(`.//td[@class="t7"]`, &t7),
        chromedp.Sleep(3*time.Second),
        chromedp.Click(`div[class="right"]`, chromedp.NodeVisible), //点击下一页
        chromedp.Sleep(5*time.Second),
    )
    checkErr(err)
    //组装数据
    for k, node := range t1 {
        r_id, _ := strconv.ParseInt(node.Children[0].NodeValue, 10, 64)
        l := Registrar{
            r_id,
            t2[k].Children[0].NodeValue,
            t3[k].Children[0].NodeValue,
            t4[k].Children[0].NodeValue,
            t6[k].Children[0].NodeValue,
            t7[k].Children[0].NodeValue,
        }
        //数据加入数据库
        addToDb(l, db)
    }

    //获取分页的数据
    for i := 0; i < 140; i++ {
        err := chromedp.Run(ctx,
            chromedp.WaitVisible(`div[class="tabb"]`),
            chromedp.Nodes(`.//td[@class="t1"]`, &t1),
            chromedp.Nodes(`.//td[@class="t2"]`, &t2),
            chromedp.Nodes(`.//td[@class="t3"]`, &t3),
            chromedp.Nodes(`.//td[@class="t4"]`, &t4),
            chromedp.Nodes(`.//td[@class="t6"]`, &t6),
            chromedp.Nodes(`.//td[@class="t7"]`, &t7),
            chromedp.Sleep(3*time.Second),
            chromedp.Click(`div[class="right"]`, chromedp.NodeVisible), //点击下一页
            chromedp.Sleep(5*time.Second),
        )
        checkErr(err)
        for k, node := range t1 {
            r_id, _ := strconv.ParseInt(node.Children[0].NodeValue, 10, 64)
            l := Registrar{
                r_id,
                t2[k].Children[0].NodeValue,
                t3[k].Children[0].NodeValue,
                t4[k].Children[0].NodeValue,
                t6[k].Children[0].NodeValue,
                t7[k].Children[0].NodeValue,
            }
            addToDb(l, db)
        }
    }
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}