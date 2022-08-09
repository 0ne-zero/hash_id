package hash_id

import "testing"

// Checks result has given algorithm
func isSliceContainsElement(s []string, element string) bool {
	for i := range s {
		if s[i] == element {
			return true
		}
	}
	return false
}

func TestValidateAlgorithms(t *testing.T) {
	var supported_format = []string{"AlphaNum", "Alpha", "Num"}
	var length int
	var format string
	var ok bool

	for _, v := range ALGORITHMS {
		if length, ok = v["Length"].(int); !ok {
			t.FailNow()
		}
		if length < 1 {
			t.FailNow()
		}
		if format, ok = v["Format"].(string); !ok {
			t.FailNow()
		}
		if !isSliceContainsElement(supported_format, format) {
			t.FailNow()
		}
	}
}

func TestSHA1(t *testing.T) {
	// Should be identify
	var hash = "3f786850e387550fdab836ed7e6dc881de23001b"
	var hash_type = "SHA1"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "23oisd2343sdfsdvb252332df"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestSha224(t *testing.T) {
	// Should be identify
	var hash = "f18276755ba8725bfe0e6bff764621dc562fdaba9d60308327ad55a2"
	var hash_type = "SHA224"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "23oisd2343sdfsdasdfsdvb252332df"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestSHA256(t *testing.T) {
	// Should be identify
	var hash = "87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7"
	var hash_type = "SHA256"
	var res []string

	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "sdfsfadklfjas8923ru8249shdoiuf89sdjvknvb"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestSha384(t *testing.T) {
	// Should be identify
	var hash = "bed4e0f8b9c0ec8bee077c2d5ffea39f5b8858458f2694cbe3bd50e137dedb806c76781c53e7cf25dd074855dbbfe3d4"
	var hash_type = "SHA384"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "23oisd2343sdfsdaaskdlfjsadlkfsjdflsdsdfsdvb2sadfsdfasf52332df"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestSha512(t *testing.T) {
	// Should be identify
	var hash = "45843648ecf9da8e513286f136e3f271e7d6dee4d29b947a50dde8c61f3e197694c13bcdc279ce459839757cd8de19c11b23b33565384a97afcf360483578cd4"
	var hash_type = "SHA512"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "23oisd2343sdfsdasdlkf;sdlbkjnvi7r834haskfklsdfsdvb252332df"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestHAVAL128(t *testing.T) {
	// Should be identify
	var hash = "d30c0ca6d154920b09fd51b8ec5a0900"
	var hash_type = "HAVAL128"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "ssdkfljfkjlsjfl2r894u98bdussdjvklasd9f8u9834"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestHAVAL160(t *testing.T) {
	// Should be identify
	var hash = "d082585ba5b8f89f625f6085290da9174818a421"
	var hash_type = "HAVAL160"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "ssdkfljfkjlsjfl2r894u98bsdfsdfsdfdussdjvklasd9f8u9834"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestHAVAL192(t *testing.T) {
	// Should be identify
	var hash = "7a5a85a93e118268b8216d2379eb1c261c75fb3e603adc5e"
	var hash_type = "HAVAL192"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "ssdkfljfkjlsjfl2r894u98bdusxfsdkl;bposfsacbqqjfsdghsdjvklasd9f8u9834"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestHAVAL224(t *testing.T) {
	// Should be identify
	var hash = "92648e8f114a199ec91b10295f3c1c6e9c53a554da72ab72b868de6a"
	var hash_type = "HAVAL224"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "ssdkfljfkjlsjfl2r894u98bdussdjsadfsdfsdfsadfvklasd9f8u9834"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestHAVAL256(t *testing.T) {
	// Should be identify
	var hash = "185d0ddbac6a5d79c34b9ea3c3db85e6f0ba222ab0f7030529f70f573cfa09a9"
	var hash_type = "HAVAL256"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "ssdkfljfkjlsjfl2r894usadfsdafasdf98bdussdjvklasd9f8u9834"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
func TestMD5(t *testing.T) {
	// Should be identify
	var hash = "401b30e3b8b5d629635a5c613cdb7919"
	var hash_type = "MD5"
	var res []string

	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}

	// Shouldn't be identify
	hash = "dsklfj23r8ous9d8dslksldkfj2klrj98bsadfsdfasfsdu9d8fwrk4n98"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}

func TestCRC16(t *testing.T) {
	// Should be identify
	var hash = "2200"
	var hash_type = "CRC16"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
	// Shouldn't be identify
	hash = "dsklfj23r8ous9d8dslksldkfj2ksdsfsafsdfdaflrj98bu9d8fwrk4n98"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}

func TestCRC32(t *testing.T) {
	// Should be identify
	var hash = "8cdc1683"
	var hash_type = "CRC32"
	var res []string
	res = Identify(hash)
	if !isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
	// Shouldn't be identify
	hash = "dsklfj23r8ous9d8sdfsddslksldkfj2ksdfsdfdaflrj98bu9d8fwrk4n98"
	res = Identify(hash)
	if isSliceContainsElement(res, hash_type) {
		t.FailNow()
	}
}
