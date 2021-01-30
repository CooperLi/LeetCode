package leetcode_635_设计日志存储系统

/*
你将获得多条日志，每条日志都有唯一的 id 和 timestamp，timestamp 是形如 Year:Month:Day:Hour:Minute:Second 的字符串
例如 2017:01:01:23:59:59，所有值域都是零填充的十进制数。

设计一个日志存储系统实现如下功能：

void Put(int id, string timestamp)：
给定日志的 id 和 timestamp，将这个日志存入你的存储系统中。

int[] Retrieve(String start, String end, String granularity)：
返回在给定时间区间内的所有日志的 id。start 、 end 和 timestamp 的格式相同，granularity 表示考虑的时间级。
比如，start = "2017:01:01:23:59:59", end = "2017:01:02:23:59:59", granularity = "Day" 代表区间 2017 年 1 月 1 日到 2017 年 1 月 2 日。

样例 1 ：
put(1, "2017:01:01:23:59:59");
put(2, "2017:01:01:22:59:59");
put(3, "2016:01:01:00:00:00");
retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Year"); // 返回值 [1,2,3]，返回从 2016 年到 2017 年所有的日志。
retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Hour"); // 返回值 [1,2], 返回从 2016:01:01:01 到 2017:01:01:23 区间内的日志，日志 3 不在区间内。

注释 ：
Put 和 Retrieve 的指令总数不超过 300。
年份的区间是 [2000,2017]，小时的区间是 [00,23]。
Retrieve 的输出顺序不作要求。
*/

type LogSystem struct {
	ids []int
	ts  []string
}

func Constructor() LogSystem {
	return LogSystem{[]int{}, []string{}}
}

func (this *LogSystem) Put(id int, timestamp string) {
	this.ids = append(this.ids, id)
	this.ts = append(this.ts, timestamp)
}

// 时间级 <---> 在字符串中的范围
var keyMap = map[string]int{
	"Year":   4,  // [0:4]
	"Month":  7,  // [0:7]
	"Day":    10, // [0:10]
	"Hour":   13, // [0:13]
	"Minute": 16, // [0:16]
	"Second": 19, // [0:19]
}

/*
此解法完全利用到了字符串大小的比较，忽略了时间转换等复杂操作，着实秀
*/
func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	var res []int
	// 通过HashMap找到对应的范围
	j := keyMap[gra]
	for i, id := range this.ids {
		// 比较的时间范围
		time := this.ts[i][:j]
		// 直接字符串比较大小，等同于时间
		// 大于等于起始时间，小于等于结束时间，符合，加入id
		if time >= s[:j] && time <= e[:j] {
			res = append(res, id)
		}
	}
	return res
}

/**
 * Your LogSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(id,timestamp);
 * param_2 := obj.Retrieve(s,e,gra);
 */
