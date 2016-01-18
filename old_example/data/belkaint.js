"use strict";
var JBelka = function() {
	console.log(arguments[0]);
	//---------vars, funcs
	var GlobalVar = {
		"true": true,
		"false": false,
		"nil": null
	};
	var GlobalFunc = {};
	//---------errors
	var errPrint = function() {
		if(arguments.length > 0) {
			var e = [
				"Не верный код программы.",
				"Не верное выражение",
				"getExprToken | Не верное кол-во аргументов",
				"getTokenType | Не верное кол-во аргументов",
				
			]
			for(var i = 0; i < arguments.length; i++) {
				var ind = arguments[i];
				if(typeof(ind)==="number"&&ind<e.length) {
					consoleWrite("Ошибка ("+ind+"): "+e[ind]);
				}
			}
		} else {
			consoleWrite("Ошибка (-1): undefined");
		}
	};
	//---------tokenIsKeyword
	var tokenIsKeyword = function(a) {
		return (["if", "then", "while", "do", "repeat", "until", "else"].indexOf(a) > -1);
	};
	//---------getTokenType
	var getTokenType = function() {
		if((arguments.length === 1)) {
			return (parseInt(arguments[0], "10").toString() === arguments[0])?"num":"word";
		} else {
			errPrint(2);
		}
	};
	//---------getExprToken
	var getExprToken = function() {
		if((arguments.length === 2)&&(arguments[0] instanceof Array)&&(typeof(arguments[1])==="number")) {
			if((arguments[1]+1) < arguments[0].length) {
				return {
					"t": arguments[0][arguments[1]+1],
					"i": arguments[1]+1
				};
			} else {
				return {"i": -1};
			}
		} else {
			errPrint(1);
		}
	};
	//---------readExpr
	var readExpr = function() {
		if((arguments.length === 2)&&(arguments[0] instanceof Array)&&(typeof(arguments[1])==="number")) {
			var expr = arguments[0],
			exInd = getExprToken(expr, arguments[1]);
			return this.exprAdd(exInd.t, exInd.i);
		} else {
			errPrint(1);
		}
	};
	//---------exprAdd
	var exprAdd = function() {
		if(arguments.length === 2) {
			if(arguments[1]==0) {
				//
				//
				// 
				//
				// обработка присваивания
				// this.getTokenType: `num` or `word`
				// this.tokenIsKeyword: true or false
				//
				//
				//
			} else {
				//
				//
				//
				//
				// всё остальное
				//
				//
				//
				//
				//
			}
		} else {
			errPrint(1);
		}
	};
	//---------start_decode
	if(arguments.length > 0) {
		var code = arguments[0];
		if(code instanceof Array) {
			//---------проверка
			if(code.length == 0) {
				consoleWrite("Ничего нет.");
				return;
			}
			var len = code.length, i = 0;
			var iserr = !1, errnum = -1;
			//---------search_func
			while(i < len) {
				if((code[i] instanceof Object)&&(code[i].type=="function")) {
					GlobalFunc[code[i].name] = {
						"args": code[i].name.split(","),
						"stmt": code[i].stmt
					}
				}
				i++;
			}
			i = 0;
			//---------start
			while(i < len) {
				if(code[i] instanceof Array) {
					//простые команды
					this.readExpr(code[i], -1);
				} else if(code[i] instanceof Object) {
					//командный блок
				} else {
					iserr = !!1;
					errnum = 1
					break;
				}
				i++;
			}
			if(iserr) {
				this.errPrint(errnum);
			}
		} else {
			consoleWrite("Нет программы.");
		}
	} else {
		consoleWrite("Нет программы.");
	}
};
