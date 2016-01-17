package main

import (
    "fmt"
    "os"
    "unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Use: belka [file's name]\n")
		os.Exit(0)
	}
	//------
	poi := 0
	f, err := os.Open(os.Args[1])
	ce(err)
	defer f.Close()
	s ,_ := f.Stat()
	bs := make([]byte, s.Size())
	_, err = f.Read(bs)
	ce(err)
	//------
	var gottok string
	/*for gottok != "#enderror#" {
		gottok = GetToken(bs, &poi)
		fmt.Println(gottok)
	}*/
	var jcode string
	jcode += "["
	GetText(gottok, bs, &jcode, &poi, "0000")
	jcode = jcode[:(len(jcode)-1)]
	jcode += "]"
	fmt.Printf("%s", jcode)
}
// FLAGS: if_wait_else (for else), if_wait_end, if_wait_end, func_wait_end
func GetText(GotToken string, wetCode []byte, jsonCode *string, pntInCode *int, flags string) {
	for GotToken != "#enderror#" {
    	GotToken = GetToken(wetCode, pntInCode)
    	if IsKeyword(gottok) {
    		if (GotToken == "then")||((GotToken == "do") {
				PrintErr("Wrong `then` or `do`")
    		} else if gottok == "if" { // ------if
    			var BlockStmt string
				BlockStmt += "{"
    			GotToken2 := GetToken(wetCode, pntInCode)
    			BlockStmt += "\"type\":\"if\",\"cond\":["
    			for (GotToken2 != "then")&&(GotToken2 != "#enderror#")&&IsKeyword(GotToken2) {
    				if GotToken2 == "\n" {
						PrintErr("Unvalid `cond`")
    				}
    				BlockStmt += ("\""+GotToken2+"\",")
					GotToken2 = GetToken(wetCode, pntInCode)
    			}
				if BlockStmt[(len(BlockStmt)-1)] == ',' {
					BlockStmt = BlockStmt[:(len(BlockStmt)-1)]
				} else {
					PrintErr("Wrong `cond`")
				}
    			BlockStmt += "],\"body\":["
				GetText(GetToken(wetCode, pntInCode), wetCode, &BlockStmt, pntInCode, "1100")
				BlockStmt = BlockStmt[:(len(BlockStmt)-1)]
    			BlockStmt += "],"
    			//in future: elseif
    			BlockStmt = BlockStmt[:(len(BlockStmt)-1)]
    			BlockStmt += "},"
    			//in future: ...
    			*jsonCode += BlockStmt
    		} else if (GotToken == "else")||(GotToken == "elseif") { // -----else(if)
    			var BlockStmt string
				if gg(flags[0]) {
					switch(GotToken) {
					//
					//
					//
					// я остановился тут
					// 
					//
					//
						case "else":
							if (*jsonCode)[(len(*jsonCode)-1)] == ',' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								*jcode += "],\"else\":["
							} else if (*jsonCode)[(len(*jsonCode)-1)] == '[' {
								*jsonCode = (*jsonCode)[:(len(*jsonCode)-1)]
								*jsonCode += "null,\"else\":["
							}
							GetText(GetToken(wetCode, pntInCode), wetCode, jsonCode, pntInCode, "0100")
						case "elseif":
							
					}
					
					
				} else {
					PrintErr("Wrong `else` or `elseif`")
    			}
    		} else if gottok == "while" { // -----while
				*jcode += "{"
    			is_start_while = true
    			gottok2 := GetToken(bs, poi)
    			*jcode += "\"type\":\"while\",\"cond\":["
    			for (gottok2 != "do")&&(gottok2 != "#enderror#") {
    				if gottok2 == "\n" {
						fmt.Printf("Wrong `cond`")
						os.Exit(0)
    				}
    				*jcode += ("\""+gottok2+"\",")
					gottok2 = GetToken(bs, poi)
    			}
				if (*jcode)[(len(*jcode)-1)] == ',' {
					*jcode = (*jcode)[:(len(*jcode)-1)]
				}
    			*jcode += "],\"stmt\":["
				GetText(GetToken(bs, poi), bs, jcode, poi, false, true, false)
				if (*jcode)[(len(*jcode)-1)] == ',' {
					*jcode = (*jcode)[:(len(*jcode)-1)]
				}
    			*jcode += "],"
    			*jcode = (*jcode)[:(len(*jcode)-1)]
    			*jcode += "},"
    		} else if gottok == "end" { // -----end
    			if (ifs_arg||ifstatement)||(whiles_arg)||(func_arg) {
					break
				} else {
					fmt.Printf("Wrong `end`")
					os.Exit(0)
				}
    		} else if gottok == "function" { // -----functions
				*jcode += "{"
    			is_start_while = true
    			gottok2 := GetToken(bs, poi)
    			*jcode += "\"type\":\"function\",\"name\":\""
    			for (gottok2 != "(")&&(gottok2 != "#enderror#") {
    				if gottok2 == "\n" {
						fmt.Printf("Unvalid `func's name`")
						os.Exit(0)
    				}
    				*jcode += gottok2
					gottok2 = GetToken(bs, poi)
    			}
				gottok2 = GetToken(bs, poi)
    			*jcode += "\",\"args\":\""
    			for (gottok2 != ")")&&(gottok2 != "#enderror#") {
    				if gottok2 == "\n" {
						fmt.Printf("Unvalid `args`")
						os.Exit(0)
    				}
    				*jcode += gottok2
					gottok2 = GetToken(bs, poi)
    			}
    			*jcode += "\",\"stmt\":["
				GetText(GetToken(bs, poi), bs, jcode, poi, false, false, false,true) /// обратить внмание на последний TRUE
				if (*jcode)[(len(*jcode)-1)] == ',' {
					*jcode = (*jcode)[:(len(*jcode)-1)]
				}
    			*jcode += "],"
    			*jcode = (*jcode)[:(len(*jcode)-1)]
    			*jcode += "},"
    		}
    	} else if (gottok != "\n")&&(gottok != "#enderror#")  {
    		//if !is_return_statement {
    			*jcode += ("[\"" + gottok + "\",")
    		//	is_return_statement = false
    		//} else {
    		//	*jcode += (gottok)
    		//}
    		gottok2 := GetToken(bs, poi)
			for (gottok2 != "#enderror#")&&(gottok2 != "\n") {
				if IsKeyword(gottok2) {
					fmt.Printf("Unvalid `stmt`")
					os.Exit(0)
				}
    			*jcode += ("\"" + gottok2 + "\",")
				gottok2 = GetToken(bs, poi)
			}
			if (*jcode)[(len(*jcode)-1)] == ',' {
				*jcode = (*jcode)[:(len(*jcode)-1)]
			}
    		*jcode += "],"
		}
    }
}

func IsKeyword(word string) bool {
	if (word == "while")||(word == "if")||(word == "else")||(word == "then")||(word == "end")||(word == "function") {
		return true
	} else {
		return false
	}
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
			for t<l&&unicode.IsDigit(rune(bs[t])) {
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
					fmt.Printf("Expected `-`")
					os.Exit(0)
							}
						} else {
							fmt.Printf("Expected `]`")
							os.Exit(0)
						}
					} else {
						fmt.Printf("Expected `[`")
						os.Exit(0)
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
			t++
			*p = t
			return string(bs[t-1])
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
						fmt.Printf("Wrong `]`, but it'll expect")
						os.Exit(0)
					}
				} else {
					for t<l&&(rune(bs[t]) != ']') {
						tok += string(bs[t])
						t++
					}
					tok += string(bs[t])
				}
			} else {
				fmt.Printf("Wrong `[`")
				os.Exit(0)
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

func gg(byte a) {
	return (string(a) == "1") 
}

func ce(e error) {
	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(0)
	}
}
