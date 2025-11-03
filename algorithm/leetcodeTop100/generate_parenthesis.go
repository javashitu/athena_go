package leetcodetop100

type GenerateParenthesis struct {
	result []string
}

func (g *GenerateParenthesis) generateParenthesis(num int) []string {
	g.dfs(num, num, "")
	return g.result
}

func (g *GenerateParenthesis) dfs(left int, right int, str string) {
	if left < 0 || right < 0 || left > right {
		return
	}
	if left == 0 && right == 0 {
		g.result = append(g.result, str)
		return
	}
	g.dfs(left-1, right, str+"(")
	g.dfs(left, right-1, str+")")
}
