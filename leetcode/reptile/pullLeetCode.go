package reptile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const url string = "https://leetcode-cn.com/api/user_submission_calendar"

type UserCommit struct {
	NickName string
	Commits  []Commit
}

func (uc UserCommit) toJsonString() string {
	jsonBytes, _ := json.Marshal(uc)

	return string(jsonBytes)
}

type Commit struct {
	Ts    int64
	Count int
}

var total = make(map[string]UserCommit)

func PullLeetCodeUser(nickName string) UserCommit {
	commits := PullLeetCodeCommit(nickName)

	return UserCommit{nickName, commits}
}

func PullLeetCodeCommit(nickName string) []Commit {
	d, err := http.Get(fmt.Sprintf("%s/%s/", url, nickName))
	if err != nil {
		panic(fmt.Sprintf("[%s] get error [%s]", url, err))
	}
	defer d.Body.Close()

	body, _ := ioutil.ReadAll(d.Body)
	bodyStr := string(body)

	var jsonStr string
	var mp map[string]int
	err = json.Unmarshal([]byte(bodyStr), &jsonStr)
	err = json.Unmarshal([]byte(jsonStr), &mp)
	if err != nil {
		panic(fmt.Sprintf("[%s] convert to %T error", bodyStr, Commit{}))
	}

	var result []Commit
	for k, v := range mp {
		ts, _ := strconv.ParseInt(k, 10, 64)
		result = append(result, Commit{ts, v})
	}

	return result
}

func formatEasyRead(userCommits map[string]UserCommit) map[int64]map[string]int {
	result := make(map[int64]map[string]int)

	for _, userCommit := range userCommits {
		for _, commit := range userCommit.Commits {
			if v, ok := result[commit.Ts]; ok {
				v[userCommit.NickName] = commit.Count
			} else {
				result[commit.Ts] = map[string]int{userCommit.NickName: commit.Count}
			}
		}
	}

	return result
}

func asyncPull() {
	for {
		for k, _ := range total {
			total[k] = PullLeetCodeUser(k)
			fmt.Println("pull", k)
		}

		time.Sleep(1 * time.Minute)
	}
}

func unRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if r.Method != "POST" {
		fmt.Fprintln(w, "error")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if len(body) == 0 {
		fmt.Fprintln(w, "error")
		return
	}

	nickName := string(body)

	delete(total, nickName)
	fmt.Fprintln(w, "success")
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if r.Method != "POST" {
		fmt.Fprintln(w, "error")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if len(body) == 0 {
		fmt.Fprintln(w, "error")
		return
	}

	nickName := string(body)

	newUc := PullLeetCodeUser(nickName)
	total[nickName] = newUc

	fmt.Fprintf(w, "success")
}

func show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := formatEasyRead(total)
	jsonBytes, _ := json.Marshal(result)

	fmt.Fprintf(w, string(jsonBytes))
}

func TestPull() {
	http.HandleFunc("/", show)
	http.HandleFunc("/register", register)
	http.HandleFunc("/unregister", unRegister)

	go asyncPull()

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
