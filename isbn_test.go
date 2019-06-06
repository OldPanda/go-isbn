package isbn

import "testing"

func TestValidate(t *testing.T) {
	stringIsbn10 := "helloworld"
	if res := Validate(stringIsbn10); res {
		t.Errorf("Validate(\"helloworld\") = %t; want false", res)
	}

	stringIsbn13 := "978helloworld"
	if res := Validate(stringIsbn13); res {
		t.Errorf("Validate(\"978helloworld\") = %t; want false", res)
	}

	wrongLength := "lengthisnotcorrect"
	if res := Validate(wrongLength); res {
		t.Errorf("Validate(\"lengthisnotcorrect\") = %t; want false", res)
	}

	validIsbn13 := "9787532736553"
	if res := Validate(validIsbn13); !res {
		t.Errorf("Validate(\"9787532736553\") = %t; want true", res)
	}

	wrongPrefixIsbn13 := "1237532736553"
	if res := Validate(wrongPrefixIsbn13); res {
		t.Errorf("Validate(\"1237532736553\") = %t; want false", res)
	}

	wrongCheckDigitIsbn13 := "9787532736557"
	if res := Validate(wrongCheckDigitIsbn13); res {
		t.Errorf("Validate(\"9787532736557\") = %t; want false", res)
	}

	validIsbn10 := "7532736555"
	if res := Validate(validIsbn10); !res {
		t.Errorf("Validate(\"7532736555\") = %t; want true", res)
	}

	validIsbn10EndX := "043942089X"
	if res := Validate(validIsbn10EndX); !res {
		t.Errorf("Validate(\"043942089X\") = %t; want true", res)
	}

	wrongCheckDigitIsbn10 := "7532736559"
	if res := Validate(wrongCheckDigitIsbn10); res {
		t.Errorf("Validate(\"7532736559\") = %t; want false", res)
	}
}

func TestConvertToIsbn13(t *testing.T) {
	wrongLength := "123456789"
	if _, err := ConvertToIsbn13(wrongLength); err == nil {
		t.Errorf("Wrong length error should be thrown")
	}

	wrongIsbn10 := "7532736559"
	if _, err := ConvertToIsbn13(wrongIsbn10); err == nil {
		t.Errorf("Invalid isbn error should be thrown")
	}

	validIsbn10 := "7532736555"
	expectedIsbn13 := "9787532736553"
	if res, _ := ConvertToIsbn13(validIsbn10); res != expectedIsbn13 {
		t.Errorf("Expected ISBN-13: %s, got: %s", expectedIsbn13, res)
	}

	wrongTypeDigit := "helloworld"
	if _, err := ConvertToIsbn13(wrongTypeDigit); err == nil {
		t.Errorf("Invalid isbn error should be thrown")
	}
}

func TestConvertToIsbn10(t *testing.T) {
	wrongLength := "123456"
	if _, err := ConvertToIsbn10(wrongLength); err == nil {
		t.Errorf("Wrong length error should be thrown")
	}

	wrongIsbn13 := "9787532736552"
	if _, err := ConvertToIsbn10(wrongIsbn13); err == nil {
		t.Errorf("Invalid isbn error should be thrown")
	}

	wrongTypeDigit := "helloworld123"
	if _, err := ConvertToIsbn10(wrongTypeDigit); err == nil {
		t.Errorf("Invalid isbn error should be thrown")
	}

	validIsbn13 := "9787532736553"
	expectedIsbn10 := "7532736555"
	if res, _ := ConvertToIsbn10(validIsbn13); res != expectedIsbn10 {
		t.Errorf("Expected ISBN-10: %s, got: %s", expectedIsbn10, res)
	}

	wrongPrefixIsbn13 := "1237532736553"
	if _, err := ConvertToIsbn10(wrongPrefixIsbn13); err == nil {
		t.Errorf("Wrong prefix error should be thrown")
	}
}
