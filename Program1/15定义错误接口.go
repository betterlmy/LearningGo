package main

import "fmt"

//定义一个除法错误的结构
type DivideError struct {
	divided float64
	divider float64
}

func (de DivideError) Error() string {
	strFmt := `
	无法处理,
	被除数:%f
	除数是0
`
	return fmt.Sprintf(strFmt, de.divided)
}

func Divide(divided float64, divider float64) (result float64, errorMsg string) {
	if divider == 0 {
		dData := DivideError{
			divided: divided,
			divider: divider,
		}
		return 0, dData.Error()
	} else {
		return divided / divider, ""
	}
}
func main() {
	if result, errorMSG := Divide(500.2, 0); errorMSG == "" {
		fmt.Println(result)
	} else {
		fmt.Println(errorMSG)
	}
}
