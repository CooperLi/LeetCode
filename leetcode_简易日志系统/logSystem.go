// package leetcode_简易日志系统
package main

import "fmt"

type LogSystem struct {
	ids []int
	ts  []int
}

func Constructor() LogSystem {
	return LogSystem{[]int{}, []int{}}
}

func (logSystem *LogSystem) add(id int, timeStamp int) {
	logSystem.ids = append(logSystem.ids, id)
	logSystem.ts = append(logSystem.ts, timeStamp)
}

func (logSystem *LogSystem) delete(id int) int {
	for i := range logSystem.ids {
		if logSystem.ids[i] == id {
			logSystem.ids = append(logSystem.ids[:i], logSystem.ids[i+1:]...)
			return 0
		}
	}
	return -1
}

func (logSystem *LogSystem) query(startTime int, endTime int) int {
	res := 0
	for i, _ := range logSystem.ids {
		time := logSystem.ts[i]
		if time >= startTime && time <= endTime {
			res++
		}
	}
	return res
}

func main() {
	log := Constructor()
	log.add(1, 5)
	log.add(2, 5)
	log.add(3, 6)
	fmt.Println(log.query(5, 6))
	fmt.Println(log.delete(2))
	fmt.Println(log.delete(4))
	fmt.Println(log.query(5, 6))
}

/*
map 做法

type LogSystem struct {
    log map[int]int
}


func Constructor() LogSystem {
    return LogSystem{log: make(map[int]int)}
}


func (logSystem *LogSystem) add(id int, timeStamp int)  {
    logSystem.log[id] = timeStamp
}


func (logSystem *LogSystem) delete(id int) int {
    if _, ok := logSystem.log[id]; ok {
        delete(logSystem.log, id)
        return 0
    }
    return -1
}


func (logSystem *LogSystem) query(startTime int, endTime int) int {
    res := 0
    for _, v := range logSystem.log {
        if v >= startTime && v <= endTime {
            res++
        }
    }
    return res
}
*/
