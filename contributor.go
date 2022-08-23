package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Contributor struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Avatar            string `json:"avatar"`
	CertificationLink string `json:"certificationLink"`
	Badge             *Badge `json:"badge"`
	DateAdded         string `json:"dateAdded"`
}
type Contributors []Contributor
type Badge struct {
	Type        string `json:"type"`
	Achievement string `json:"achievement"`
}

func getContributorsFromCsv(filepath string) (contributors Contributors) {
	opencast, err := os.Open(filepath)
	if err != nil {
		log.Fatal("csv文件打开失败！")
	}
	defer opencast.Close()

	//创建csv读取接口实例
	ReadCsv := csv.NewReader(opencast)

	//获取一行内容，一般为第一行内容
	read, _ := ReadCsv.Read() //返回切片类型：[chen  hai wei]
	log.Println(read)

	//读取所有内容
	// "GitHub ID","人员","微信昵称（ID）","证书日期","证书类型","证书顺序","证书发放","礼物寄送","邮箱","证书英文名"
	all, err := ReadCsv.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range all {
		// 证书名字, 默认用「证书英文名」
		name := strings.TrimSpace(line[9])
		if name == "" {
			// 否则使用中文拼音
			name = nameToPinyin(strings.TrimSpace(line[1]))
		}
		// GitHub ID(昵称)
		github := removeBracket(strings.TrimSpace(line[0]))
		// 证书日期, 2022/03/05
		date := line[3]
		contributor := Contributor{
			Name:              name,
			Email:             line[8],
			Avatar:            "https://avatars.githubusercontent.com/" + github,
			CertificationLink: "https://www.credly.com/organizations/devstream/badge", // 这个链接后面手动一个个改
			//Badge:             badge,
			DateAdded: date,
		}

		certifyType := line[4]
		var badge *Badge
		switch certifyType {
		case "Contributor - Associate":
			badge = &Badge{Type: "Open-Source Contributor", Achievement: "Associate"}
		case "Contributor - Professional":
			badge = &Badge{Type: "Open-Source Contributor", Achievement: "Professional"}
		case "Evangelist - Associate":
			badge = &Badge{Type: "Evangelist", Achievement: "Associate"}
		case "Evangelist - Professional":
			badge = &Badge{Type: "Evangelist", Achievement: "Professional"}
		case "Speaker - Associate":
			badge = &Badge{Type: "Talented Speaker", Achievement: "Associate"}
		case "Speaker - Professional":
			badge = &Badge{Type: "Talented Speaker", Achievement: "Professional"}
		case "Month Top 3":
			badge = &Badge{Type: "TopN", Achievement: "Top 3 of Month"}
		case "Year Top 3":
			badge = &Badge{Type: "TopN", Achievement: "Top 3 of Year"}
		case "Year Top 10":
			badge = &Badge{Type: "TopN", Achievement: "Top 10 of Year"}
		default:
			log.Fatalf("非法的证书类型 %s", certifyType)
		}
		contributor.Badge = badge

		contributors = append(contributors, contributor)
	}

	return contributors
}

func (cs Contributors) ToJson() string {
	bytes, err := json.MarshalIndent(cs, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

// 去除括号里的GitHub昵称
func removeBracket(src string) string {
	src = strings.TrimSpace(src)
	pattern1 := `\([^)]*\)` // 英文括号
	pattern2 := `（[^)]*）`   // 中文括号
	reg1, _ := regexp.Compile(pattern1)
	reg2, _ := regexp.Compile(pattern2)

	destBytes := reg1.ReplaceAll([]byte(src), []byte(""))
	destBytes = reg2.ReplaceAll(destBytes, []byte(""))

	return string(destBytes)
}

func GenContributorsJson(csvDir, outputPath string) {
	files, err := getFilesOfDir(csvDir)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		log.Fatalf("%s 文件夹下未检测到任何 csv 文件", csvDir)
	}

	var contributors Contributors
	for _, file := range files {
		contributors = append(contributors, getContributorsFromCsv(file)...)
	}
	fmt.Println(contributors[:3:len(contributors)].ToJson())

	// 创建文件夹
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		log.Fatalf("创建文件夹 %s 失败: %v", filepath.Dir(outputPath), err)
	}
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(contributors.ToJson())
	if err != nil {
		log.Fatal(err)
	}
}

func getFilesOfDir(dir string) (files []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".csv" {
			files = append(files, path)
		}
		return nil
	})

	return
}
