/**
 * @Author:      leafney
 * @Date:        2022-10-09 10:03
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *DingTalk) send(msg interface{}) error {
	if utils.IsEmpty(r.token) {
		return vars.ErrTokenEmpty
	}

	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if r.debug {
		fmt.Printf("[debug] body: [%v] \n", string(m))
	}

	webURL := r.host
	value := url.Values{}
	value.Set("access_token", r.token)

	if utils.IsNotEmpty(r.secret) {
		timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
		signMsg := fmt.Sprintf("%s\n%s", timestamp, r.secret)
		value.Set("timestamp", timestamp)
		value.Set("sign", sign(signMsg, r.secret))
	}

	req, err := http.NewRequest(http.MethodPost, webURL, bytes.NewReader(m))
	if err != nil {
		return err
	}

	req.URL.RawQuery = value.Encode()
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if r.debug {
		fmt.Printf("[debug] url: [%v] \n", req.URL.String())
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if r.debug {
		fmt.Printf("[debug] response: [%v] \n", string(data))
	}

	var dr Response
	err = json.Unmarshal(data, &dr)
	if err != nil {
		return err
	}
	if dr.ErrCode != 0 {
		return fmt.Errorf("dingtalk notification send failed: [%v]", dr.ErrMsg)
	}

	return nil
}

func sign(message, secret string) string {
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
}
