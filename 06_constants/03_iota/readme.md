由以上五个例子可以得出一下结论：

1. `iota`只能在常量表达式中使用
2. 每次出现`const`都会将iota初始化为0
3. iota从出现const后第一个常量开始，从0递增。从第二个iota开始可以省略
4. iota可以进行算数运算
5. 可以用 _ 跳过不想要的iota值，只需将 _ 单独放在一行即可
6. 如果iota序列被插队，那么iota依然会增加（根据被插队行数）

you are smart .you can draw a conclusion by my examples.
