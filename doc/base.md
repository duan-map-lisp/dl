1、首先来个HelloWorld，
"HelloWorld"
[点击code](../test/base_sentence/hello_world.json)
嗯，写完了。
熟悉的lisp版HelloWorld

这里介绍一下dl语言的数据类型：
"int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "float32" "float64" "int" "uint" "rune" "byte" "uintptr" "string"
具体的可以直接参考go语言教程，扣的代码，啥都没变。
因为扣的go语言的json解释代码，无类型自动数据推导类型如下：
Bool                   对应JSON布尔类型
float64                对应JSON数字类型
string                 对应JSON字符串类型
[]interface{}          对应JSON数组
map[string]interface{} 对应JSON对象
nil                    对应JSON的null

复合数据类型有两种object和list。
对应golang的slice和map概念，实体为[]interface{}和map[string]interface{}。

2、代码块，block语句
是原版lisp的cond语句的变种，虽然block能被cond用宏实现出来，但是我为了回避造就了无数括号的官方的let语句的写法，需要一个仅有一层的代码块，block语句因此而生。
``` json
[
	"block",
	123,
	234,
	456,
	"jsidfo",
	"reoj"
]
```

[点击code](../test/base_sentence/block.json)
就是顺序执行，返回最后一条结果。
因为block可能有无数个参数，完全遵循lisp的模型，没有object版调用方式。

3、eval语句，解释器解释一段类型为[]byte的数据，得到解释器返回的结果。
``` json
[
	"block",
	"list_mode",
	["eval", ["import", "./test/base_sentence/block.json"]],
	"object_mode",
	{
		"name": "eval",
		"data": ["import", "./test/base_sentence/block.json"]
	}
]
```

[点击code](../test/base_sentence/eval.json)

4、load语句，解释器解释一段类型为[]byte的数据，直接将解释的数据返回，这也体现了数据与代码可互相转换的思想。
也是lisp写的代码可以写代码，自动写的代码产生的数据又可以编写自己，从而达到写的代码自己写自己，最终代码可自我进化的目的。(听起来有点儿玄幻，熟悉lisp的同学应该都清楚，自我进化的代码麦卡锡老师1958年就实现过了。)
``` json
[
	"block",
	"list_mode",
	["load", ["import", "./test/base_sentence/block.json"]],
	"object_mode",
	{
		"name": "load",
		"data": ["import", "./test/base_sentence/block.json"]
	}
]
```

[点击code](../test/base_sentence/eval.json)

5、qoute语句，lisp的经典语句，将传入参数当做数据来理解。用于将代码转换成数据。
``` json
[
	"block",
	"list_mode",
	["quote", {
			"fjeiowji": "jfioewj",
			"jfeiow": 1235
	}],
	"object_mode",
	{
		"name": "quote",
		"data": {
			"fjeiowji": "jfioewj",
			"jfeiow": 1235
		}
	}
]
```

[点击code](../test/base_sentence/quote.json)

6、7、let和get，定义参数和获取参数的值。
``` json
[
	"block",
	"list_mode",
	["let", "he", "uint64", 12345],
	["get", "he"],
	"object_mode",
    {   
        "name":"let",
        "symbol": "heheda",
        "type": "uint64",
        "value": 12345
    },  
	{
        "name": "get",
        "symbol": "heheda"
    }
]
```

[点击code](../test/base_sentence/let.json)

这个语法看起来已经跟lisp原生的let不同了，这是个有副作用的let语句。详情请看 [这里](./let.md) ，lisp的let会直接导致括号膨胀，我换了一种写法。
注意：整个机制都与lisp的let机制不同。get未定义的参数时返回空，对应json是null，对应golang是nil。

8、set，数据类型自动推导，赋值给已经定义过的参数。推导机制请参考go语言，扣（抄）的代码。
支持参数状态值修改，就不是纯函数式语言了，我觉得支持修改挺好，纯函数式语言写个循环都能疯。我还是普通一点吧。
``` json
[
	"block",
	["let", "he", "uint64", 12345],
	["get", "he"],
	"list_mode",
	["set", "he", 9],
	["get", "he"],
	"object_mode",
	{
		"name": "set",
		"symbol": "he",
		"value": 99
	},
	["get", "he"]
]
```

只允许基础类型赋值，不支持list和object类型赋值。list和object的赋值可以用函数或者宏自己写。

[code](../test/base_sentence/set.json)

9、del，删除一个symbol。从当前位置往上找作用域，将最近一个能找到的作用域的该symbol参数删除。
``` json
[
	"block",
	["let", "he", "uint64", 12345],
	["get", "he"],
	"list_mode",
	["del", "he"],
	["get", "he"]
]
```

