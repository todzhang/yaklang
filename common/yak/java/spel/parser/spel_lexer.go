// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package spelparser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SpelLexer struct {
	SpelLexerBase
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var spellexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func spellexerLexerInit() {
	staticData := &spellexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE", "IN_TEMPLATE_STRING",
	}
	staticData.literalNames = []string{
		"", "';'", "", "'++'", "'+'", "'--'", "'-'", "':'", "'.'", "','", "'*'",
		"'/'", "'%'", "'('", "')'", "'['", "']'", "'#'", "'@'", "'^['", "'^'",
		"'!='", "'!['", "'!'", "'=='", "'='", "'&&'", "'&'", "'||'", "'?['",
		"'?:'", "'?.'", "'?'", "'$['", "'>='", "'>'", "'<='", "'<'", "", "",
		"", "'or'", "'and'", "'true'", "'false'", "'new'", "'null'", "'T'",
		"'matches'", "'gt'", "'ge'", "'le'", "'lt'", "'eq'", "'ne'", "", "",
		"", "", "", "", "", "'``'",
	}
	staticData.symbolicNames = []string{
		"", "SEMICOLON", "WS", "INC", "PLUS", "DEC", "MINUS", "COLON", "DOT",
		"COMMA", "STAR", "DIV", "MOD", "LPAREN", "RPAREN", "LSQUARE", "RSQUARE",
		"HASH", "BEAN_REF", "SELECT_FIRST", "POWER", "NE", "PROJECT", "NOT",
		"EQ", "ASSIGN", "SYMBOLIC_AND", "FACTORY_BEAN_REF", "SYMBOLIC_OR", "SELECT",
		"ELVIS", "SAFE_NAVI", "QMARK", "SELECT_LAST", "GE", "GT", "LE", "LT",
		"LCURLY", "RCURLY", "BACKTICK", "OR", "AND", "TRUE", "FALSE", "NEW",
		"NULL", "T", "MATCHES", "GT_KEYWORD", "GE_KEYWORD", "LE_KEYWORD", "LT_KEYWORD",
		"EQ_KEYWORD", "NE_KEYWORD", "IDENTIFIER", "REAL_LITERAL", "INTEGER_LITERAL",
		"STRING_LITERAL", "SINGLE_QUOTED_STRING", "DOUBLE_QUOTED_STRING", "PROPERTY_PLACE_HOLDER",
		"ESCAPED_BACKTICK", "SPEL_IN_TEMPLATE_STRING_OPEN", "TEMPLATE_TEXT",
	}
	staticData.ruleNames = []string{
		"SEMICOLON", "NEWLINE", "WS", "INC", "PLUS", "DEC", "MINUS", "COLON",
		"DOT", "COMMA", "STAR", "DIV", "MOD", "LPAREN", "RPAREN", "LSQUARE",
		"RSQUARE", "HASH", "BEAN_REF", "SELECT_FIRST", "POWER", "NE", "PROJECT",
		"NOT", "EQ", "ASSIGN", "SYMBOLIC_AND", "FACTORY_BEAN_REF", "SYMBOLIC_OR",
		"SELECT", "ELVIS", "SAFE_NAVI", "QMARK", "SELECT_LAST", "GE", "GT",
		"LE", "LT", "LCURLY", "RCURLY", "BACKTICK", "OR", "AND", "TRUE", "FALSE",
		"NEW", "NULL", "T", "MATCHES", "GT_KEYWORD", "GE_KEYWORD", "LE_KEYWORD",
		"LT_KEYWORD", "EQ_KEYWORD", "NE_KEYWORD", "IDENTIFIER", "REAL_LITERAL",
		"INTEGER_LITERAL", "STRING_LITERAL", "SINGLE_QUOTED_STRING", "DOUBLE_QUOTED_STRING",
		"PROPERTY_PLACE_HOLDER", "INTEGER_TYPE_SUFFIX", "HEX_DIGIT", "DECIMAL_DIGIT",
		"EXPONENT_PART", "SIGN", "REAL_TYPE_SUFFIX", "ALPHABETIC", "DIGIT",
		"ESCAPED_BACKTICK", "SPEL_IN_TEMPLATE_STRING_OPEN", "TEMPLATE_TEXT",
		"BACKTICK_IN_TEMPLATE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 64, 479, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7,
		14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19,
		2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2,
		25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30,
		7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7,
		35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40,
		2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2,
		46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51,
		7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7,
		56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61,
		2, 62, 7, 62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2,
		67, 7, 67, 2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72,
		7, 72, 2, 73, 7, 73, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 156, 8, 1, 1,
		2, 4, 2, 159, 8, 2, 11, 2, 12, 2, 160, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1,
		51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 54, 1, 55, 1, 55, 3, 55, 320, 8, 55, 1, 55, 1, 55, 1, 55, 5, 55, 325,
		8, 55, 10, 55, 12, 55, 328, 9, 55, 1, 56, 1, 56, 4, 56, 332, 8, 56, 11,
		56, 12, 56, 333, 1, 56, 3, 56, 337, 8, 56, 1, 56, 3, 56, 340, 8, 56, 1,
		56, 4, 56, 343, 8, 56, 11, 56, 12, 56, 344, 1, 56, 1, 56, 4, 56, 349, 8,
		56, 11, 56, 12, 56, 350, 1, 56, 3, 56, 354, 8, 56, 1, 56, 3, 56, 357, 8,
		56, 1, 56, 4, 56, 360, 8, 56, 11, 56, 12, 56, 361, 1, 56, 1, 56, 3, 56,
		366, 8, 56, 1, 56, 4, 56, 369, 8, 56, 11, 56, 12, 56, 370, 1, 56, 1, 56,
		3, 56, 375, 8, 56, 1, 57, 4, 57, 378, 8, 57, 11, 57, 12, 57, 379, 1, 57,
		3, 57, 383, 8, 57, 1, 58, 1, 58, 3, 58, 387, 8, 58, 1, 59, 1, 59, 1, 59,
		1, 59, 5, 59, 393, 8, 59, 10, 59, 12, 59, 396, 9, 59, 1, 59, 1, 59, 1,
		60, 1, 60, 1, 60, 1, 60, 5, 60, 404, 8, 60, 10, 60, 12, 60, 407, 9, 60,
		1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 5, 61, 415, 8, 61, 10, 61, 12,
		61, 418, 9, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64,
		1, 65, 1, 65, 5, 65, 430, 8, 65, 10, 65, 12, 65, 433, 9, 65, 1, 65, 4,
		65, 436, 8, 65, 11, 65, 12, 65, 437, 1, 65, 1, 65, 5, 65, 442, 8, 65, 10,
		65, 12, 65, 445, 9, 65, 1, 65, 4, 65, 448, 8, 65, 11, 65, 12, 65, 449,
		3, 65, 452, 8, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 3, 68, 459, 8, 68,
		1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1,
		71, 1, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 416, 0,
		74, 2, 1, 4, 0, 6, 2, 8, 3, 10, 4, 12, 5, 14, 6, 16, 7, 18, 8, 20, 9, 22,
		10, 24, 11, 26, 12, 28, 13, 30, 14, 32, 15, 34, 16, 36, 17, 38, 18, 40,
		19, 42, 20, 44, 21, 46, 22, 48, 23, 50, 24, 52, 25, 54, 26, 56, 27, 58,
		28, 60, 29, 62, 30, 64, 31, 66, 32, 68, 33, 70, 34, 72, 35, 74, 36, 76,
		37, 78, 38, 80, 39, 82, 40, 84, 41, 86, 42, 88, 43, 90, 44, 92, 45, 94,
		46, 96, 47, 98, 48, 100, 49, 102, 50, 104, 51, 106, 52, 108, 53, 110, 54,
		112, 55, 114, 56, 116, 57, 118, 58, 120, 59, 122, 60, 124, 61, 126, 0,
		128, 0, 130, 0, 132, 0, 134, 0, 136, 0, 138, 0, 140, 0, 142, 62, 144, 63,
		146, 64, 148, 0, 2, 0, 1, 12, 3, 0, 10, 10, 13, 13, 8232, 8233, 3, 0, 9,
		10, 13, 13, 32, 32, 2, 0, 36, 36, 95, 95, 2, 0, 10, 10, 39, 39, 2, 0, 10,
		10, 34, 34, 2, 0, 76, 76, 108, 108, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0,
		48, 57, 2, 0, 43, 43, 45, 45, 4, 0, 68, 68, 70, 70, 100, 100, 102, 102,
		648, 0, 65, 90, 97, 122, 170, 170, 181, 181, 186, 186, 192, 214, 216, 246,
		248, 705, 710, 721, 736, 740, 748, 748, 750, 750, 880, 884, 886, 887, 890,
		893, 895, 895, 902, 902, 904, 906, 908, 908, 910, 929, 931, 1013, 1015,
		1153, 1162, 1327, 1329, 1366, 1369, 1369, 1376, 1416, 1488, 1514, 1519,
		1522, 1568, 1610, 1646, 1647, 1649, 1747, 1749, 1749, 1765, 1766, 1774,
		1775, 1786, 1788, 1791, 1791, 1808, 1808, 1810, 1839, 1869, 1957, 1969,
		1969, 1994, 2026, 2036, 2037, 2042, 2042, 2048, 2069, 2074, 2074, 2084,
		2084, 2088, 2088, 2112, 2136, 2144, 2154, 2160, 2183, 2185, 2190, 2208,
		2249, 2308, 2361, 2365, 2365, 2384, 2384, 2392, 2401, 2417, 2432, 2437,
		2444, 2447, 2448, 2451, 2472, 2474, 2480, 2482, 2482, 2486, 2489, 2493,
		2493, 2510, 2510, 2524, 2525, 2527, 2529, 2544, 2545, 2556, 2556, 2565,
		2570, 2575, 2576, 2579, 2600, 2602, 2608, 2610, 2611, 2613, 2614, 2616,
		2617, 2649, 2652, 2654, 2654, 2674, 2676, 2693, 2701, 2703, 2705, 2707,
		2728, 2730, 2736, 2738, 2739, 2741, 2745, 2749, 2749, 2768, 2768, 2784,
		2785, 2809, 2809, 2821, 2828, 2831, 2832, 2835, 2856, 2858, 2864, 2866,
		2867, 2869, 2873, 2877, 2877, 2908, 2909, 2911, 2913, 2929, 2929, 2947,
		2947, 2949, 2954, 2958, 2960, 2962, 2965, 2969, 2970, 2972, 2972, 2974,
		2975, 2979, 2980, 2984, 2986, 2990, 3001, 3024, 3024, 3077, 3084, 3086,
		3088, 3090, 3112, 3114, 3129, 3133, 3133, 3160, 3162, 3165, 3165, 3168,
		3169, 3200, 3200, 3205, 3212, 3214, 3216, 3218, 3240, 3242, 3251, 3253,
		3257, 3261, 3261, 3293, 3294, 3296, 3297, 3313, 3314, 3332, 3340, 3342,
		3344, 3346, 3386, 3389, 3389, 3406, 3406, 3412, 3414, 3423, 3425, 3450,
		3455, 3461, 3478, 3482, 3505, 3507, 3515, 3517, 3517, 3520, 3526, 3585,
		3632, 3634, 3635, 3648, 3654, 3713, 3714, 3716, 3716, 3718, 3722, 3724,
		3747, 3749, 3749, 3751, 3760, 3762, 3763, 3773, 3773, 3776, 3780, 3782,
		3782, 3804, 3807, 3840, 3840, 3904, 3911, 3913, 3948, 3976, 3980, 4096,
		4138, 4159, 4159, 4176, 4181, 4186, 4189, 4193, 4193, 4197, 4198, 4206,
		4208, 4213, 4225, 4238, 4238, 4256, 4293, 4295, 4295, 4301, 4301, 4304,
		4346, 4348, 4680, 4682, 4685, 4688, 4694, 4696, 4696, 4698, 4701, 4704,
		4744, 4746, 4749, 4752, 4784, 4786, 4789, 4792, 4798, 4800, 4800, 4802,
		4805, 4808, 4822, 4824, 4880, 4882, 4885, 4888, 4954, 4992, 5007, 5024,
		5109, 5112, 5117, 5121, 5740, 5743, 5759, 5761, 5786, 5792, 5866, 5873,
		5880, 5888, 5905, 5919, 5937, 5952, 5969, 5984, 5996, 5998, 6000, 6016,
		6067, 6103, 6103, 6108, 6108, 6176, 6264, 6272, 6276, 6279, 6312, 6314,
		6314, 6320, 6389, 6400, 6430, 6480, 6509, 6512, 6516, 6528, 6571, 6576,
		6601, 6656, 6678, 6688, 6740, 6823, 6823, 6917, 6963, 6981, 6988, 7043,
		7072, 7086, 7087, 7098, 7141, 7168, 7203, 7245, 7247, 7258, 7293, 7296,
		7304, 7312, 7354, 7357, 7359, 7401, 7404, 7406, 7411, 7413, 7414, 7418,
		7418, 7424, 7615, 7680, 7957, 7960, 7965, 7968, 8005, 8008, 8013, 8016,
		8023, 8025, 8025, 8027, 8027, 8029, 8029, 8031, 8061, 8064, 8116, 8118,
		8124, 8126, 8126, 8130, 8132, 8134, 8140, 8144, 8147, 8150, 8155, 8160,
		8172, 8178, 8180, 8182, 8188, 8305, 8305, 8319, 8319, 8336, 8348, 8450,
		8450, 8455, 8455, 8458, 8467, 8469, 8469, 8473, 8477, 8484, 8484, 8486,
		8486, 8488, 8488, 8490, 8493, 8495, 8505, 8508, 8511, 8517, 8521, 8526,
		8526, 8579, 8580, 11264, 11492, 11499, 11502, 11506, 11507, 11520, 11557,
		11559, 11559, 11565, 11565, 11568, 11623, 11631, 11631, 11648, 11670, 11680,
		11686, 11688, 11694, 11696, 11702, 11704, 11710, 11712, 11718, 11720, 11726,
		11728, 11734, 11736, 11742, 11823, 11823, 12293, 12294, 12337, 12341, 12347,
		12348, 12353, 12438, 12445, 12447, 12449, 12538, 12540, 12543, 12549, 12591,
		12593, 12686, 12704, 12735, 12784, 12799, 13312, 19903, 19968, 42124, 42192,
		42237, 42240, 42508, 42512, 42527, 42538, 42539, 42560, 42606, 42623, 42653,
		42656, 42725, 42775, 42783, 42786, 42888, 42891, 42954, 42960, 42961, 42963,
		42963, 42965, 42969, 42994, 43009, 43011, 43013, 43015, 43018, 43020, 43042,
		43072, 43123, 43138, 43187, 43250, 43255, 43259, 43259, 43261, 43262, 43274,
		43301, 43312, 43334, 43360, 43388, 43396, 43442, 43471, 43471, 43488, 43492,
		43494, 43503, 43514, 43518, 43520, 43560, 43584, 43586, 43588, 43595, 43616,
		43638, 43642, 43642, 43646, 43695, 43697, 43697, 43701, 43702, 43705, 43709,
		43712, 43712, 43714, 43714, 43739, 43741, 43744, 43754, 43762, 43764, 43777,
		43782, 43785, 43790, 43793, 43798, 43808, 43814, 43816, 43822, 43824, 43866,
		43868, 43881, 43888, 44002, 44032, 55203, 55216, 55238, 55243, 55291, 63744,
		64109, 64112, 64217, 64256, 64262, 64275, 64279, 64285, 64285, 64287, 64296,
		64298, 64310, 64312, 64316, 64318, 64318, 64320, 64321, 64323, 64324, 64326,
		64433, 64467, 64829, 64848, 64911, 64914, 64967, 65008, 65019, 65136, 65140,
		65142, 65276, 65313, 65338, 65345, 65370, 65382, 65470, 65474, 65479, 65482,
		65487, 65490, 65495, 65498, 65500, 65536, 65547, 65549, 65574, 65576, 65594,
		65596, 65597, 65599, 65613, 65616, 65629, 65664, 65786, 66176, 66204, 66208,
		66256, 66304, 66335, 66349, 66368, 66370, 66377, 66384, 66421, 66432, 66461,
		66464, 66499, 66504, 66511, 66560, 66717, 66736, 66771, 66776, 66811, 66816,
		66855, 66864, 66915, 66928, 66938, 66940, 66954, 66956, 66962, 66964, 66965,
		66967, 66977, 66979, 66993, 66995, 67001, 67003, 67004, 67072, 67382, 67392,
		67413, 67424, 67431, 67456, 67461, 67463, 67504, 67506, 67514, 67584, 67589,
		67592, 67592, 67594, 67637, 67639, 67640, 67644, 67644, 67647, 67669, 67680,
		67702, 67712, 67742, 67808, 67826, 67828, 67829, 67840, 67861, 67872, 67897,
		67968, 68023, 68030, 68031, 68096, 68096, 68112, 68115, 68117, 68119, 68121,
		68149, 68192, 68220, 68224, 68252, 68288, 68295, 68297, 68324, 68352, 68405,
		68416, 68437, 68448, 68466, 68480, 68497, 68608, 68680, 68736, 68786, 68800,
		68850, 68864, 68899, 69248, 69289, 69296, 69297, 69376, 69404, 69415, 69415,
		69424, 69445, 69488, 69505, 69552, 69572, 69600, 69622, 69635, 69687, 69745,
		69746, 69749, 69749, 69763, 69807, 69840, 69864, 69891, 69926, 69956, 69956,
		69959, 69959, 69968, 70002, 70006, 70006, 70019, 70066, 70081, 70084, 70106,
		70106, 70108, 70108, 70144, 70161, 70163, 70187, 70272, 70278, 70280, 70280,
		70282, 70285, 70287, 70301, 70303, 70312, 70320, 70366, 70405, 70412, 70415,
		70416, 70419, 70440, 70442, 70448, 70450, 70451, 70453, 70457, 70461, 70461,
		70480, 70480, 70493, 70497, 70656, 70708, 70727, 70730, 70751, 70753, 70784,
		70831, 70852, 70853, 70855, 70855, 71040, 71086, 71128, 71131, 71168, 71215,
		71236, 71236, 71296, 71338, 71352, 71352, 71424, 71450, 71488, 71494, 71680,
		71723, 71840, 71903, 71935, 71942, 71945, 71945, 71948, 71955, 71957, 71958,
		71960, 71983, 71999, 71999, 72001, 72001, 72096, 72103, 72106, 72144, 72161,
		72161, 72163, 72163, 72192, 72192, 72203, 72242, 72250, 72250, 72272, 72272,
		72284, 72329, 72349, 72349, 72368, 72440, 72704, 72712, 72714, 72750, 72768,
		72768, 72818, 72847, 72960, 72966, 72968, 72969, 72971, 73008, 73030, 73030,
		73056, 73061, 73063, 73064, 73066, 73097, 73112, 73112, 73440, 73458, 73648,
		73648, 73728, 74649, 74880, 75075, 77712, 77808, 77824, 78894, 82944, 83526,
		92160, 92728, 92736, 92766, 92784, 92862, 92880, 92909, 92928, 92975, 92992,
		92995, 93027, 93047, 93053, 93071, 93760, 93823, 93952, 94026, 94032, 94032,
		94099, 94111, 94176, 94177, 94179, 94179, 94208, 100343, 100352, 101589,
		101632, 101640, 110576, 110579, 110581, 110587, 110589, 110590, 110592,
		110882, 110928, 110930, 110948, 110951, 110960, 111355, 113664, 113770,
		113776, 113788, 113792, 113800, 113808, 113817, 119808, 119892, 119894,
		119964, 119966, 119967, 119970, 119970, 119973, 119974, 119977, 119980,
		119982, 119993, 119995, 119995, 119997, 120003, 120005, 120069, 120071,
		120074, 120077, 120084, 120086, 120092, 120094, 120121, 120123, 120126,
		120128, 120132, 120134, 120134, 120138, 120144, 120146, 120485, 120488,
		120512, 120514, 120538, 120540, 120570, 120572, 120596, 120598, 120628,
		120630, 120654, 120656, 120686, 120688, 120712, 120714, 120744, 120746,
		120770, 120772, 120779, 122624, 122654, 123136, 123180, 123191, 123197,
		123214, 123214, 123536, 123565, 123584, 123627, 124896, 124902, 124904,
		124907, 124909, 124910, 124912, 124926, 124928, 125124, 125184, 125251,
		125259, 125259, 126464, 126467, 126469, 126495, 126497, 126498, 126500,
		126500, 126503, 126503, 126505, 126514, 126516, 126519, 126521, 126521,
		126523, 126523, 126530, 126530, 126535, 126535, 126537, 126537, 126539,
		126539, 126541, 126543, 126545, 126546, 126548, 126548, 126551, 126551,
		126553, 126553, 126555, 126555, 126557, 126557, 126559, 126559, 126561,
		126562, 126564, 126564, 126567, 126570, 126572, 126578, 126580, 126583,
		126585, 126588, 126590, 126590, 126592, 126601, 126603, 126619, 126625,
		126627, 126629, 126633, 126635, 126651, 131072, 173791, 173824, 177976,
		177984, 178205, 178208, 183969, 183984, 191456, 194560, 195101, 196608,
		201546, 2, 0, 10, 10, 96, 96, 500, 0, 2, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0,
		0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0,
		0, 0, 16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0,
		0, 0, 0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1,
		0, 0, 0, 0, 32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38,
		1, 0, 0, 0, 0, 40, 1, 0, 0, 0, 0, 42, 1, 0, 0, 0, 0, 44, 1, 0, 0, 0, 0,
		46, 1, 0, 0, 0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0,
		0, 54, 1, 0, 0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0, 0, 60, 1, 0, 0,
		0, 0, 62, 1, 0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0, 0, 0, 68, 1, 0,
		0, 0, 0, 70, 1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0, 0, 76, 1,
		0, 0, 0, 0, 78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1, 0, 0, 0, 0, 84,
		1, 0, 0, 0, 0, 86, 1, 0, 0, 0, 0, 88, 1, 0, 0, 0, 0, 90, 1, 0, 0, 0, 0,
		92, 1, 0, 0, 0, 0, 94, 1, 0, 0, 0, 0, 96, 1, 0, 0, 0, 0, 98, 1, 0, 0, 0,
		0, 100, 1, 0, 0, 0, 0, 102, 1, 0, 0, 0, 0, 104, 1, 0, 0, 0, 0, 106, 1,
		0, 0, 0, 0, 108, 1, 0, 0, 0, 0, 110, 1, 0, 0, 0, 0, 112, 1, 0, 0, 0, 0,
		114, 1, 0, 0, 0, 0, 116, 1, 0, 0, 0, 0, 118, 1, 0, 0, 0, 0, 120, 1, 0,
		0, 0, 0, 122, 1, 0, 0, 0, 0, 124, 1, 0, 0, 0, 1, 142, 1, 0, 0, 0, 1, 144,
		1, 0, 0, 0, 1, 146, 1, 0, 0, 0, 1, 148, 1, 0, 0, 0, 2, 150, 1, 0, 0, 0,
		4, 155, 1, 0, 0, 0, 6, 158, 1, 0, 0, 0, 8, 164, 1, 0, 0, 0, 10, 167, 1,
		0, 0, 0, 12, 169, 1, 0, 0, 0, 14, 172, 1, 0, 0, 0, 16, 174, 1, 0, 0, 0,
		18, 176, 1, 0, 0, 0, 20, 178, 1, 0, 0, 0, 22, 180, 1, 0, 0, 0, 24, 182,
		1, 0, 0, 0, 26, 184, 1, 0, 0, 0, 28, 186, 1, 0, 0, 0, 30, 188, 1, 0, 0,
		0, 32, 190, 1, 0, 0, 0, 34, 192, 1, 0, 0, 0, 36, 194, 1, 0, 0, 0, 38, 196,
		1, 0, 0, 0, 40, 198, 1, 0, 0, 0, 42, 201, 1, 0, 0, 0, 44, 203, 1, 0, 0,
		0, 46, 206, 1, 0, 0, 0, 48, 209, 1, 0, 0, 0, 50, 211, 1, 0, 0, 0, 52, 214,
		1, 0, 0, 0, 54, 216, 1, 0, 0, 0, 56, 219, 1, 0, 0, 0, 58, 221, 1, 0, 0,
		0, 60, 224, 1, 0, 0, 0, 62, 227, 1, 0, 0, 0, 64, 230, 1, 0, 0, 0, 66, 233,
		1, 0, 0, 0, 68, 235, 1, 0, 0, 0, 70, 238, 1, 0, 0, 0, 72, 241, 1, 0, 0,
		0, 74, 243, 1, 0, 0, 0, 76, 246, 1, 0, 0, 0, 78, 248, 1, 0, 0, 0, 80, 253,
		1, 0, 0, 0, 82, 258, 1, 0, 0, 0, 84, 262, 1, 0, 0, 0, 86, 265, 1, 0, 0,
		0, 88, 269, 1, 0, 0, 0, 90, 274, 1, 0, 0, 0, 92, 280, 1, 0, 0, 0, 94, 284,
		1, 0, 0, 0, 96, 289, 1, 0, 0, 0, 98, 291, 1, 0, 0, 0, 100, 299, 1, 0, 0,
		0, 102, 302, 1, 0, 0, 0, 104, 305, 1, 0, 0, 0, 106, 308, 1, 0, 0, 0, 108,
		311, 1, 0, 0, 0, 110, 314, 1, 0, 0, 0, 112, 319, 1, 0, 0, 0, 114, 374,
		1, 0, 0, 0, 116, 377, 1, 0, 0, 0, 118, 386, 1, 0, 0, 0, 120, 388, 1, 0,
		0, 0, 122, 399, 1, 0, 0, 0, 124, 410, 1, 0, 0, 0, 126, 421, 1, 0, 0, 0,
		128, 423, 1, 0, 0, 0, 130, 425, 1, 0, 0, 0, 132, 451, 1, 0, 0, 0, 134,
		453, 1, 0, 0, 0, 136, 455, 1, 0, 0, 0, 138, 458, 1, 0, 0, 0, 140, 460,
		1, 0, 0, 0, 142, 462, 1, 0, 0, 0, 144, 465, 1, 0, 0, 0, 146, 472, 1, 0,
		0, 0, 148, 474, 1, 0, 0, 0, 150, 151, 5, 59, 0, 0, 151, 3, 1, 0, 0, 0,
		152, 153, 5, 13, 0, 0, 153, 156, 5, 10, 0, 0, 154, 156, 7, 0, 0, 0, 155,
		152, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 5, 1, 0, 0, 0, 157, 159, 7,
		1, 0, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0,
		0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 6, 2, 0, 0, 163,
		7, 1, 0, 0, 0, 164, 165, 5, 43, 0, 0, 165, 166, 5, 43, 0, 0, 166, 9, 1,
		0, 0, 0, 167, 168, 5, 43, 0, 0, 168, 11, 1, 0, 0, 0, 169, 170, 5, 45, 0,
		0, 170, 171, 5, 45, 0, 0, 171, 13, 1, 0, 0, 0, 172, 173, 5, 45, 0, 0, 173,
		15, 1, 0, 0, 0, 174, 175, 5, 58, 0, 0, 175, 17, 1, 0, 0, 0, 176, 177, 5,
		46, 0, 0, 177, 19, 1, 0, 0, 0, 178, 179, 5, 44, 0, 0, 179, 21, 1, 0, 0,
		0, 180, 181, 5, 42, 0, 0, 181, 23, 1, 0, 0, 0, 182, 183, 5, 47, 0, 0, 183,
		25, 1, 0, 0, 0, 184, 185, 5, 37, 0, 0, 185, 27, 1, 0, 0, 0, 186, 187, 5,
		40, 0, 0, 187, 29, 1, 0, 0, 0, 188, 189, 5, 41, 0, 0, 189, 31, 1, 0, 0,
		0, 190, 191, 5, 91, 0, 0, 191, 33, 1, 0, 0, 0, 192, 193, 5, 93, 0, 0, 193,
		35, 1, 0, 0, 0, 194, 195, 5, 35, 0, 0, 195, 37, 1, 0, 0, 0, 196, 197, 5,
		64, 0, 0, 197, 39, 1, 0, 0, 0, 198, 199, 5, 94, 0, 0, 199, 200, 5, 91,
		0, 0, 200, 41, 1, 0, 0, 0, 201, 202, 5, 94, 0, 0, 202, 43, 1, 0, 0, 0,
		203, 204, 5, 33, 0, 0, 204, 205, 5, 61, 0, 0, 205, 45, 1, 0, 0, 0, 206,
		207, 5, 33, 0, 0, 207, 208, 5, 91, 0, 0, 208, 47, 1, 0, 0, 0, 209, 210,
		5, 33, 0, 0, 210, 49, 1, 0, 0, 0, 211, 212, 5, 61, 0, 0, 212, 213, 5, 61,
		0, 0, 213, 51, 1, 0, 0, 0, 214, 215, 5, 61, 0, 0, 215, 53, 1, 0, 0, 0,
		216, 217, 5, 38, 0, 0, 217, 218, 5, 38, 0, 0, 218, 55, 1, 0, 0, 0, 219,
		220, 5, 38, 0, 0, 220, 57, 1, 0, 0, 0, 221, 222, 5, 124, 0, 0, 222, 223,
		5, 124, 0, 0, 223, 59, 1, 0, 0, 0, 224, 225, 5, 63, 0, 0, 225, 226, 5,
		91, 0, 0, 226, 61, 1, 0, 0, 0, 227, 228, 5, 63, 0, 0, 228, 229, 5, 58,
		0, 0, 229, 63, 1, 0, 0, 0, 230, 231, 5, 63, 0, 0, 231, 232, 5, 46, 0, 0,
		232, 65, 1, 0, 0, 0, 233, 234, 5, 63, 0, 0, 234, 67, 1, 0, 0, 0, 235, 236,
		5, 36, 0, 0, 236, 237, 5, 91, 0, 0, 237, 69, 1, 0, 0, 0, 238, 239, 5, 62,
		0, 0, 239, 240, 5, 61, 0, 0, 240, 71, 1, 0, 0, 0, 241, 242, 5, 62, 0, 0,
		242, 73, 1, 0, 0, 0, 243, 244, 5, 60, 0, 0, 244, 245, 5, 61, 0, 0, 245,
		75, 1, 0, 0, 0, 246, 247, 5, 60, 0, 0, 247, 77, 1, 0, 0, 0, 248, 249, 5,
		123, 0, 0, 249, 250, 6, 38, 1, 0, 250, 251, 1, 0, 0, 0, 251, 252, 6, 38,
		2, 0, 252, 79, 1, 0, 0, 0, 253, 254, 5, 125, 0, 0, 254, 255, 6, 39, 3,
		0, 255, 256, 1, 0, 0, 0, 256, 257, 6, 39, 4, 0, 257, 81, 1, 0, 0, 0, 258,
		259, 5, 96, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 6, 40, 5, 0, 261, 83,
		1, 0, 0, 0, 262, 263, 5, 111, 0, 0, 263, 264, 5, 114, 0, 0, 264, 85, 1,
		0, 0, 0, 265, 266, 5, 97, 0, 0, 266, 267, 5, 110, 0, 0, 267, 268, 5, 100,
		0, 0, 268, 87, 1, 0, 0, 0, 269, 270, 5, 116, 0, 0, 270, 271, 5, 114, 0,
		0, 271, 272, 5, 117, 0, 0, 272, 273, 5, 101, 0, 0, 273, 89, 1, 0, 0, 0,
		274, 275, 5, 102, 0, 0, 275, 276, 5, 97, 0, 0, 276, 277, 5, 108, 0, 0,
		277, 278, 5, 115, 0, 0, 278, 279, 5, 101, 0, 0, 279, 91, 1, 0, 0, 0, 280,
		281, 5, 110, 0, 0, 281, 282, 5, 101, 0, 0, 282, 283, 5, 119, 0, 0, 283,
		93, 1, 0, 0, 0, 284, 285, 5, 110, 0, 0, 285, 286, 5, 117, 0, 0, 286, 287,
		5, 108, 0, 0, 287, 288, 5, 108, 0, 0, 288, 95, 1, 0, 0, 0, 289, 290, 5,
		84, 0, 0, 290, 97, 1, 0, 0, 0, 291, 292, 5, 109, 0, 0, 292, 293, 5, 97,
		0, 0, 293, 294, 5, 116, 0, 0, 294, 295, 5, 99, 0, 0, 295, 296, 5, 104,
		0, 0, 296, 297, 5, 101, 0, 0, 297, 298, 5, 115, 0, 0, 298, 99, 1, 0, 0,
		0, 299, 300, 5, 103, 0, 0, 300, 301, 5, 116, 0, 0, 301, 101, 1, 0, 0, 0,
		302, 303, 5, 103, 0, 0, 303, 304, 5, 101, 0, 0, 304, 103, 1, 0, 0, 0, 305,
		306, 5, 108, 0, 0, 306, 307, 5, 101, 0, 0, 307, 105, 1, 0, 0, 0, 308, 309,
		5, 108, 0, 0, 309, 310, 5, 116, 0, 0, 310, 107, 1, 0, 0, 0, 311, 312, 5,
		101, 0, 0, 312, 313, 5, 113, 0, 0, 313, 109, 1, 0, 0, 0, 314, 315, 5, 110,
		0, 0, 315, 316, 5, 101, 0, 0, 316, 111, 1, 0, 0, 0, 317, 320, 3, 138, 68,
		0, 318, 320, 5, 95, 0, 0, 319, 317, 1, 0, 0, 0, 319, 318, 1, 0, 0, 0, 320,
		326, 1, 0, 0, 0, 321, 325, 3, 138, 68, 0, 322, 325, 3, 140, 69, 0, 323,
		325, 7, 2, 0, 0, 324, 321, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 323,
		1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0,
		0, 0, 327, 113, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 329, 331, 5, 46, 0, 0,
		330, 332, 3, 130, 64, 0, 331, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333,
		331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335, 337,
		3, 132, 65, 0, 336, 335, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 339, 1,
		0, 0, 0, 338, 340, 3, 136, 67, 0, 339, 338, 1, 0, 0, 0, 339, 340, 1, 0,
		0, 0, 340, 375, 1, 0, 0, 0, 341, 343, 3, 130, 64, 0, 342, 341, 1, 0, 0,
		0, 343, 344, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345,
		346, 1, 0, 0, 0, 346, 348, 5, 46, 0, 0, 347, 349, 3, 130, 64, 0, 348, 347,
		1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0,
		0, 0, 351, 353, 1, 0, 0, 0, 352, 354, 3, 132, 65, 0, 353, 352, 1, 0, 0,
		0, 353, 354, 1, 0, 0, 0, 354, 356, 1, 0, 0, 0, 355, 357, 3, 136, 67, 0,
		356, 355, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 375, 1, 0, 0, 0, 358,
		360, 3, 130, 64, 0, 359, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 359,
		1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 365, 3, 132,
		65, 0, 364, 366, 3, 136, 67, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0,
		0, 366, 375, 1, 0, 0, 0, 367, 369, 3, 130, 64, 0, 368, 367, 1, 0, 0, 0,
		369, 370, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371,
		372, 1, 0, 0, 0, 372, 373, 3, 136, 67, 0, 373, 375, 1, 0, 0, 0, 374, 329,
		1, 0, 0, 0, 374, 342, 1, 0, 0, 0, 374, 359, 1, 0, 0, 0, 374, 368, 1, 0,
		0, 0, 375, 115, 1, 0, 0, 0, 376, 378, 3, 130, 64, 0, 377, 376, 1, 0, 0,
		0, 378, 379, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380,
		382, 1, 0, 0, 0, 381, 383, 3, 126, 62, 0, 382, 381, 1, 0, 0, 0, 382, 383,
		1, 0, 0, 0, 383, 117, 1, 0, 0, 0, 384, 387, 3, 120, 59, 0, 385, 387, 3,
		122, 60, 0, 386, 384, 1, 0, 0, 0, 386, 385, 1, 0, 0, 0, 387, 119, 1, 0,
		0, 0, 388, 394, 5, 39, 0, 0, 389, 390, 5, 39, 0, 0, 390, 393, 5, 39, 0,
		0, 391, 393, 8, 3, 0, 0, 392, 389, 1, 0, 0, 0, 392, 391, 1, 0, 0, 0, 393,
		396, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 397,
		1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 397, 398, 5, 39, 0, 0, 398, 121, 1, 0,
		0, 0, 399, 405, 5, 34, 0, 0, 400, 401, 5, 34, 0, 0, 401, 404, 5, 34, 0,
		0, 402, 404, 8, 4, 0, 0, 403, 400, 1, 0, 0, 0, 403, 402, 1, 0, 0, 0, 404,
		407, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 408,
		1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 408, 409, 5, 34, 0, 0, 409, 123, 1, 0,
		0, 0, 410, 411, 5, 36, 0, 0, 411, 412, 5, 123, 0, 0, 412, 416, 1, 0, 0,
		0, 413, 415, 9, 0, 0, 0, 414, 413, 1, 0, 0, 0, 415, 418, 1, 0, 0, 0, 416,
		417, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 417, 419, 1, 0, 0, 0, 418, 416,
		1, 0, 0, 0, 419, 420, 5, 125, 0, 0, 420, 125, 1, 0, 0, 0, 421, 422, 7,
		5, 0, 0, 422, 127, 1, 0, 0, 0, 423, 424, 7, 6, 0, 0, 424, 129, 1, 0, 0,
		0, 425, 426, 7, 7, 0, 0, 426, 131, 1, 0, 0, 0, 427, 431, 5, 101, 0, 0,
		428, 430, 3, 134, 66, 0, 429, 428, 1, 0, 0, 0, 430, 433, 1, 0, 0, 0, 431,
		429, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 435, 1, 0, 0, 0, 433, 431,
		1, 0, 0, 0, 434, 436, 3, 130, 64, 0, 435, 434, 1, 0, 0, 0, 436, 437, 1,
		0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 452, 1, 0, 0,
		0, 439, 443, 5, 69, 0, 0, 440, 442, 3, 134, 66, 0, 441, 440, 1, 0, 0, 0,
		442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444,
		447, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 448, 3, 130, 64, 0, 447, 446,
		1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 449, 450, 1, 0,
		0, 0, 450, 452, 1, 0, 0, 0, 451, 427, 1, 0, 0, 0, 451, 439, 1, 0, 0, 0,
		452, 133, 1, 0, 0, 0, 453, 454, 7, 8, 0, 0, 454, 135, 1, 0, 0, 0, 455,
		456, 7, 9, 0, 0, 456, 137, 1, 0, 0, 0, 457, 459, 7, 10, 0, 0, 458, 457,
		1, 0, 0, 0, 459, 139, 1, 0, 0, 0, 460, 461, 7, 7, 0, 0, 461, 141, 1, 0,
		0, 0, 462, 463, 5, 96, 0, 0, 463, 464, 5, 96, 0, 0, 464, 143, 1, 0, 0,
		0, 465, 466, 5, 35, 0, 0, 466, 467, 5, 123, 0, 0, 467, 468, 1, 0, 0, 0,
		468, 469, 6, 71, 6, 0, 469, 470, 1, 0, 0, 0, 470, 471, 6, 71, 2, 0, 471,
		145, 1, 0, 0, 0, 472, 473, 8, 11, 0, 0, 473, 147, 1, 0, 0, 0, 474, 475,
		5, 96, 0, 0, 475, 476, 1, 0, 0, 0, 476, 477, 6, 73, 7, 0, 477, 478, 6,
		73, 4, 0, 478, 149, 1, 0, 0, 0, 32, 0, 1, 155, 160, 319, 324, 326, 333,
		336, 339, 344, 350, 353, 356, 361, 365, 370, 374, 379, 382, 386, 392, 394,
		403, 405, 416, 431, 437, 443, 449, 451, 458, 8, 0, 1, 0, 1, 38, 0, 5, 0,
		0, 1, 39, 1, 4, 0, 0, 5, 1, 0, 1, 71, 2, 7, 40, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SpelLexerInit initializes any static state used to implement SpelLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSpelLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SpelLexerInit() {
	staticData := &spellexerLexerStaticData
	staticData.once.Do(spellexerLexerInit)
}

// NewSpelLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSpelLexer(input antlr.CharStream) *SpelLexer {
	SpelLexerInit()
	l := new(SpelLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &spellexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SpelLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SpelLexer tokens.
const (
	SpelLexerSEMICOLON                    = 1
	SpelLexerWS                           = 2
	SpelLexerINC                          = 3
	SpelLexerPLUS                         = 4
	SpelLexerDEC                          = 5
	SpelLexerMINUS                        = 6
	SpelLexerCOLON                        = 7
	SpelLexerDOT                          = 8
	SpelLexerCOMMA                        = 9
	SpelLexerSTAR                         = 10
	SpelLexerDIV                          = 11
	SpelLexerMOD                          = 12
	SpelLexerLPAREN                       = 13
	SpelLexerRPAREN                       = 14
	SpelLexerLSQUARE                      = 15
	SpelLexerRSQUARE                      = 16
	SpelLexerHASH                         = 17
	SpelLexerBEAN_REF                     = 18
	SpelLexerSELECT_FIRST                 = 19
	SpelLexerPOWER                        = 20
	SpelLexerNE                           = 21
	SpelLexerPROJECT                      = 22
	SpelLexerNOT                          = 23
	SpelLexerEQ                           = 24
	SpelLexerASSIGN                       = 25
	SpelLexerSYMBOLIC_AND                 = 26
	SpelLexerFACTORY_BEAN_REF             = 27
	SpelLexerSYMBOLIC_OR                  = 28
	SpelLexerSELECT                       = 29
	SpelLexerELVIS                        = 30
	SpelLexerSAFE_NAVI                    = 31
	SpelLexerQMARK                        = 32
	SpelLexerSELECT_LAST                  = 33
	SpelLexerGE                           = 34
	SpelLexerGT                           = 35
	SpelLexerLE                           = 36
	SpelLexerLT                           = 37
	SpelLexerLCURLY                       = 38
	SpelLexerRCURLY                       = 39
	SpelLexerBACKTICK                     = 40
	SpelLexerOR                           = 41
	SpelLexerAND                          = 42
	SpelLexerTRUE                         = 43
	SpelLexerFALSE                        = 44
	SpelLexerNEW                          = 45
	SpelLexerNULL                         = 46
	SpelLexerT                            = 47
	SpelLexerMATCHES                      = 48
	SpelLexerGT_KEYWORD                   = 49
	SpelLexerGE_KEYWORD                   = 50
	SpelLexerLE_KEYWORD                   = 51
	SpelLexerLT_KEYWORD                   = 52
	SpelLexerEQ_KEYWORD                   = 53
	SpelLexerNE_KEYWORD                   = 54
	SpelLexerIDENTIFIER                   = 55
	SpelLexerREAL_LITERAL                 = 56
	SpelLexerINTEGER_LITERAL              = 57
	SpelLexerSTRING_LITERAL               = 58
	SpelLexerSINGLE_QUOTED_STRING         = 59
	SpelLexerDOUBLE_QUOTED_STRING         = 60
	SpelLexerPROPERTY_PLACE_HOLDER        = 61
	SpelLexerESCAPED_BACKTICK             = 62
	SpelLexerSPEL_IN_TEMPLATE_STRING_OPEN = 63
	SpelLexerTEMPLATE_TEXT                = 64
)

// SpelLexerIN_TEMPLATE_STRING is the SpelLexer mode.
const SpelLexerIN_TEMPLATE_STRING = 1

func (l *SpelLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 38:
		l.LCURLY_Action(localctx, actionIndex)

	case 39:
		l.RCURLY_Action(localctx, actionIndex)

	case 71:
		l.SPEL_IN_TEMPLATE_STRING_OPEN_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *SpelLexer) LCURLY_Action(localctx antlr.RuleContext, actionIndex int) {
	this := l
	_ = this

	switch actionIndex {
	case 0:
		this.indent++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SpelLexer) RCURLY_Action(localctx antlr.RuleContext, actionIndex int) {
	this := l
	_ = this

	switch actionIndex {
	case 1:

		if this.indent > 0 {
			this.indent--
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *SpelLexer) SPEL_IN_TEMPLATE_STRING_OPEN_Action(localctx antlr.RuleContext, actionIndex int) {
	this := l
	_ = this

	switch actionIndex {
	case 2:
		this.indent++

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}