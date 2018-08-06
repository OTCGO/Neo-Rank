package redis

import (
	"Neo-Rank/utils/redis"
	"testing"
)

func TestRedisSet(t *testing.T) {
	testCases := []struct {
		key    string
		value  string
		expect bool
	}{
		{"key1", "value1", true},
	}

	var redis = &redis.RedisCon{}
	for _, tc := range testCases {

		ok, err := redis.Set(tc.key, tc.value)
		if err != nil {
			t.Log("err", err)
			t.Fail()
			return
		}

		if ok != tc.expect {
			t.Log("ok", ok)
			t.Fail()
			return
		}

		t.Log("ok", ok)

	}
}

func TestRedisGet(t *testing.T) {
	testCases := []struct {
		key    string
		expect string
	}{
		{"key444", "value1"},
	}

	var redis = &redis.RedisCon{}
	for _, tc := range testCases {

		ok, err := redis.Get(tc.key)
		if err != nil {
			t.Log("err", err)
			t.Fail()
			return
		}

		if ok != tc.expect {
			t.Log("ok", ok)
			t.Fail()
			return
		}

		t.Log("ok", ok)

	}
}
