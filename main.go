package main

import (
	"fmt"
	"os"
)

func findslashn(input string) []string {
	result := []string{}
	str := ""
	for i := 0; i < len(input)-1; i++ {
		if input[i] == '\\' && input[i+1] == 'n' {
			if i == 0 {
				result = append(result, "")
				i++
			}else {
				if i < len(input)-3  && input[i+2] == '\\' && input[i+3] == 'n'{
					if str != "" {
						result = append(result, str)
						str = ""
						result = append(result, "")
						i += 3
					}
				}else {
					if str != "" {
						result = append(result, str)
						str = ""
						i+=1
					}
				}
			}
		} else {
			str += string(input[i])
		}
	}
	if input[len(input)-1] != 'n' {
		str += string(input[len(input)-1])
	}else {
		if input[len(input)-2] != '\\' {
			str += string(input[len(input)-1])
		}else {
			result = append(result, "")
		}
	}
	if str != "" {
		result = append(result, str)
		str = ""
	}

	return result
}

func ascii(fil []byte, input string) string {
	cheker := 0
	rr := 0
	str := ""
	for j := 0; j < 8; j++ {
		for k := 0; k < len(input); k++ {
			rr = int(rune(input[k])) - 32
			cheker = j
			for i := 1; i < len(fil)-1; i++ {
				if fil[i] == '\n' && fil[i+1] == '\n' {
					rr--
					i++
				}
				if rr == 0 {
					if cheker == -1 {
						if fil[i] != '\n' {
							str += string(fil[i])
							// fmt.Print(str)
						} else {
							break
						}
					} else {
						if fil[i] == '\n' {
							cheker--
						}
					}
				}
			}
		}
		str += "\n"
	}
	return str
}

func main() {
	fil, _ := os.ReadFile("standard.txt")

	input := os.Args[1]
	aa := findslashn(input)
//	fmt.Println(aa)
	 result := ""
	for i := 0; i < len(aa); i++ {
		if aa[i] == "" {
			result+="\n"
			
		}else {
			result+=(ascii(fil, aa[i]))
			
		}
		//fmt.Println(ascii(fil, aa[i]))
		// result = ascii(fil, input)
		//  err := os.WriteFile("expl.txt", []byte(result), 0777)
		// if err != nil {
		// 	panic(err)
		// }
		// 7 7 7
		// 7 = 4 + 2 + 1
		// f = r + w + E
	}
	fmt.Print(result)
}
