/*
#
# Name: Belka Translator
#
# Life cycle:
# 	1. Get data (sequence of chars) from file / Получение данных (последовательность символов) из файла
# 	2. Generation tokens from the received sequence, recursive analysing it / Генерация токены из полученной последовательности символов, рекурсивный разбор токен
#	3. Gradual generation of a Simple Lua's JSON AST / Постепенная генерация JSON AST простого Lua
#
# Author: Emil Khayrullin/Эмиль Хайруллин
#		
#######################################################################
#                                                                     #
#  #### ###   ##     ####### ##   ## ####### ##   ## ######  #######  #
#   ##  ####  ##     ##      ##   ##   ###   ##   ## ##   ## ##       #
#   ##  ## ## ##     ######  ##   ##   ###   ##   ## ######  ######   #
#   ##  ##  ####     ##      ##   ##   ###   ##   ## ##  ##  ##       #
#  ###  ##   ###     ##        ###     ###     ###   ##   ## #######  #
#                                                                     #
#######################################################################
#
# 1. for/each
# 2. доработать вложенность
# -1. optimization AST
*/
package main

import (
    "fmt"
    "os"
    "unicode"
    _"encoding/json"
    "strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Use: belka [file's name]\n")
		os.Exit(0)
	}
	//------
	f, err := os.Open(os.Args[1])
	ce(err)
	defer f.Close()
	s ,_ := f.Stat()
	bs := make([]byte, s.Size())
	_, err = f.Read(bs)
	ce(err)
	//------
	var jcode string
	poi := 0
	jcode += "["
	GetText("", bs, &jcode, &poi, "00000")
	jcode = jcode[:(len(jcode)-1)]
	jcode += "]"
	fmt.Printf("%s", jcode)
}
// FLAGS: if_wait_else (for else), if_wait_end, if_wait_end, func_wait_end, repeat_wait_until
func GetText(GotToken string, wetCode []byte, jsonCode *string, pntInCode *int, flags string) {
	for GotToken != "#enderror#" {
    	GotToken = GetToken(wetCode, pntInCode)
    	if IsKeyword(GotToken) {
    		if (GotToken == "then")||(GotToken == "do") {
				PrintErr("Wrong `then` or `do`")
    		} else if GotToken == "if" { // ------if
    			*jsonCode += "{\"type\":\"if\",\"cond\":["
    			GotToken2 := GetToken(wetCode, pntInCode)
    			for (GotToken2 != "then")&&(GotToken2 != "#enderror#")&&!IsKeyword(GotToken2) {
    				if GotToken2 == "\n" {
						PrintErr("Unvalid `cond` in `if` block")
    				}
    				*jsonCode += ("\""+GotToken2+"\",")
					GotToken2 = GetToken(wetCode, pntInCode)
    			}
				if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
					*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
				} else {
					PrintErr("Empty `cond` in `if` block")
				}
    			*jsonCode += "],\"body\":["
				GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, gc(gc(flags,0,true),1,true))
    		} else if (GotToken == "else")||(GotToken == "elseif") { // -----else(if)
				if gg(flags,0) {
					switch(GotToken) {
						case "else":
							if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								*jsonCode += "],\"else\":["
							} else if (*jsonCode)[(len(*jsonCode)-1)] == '[' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								*jsonCode += "null,\"else\":["
							}
							GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, gc(gc(flags,1,true),0,false))
							*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
							*jsonCode += "],"
						case "elseif":
							if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								*jsonCode += "],\"elseif\":"
							} else if (*jsonCode)[(len(*jsonCode)-1)] == '[' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								*jsonCode += "null,\"elseif\":"
							}
    						*jsonCode += "{\"cond\":["
							GotToken2 := GetToken(wetCode, pntInCode)
							for (GotToken2 != "then")&&(GotToken2 != "#enderror#")&&!IsKeyword(GotToken2) {
								if GotToken2 == "\n" {
									PrintErr("Unvalid `cond` in `elseif` block")
								}
								*jsonCode += ("\""+GotToken2+"\",")
								GotToken2 = GetToken(wetCode, pntInCode)
							}
							if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
							} else {
								PrintErr("Empty `cond` in `elseif` block")
							}
    						*jsonCode += "],\"stmt\":["
							GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, flags)
    						if (*jsonCode)[(len(*jsonCode)-2)] == ']' {
								if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
									*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								}
								*jsonCode += "],"
							}
							*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
							*jsonCode+= "},"
							//in future: ...
					}
				} else {
					PrintErr("Wrong `else` or `elseif`")
    			}
    		} else if GotToken == "while" { // -----while
    			*jsonCode += "{\"type\":\"while\",\"cond\":["
    			GotToken2 := GetToken(wetCode, pntInCode)
    			for (GotToken2 != "do")&&(GotToken2 != "#enderror#")&&!IsKeyword(GotToken2) {
    				if GotToken2 == "\n" {
						PrintErr("Wrong `cond` in `while` block")
    				}
    				*jsonCode += ("\""+GotToken2+"\",")
					GotToken2 = GetToken(wetCode, pntInCode)
    			}
				if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
					*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
				}
    			*jsonCode += "],\"stmt\":["
				GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, gc(flags,2,true))
    			if (*jsonCode)[(len(*jsonCode)-2)] == ']' {
					if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
						*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
					}
					*jsonCode += "],"
				}
    			*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
    			*jsonCode += "},"
				//in future: ...
    		} else if GotToken == "repeat" { // -----repeat
    			*jsonCode += "{\"type\":\"repeat\",\"stmt\":["
				GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, gc(flags,4,true))
				//in future: ...
    		} else if GotToken == "until" {
    			if gg(flags,4) {
					if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
						*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
					}
    				*jsonCode += "],\"cond\":["
    				GotToken2 := GetToken(wetCode, pntInCode)
					for (GotToken2 != "#enderror#")&&!IsKeyword(GotToken2)&&(GotToken2 != "\n" ) {
						*jsonCode += ("\""+GotToken2+"\",")
						GotToken2 = GetToken(wetCode, pntInCode)
					}
					if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
						*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
					}
					*jsonCode += "],"
					*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
    				*jsonCode += "},"
    			} else {
					PrintErr("Wrong `until`")
    			}
    		} else if GotToken == "end" { // -----end
    			if gg(flags,1)||gg(flags,2)||gg(flags,3) {
    				if gg(flags,1)||(*jsonCode)[(len(*jsonCode)-2)]=='}' {
    					if (*jsonCode)[(len(*jsonCode)-2)] == ']' {
							if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
							}
							*jsonCode += "],"
						}
						*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
						*jsonCode += "},"
    				}
					return
				} else {
					PrintErr("Wrong `end`")
				}
    		} else if GotToken == "function" { // -----function
    			*jsonCode += "{\"type\":\"func\",\"name\":\""
    			GotToken2 := GetToken(wetCode, pntInCode)
    			for (GotToken2 != "(")&&(GotToken2 != "#enderror#") {
    				if GotToken2 == "\n" {
						PrintErr("Unvalid name in `function` block")
    				}
    				*jsonCode += GotToken2
					GotToken2 = GetToken(wetCode, pntInCode)
    			}
				GotToken2 = GetToken(wetCode, pntInCode)
    			*jsonCode += "\",\"args\":\""
    			for (GotToken2 != ")")&&(GotToken2 != "#enderror#") {
    				if GotToken2 == "\n" {
						PrintErr("Unvalid `args` in `function` block")
    				}
    				*jsonCode += GotToken2
					GotToken2 = GetToken(wetCode, pntInCode)
    			}
    			*jsonCode += "\",\"stmt\":["
				GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, gc(flags,3,true))
				if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
					*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
				}
    			*jsonCode += "],"
    			*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
    			*jsonCode += "},"
				//in future: ...
    		}
    	} else if (GotToken != "\n")&&(GotToken != "#enderror#")  {
    		*jsonCode += ("[\"" + GotToken + "\",")
    		GotToken2 := GetToken(wetCode, pntInCode)
			for (GotToken2 != "#enderror#")&&(GotToken2 != "\n") {
				if IsKeyword(GotToken2) {
					PrintErr("Unvalid `stmt`")
				}
    			*jsonCode += ("\"" + GotToken2 + "\",")
				GotToken2 = GetToken(wetCode, pntInCode)
			}
			if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
				*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
			}
    		*jsonCode += "],"
		}
    }
}

