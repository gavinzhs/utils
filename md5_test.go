package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMd5(t *testing.T) {
	// 测试正常情况
	str := "test"
	expectedMd5 := "098f6bcd4621d373cade4e832627b4f6"
	actualMd5 := StringMd5(str)
	assert.Equal(t, expectedMd5, actualMd5, "MD5 计算结果不一致")

	// 测试空字符串
	str = ""
	expectedMd5 = "d41d8cd98f00b204e9800998ecf8427e"
	actualMd5 = StringMd5(str)
	assert.Equal(t, expectedMd5, actualMd5, "空字符串的 MD5 计算结果不一致")
}