[code](../test/base_sentence/del.json)

10、lambda，定义一个lambda，或者说是函数。传统函数概念其实是定义lambda并赋值给一个symbol。我这里采用lisp的方案，将其拆为两个步骤。
函数写起来一般都很长，需要的结构也很复杂，按照lisp的表格式写，其实非常头大。因此我决定lambda定义不允许使用lisp风格的写法，只允许object格式，不允许list格式。
lambda模型支持闭包，支持匿名函数，支持函数式编程，支持不存在的闭包变量的修改，支持保存lambda生成的当前参数域环境，因为说起来比较复杂，详情请看 [这里](./lambda.md) 。
``` json
[
	"block",
	{
		"name": "lambda",
		"body": {
			"name": "import",
			"filepath": {
				"name": "get",
				"symbol": "heheda"
			}
		}
	},
	{
		"name": "lambda",
		"body": [
			"block",
			["quote", {
				"heheda": 123,
				"jifoew": "jfioew"
			}]
		]
	}
]
```

[code](../test/base_sentence/lambda.json)

11、call，使用symbol调用一个lambda。因为lambda生成后暴露的就是个string，call传入的参数lambda其本质也就是一个string。
``` json
[
	"block",
	"list_mode",
	[
		"call", {
			"name": "lambda",
			"body": ["eval", {
				"name": "import",
				"filepath": {
					"name": "get",
					"symbol": "heheda"
				}
			}]
		},
		{
			"heheda": "./test/base_sentence/hello_world.json"
		}
	],
	"object_mode",
	{
		"name": "call",
		"lambda": {
			"name": "lambda",
			"body": [
				"block",
				["quote", {
					"heheda": 123,
					"jifoew": "jfioew"
				}]
			]
		}
	}
]
```

[code](../test/base_sentence/call.json)

12、string，将当前代码结点转换成string类型的json字符串返回。
``` json
[
	"block",
	["string", ["quote", {
		"fjeiowji": "jfioewj",
		"jfeiow": 1235
	}]],
	{
		"name": "string",
		"data": ["quote", {
			"fjeiowji": "jfioewj",
			"jfeiow": 1235
		}]
	}
]
```

[code](../test/base_sentence/string.json)

13、lisp的标准命令，这里不再赘述。[参考点这里](./lisp.md)

14、build，编译代码，并保存到该路径生成可执行程序。这里编译器代码没放出来，暂时没加，是这么个功能，先放这儿吧。

``` json
["build", ["get", "code"], "./heheda.elf"]
```

14、macro宏操作，这个更复杂， [单开一页讲macro](./macro.md)。代码分为执行期和预处理期。预处理期只要被执行过，就会反复重新执行预处理期，直到预处理期不再被触发为止，进入执行期。
编译的时候会将所有预处理期做完后得到稳定的执行期代码，编译执行期代码。
如果在执行期代码执行万恶的eval，eval会根据传入的数据，重新进入预处理其得到预处理期的执行过程，代码就可能在执行期根据传入数据不同重新编译执行自己，从而达到自己写自己的自我进化。

quote可以将代码转换为数据，
lambda可以将数据转换为代码，
eval可以在执行期进入预处理期，在预处理期修改自己。
build可以将预处理期的执行结果，编译成可执行程序，从而从代码转换成可执行的程序。

看看麦卡锡lisp的设计，真的发自肺腑的感叹，lisp是世界上最好的语言，当代各种编程语言的思想设计，从来就没有超越过lisp。

15、plugin，加载插件。
可以发现基本语法就这些了，但比如"+", "-", "*"之类的正常数学运算符都没有。
毕竟哪怕是数学运算，点乘叉乘之类的，有很多规范，我觉得应该是个第三方数学运算库，而不应该是语言内置的。
参考我提供的go语言代码，编写中间件，可以添加功能，添加任何go语言能调用的代码，依赖，并添加symbol标号，lambda标号，或者说功能。
我这里标准库里先写了一个加法"+"，请求标准库代码。make all编译之后生成std.so插件，使用plugin命令找到插件加载。该功能暂未完善，后续有空了再慢慢写。
使用plugin命令加载第三方插件，可以引入新命令。

``` json
[
	"block",
	{
		"name": "plugin",
		"filepath": "./lib/std.so"
	},
	[
		"+",
		6,
		3,
		1,
		7.7
	]
]
```

[code](../test/base_sentence/math_add.json)


教程先写这么多了，有时间了再写吧。
