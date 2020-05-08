PS: 有空了再补这个详细文档，先写这些吧。宏这个太复杂，我简单写一写吧，等标准库完善了，有空了再写。

1、macro
定义一个宏，其实就是lambda，代码写的是直接调用lambda，换个名字区分执行期与预处理期。这也说明了预处理期可使用的语法代码跟执行期一模一样，代码是自递归的。
加了一个参数symbol，毕竟非执行期，let无法在最外层使用，定义宏的名字就成问题了。
如果symbol字段存在，将生成的宏的去重hash写入symbol中，按照提供的symbol字段换名。
如果symbol字段不存在或为空，认为是匿名宏，直接返回生成的去重hash。

macro定义后需要自我销毁，否则下一轮macro又会被执行。

而且这里注意一点，跟let命令一样，定义symbol的时候不是定义在当前位置，而是定义到了当前节点的父节点上。
[
	"block",
	{
		"name": "macro",
		"body": {
			"name": "import",
			"filepath": {
				"name": "get",
				"symbol": "heheda"
			}
		}
	},
	{
		"name": "macro",
		"symbol": "heheda",
		"body": {
			"name": "import",
			"filepath": {
				"name": "get",
				"symbol": "heheda"
			}
		}
	},
	["get", "heheda"]
]
[点击code](../test/base_sentence/macro.json)

2、exmacro
假设exmacro所在作用域为A域，假设生成macro时所在作用域为B。exmacro宏展开实质上是在调用生成lambda时的状态，假设lambda该作用域为B。
(1)获取args所在A作用域的一组参数，以symbol的方式，写入到B作用域，提前会读取B所在作用域内的body节点，是否存在该symbol名字，如果存在报重定义错误。
(2)lambda执行B作用域的body，如果需要参数，从刚刚设置过的body里可以获取到传入的symbol参数，获取的参数的fatherNode指向的又是A作用域，所以lambda在执行的时候，又可以使用来自A作用域的参数。
(3)预处理期的宏跟执行期的lambda最大的区别就是这一步。lambda的执行结果是返回了。宏展开得到的执行结果，会将当前所在节点替换掉，从而修改自身。
(4)在执行结果覆盖自身之后，宏展开自身便消失了。如果预处理期展开过宏，重新加载展开后的所有代码，可能展开的宏本身还可能继续展开，直到遍历所有节点，已经不再有宏被展开了。预处理期结束，代码里不再有宏代码。
(5)预处理结束后，为防止预处理期垃圾干扰执行期。从根节点遍历所有节点，清理所有预处理期执行痕迹。
(6)进入执行期，执行所有代码。得到执行结果。

3、remacro，正则宏。
lisp里有一些缩写，如:aaa, 'bbb, '(123, 345)，之类的。用单引号等简化书写，类似用.来寻找类内成员一样。
一般这种简写都是语法内置的。我这里提出一个玩法，注册正则宏后，被正则匹配的字符串以及正则子句会传递给已经注册的某个宏。然后达到自定义缩写的写法。
这里列举两个简单的例子。

[
	"block",
	"begin_define_macro",
	{
		"name": "macro",
		"symbol": "yo",
		"body": [
			"block",
			"begin宏展开期",
			["get", "reargs"],
			["car", ["get", "reargs"]],
			["car", ["car", ["get", "reargs"]]],
			["cdr", ["car", ["get", "reargs"]]],
			["car", ["cdr", ["car", ["get", "reargs"]]]],
			["cons", ["car", ["cdr", ["car", ["get", "reargs"]]]],
				["quote", {
					"name": "let",
					"type": "uint64",
					"value": 999
				}],
				"symbol"
			]
		]
	},
	"end_define_macro",
	"begin_define_remacro",
	{
		"name": "remacro",
		"regexp": "^:(.*)$",
		"macro": "yo"
	},
	"end_define_remacro",
	"begin_exmacro",
	":heheda",
	["get", "heheda"],
	"end_exmacro"
]
[点击code](../test/base_sentence/remacro.json)

[
	"block",
	{
		"name": "macro",
		"symbol": "m_get",
		"body": [
			"cons",
			["car", ["cdr", ["car", ["get", "reargs"]]]],
			["quote", {"name": "get"}],
			"symbol"
		]
	},
	{
		"name": "remacro",
		"regexp": "^@(.*)$",
		"macro": "m_get"
	},

	["let", "heheda", "uint64", 897],
	"@heheda"
]
[点击code](../test/base_sentence/remacro_get.json)
这里的"@heheda"会自动展开为["get", "heheda"]

4、safe，保护标记。
因为上面的remacro正则宏太强大了，正常的字符串一旦符合正则规则就会被展开。容易损坏数据。
所以我设计了一个保护函数"safe"，预处理期一旦见到这个名字，无论是list模式下的第一个字符串参数，还是object模式下的name。
safe后面的代码块都不会在预处理期被处理。
执行期遇到safe的时候，执行效果是将data区域的数据，替换掉自己所在的位置。代码层级不会因为使用了safe而降级。
