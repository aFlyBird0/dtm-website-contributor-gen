package main

import (
	"flag"
)

// 速冲的小脚本，主要是一次性使用，就别纠结优雅性了哈哈哈
func main() {
	var csvDir, outputPath string

	flag.StringVar(&csvDir, "data", "./secret", "源数据文件夹，里面应当至少有一个 csv 文件")
	flag.StringVar(&outputPath, "output", "Contributors.json", "输出文件")

	flag.Parse()

	GenContributorsJson(csvDir, outputPath)
}
