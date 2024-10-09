package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func main() {

	{
		fmt.Println("四則演算 + カッコ使えそうか")
		code := `((1+100-1)*5 )/5 + Num`

		type Env struct {
			Num int
		}
		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			panic(err)
		}

		env := Env{
			Num: 99,
		}
		output, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}

		fmt.Println(code, "=>", output)
		fmt.Println()
	}

	{
		fmt.Println("if文")
		code := `1 + (2 == 2 ? 100 : 0)`

		type Env struct{}
		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			panic(err)
		}

		env := Env{}
		output, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}

		fmt.Println(code, "=>", output)
		fmt.Println()
	}

	{
		fmt.Println("比較演算子(nullかどうか)(数字が0かどうか)(文字が空文字かどうか)")
		code := `1 + (Num1 == nil ? 10 : 20) + (Num2 == 0 ? 100 : 200) + (Str1 == "" ? 1000 : 2000)`

		type Env struct {
			Num1 *int
			Num2 int
			Str1 string
		}
		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			panic(err)
		}

		env := Env{
			Num2: 0,
			Str1: "",
		}
		output, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}

		fmt.Println(code, "=>", output)
		fmt.Println()
	}

	{
		fmt.Println("割り算の部分でゼロ除算したらどうなるか")
		code := `1 + (0 / 0) + 10`

		type Env struct {
		}
		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			panic(err)
		}

		env := Env{}
		output, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}

		fmt.Println(code, "=>", output)
		fmt.Println()
	}

	{
		fmt.Println("[型違い](整数 & 小数)")
		code := `1 + 1.0`

		type Env struct {
		}
		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			panic(err)
		}

		env := Env{}
		output, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}

		fmt.Println(code, "=>", output)
		fmt.Println()
	}

	// Panic
	// {
	// 	// panic: invalid operation: + (mismatched types int and string) (1:3)
	// 	//
	// 	fmt.Println("[型違い](整数 & 文字列 など)")
	// 	code := `1 + "ok"`

	// 	type Env struct {
	// 	}
	// 	program, err := expr.Compile(code, expr.Env(Env{}))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	env := Env{}
	// 	output, err := expr.Run(program, env)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(code, "=>", output)
	// 	fmt.Println()
	// }

	{
		fmt.Println("[built-in関数]要素のSUM")

		code := `sum(map(Tweets, .Len))`

		type Tweet struct {
			Len int
		}

		type Env struct {
			Tweets []Tweet
		}
		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			panic(err)
		}

		env := Env{
			Tweets: []Tweet{{42}, {98}, {69}},
		}
		output, err := expr.Run(program, env)
		if err != nil {
			panic(err)
		}

		fmt.Println(code, "=>", output)
		fmt.Println()

	}

}
