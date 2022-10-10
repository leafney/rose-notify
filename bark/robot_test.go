/**
 * @Author:      leafney
 * @Date:        2022-10-09 17:55
 * @Project:     rose-notice
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package bark

import "testing"

func TestNewRobot(t *testing.T) {

	bot := NewRobot("https://api.day.app", "")
	//bot.SetKey("") // reset key
	err := bot.
		SetGroup("test").
		UseGet().
		//SetBadge(5).
		//SetLevel(1).
		//SetSound(SoundBell).
		SetCopy("吃饭啦").
		SetUrl("weixin://").
		SendText("中午啦")
	t.Log(err)
}
