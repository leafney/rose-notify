/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:18
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

import "testing"

func TestNewRobot(t *testing.T) {
	bot := NewFeiShu("")
	bot.UseSecret("")

	err := bot.
		SetDebug(true).
		SendText("消息发送测试\n测试")
	t.Log(err)
}
