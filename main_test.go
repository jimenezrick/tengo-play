package main

import (
	"testing"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

func prepareExpr(code string) *vm.Program {
	program, err := expr.Compile(code)
	if err != nil {
		panic(err)
	}

	_, err = expr.Run(program, nil)
	if err != nil {
		panic(err)
	}
	// fmt.Println(output)

	return program
}

func prepareTengo(code string, vars map[string]interface{}) *tengo.Compiled {
	scr := tengo.NewScript([]byte(code))
	scr.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))

	for name, value := range vars {
		scr.Add(name, value)
	}

	program, err := scr.Run()
	if err != nil {
		panic(err)
	}

	return program
}

func BenchmarkExpr_true(b *testing.B) {
	code := `true`
	program := prepareExpr(code)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Run(program, nil)
	}
}

func BenchmarkExpr_50_array_lookup(b *testing.B) {
	code := `1001 in [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50]`
	program := prepareExpr(code)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Run(program, nil)
	}
}

func BenchmarkExpr_50_array_build(b *testing.B) {
	code := `[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50]`
	program := prepareExpr(code)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Run(program, nil)
	}
}

func BenchmarkExpr_500_array_lookup(b *testing.B) {
	code := `1001 in [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,271,272,273,274,275,276,277,278,279,280,281,282,283,284,285,286,287,288,289,290,291,292,293,294,295,296,297,298,299,300,301,302,303,304,305,306,307,308,309,310,311,312,313,314,315,316,317,318,319,320,321,322,323,324,325,326,327,328,329,330,331,332,333,334,335,336,337,338,339,340,341,342,343,344,345,346,347,348,349,350,351,352,353,354,355,356,357,358,359,360,361,362,363,364,365,366,367,368,369,370,371,372,373,374,375,376,377,378,379,380,381,382,383,384,385,386,387,388,389,390,391,392,393,394,395,396,397,398,399,400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,421,422,423,424,425,426,427,428,429,430,431,432,433,434,435,436,437,438,439,440,441,442,443,444,445,446,447,448,449,450,451,452,453,454,455,456,457,458,459,460,461,462,463,464,465,466,467,468,469,470,471,472,473,474,475,476,477,478,479,480,481,482,483,484,485,486,487,488,489,490,491,492,493,494,495,496,497,498,499,500]`
	program := prepareExpr(code)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Run(program, nil)
	}
}

func BenchmarkExpr_500_array_build(b *testing.B) {
	code := `[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,271,272,273,274,275,276,277,278,279,280,281,282,283,284,285,286,287,288,289,290,291,292,293,294,295,296,297,298,299,300,301,302,303,304,305,306,307,308,309,310,311,312,313,314,315,316,317,318,319,320,321,322,323,324,325,326,327,328,329,330,331,332,333,334,335,336,337,338,339,340,341,342,343,344,345,346,347,348,349,350,351,352,353,354,355,356,357,358,359,360,361,362,363,364,365,366,367,368,369,370,371,372,373,374,375,376,377,378,379,380,381,382,383,384,385,386,387,388,389,390,391,392,393,394,395,396,397,398,399,400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,421,422,423,424,425,426,427,428,429,430,431,432,433,434,435,436,437,438,439,440,441,442,443,444,445,446,447,448,449,450,451,452,453,454,455,456,457,458,459,460,461,462,463,464,465,466,467,468,469,470,471,472,473,474,475,476,477,478,479,480,481,482,483,484,485,486,487,488,489,490,491,492,493,494,495,496,497,498,499,500]`
	program := prepareExpr(code)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Run(program, nil)
	}
}

func BenchmarkTengo_true(b *testing.B) {
	code := `true`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_true_clone(b *testing.B) {
	m := map[string]interface{}{"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true, "11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true, "20": true, "21": true, "22": true, "23": true, "24": true, "25": true, "26": true, "27": true, "28": true, "29": true, "30": true, "31": true, "32": true, "33": true, "34": true, "35": true, "36": true, "37": true, "38": true, "39": true, "40": true, "41": true, "42": true, "43": true, "44": true, "45": true, "46": true, "47": true, "48": true, "49": true, "50": true}
	code := `true`
	program := prepareTengo(code, map[string]interface{}{"m": m})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Clone().Run()
	}
}

