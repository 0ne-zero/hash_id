package hash_id

// Available formats:
// AlphaNum
// Alpha
// Num
var ALGORITHMS = map[string]map[string]any{
	"MD5": {
		"Length": 32,
		"Format": "AlphaNum",
	},
	"SHA1": {
		"Length": 40,
		"Format": "AlphaNum",
	},
	"SHA224": {
		"Length": 56,
		"Format": "AlphaNum",
	},
	"SHA256": {
		"Length": 64,
		"Format": "AlphaNum",
	},
	"SHA384": {
		"Length": 96,
		"Format": "AlphaNum",
	},
	"SHA512": {
		"Length": 128,
		"Format": "AlphaNum",
	},
	"HAVAL128": {
		"Length": 32,
		"Format": "AlphaNum",
	},
	"HAVAL160": {
		"Length": 40,
		"Format": "AlphaNum",
	},
	"HAVAL192": {
		"Length": 48,
		"Format": "AlphaNum",
	},
	"HAVAL224": {
		"Length": 56,
		"Format": "AlphaNum",
	},
	"HAVAL256": {
		"Length": 64,
		"Format": "AlphaNum",
	},
	"CRC16": {
		"Length": 4,
		"Format": "Num",
	},
	"CRC32": {
		"Length": 8,
		"Format": "AlphaNum",
	},
}
