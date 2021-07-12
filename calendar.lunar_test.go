package carbon

import "testing"

func TestLunar_Animal(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "鼠"},
		{"2020-08-05", "鼠"},
		{"2021-07-07", "牛"},

		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2010-08-05", "虎"},
		{"2011-08-05", "兔"},
		{"2012-08-05", "龙"},
		{"2013-08-05", "蛇"},
		{"2014-08-05", "马"},
		{"2015-08-05", "羊"},
		{"2016-08-05", "猴"},
		{"2017-08-05", "鸡"},
		{"2018-08-05", "狗"},
		{"2019-08-05", "猪"},
		{"2020-08-05", "鼠"},
		{"2021-08-05", "牛"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Animal()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Festival(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2021-02-12", "春节"},
		{"2021-02-26", "元宵节"},
		{"2021-06-14", "端午节"},
		{"2021-08-14", "七夕节"},
		{"2021-08-22", "中元节"},
		{"2021-09-21", "中秋节"},
		{"2021-10-14", "重阳节"},
		{"2021-10-14", "重阳节"},
		{"2021-11-05", "寒衣节"},
		{"2021-11-19", "下元节"},
		{"2022-01-10", "腊八节"},
		{"2022-01-25", "小年"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Festival()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Year(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 2020},
		{"2020-08-05", 2020},
		{"2021-01-01", 2020},
		{"2021-07-07", 2021},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Year()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Month(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 4},
		{"2020-08-05", 6},
		{"2021-07-07", 5},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Month()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_Day(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 9},
		{"2020-08-05", 16},
		{"2021-07-07", 28},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().Day()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 4},
		{"2020-08-05", 4},
		{"2021-07-07", 0},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().LeapMonth()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToChinaYearString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "二零二零"},
		{"2020-08-05", "二零二零"},
		{"2021-01-01", "二零二零"},
		{"2021-07-07", "二零二一"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToChineseYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToChinaMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "四月"},
		{"2020-08-05", "六月"},
		{"2021-07-07", "五月"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToChineseMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToChinaDayString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "初九"},
		{"2020-08-05", "十六"},
		{"2021-07-07", "廿八"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToChineseDayString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_GanYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 7},
		{"2020-08-05", 7},
		{"2021-01-01", 7},
		{"2021-07-07", 8},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().GanYear()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_GanMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		//{"2020-05-01", 7},
		{"2020-08-05", 10},
		{"2021-01-01", 5},
		//{"2021-07-07", 2},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().GanMonth()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_GanDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 1},
		{"2020-08-05", 7},
		{"2021-01-01", 6},
		{"2021-07-07", 3},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().GanDay()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ZhiYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		{"2020-05-01", 1},
		{"2020-08-05", 1},
		{"2021-01-01", 1},
		{"2021-07-07", 2},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ZhiYear()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ZhiMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		//{"2020-05-01", 5},
		{"2020-08-05", 8},
		{"2021-01-01", 1},
		//{"2021-07-07", 8},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ZhiMonth()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ZhiDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output int    // 期望输出值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},
		//{"2020-05-01", 5},
		{"2020-08-05", 5},
		{"2021-01-01", 10},
		//{"2021-07-07", 5},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ZhiDay()

		if output != v.output {
			t.Errorf("Input %s, expected %d, but got %d\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanYearString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "庚"},
		{"2020-08-05", "庚"},
		{"2021-01-01", "庚"},
		{"2021-07-07", "辛"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"1988-12-01", "癸"},
		{"1999-06-15", "庚"},
		{"2013-10-18", "壬"},
		{"2019-07-02", "庚"},
		//{"2020-05-01", "庚"},
		{"2020-08-05", "癸"},
		{"2021-01-01", "戊"},
		//{"2021-07-07", "乙"},
		{"2021-07-12", "乙"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanDayString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "甲"},
		{"2020-08-05", "庚"},
		{"2021-01-01", "己"},
		{"2021-07-07", "丙"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanDayString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToZhiYearString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "子"},
		{"2020-08-05", "子"},
		{"2021-01-01", "子"},
		{"2021-07-07", "丑"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToZhiYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToZhiMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"1988-12-01", "亥"},
		{"1999-06-15", "午"},
		{"2013-10-18", "戌"},
		{"2019-07-02", "午"},
		//{"2020-05-01", "辰"},
		{"2020-08-05", "未"},
		{"2021-01-01", "子"},
		//{"2021-07-07", "未"},
		{"2021-07-12", "未"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToZhiMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToZhiDayString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "辰"},
		{"2020-08-05", "辰"},
		{"2021-01-01", "酉"},
		{"2021-07-07", "辰"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToZhiDayString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanZhiYearString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "庚子"},
		{"2020-08-05", "庚子"},
		{"2021-01-01", "庚子"},
		{"2021-07-07", "辛丑"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanZhiYearString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanZhiMonthString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"1988-12-01", "癸亥"},
		{"1999-06-15", "庚午"},
		{"2013-10-18", "壬戌"},
		{"2019-07-02", "庚午"},
		//{"2020-05-01", "庚辰"},
		{"2020-08-05", "癸未"},
		{"2021-01-01", "戊子"},
		//{"2021-07-07", "乙未"},
		{"2021-07-12", "乙未"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanZhiMonthString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_ToGanZhiDayString(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-05-01", "甲辰"},
		{"2020-08-05", "庚辰"},
		{"2021-01-01", "己酉"},
		{"2021-07-07", "丙辰"},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().ToGanZhiDayString()

		if output != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", true},
		{"2020-08-05", true},
		{"2021-01-01", true},
		{"2021-07-07", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsLeapYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is leap year, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is leap month, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", true},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsLeapMonth()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is leap month, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is leap month, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsRatYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", true},
		{"2020-08-05", true},
		{"2021-01-01", true},
		{"2021-07-07", false},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsRatYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rat, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rat, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsOxYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsOxYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of ox, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of ox, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2022-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsTigerYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of tiger, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of tiger, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsRabbitYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2023-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsRabbitYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rabbit, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rabbit, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2024-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsDragonYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of dragon, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of dragon, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2025-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsSnakeYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of snake, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of snake, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2026-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsHorseYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of horse, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of horse, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2027-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsGoatYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of goat, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of goat, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2028-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsMonkeyYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of monkey, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of monkey, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2029-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsRoosterYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of rooster, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of rooster, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsDogYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2030-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsDogYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of dog, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of dog, expected false, but got true\n", v.input)
			}
		}
	}
}

func TestLunar_IsPigYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output bool   // 期望输出值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},
		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2031-08-05", true},
	}

	for _, v := range Tests {
		output := Parse(v.input).Lunar().IsPigYear()

		if output == true {
			if v.output == false {
				t.Errorf("Input %s is year of pig, expected true, but got false\n", v.input)
			}
		} else {
			if v.output == true {
				t.Errorf("Input %s is year of pig, expected false, but got true\n", v.input)
			}
		}
	}
}