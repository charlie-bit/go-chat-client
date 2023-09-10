package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func OperationIDGenerator() string {
	return strconv.FormatInt(time.Now().UnixNano()+int64(rand.Uint32()), 10)
}
func GetMsgID(sendID string) string {
	t := Int64ToString(GetCurrentTimestampByNano())
	return Md5(t + sendID + Int64ToString(rand.Int63n(GetCurrentTimestampByNano())))
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

// Get the current timestamp by Second

func GetCurrentTimestampBySecond() int64 {
	return time.Now().Unix()
}

// Get the current timestamp by Mill
func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Get the current timestamp by Nano
func GetCurrentTimestampByNano() int64 {
	return time.Now().UnixNano()
}

func StructToJsonString(param interface{}) string {
	dataType, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	dataString := string(dataType)
	return dataString
}
