package payload

import (
	"bufio"
	"os"
	"strings"
)

// 解析 payloads
func ParsePayloads(payloads string) ([]string, error) {
	var payloadList []string

	// 判断是否是文件路径
	if _, err := os.Stat(payloads); err == nil {
		// 从文件中读取 payload
		file, err := os.Open(payloads)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			payloadList = append(payloadList, scanner.Text())
		}
	} else {
		// 从命令行参数中读取 payload
		if strings.Contains(payloads, ",") {
			payloadList = strings.Split(payloads, ",")
		} else {
			payloadList = append(payloadList, payloads)
		}
	}

	return payloadList, nil
}
