package base58_test

import (
	"testing"

	"github.com/checksum0/go-cryptoutils/base58"
)

var checkEncodingStringTests = []struct {
	version byte
	in      string
	out     string
}{
	{20, "", "3MNQE1X"},
	{20, " ", "B2Kr6dBE"},
	{20, "-", "B3jv1Aft"},
	{20, "0", "B482yuaX"},
	{20, "1", "B4CmeGAC"},
	{20, "-1", "mM7eUf6kB"},
	{20, "11", "mP7BMTDVH"},
	{20, "abc", "4QiVtDjUdeq"},
	{20, "1234598760", "ZmNb8uQn5zvnUohNCEPP"},
	{20, "abcdefghijklmnopqrstuvwxyz", "K2RYDcKfupxwXdWhSAxQPCeiULntKm63UXyx5MvEH2"},
	{20, "00000000000000000000000000000000000000000000000000000000000000", "bi1EWXwJay2udZVxLJozuTb8Meg4W9c6xnmJaRDjg6pri5MBAxb9XwrpQXbtnqEoRV5U2pixnFfwyXC8tRAVC8XxnjK"},
}

func TestBase58Check(t *testing.T) {
	for x, test := range checkEncodingStringTests {
		if result := base58.CheckEncode([]byte(test.in), test.version); result != test.out {
			t.Errorf("CheckEncode(%d) = %s (want: %s)", x, result, test.out)
		}

		result, version, err := base58.CheckDecode(test.out)
		if err != nil {
			t.Errorf("CheckDecode(%d) = err %v", x, err)
		} else if version != test.version {
			t.Errorf("CheckDecode(%d) = version %d (want: %d)", x, version, test.version)
		} else if string(result) != test.in {
			t.Errorf("CheckDecode(%d) = string %s (want: %s)", x, result, test.in)
		}
	}

	_, _, err := base58.CheckDecode("3MNQE1Y")
	if err != base58.ErrChecksum {
		t.Error("CheckDecode test failed, expected ErrChecksum")
	}

	testString := ""
	for length := 0; length < 4; length++ {
		_, _, err = base58.CheckDecode(testString)
		if err != base58.ErrInvalidFormat {
			t.Error("CheckDecode test failed, expected ErrInvalidFormat")
		}
	}
}
