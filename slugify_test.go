package slugify

import (
	"testing"
	"reflect"
)

func TestSlugify(t *testing.T) {
	checkTable := []string{
		"Şok! Silivri ve Sason ilçe jandarma komutanları 'FETÖ'den tutuklandı", "sok-silivri-ve-sason-ilce-jandarma-komutanlari-fetoden-tutuklandi",
		"Erdoğan'ın tek adamlığı, Trump'ın menfaatine", "erdoganin-tek-adamligi-trumpin-menfaatine",
		"5 simple tips and tricks for writing unit tests in #golang", "5-simple-tips-and-tricks-for-writing-unit-tests-in-golang",
	}

	for i := 0; i < len(checkTable); i += 2 {
		response, err := Slugify(checkTable[i])
		if err != nil {
			t.Errorf("Error on Slugify: %e", err)
		}
		expect := checkTable[i+1]
		t.Logf("checking for: %s,\n result: %s,\n expected: %s\n", checkTable[i], response, expect)
		if response != expect {
			t.Errorf("Slugify failed. Given: %s, Expected: %s", response, expect)
		}
	}
}

func TestSetReplaceSet(t *testing.T) {
	checkTable := []string{
		"Şok! Silivri ve Sason ilçe jandarma komutanları 'FETÖ'den tutuklandı", "sok!-silivri-ve-sason-ilce-jandarma-komutanlari-fetoden-tutuklandi",
		"Erdoğan'ın tek adamlığı, Trump'ın menfaatine", "erdoganin-tek-adamligi-trumpin-menfaatine",
		"5 simple tips and tricks for writing unit tests in #golang", "5-simple-tips-and-tricks-for-writing-unit-tests-in-#golang",
	}

	slugify := GetWithCustomReplacer([]string{
		" ", "-",
		"'", "",
		"ı", "i",
		",", "",
		".", "",
	})

	for i := 0; i < len(checkTable); i += 2 {
		response, err := slugify.Slugify(checkTable[i])
		if err != nil {
			t.Errorf("Error on Slugify: %e", err)
		}
		expect := checkTable[i+1]
		t.Logf("checking for: %s,\n result: %s,\n expected: %s\n", checkTable[i], response, expect)
		if response != expect {
			t.Errorf("Slugify failed. Given: %s, Expected: %s", response, expect)
		}
	}
}

func TestGetWithCustomReplacer(t *testing.T) {
	slugifyInstance := GetWithCustomReplacer(nil)
	if instanceType := reflect.TypeOf(slugifyInstance).String(); instanceType != "*slugify.API" {
		t.Errorf("Returned instance type is not slugify.API but %s", instanceType)
	}
}