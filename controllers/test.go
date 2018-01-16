package controllers

import (
	"bytes"
	_ "ehome/trade"
	_ "fmt"
	"github.com/astaxie/beego"
	"os"
)

type TestController struct {
	beego.Controller
}

// URLMapping ...
func (c *TestController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

// @param   num   num false
// @Failure 403 body is empty
// @router / [get]
func (c *TestController) Get() {
	strs, err := tail("/root/nohup.out", 30)
	beego.Error(len(strs))
	if err != nil {
		beego.Error("nohup out")
		c.Ctx.WriteString("error")
		return
	}

	for i := range strs {
		//beego.Info(strs[i])
		c.Ctx.WriteString(strs[i])
		c.Ctx.WriteString("\r\n")
		//	c.Ctx.WriteString("<br>")
	}

}

// @param   num   num false
// @Failure 403 body is empty
// @router / [post]
func (c *TestController) Post() {
	beego.Error(c.Ctx.Input.Context.Input.Context.Input)

	beego.Error("test post Data")
	beego.Error(c.Input())

	beego.Error("test post Body")
	beego.Error(string(c.Ctx.Input.RequestBody))
}

const (
	defaultBufSize = 4096
)

func tail(filename string, n int) (lines []string, err error) {
	f, e := os.Stat(filename)
	if e == nil {
		size := f.Size()
		var fi *os.File
		fi, err = os.Open(filename)
		if err == nil {
			b := make([]byte, defaultBufSize)
			sz := int64(defaultBufSize)
			nn := n
			bTail := bytes.NewBuffer([]byte{})
			istart := size
			flag := true
			for flag {
				if istart < defaultBufSize {
					sz = istart
					istart = 0
				} else {
					istart -= sz
				}
				_, err = fi.Seek(istart, os.SEEK_SET)
				if err == nil {
					mm, e := fi.Read(b)
					if e == nil && mm > 0 {
						j := mm
						for i := mm - 1; i >= 0; i-- {
							if b[i] == '\n' {
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[i+1 : j])
								j = i
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()
								}

								if (nn == n && bLine.Len() > 0) || nn < n { //skip last "\n"
									lines = append(lines, bLine.String())
									nn--
								}
								if nn == 0 {
									flag = false
									break
								}
							}
						}
						if flag && j > 0 {
							if istart == 0 {
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[:j])
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()
								}
								lines = append(lines, bLine.String())
								flag = false
							} else {
								bb := make([]byte, bTail.Len())
								copy(bb, bTail.Bytes())
								bTail.Reset()
								bTail.Write(b[:j])
								bTail.Write(bb)
							}
						}
					}
				}
			}
		}
		defer fi.Close()
	}
	return
}
