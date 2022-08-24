package main

import (
	"regexp"
)

const defaultBadgeLink = "https://www.credly.com/organizations/devstream/badges"

// 在 credly 后台能直接复制的链接是这个 https://www.credly.com/mgmt/organizations/19f0ed81-a1a5-4df4-bd76-e3d40e23c328/badges/earners/59c69146-58a7-471d-a908-bc5a0b7f5f6f/details
// 我们需要转换成这个 https://www.credly.com/badges/59c69146-58a7-471d-a908-bc5a0b7f5f6f
func handleCredlyLink(origin string) (res string) {
	if origin == "" {
		return defaultBadgeLink
	}

	// 如果 origin 已经是需要转换的格式了，就不用转换了
	re1 := regexp.MustCompile(`https://www.credly.com/badges/([0-9A-Za-z-]{6,50})/?`)
	matches := re1.FindStringSubmatch(origin)
	if len(matches) >= 2 {
		return origin
	}

	// 否则尝试匹配 https://www.credly.com/mgmt/organizations/19f0ed81-a1a5-4df4-bd76-e3d40e23c328/badges/earners/59c69146-58a7-471d-a908-bc5a0b7f5f6f/details
	re2 := regexp.MustCompile(`https://www.credly.com/mgmt/organizations/[0-9A-Za-z-]{6,50}/badges/earners/([0-9A-Za-z-]{6,50})/details/?`)
	matches = re2.FindStringSubmatch(origin)
	if len(matches) >= 2 {
		// matches[0] 是整个字符串
		return "https://www.credly.com/badges/" + matches[1]
	}

	return defaultBadgeLink
}