func IsKeyword(word string) bool {
	return (word == "while")||(word == "for")||(word == "repeat")||(word == "until")||(word == "in")||(word == "if")||(word == "else")||(word == "elseif")||(word == "then")||(word == "end")||(word == "function")
}

func PrintErr(a string) {
	fmt.Printf("{\"error\":\"%s\"}", a);
	os.Exit(0)
}

func GetToken(bs []byte, p *int) string {
	l := len(bs)
	tok := ""
	t := *p
	for t < l {
		if unicode.IsLetter(rune(bs[t])) {
			for t<l&&unicode.IsLetter(rune(bs[t])) { // -----words
				tok += string(bs[t])
				t++
			}
			*p = t
			return tok
		} else if unicode.IsDigit(rune(bs[t])) { // -----numbers
			for t<l&&(unicode.IsDigit(rune(bs[t]))||(bs[t]=='.')) {
				if len(tok)>0&&(tok[len(tok)-1]=='.')&&(bs[t]=='.') {
					PrintErr("Wrong number")
				}
				tok += string(bs[t])
				t++
			}
			*p = t
			return tok
		} else if unicode.IsSpace(rune(bs[t])) { // -----spaces
			if bs[t] == '\n' {
				t++
				*p = t
				return string(bs[t-1])
			} else {
				for t<l&&unicode.IsSpace(rune(bs[t])) {
					t++
				}
			}
		} else if (rune(bs[t]) == '(')||(rune(bs[t]) == ')') { // ----- ()
			t++
			*p = t
			return string(bs[t-1])
		} else if (rune(bs[t]) == '"')||(rune(bs[t]) == '\'') {// ----- "" or ''
			if (string(bs[t]) == "\"") {
				tok += ("\\"+string(bs[t]))
			} else {
				tok += (string(bs[t]))
			}
			t++
			for t<l {
				tok += string(bs[t])
				t++
				if (rune(bs[t]) == '"')||(rune(bs[t]) == '\'') {
					break
				}
			}
			if (string(bs[t]) == "\"") {
				tok += ("\\"+string(bs[t]))
			} else {
				tok += (string(bs[t]))
			}
			t++
			*p = t
			return tok
		} else if (rune(bs[t]) == '-') { // ----- oneline and multiline comment
			if t<l-1&&(rune(bs[t+1]) == '-') {
				t++
				if t<l-1&&(rune(bs[t+1]) == '[') {
					t++
					if t<l-1&&(rune(bs[t+1]) == '[') {
						for t<l&&(rune(bs[t]) != ']') {
							t++
						}
						t++
						if t<l&&(rune(bs[t]) == ']') {
							t++
							if t<l&&t<l-1&&(rune(bs[t]) == '-')&&(rune(bs[t+1]) == '-') {
								t++
								bs[t] = byte('\n')
							} else {
								PrintErr("Expected `-`")
							}
						} else {
							PrintErr("Expected `]`")
						}
					} else {
						PrintErr("Expected `[`")
					}
				} else {
					for t<l&&(rune(bs[t]) != '\n') {
						t++
					}
				}
			} else {
				t++
				*p = t
				return string(bs[t-1])
			}
		} else if (rune(bs[t]) == '+')||(rune(bs[t]) == '*')||(rune(bs[t]) == '/') { // -----operator*+-/
			t++
			*p = t
			return string(bs[t-1])
		} else if (rune(bs[t]) == ',') { // -----operator,
			t++
			*p = t
			return string(bs[t-1])
		} else if (rune(bs[t]) == '.') { // -----operator.
			if (t+1<l)&&unicode.IsDigit(rune(bs[t+1])) {
				tok += string(bs[t])
				t++
				for t<l&&(unicode.IsDigit(rune(bs[t]))||(bs[t]=='.')) {
					if (bs[t]=='.') {
						PrintErr("Wrong number")
					}
					tok += string(bs[t])
					t++
				}
				*p = t
				return tok
			} else {
				t++
				*p = t
				return string(bs[t-1])
			}
		} else if rune(bs[t]) == '[' { // -----multiline string or map[index]
			if t<l-1 {
				if (rune(bs[t+1]) == '[') {
					t = t + 2
					tok += "\\\""
					for t<l&&(rune(bs[t]) != ']') {
						if string(bs[t]) == "\n" {
							tok += "\\n"
						} else if unicode.IsSpace(rune(bs[t]))  {
							//nothing
						} else {
							tok += string(bs[t])
						}
						t++
					}
					if t<l-1&&(rune(bs[t+1]) == ']') {
						tok += "\\\""
						t++
					} else {
						PrintErr("Wrong `]`, but it'll expect")
					}
				} else {
					for t<l&&(rune(bs[t]) != ']') {
						tok += string(bs[t])
						t++
					}
					tok += string(bs[t])
				}
			} else {
				PrintErr("Wrong `[`")
			}
			t++
			*p = t
			return tok
		} else if (rune(bs[t]) == '>')||(rune(bs[t]) == '=')||(rune(bs[t]) == '<') { // -----ops > >= < <= ==
			tok += string(bs[t])
			if t<l-1&&(rune(bs[t+1]) == '=') {
				tok += string(bs[t+1])
				t+=2
			} else {
				t++
			}
			*p = t
			return tok
		} else if (rune(bs[t]) == '~') { // -----op ~=
			tok += string(bs[t])
			if t<l-1&&(rune(bs[t+1]) == '=') {
				tok += string(bs[t+1])
				t+=2
			} else {
				break
			}
			*p = t
			return tok
		} else {
			return "#enderror#"
		}
	}
	return "#enderror#"
}

func gg(a string, b int) bool {
	/*fmt.Println(a)
	fmt.Println(b)*/
	t := int(rune(a[b]))
	//ce(e)
	return (len(a)>b)&&(t >= 1) 
}

func gc(a string, b int, d bool) string {
	if len(a)>b {
		c := make([]byte, len(a))
		c = []byte(a)
		t, e := strconv.Atoi(string(c[b]))
		ce3(e)
		if d {
			c[b] = byte(t+1)
		} else {
			if t > 0 {
				c[b] = byte(t-1)
			} else {
				PrintErr("error on `gc`")
			}
		}
		return string(c)
	} else {
		return a
	}
}
func ce(e error) {
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(0)
	}
}
func ce3(e error) {
	if e != nil {
		fmt.Printf("----%s----", e.Error())
		os.Exit(0)
	}
}
