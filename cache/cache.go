package cache

import (
	c "aibot/constant"
)

type AIInfo struct {
	AIMode c.AIMode
	Data   any
}

var (
	userCache = make(map[string]*AIInfo)
)

func InitUserCache(lineId string) {
	if _, ok := userCache[lineId]; !ok {
		// default is gpt
		userCache[lineId] = &AIInfo{AIMode: c.GPT}
	}
}

func RemoveUserCache(lineId string) {
	delete(userCache, lineId)
}

func SetAIMode(lineId string, mode c.AIMode) {
	userCache[lineId].AIMode = mode
}

func GetAIMode(lineId string) c.AIMode {
	return userCache[lineId].AIMode
}
