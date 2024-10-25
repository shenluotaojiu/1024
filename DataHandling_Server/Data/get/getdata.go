package getData

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 该包主要用于主动从github获取数据

// 通过账户名获取数据并返回map类型
func GetDatabyUser(user string) (map[string]any, error) {
	szbuf := fmt.Sprintf("https://api.github.com/search/users?q=%s", user)
	// 发送GET请求
	resp, err := http.Get(szbuf)
	if err != nil {
		fmt.Println("请求失败:", err)
		return map[string]any{}, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return map[string]any{}, err
	}

	// 将json格式数据转化map
	var ans map[string]any
	err = json.Unmarshal([]byte(body), &ans)
	if err != nil {
		return map[string]any{}, err
	}
	fmt.Println("响应内容:", ans)
	fmt.Println("11111111", string(body))
	return ans, nil
}
