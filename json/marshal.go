package json

import (
	"encoding/json"
	"fmt"
	"time"
)

// User 用户信息结构体
type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Age            int32     `json:"age"`
	Addr           string    `json:"addr"`
	OldExpiredTime time.Time `json:"-"`
}

// MarshalHelper marshal test
func MarshalHelper() {
	user1 := &User{
		ID:             1,
		Name:           "zj",
		OldExpiredTime: time.Now(),
	}
	b, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("marshal error")
	}
	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	fmt.Printf("%+v", u)
}

// Student 在处理json格式字符串的时候，经常会看到声明struct结构的时候，属性的右侧还有小米点括起来的内容。`TAB键左上角的按键，～线同一个键盘`
type Student struct {
	StudentID      string `json:"sid"`
	StudentName    string `json:"sname"`
	StudentClass   string `json:"class"`
	StudentTeacher string `json:"teacher"`
}

// StudentNoJSON 学生tag没有json
type StudentNoJSON struct {
	StudentID      string
	StudentName    string
	StudentClass   string
	StudentTeacher string
}

// StudentWithOption 可以选择的控制字段有三种：
// -：不要解析这个字段
// omitempty：当字段为空（默认值）时，不要解析这个字段。比如 false、0、nil、长度为 0 的 array，map，slice，string
// FieldName：当解析 json 的时候，使用这个名字
type StudentWithOption struct {
	StudentID      string //默认使用原定义中的值
	StudentName    string `json:"sname"`           // 解析（encode/decode） 的时候，使用 `sname`，而不是 `Field`
	StudentClass   string `json:"class,omitempty"` // 解析的时候使用 `class`，如果struct 中这个值为空，就忽略它
	StudentTeacher string `json:"-"`               // 解析的时候忽略该字段。默认情况下会解析这个字段，因为它是大写字母开头的
}

// StructJSON json marshal test
func StructJSON() {
	//NO.1 with json struct tag
	s := &Student{StudentID: "1", StudentName: "fengxm", StudentClass: "0903", StudentTeacher: "feng"}
	jsonString, _ := json.Marshal(s)

	fmt.Println(string(jsonString))
	//{"sid":"1","sname":"fengxm","class":"0903","teacher":"feng"}
	newStudent := new(Student)
	json.Unmarshal(jsonString, newStudent)
	fmt.Println(newStudent)
	//&{1 fengxm 0903 feng}
	//Unmarshal 是怎么找到结构体中对应的值呢？比如给定一个 JSON key Filed，它是这样查找的：
	// 首先查找 tag 名字（关于 JSON tag 的解释参看下一节）为 Field 的字段
	// 然后查找名字为 Field 的字段
	// 最后再找名字为 FiElD 等大小写不敏感的匹配字段。
	// 如果都没有找到，就直接忽略这个 key，也不会报错。这对于要从众多数据中只选择部分来使用非常方便。

	//NO.2 without json struct tag
	so := &StudentNoJSON{StudentID: "1", StudentName: "fengxm", StudentClass: "0903", StudentTeacher: "feng"}
	jsonStringO, _ := json.Marshal(so)

	fmt.Println(string(jsonStringO))
	//{"StudentId":"1","StudentName":"fengxm","StudentClass":"0903","StudentTeacher":"feng"}

	//NO.3 StudentWithOption
	studentWO := new(StudentWithOption)
	js, _ := json.Marshal(studentWO)

	fmt.Println(string(js))
	//{"StudentId":"","sname":""}

	studentWO2 := &StudentWithOption{StudentID: "1", StudentName: "fengxm", StudentClass: "0903", StudentTeacher: "feng"}
	js2, _ := json.Marshal(studentWO2)

	fmt.Println(string(js2))
	//{"StudentId":"1","sname":"fengxm","class":"0903"}
}
