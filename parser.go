//line parser.go.y:2
package main

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:6
type yySymType struct {
	yys     int
	token   Token
	block   Block
	blocks  []Block
	inline  Inline
	inlines []Inline
	url     string
	depth   int
}

const TEXT = 57346
const UNORDERED_LIST_MARKER = 57347
const ORDERED_LIST_MARKER = 57348
const CR = 57349
const LBRACKET = 57350
const RBRACKET = 57351
const LT = 57352
const GT = 57353
const HEADING_MARKER = 57354

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TEXT",
	"UNORDERED_LIST_MARKER",
	"ORDERED_LIST_MARKER",
	"CR",
	"LBRACKET",
	"RBRACKET",
	"LT",
	"GT",
	"HEADING_MARKER",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:191

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 34
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 57

var yyAct = [...]int{

	34, 13, 14, 25, 23, 15, 18, 24, 11, 17,
	16, 17, 36, 30, 46, 49, 32, 33, 4, 35,
	42, 38, 48, 39, 1, 40, 37, 26, 28, 25,
	3, 51, 15, 24, 50, 29, 47, 45, 44, 27,
	43, 31, 23, 36, 41, 20, 12, 10, 21, 22,
	19, 6, 5, 7, 9, 8, 2,
}
var yyPact = [...]int{

	-1, -1000, -1, -1000, -1000, -1000, -1000, -1000, 37, 26,
	-1, 25, -1000, 34, 25, 25, 8, -3, -1000, 25,
	37, -1000, -1000, -1000, 39, -1000, -1000, -1000, -1000, 10,
	33, -1000, 31, 30, 3, 29, -1000, -1000, -1000, -1000,
	13, -1000, 5, -1000, -1000, -1000, 27, -1000, -1000, 24,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 56, 55, 30, 18, 54, 53, 52, 51, 24,
	50, 49, 48, 1, 0, 47, 2, 8, 46, 45,
	44,
}
var yyR1 = [...]int{

	0, 9, 9, 1, 1, 1, 1, 1, 6, 6,
	18, 13, 13, 10, 10, 11, 12, 14, 3, 3,
	2, 16, 16, 19, 4, 4, 5, 7, 15, 15,
	20, 8, 17, 17,
}
var yyR2 = [...]int{

	0, 1, 2, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 2, 1, 1, 1, 3, 1, 1, 2,
	3, 1, 2, 1, 1, 2, 3, 3, 4, 3,
	3, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -9, -1, -3, -4, -7, -8, -6, -2, -5,
	-15, -17, -18, -13, -16, 6, 11, 12, 7, -10,
	-19, -12, -11, 5, 8, 4, -9, -3, -4, -9,
	-13, 7, -13, -13, -14, 11, 4, -17, -13, -16,
	-14, -20, 10, 7, 7, 7, 11, 7, 9, 10,
	7, 7,
}
var yyDef = [...]int{

	0, -2, 1, 3, 4, 5, 6, 7, 18, 24,
	0, 0, 8, 0, 0, 0, 0, 32, 10, 11,
	21, 13, 14, 23, 0, 15, 2, 19, 25, 0,
	0, 9, 0, 0, 0, 0, 17, 33, 12, 22,
	0, 27, 0, 31, 20, 26, 0, 29, 16, 0,
	28, 30,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:30
		{
			yyVAL.blocks = []Block{yyDollar[1].block}
			yylex.(*Lexer).result = yyVAL.blocks
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:35
		{
			yyVAL.blocks = append([]Block{yyDollar[1].block}, yyDollar[2].blocks...)
			yylex.(*Lexer).result = yyVAL.blocks
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:42
		{
			yyVAL.block = yyDollar[1].block
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:46
		{
			yyVAL.block = yyDollar[1].block
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:50
		{
			yyVAL.block = yyDollar[1].block
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:54
		{
			yyVAL.block = yyDollar[1].block
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:58
		{
			yyVAL.block = yyDollar[1].block
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:64
		{
			yyVAL.block = Line{Inlines: []Inline{}}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:68
		{
			yyVAL.block = Line{Inlines: yyDollar[1].inlines}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.inlines = []Inline{yyDollar[1].inline}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.inlines = append([]Inline{yyDollar[1].inline}, yyDollar[2].inlines...)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.inline = InlineText{Literal: yyDollar[1].token.literal}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.inline = InlineHttp{Url: yyDollar[2].url}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.url = yyDollar[1].token.literal
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.block = UnorderedList{Items: []UnorderedListItem{yyDollar[1].block.(UnorderedListItem)}}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:112
		{
			items := yyDollar[2].block.(UnorderedList).Items
			list := UnorderedList{Items: append([]UnorderedListItem{yyDollar[1].block.(UnorderedListItem)}, items...)}
			yyVAL.block = list
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.block = UnorderedListItem{Depth: yyDollar[1].depth, Inlines: yyDollar[2].inlines}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.depth = 1
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:130
		{
			yyVAL.depth = yyDollar[2].depth + 1
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:140
		{
			yyVAL.block = OrderedList{Items: []OrderedListItem{yyDollar[1].block.(OrderedListItem)}}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:144
		{
			items := yyDollar[2].block.(OrderedList).Items
			list := OrderedList{Items: append([]OrderedListItem{yyDollar[1].block.(OrderedListItem)}, items...)}
			yyVAL.block = list
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.block = OrderedListItem{Inlines: yyDollar[2].inlines}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.block = Quotation{Cite: yyDollar[1].url, Content: yyDollar[2].blocks}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.url = yyDollar[2].url
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.url = ""
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.block = Heading{Level: yyDollar[1].depth, Content: yyDollar[2].inlines}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.depth = 1
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.depth = yyDollar[2].depth + 1
		}
	}
	goto yystack /* stack new state and value */
}
