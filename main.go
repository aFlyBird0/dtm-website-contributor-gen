package main

// 速冲的小脚本，主要是一次性使用，就别纠结优雅性了哈哈哈
func main() {
	csvPath := "./secret/Contributors.csv"
	outputPath := "./output/Contributors.json"
	GenContributorsJson(csvPath, outputPath)
}
