package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func getValue(x string) string {
	x = strings.TrimSpace(x)
	x = x + "     "
	if x[0:2] == "10" {
		if x[0:3] != "10 " {
			return "-1"
		}
		return "10"
	}
	if x[0:1] == "I" {
		if x[0:2] == "II" {
			if x[0:3] == "III" {
				return "III"
			}
			return "II"
		}
		if x[0:2] == "IV" {
			return "IV"
		}
		if x[0:2] == "IX" {
			return "IX"
		}
		return "I"
	}
	if x[0:2] == "VI" {
		if x[0:3] == "VII" {
			if x[0:4] == "VIII" {
				return "VIII"
			}
			return "VII"
		}
		return "VI"
	}
	if x[0:1] == "X" {
		if x[0:2] == "XI" || x[0:2] == "XV" || x[0:2] == "XX" {
			return "-1"
		}
		return "X"
	}
	if ArabToInt(x[0:2]) < 11 && ArabToInt(x[0:1]) > 0 {
		return x[0:1]
	}
	return "-1"
}

func getOperation(x string) string {
	x = strings.TrimSpace(x)
	x = x + "     "
	if x[0:1] == "+" || x[0:1] == "-" || x[0:1] == "*" || x[0:1] == "/" {
		return x[0:1]
	}
	return "-1"
}

func getValueType(x string) string {
	if x == "1" || x == "2" || x == "3" || x == "4" || x == "5" || x == "6" || x == "7" || x == "8" || x == "9" || x == "10" {
		return "arabian"
	}
	if x == "I" || x == "II" || x == "III" || x == "IV" || x == "V" || x == "VI" || x == "VII" || x == "VIII" || x == "IX" || x == "X" {
		return "roman"
	}
	return "-1"
}

func RomeToArab(x string) int {
	if x == "I" {
		return 1
	}
	if x == "II" {
		return 2
	}
	if x == "III" {
		return 3
	}
	if x == "IV" {
		return 4
	}
	if x == "V" {
		return 5
	}
	if x == "VI" {
		return 6
	}
	if x == "VII" {
		return 7
	}
	if x == "VIII" {
		return 8
	}
	if x == "IX" {
		return 9
	}
	if x == "X" {
		return 10
	}
	return -1
}

func ArabToInt(x string) int {
	ans, _ := strconv.Atoi(x)
	return ans
}

func ArabToRome10(x int) string {
	result := ""
	if x == 10 {
		return "X"
	}
	if x/5 == 1 {
		if x%5 == 4 {
			return "IX"
		}
		result = "V" + result
	} else {
		if x%5 == 4 {
			return "IV"
		}
	}
	for i := 0; i < x%5; i++ {
		result = result + "I"
	}
	return result
}

func ArabToRome100(x int) string {
	result := ""
	if x == 10 {
		return "C"
	}
	if x/5 == 1 {
		if x%5 == 4 {
			return "XC"
		}
		result = "L" + result
	} else {
		if x%5 == 4 {
			return "XL"
		}
	}
	for i := 0; i < x%5; i++ {
		result = result + "X"
	}
	return result
}

func ArabToRome(x int) string {
	return ArabToRome100(x/10) + ArabToRome10(x%10)
}

func Calculate(a int, b int, o string) int {
	if o == "+" {
		return a + b
	}
	if o == "-" {
		return a - b
	}
	if o == "*" {
		return a * b
	}
	if o == "/" {
		return a / b
	}
	return -1
}

func Check2Short(x string) bool {
	if len(x) == 0 {
		println("Error. Expression is too short to being mathematical")
		return true
	}
	return false
}

func Check2Long(x string) bool {
	x = strings.TrimSpace(x)
	if len(x) != 0 {
		println("Error.Expression is too long, Only use 2 values and single operation")
		return true
	}
	return false
}

func CheckValues(a, b, o string) bool {
	if getValueType(a) == "-1" || getValueType(b) == "-1" {
		println("Error: value is not supported")
		return false
	}
	if getValueType(a) != getValueType(b) {
		println("Error. Data Type Mismatch")
		return false
	}
	if o == "-1" {
		println("Error. Operation is not recognised")
		return false
	}
	if getValueType(a) == "roman" && RomeToArab(a) <= RomeToArab(b) {
		if getOperation(o) == "-" {
			println("Error. Result is not a Roman number")
			return false
		}
		if getOperation(o) == "/" && RomeToArab(a) != RomeToArab(b) {
			println("Error. Result is not a Roman number")
			return false
		}
	}
	return true
}

func main() {
	var Data = getText()
	var Value1 = getValue(Data)
	Data = strings.TrimSpace(Data)
	Data = Data[len(Value1):]
	if Check2Short(Data) == false {
		var Operation = getOperation(Data)
		Data = strings.TrimSpace(Data)
		Data = Data[1:]
		if Check2Short(Data) == false {
			var Value2 = getValue(Data)
			if CheckValues(Value1, Value2, Operation) == true {
				Data = strings.TrimSpace(Data)
				Data = Data[len(Value2):]
				if Check2Long(Data) == false {
					if getValueType(Value1) == "arabian" {
						println(Calculate(ArabToInt(Value1), ArabToInt(Value2), Operation))
					}
					if getValueType(Value1) == "roman" {
						println(ArabToRome(Calculate(RomeToArab(Value1), RomeToArab(Value2), Operation)))
					}
				}
			}
		}
	}
}
