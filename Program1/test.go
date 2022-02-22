package main

import (
	"fmt"
)

func formattedOutput() string {
	var stockcode = 123
	var enddate = "2022-02-22"
	var url = "Code=%d,endDate=%s,%s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate, "来自同一个package下的另一个go文件的输出")
	return targetUrl
}
