package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID:%d | email:%s\nActivity log:\n", u.id, u.email)
	for i, log := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, log.action, log.timestamp)
	}

	return out
}

func main() {
	fmt.Println("Hi. Wait groups")

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	go fmt.Println("Konkurentnij in anonim function")
	// }()
	// go fmt.Println("Konkurentnij 1")
	// go fmt.Println("Konkurentnij 2")
	// go fmt.Println("Konkurentnij 3")
	// go fmt.Println("Konkurentnij 4")
	// time.Sleep(time.Second * 3)
	// fmt.Println("NE Konkurentnij")

	// user := User{
	// 	id:    1,
	// 	email: "ivan@gmail.com",
	// 	logs: []logItem{
	// 		{action: actions[0], timestamp: time.Now()},
	// 		{action: actions[1], timestamp: time.Now()},
	// 		{action: actions[2], timestamp: time.Now()},
	// 	},
	// }
	// fmt.Println(user.getActivityInfo())

	rand.Seed(time.Now().Unix())
	for _, user := range generateUsers(3) {
		saveUserInfo(user)
	}
}

func saveUserInfo(user User) error {
	fmt.Printf("WRITING FILE FOR USER %d\n", user.id)

	fileName := fmt.Sprintf("logs/uuid_%d.txt", user.id)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(user.getActivityInfo())

	return nil
}

func generateUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs:  generateLogs(rand.Intn(3)),
		}
	}

	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
