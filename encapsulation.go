package gote

import (
	"fmt"
	"strings"
)

const (
	// BuffSize BuffSize
	BuffSize = 65 * 1024
	HUAWEI = "huawei"
	H3C = "h3c"
	CISCO = "cisco"
)

// RunCommands 封装执行命令
func RunCommands(user, password, ipPort string, cmds ...string) (string, error) {
	var (
		result string
		str    string
		err    error
	)
	conn, err := Dial("tcp", ipPort)
	if err != nil {
		panic("Unable to connect.")
	}
	// Read 30 bytes from the stream
	buf := make([]byte, BuffSize)
	_, err = conn.Read(buf)
	if err != nil {
		err = fmt.Errorf("Unable to read from stream")
		return str, err
	}
	_, err = conn.Write([]byte(fmt.Sprintf("%s\n",user)))
	buf = make([]byte, BuffSize)
	_, err = conn.Read(buf)
	result += string(buf)
	fmt.Println(string(buf))
	_, err = conn.Write([]byte(fmt.Sprintf("%s\n",password)))
	buf = make([]byte, BuffSize)
	_, err = conn.Read(buf)
	result += string(buf)
	fmt.Println(string(buf))

	for _, cmd := range cmds {
		_, err = conn.Write([]byte(fmt.Sprintf("%s\n", cmd)))
		buf = make([]byte, BuffSize)
		_, err = conn.Read(buf)
		result += string(buf)
	}

	return result, err
}

func GetBrand(user, passw, ipPort string)(string, error){
	var (
		cmds []string
		res string
		err error
		)
	cmds=append(cmds,"display version")
	cmds=append(cmds,"show version")
	res,err = RunCommands(user,passw,ipPort,cmds...)
	res = strings.ToLower(res)
	if strings.Contains(res, HUAWEI) {
		fmt.Sprintln("The switch brand is <huawei>.")
		res = HUAWEI
	} else if strings.Contains(res, H3C) {
		fmt.Sprintln("The switch brand is <h3c>.")
		res = H3C
	} else if strings.Contains(res, CISCO) {
		fmt.Sprintln("The switch brand is <cisco>.")
		res = CISCO
	}
	return res,err
}
