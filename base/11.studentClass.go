package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Student struct {
	Name string
	Age int
	Score int
}

func AddStudent(student []Student) []Student{
	var name string
	var age int
	var score int
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入学生年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入学生成绩：")
	fmt.Scanln(&score)
	student = append(student, Student{Name: name, Age: age, Score: score})
	return student
}

func QueryStudent(student []Student, name string) {
	for _, student := range student {
		if student.Name == name {
			fmt.Println(student.Name, student.Age, student.Score)
		}
	}
}

func CalculateAverage(student []Student) {
	var calc int
	for _, student := range student {
		calc += student.Score
	}
	fmt.Println("平均分：", calc/len(student))
}

func FindMaxScore(student []Student) {
	max := 0
	for _, student := range student {
		if student.Score > max {
			max = student.Score
		}
	}
	fmt.Println("最高分：", max)
}

func main() {
	students := []Student{}
	reader := bufio.NewReader(os.Stdin)
	// 添加学生
	for {
		fmt.Println("这是第%d个学生", len(students)+1)
		line, _:= reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		parts := strings.Fields(line)
		if (len(parts) != 3) {
			fmt.Println("格式错误，请输入：姓名 年龄 分数")
			continue
		}
		name := parts[0]
		age, _ := strconv.Atoi(parts[1])
		score, _ := strconv.Atoi(parts[2])
		students = append(students, Student{Name: name, Age: age, Score: score})
	}

	for {
		option := 0
		fmt.Println("-----请问你想干什么？-----")
		fmt.Println("1. 添加学生")
		fmt.Println("2. 查询学生")
		fmt.Println("3. 计算平均分")
		fmt.Println("4. 找出最高分")
		fmt.Println("5. 退出")
		fmt.Scanln(&option)
		switch option {
		case 1:
			students = AddStudent(students)
		case 2:
			var name string
			fmt.Println("请输入学生的名字")
			fmt.Scanln(&name)
			QueryStudent(students, name)
		case 3:
			CalculateAverage(students)
		case 4:
			FindMaxScore(students)
		case 5:
			return
		default:
			fmt.Println("无效的选项")
		}
	}
}