## 钉钉 Webhook Robot

- [自定义机器人接入 - 钉钉开放平台](https://open.dingtalk.com/document/robots/custom-robot-access) 


## Feature

- `NewDingTalk(token string)`
- `UseToken(token string)`
- `UseSecret(secret string)`
- `SendText(text string)`
- `SendTextAt(text string, atMobiles []string, isAtAll bool)`
- `SendLink(title, text, messageURL, picURL string)`
- `SendMarkdown(title, body string)`
- `SendMarkdownAt(title, body string, atMobiles []string, isAtAll bool)`
- `SetDebug(debug bool)`

----