func BenchmarkTengo_50_array_lookup(b *testing.B) {
	code := `
fmt := import("fmt")
enum := import("enum")

l := [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50]
r := enum.any(l, func(k, v) { return v == 1001})
`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_50_array_build(b *testing.B) {
	code := `
fmt := import("fmt")
enum := import("enum")

l := [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50]
`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_500_array_lookup(b *testing.B) {
	code := `
fmt := import("fmt")
enum := import("enum")

l := [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,271,272,273,274,275,276,277,278,279,280,281,282,283,284,285,286,287,288,289,290,291,292,293,294,295,296,297,298,299,300,301,302,303,304,305,306,307,308,309,310,311,312,313,314,315,316,317,318,319,320,321,322,323,324,325,326,327,328,329,330,331,332,333,334,335,336,337,338,339,340,341,342,343,344,345,346,347,348,349,350,351,352,353,354,355,356,357,358,359,360,361,362,363,364,365,366,367,368,369,370,371,372,373,374,375,376,377,378,379,380,381,382,383,384,385,386,387,388,389,390,391,392,393,394,395,396,397,398,399,400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,421,422,423,424,425,426,427,428,429,430,431,432,433,434,435,436,437,438,439,440,441,442,443,444,445,446,447,448,449,450,451,452,453,454,455,456,457,458,459,460,461,462,463,464,465,466,467,468,469,470,471,472,473,474,475,476,477,478,479,480,481,482,483,484,485,486,487,488,489,490,491,492,493,494,495,496,497,498,499,500]
r := enum.any(l, func(k, v) { return v == 1001})
`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_500_array_build(b *testing.B) {
	code := `
fmt := import("fmt")
enum := import("enum")

l := [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,226,227,228,229,230,231,232,233,234,235,236,237,238,239,240,241,242,243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,271,272,273,274,275,276,277,278,279,280,281,282,283,284,285,286,287,288,289,290,291,292,293,294,295,296,297,298,299,300,301,302,303,304,305,306,307,308,309,310,311,312,313,314,315,316,317,318,319,320,321,322,323,324,325,326,327,328,329,330,331,332,333,334,335,336,337,338,339,340,341,342,343,344,345,346,347,348,349,350,351,352,353,354,355,356,357,358,359,360,361,362,363,364,365,366,367,368,369,370,371,372,373,374,375,376,377,378,379,380,381,382,383,384,385,386,387,388,389,390,391,392,393,394,395,396,397,398,399,400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,421,422,423,424,425,426,427,428,429,430,431,432,433,434,435,436,437,438,439,440,441,442,443,444,445,446,447,448,449,450,451,452,453,454,455,456,457,458,459,460,461,462,463,464,465,466,467,468,469,470,471,472,473,474,475,476,477,478,479,480,481,482,483,484,485,486,487,488,489,490,491,492,493,494,495,496,497,498,499,500]
`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_500_map_lookup(b *testing.B) {
	code := `
fmt := import("fmt")
enum := import("enum")

m := {"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true, "11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true, "20": true, "21": true, "22": true, "23": true, "24": true, "25": true, "26": true, "27": true, "28": true, "29": true, "30": true, "31": true, "32": true, "33": true, "34": true, "35": true, "36": true, "37": true, "38": true, "39": true, "40": true, "41": true, "42": true, "43": true, "44": true, "45": true, "46": true, "47": true, "48": true, "49": true, "50": true, "51": true, "52": true, "53": true, "54": true, "55": true, "56": true, "57": true, "58": true, "59": true, "60": true, "61": true, "62": true, "63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true, "70": true, "71": true, "72": true, "73": true, "74": true, "75": true, "76": true, "77": true, "78": true, "79": true, "80": true, "81": true, "82": true, "83": true, "84": true, "85": true, "86": true, "87": true, "88": true, "89": true, "90": true, "91": true, "92": true, "93": true, "94": true, "95": true, "96": true, "97": true, "98": true, "99": true, "100": true, "101": true, "102": true, "103": true, "104": true, "105": true, "106": true, "107": true, "108": true, "109": true, "110": true, "111": true, "112": true, "113": true, "114": true, "115": true, "116": true, "117": true, "118": true, "119": true, "120": true, "121": true, "122": true, "123": true, "124": true, "125": true, "126": true, "127": true, "128": true, "129": true, "130": true, "131": true, "132": true, "133": true, "134": true, "135": true, "136": true, "137": true, "138": true, "139": true, "140": true, "141": true, "142": true, "143": true, "144": true, "145": true, "146": true, "147": true, "148": true, "149": true, "150": true, "151": true, "152": true, "153": true, "154": true, "155": true, "156": true, "157": true, "158": true, "159": true, "160": true, "161": true, "162": true, "163": true, "164": true, "165": true, "166": true, "167": true, "168": true, "169": true, "170": true, "171": true, "172": true, "173": true, "174": true, "175": true, "176": true, "177": true, "178": true, "179": true, "180": true, "181": true, "182": true, "183": true, "184": true, "185": true, "186": true, "187": true, "188": true, "189": true, "190": true, "191": true, "192": true, "193": true, "194": true, "195": true, "196": true, "197": true, "198": true, "199": true, "200": true, "201": true, "202": true, "203": true, "204": true, "205": true, "206": true, "207": true, "208": true, "209": true, "210": true, "211": true, "212": true, "213": true, "214": true, "215": true, "216": true, "217": true, "218": true, "219": true, "220": true, "221": true, "222": true, "223": true, "224": true, "225": true, "226": true, "227": true, "228": true, "229": true, "230": true, "231": true, "232": true, "233": true, "234": true, "235": true, "236": true, "237": true, "238": true, "239": true, "240": true, "241": true, "242": true, "243": true, "244": true, "245": true, "246": true, "247": true, "248": true, "249": true, "250": true, "251": true, "252": true, "253": true, "254": true, "255": true, "256": true, "257": true, "258": true, "259": true, "260": true, "261": true, "262": true, "263": true, "264": true, "265": true, "266": true, "267": true, "268": true, "269": true, "270": true, "271": true, "272": true, "273": true, "274": true, "275": true, "276": true, "277": true, "278": true, "279": true, "280": true, "281": true, "282": true, "283": true, "284": true, "285": true, "286": true, "287": true, "288": true, "289": true, "290": true, "291": true, "292": true, "293": true, "294": true, "295": true, "296": true, "297": true, "298": true, "299": true, "300": true, "301": true, "302": true, "303": true, "304": true, "305": true, "306": true, "307": true, "308": true, "309": true, "310": true, "311": true, "312": true, "313": true, "314": true, "315": true, "316": true, "317": true, "318": true, "319": true, "320": true, "321": true, "322": true, "323": true, "324": true, "325": true, "326": true, "327": true, "328": true, "329": true, "330": true, "331": true, "332": true, "333": true, "334": true, "335": true, "336": true, "337": true, "338": true, "339": true, "340": true, "341": true, "342": true, "343": true, "344": true, "345": true, "346": true, "347": true, "348": true, "349": true, "350": true, "351": true, "352": true, "353": true, "354": true, "355": true, "356": true, "357": true, "358": true, "359": true, "360": true, "361": true, "362": true, "363": true, "364": true, "365": true, "366": true, "367": true, "368": true, "369": true, "370": true, "371": true, "372": true, "373": true, "374": true, "375": true, "376": true, "377": true, "378": true, "379": true, "380": true, "381": true, "382": true, "383": true, "384": true, "385": true, "386": true, "387": true, "388": true, "389": true, "390": true, "391": true, "392": true, "393": true, "394": true, "395": true, "396": true, "397": true, "398": true, "399": true, "400": true, "401": true, "402": true, "403": true, "404": true, "405": true, "406": true, "407": true, "408": true, "409": true, "410": true, "411": true, "412": true, "413": true, "414": true, "415": true, "416": true, "417": true, "418": true, "419": true, "420": true, "421": true, "422": true, "423": true, "424": true, "425": true, "426": true, "427": true, "428": true, "429": true, "430": true, "431": true, "432": true, "433": true, "434": true, "435": true, "436": true, "437": true, "438": true, "439": true, "440": true, "441": true, "442": true, "443": true, "444": true, "445": true, "446": true, "447": true, "448": true, "449": true, "450": true, "451": true, "452": true, "453": true, "454": true, "455": true, "456": true, "457": true, "458": true, "459": true, "460": true, "461": true, "462": true, "463": true, "464": true, "465": true, "466": true, "467": true, "468": true, "469": true, "470": true, "471": true, "472": true, "473": true, "474": true, "475": true, "476": true, "477": true, "478": true, "479": true, "480": true, "481": true, "482": true, "483": true, "484": true, "485": true, "486": true, "487": true, "488": true, "489": true, "490": true, "491": true, "492": true, "493": true, "494": true, "495": true, "496": true, "497": true, "498": true, "499": true, "500": true}
r := m["1001"] != undefined
`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_500_map_build(b *testing.B) {
	code := `
fmt := import("fmt")
enum := import("enum")

m := {"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true, "11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true, "20": true, "21": true, "22": true, "23": true, "24": true, "25": true, "26": true, "27": true, "28": true, "29": true, "30": true, "31": true, "32": true, "33": true, "34": true, "35": true, "36": true, "37": true, "38": true, "39": true, "40": true, "41": true, "42": true, "43": true, "44": true, "45": true, "46": true, "47": true, "48": true, "49": true, "50": true, "51": true, "52": true, "53": true, "54": true, "55": true, "56": true, "57": true, "58": true, "59": true, "60": true, "61": true, "62": true, "63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true, "70": true, "71": true, "72": true, "73": true, "74": true, "75": true, "76": true, "77": true, "78": true, "79": true, "80": true, "81": true, "82": true, "83": true, "84": true, "85": true, "86": true, "87": true, "88": true, "89": true, "90": true, "91": true, "92": true, "93": true, "94": true, "95": true, "96": true, "97": true, "98": true, "99": true, "100": true, "101": true, "102": true, "103": true, "104": true, "105": true, "106": true, "107": true, "108": true, "109": true, "110": true, "111": true, "112": true, "113": true, "114": true, "115": true, "116": true, "117": true, "118": true, "119": true, "120": true, "121": true, "122": true, "123": true, "124": true, "125": true, "126": true, "127": true, "128": true, "129": true, "130": true, "131": true, "132": true, "133": true, "134": true, "135": true, "136": true, "137": true, "138": true, "139": true, "140": true, "141": true, "142": true, "143": true, "144": true, "145": true, "146": true, "147": true, "148": true, "149": true, "150": true, "151": true, "152": true, "153": true, "154": true, "155": true, "156": true, "157": true, "158": true, "159": true, "160": true, "161": true, "162": true, "163": true, "164": true, "165": true, "166": true, "167": true, "168": true, "169": true, "170": true, "171": true, "172": true, "173": true, "174": true, "175": true, "176": true, "177": true, "178": true, "179": true, "180": true, "181": true, "182": true, "183": true, "184": true, "185": true, "186": true, "187": true, "188": true, "189": true, "190": true, "191": true, "192": true, "193": true, "194": true, "195": true, "196": true, "197": true, "198": true, "199": true, "200": true, "201": true, "202": true, "203": true, "204": true, "205": true, "206": true, "207": true, "208": true, "209": true, "210": true, "211": true, "212": true, "213": true, "214": true, "215": true, "216": true, "217": true, "218": true, "219": true, "220": true, "221": true, "222": true, "223": true, "224": true, "225": true, "226": true, "227": true, "228": true, "229": true, "230": true, "231": true, "232": true, "233": true, "234": true, "235": true, "236": true, "237": true, "238": true, "239": true, "240": true, "241": true, "242": true, "243": true, "244": true, "245": true, "246": true, "247": true, "248": true, "249": true, "250": true, "251": true, "252": true, "253": true, "254": true, "255": true, "256": true, "257": true, "258": true, "259": true, "260": true, "261": true, "262": true, "263": true, "264": true, "265": true, "266": true, "267": true, "268": true, "269": true, "270": true, "271": true, "272": true, "273": true, "274": true, "275": true, "276": true, "277": true, "278": true, "279": true, "280": true, "281": true, "282": true, "283": true, "284": true, "285": true, "286": true, "287": true, "288": true, "289": true, "290": true, "291": true, "292": true, "293": true, "294": true, "295": true, "296": true, "297": true, "298": true, "299": true, "300": true, "301": true, "302": true, "303": true, "304": true, "305": true, "306": true, "307": true, "308": true, "309": true, "310": true, "311": true, "312": true, "313": true, "314": true, "315": true, "316": true, "317": true, "318": true, "319": true, "320": true, "321": true, "322": true, "323": true, "324": true, "325": true, "326": true, "327": true, "328": true, "329": true, "330": true, "331": true, "332": true, "333": true, "334": true, "335": true, "336": true, "337": true, "338": true, "339": true, "340": true, "341": true, "342": true, "343": true, "344": true, "345": true, "346": true, "347": true, "348": true, "349": true, "350": true, "351": true, "352": true, "353": true, "354": true, "355": true, "356": true, "357": true, "358": true, "359": true, "360": true, "361": true, "362": true, "363": true, "364": true, "365": true, "366": true, "367": true, "368": true, "369": true, "370": true, "371": true, "372": true, "373": true, "374": true, "375": true, "376": true, "377": true, "378": true, "379": true, "380": true, "381": true, "382": true, "383": true, "384": true, "385": true, "386": true, "387": true, "388": true, "389": true, "390": true, "391": true, "392": true, "393": true, "394": true, "395": true, "396": true, "397": true, "398": true, "399": true, "400": true, "401": true, "402": true, "403": true, "404": true, "405": true, "406": true, "407": true, "408": true, "409": true, "410": true, "411": true, "412": true, "413": true, "414": true, "415": true, "416": true, "417": true, "418": true, "419": true, "420": true, "421": true, "422": true, "423": true, "424": true, "425": true, "426": true, "427": true, "428": true, "429": true, "430": true, "431": true, "432": true, "433": true, "434": true, "435": true, "436": true, "437": true, "438": true, "439": true, "440": true, "441": true, "442": true, "443": true, "444": true, "445": true, "446": true, "447": true, "448": true, "449": true, "450": true, "451": true, "452": true, "453": true, "454": true, "455": true, "456": true, "457": true, "458": true, "459": true, "460": true, "461": true, "462": true, "463": true, "464": true, "465": true, "466": true, "467": true, "468": true, "469": true, "470": true, "471": true, "472": true, "473": true, "474": true, "475": true, "476": true, "477": true, "478": true, "479": true, "480": true, "481": true, "482": true, "483": true, "484": true, "485": true, "486": true, "487": true, "488": true, "489": true, "490": true, "491": true, "492": true, "493": true, "494": true, "495": true, "496": true, "497": true, "498": true, "499": true, "500": true}
`
	program := prepareTengo(code, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_50_passed_map_lookup(b *testing.B) {
	m := map[string]interface{}{"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true, "11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true, "20": true, "21": true, "22": true, "23": true, "24": true, "25": true, "26": true, "27": true, "28": true, "29": true, "30": true, "31": true, "32": true, "33": true, "34": true, "35": true, "36": true, "37": true, "38": true, "39": true, "40": true, "41": true, "42": true, "43": true, "44": true, "45": true, "46": true, "47": true, "48": true, "49": true, "50": true}
	code := `
fmt := import("fmt")
enum := import("enum")

r := m["1001"] != undefined
`
	program := prepareTengo(code, map[string]interface{}{"m": m})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}

func BenchmarkTengo_500_passed_map_lookup(b *testing.B) {
	m := map[string]interface{}{"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true, "11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true, "20": true, "21": true, "22": true, "23": true, "24": true, "25": true, "26": true, "27": true, "28": true, "29": true, "30": true, "31": true, "32": true, "33": true, "34": true, "35": true, "36": true, "37": true, "38": true, "39": true, "40": true, "41": true, "42": true, "43": true, "44": true, "45": true, "46": true, "47": true, "48": true, "49": true, "50": true, "51": true, "52": true, "53": true, "54": true, "55": true, "56": true, "57": true, "58": true, "59": true, "60": true, "61": true, "62": true, "63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true, "70": true, "71": true, "72": true, "73": true, "74": true, "75": true, "76": true, "77": true, "78": true, "79": true, "80": true, "81": true, "82": true, "83": true, "84": true, "85": true, "86": true, "87": true, "88": true, "89": true, "90": true, "91": true, "92": true, "93": true, "94": true, "95": true, "96": true, "97": true, "98": true, "99": true, "100": true, "101": true, "102": true, "103": true, "104": true, "105": true, "106": true, "107": true, "108": true, "109": true, "110": true, "111": true, "112": true, "113": true, "114": true, "115": true, "116": true, "117": true, "118": true, "119": true, "120": true, "121": true, "122": true, "123": true, "124": true, "125": true, "126": true, "127": true, "128": true, "129": true, "130": true, "131": true, "132": true, "133": true, "134": true, "135": true, "136": true, "137": true, "138": true, "139": true, "140": true, "141": true, "142": true, "143": true, "144": true, "145": true, "146": true, "147": true, "148": true, "149": true, "150": true, "151": true, "152": true, "153": true, "154": true, "155": true, "156": true, "157": true, "158": true, "159": true, "160": true, "161": true, "162": true, "163": true, "164": true, "165": true, "166": true, "167": true, "168": true, "169": true, "170": true, "171": true, "172": true, "173": true, "174": true, "175": true, "176": true, "177": true, "178": true, "179": true, "180": true, "181": true, "182": true, "183": true, "184": true, "185": true, "186": true, "187": true, "188": true, "189": true, "190": true, "191": true, "192": true, "193": true, "194": true, "195": true, "196": true, "197": true, "198": true, "199": true, "200": true, "201": true, "202": true, "203": true, "204": true, "205": true, "206": true, "207": true, "208": true, "209": true, "210": true, "211": true, "212": true, "213": true, "214": true, "215": true, "216": true, "217": true, "218": true, "219": true, "220": true, "221": true, "222": true, "223": true, "224": true, "225": true, "226": true, "227": true, "228": true, "229": true, "230": true, "231": true, "232": true, "233": true, "234": true, "235": true, "236": true, "237": true, "238": true, "239": true, "240": true, "241": true, "242": true, "243": true, "244": true, "245": true, "246": true, "247": true, "248": true, "249": true, "250": true, "251": true, "252": true, "253": true, "254": true, "255": true, "256": true, "257": true, "258": true, "259": true, "260": true, "261": true, "262": true, "263": true, "264": true, "265": true, "266": true, "267": true, "268": true, "269": true, "270": true, "271": true, "272": true, "273": true, "274": true, "275": true, "276": true, "277": true, "278": true, "279": true, "280": true, "281": true, "282": true, "283": true, "284": true, "285": true, "286": true, "287": true, "288": true, "289": true, "290": true, "291": true, "292": true, "293": true, "294": true, "295": true, "296": true, "297": true, "298": true, "299": true, "300": true, "301": true, "302": true, "303": true, "304": true, "305": true, "306": true, "307": true, "308": true, "309": true, "310": true, "311": true, "312": true, "313": true, "314": true, "315": true, "316": true, "317": true, "318": true, "319": true, "320": true, "321": true, "322": true, "323": true, "324": true, "325": true, "326": true, "327": true, "328": true, "329": true, "330": true, "331": true, "332": true, "333": true, "334": true, "335": true, "336": true, "337": true, "338": true, "339": true, "340": true, "341": true, "342": true, "343": true, "344": true, "345": true, "346": true, "347": true, "348": true, "349": true, "350": true, "351": true, "352": true, "353": true, "354": true, "355": true, "356": true, "357": true, "358": true, "359": true, "360": true, "361": true, "362": true, "363": true, "364": true, "365": true, "366": true, "367": true, "368": true, "369": true, "370": true, "371": true, "372": true, "373": true, "374": true, "375": true, "376": true, "377": true, "378": true, "379": true, "380": true, "381": true, "382": true, "383": true, "384": true, "385": true, "386": true, "387": true, "388": true, "389": true, "390": true, "391": true, "392": true, "393": true, "394": true, "395": true, "396": true, "397": true, "398": true, "399": true, "400": true, "401": true, "402": true, "403": true, "404": true, "405": true, "406": true, "407": true, "408": true, "409": true, "410": true, "411": true, "412": true, "413": true, "414": true, "415": true, "416": true, "417": true, "418": true, "419": true, "420": true, "421": true, "422": true, "423": true, "424": true, "425": true, "426": true, "427": true, "428": true, "429": true, "430": true, "431": true, "432": true, "433": true, "434": true, "435": true, "436": true, "437": true, "438": true, "439": true, "440": true, "441": true, "442": true, "443": true, "444": true, "445": true, "446": true, "447": true, "448": true, "449": true, "450": true, "451": true, "452": true, "453": true, "454": true, "455": true, "456": true, "457": true, "458": true, "459": true, "460": true, "461": true, "462": true, "463": true, "464": true, "465": true, "466": true, "467": true, "468": true, "469": true, "470": true, "471": true, "472": true, "473": true, "474": true, "475": true, "476": true, "477": true, "478": true, "479": true, "480": true, "481": true, "482": true, "483": true, "484": true, "485": true, "486": true, "487": true, "488": true, "489": true, "490": true, "491": true, "492": true, "493": true, "494": true, "495": true, "496": true, "497": true, "498": true, "499": true, "500": true}
	code := `
fmt := import("fmt")
enum := import("enum")

r := m["1001"] != undefined
`
	program := prepareTengo(code, map[string]interface{}{"m": m})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		program.Run()
	}
}
