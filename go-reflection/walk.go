package reflect

func walk(x interface{}, fn func(input string)) {
	fn("I still can't believe South Korea beat Germany 2-0")
}