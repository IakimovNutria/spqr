// Code generated by goyacc -o yacc/console/sql.go -p yy yacc/console/sql.y. DO NOT EDIT.

//line yacc/console/sql.y:3

package spqrparser

import __yyfmt__ "fmt"

//line yacc/console/sql.y:4

//line yacc/console/sql.y:11
type yySymType struct {
	yys               int
	empty             struct{}
	statement         Statement
	show              *Show
	kr                *AddKeyRange
	shard             *AddShard
	register_router   *RegisterRouter
	unregister_router *UnregisterRouter
	kill              *Kill
	drop              *Drop
	add               *Add
	dropAll           *DropAll
	lock              *Lock
	shutdown          *Shutdown
	listen            *Listen
	unlock            *Unlock
	split             *SplitKeyRange
	move              *MoveKeyRange
	unite             *UniteKeyRange
	str               string
	byte              byte
	bytes             []byte
	int               int
	bool              bool
	entrieslist       []ShardingRuleEntry
	shruleEntry       ShardingRuleEntry
}

const STRING = 57346
const COMMAND = 57347
const SHOW = 57348
const KILL = 57349
const POOLS = 57350
const STATS = 57351
const LISTS = 57352
const SERVERS = 57353
const CLIENTS = 57354
const DATABASES = 57355
const SHUTDOWN = 57356
const LISTEN = 57357
const REGISTER = 57358
const UNREGISTER = 57359
const ROUTER = 57360
const ROUTE = 57361
const CREATE = 57362
const ADD = 57363
const DROP = 57364
const LOCK = 57365
const UNLOCK = 57366
const SPLIT = 57367
const MOVE = 57368
const COMPOSE = 57369
const SHARDING = 57370
const COLUMN = 57371
const TABLE = 57372
const HASH = 57373
const FUNCTION = 57374
const KEY = 57375
const RANGE = 57376
const DATASPACE = 57377
const SHARDS = 57378
const KEY_RANGES = 57379
const ROUTERS = 57380
const SHARD = 57381
const HOST = 57382
const SHARDING_RULES = 57383
const RULE = 57384
const COLUMNS = 57385
const BY = 57386
const FROM = 57387
const TO = 57388
const WITH = 57389
const UNITE = 57390
const ALL = 57391
const ADDRESS = 57392

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"STRING",
	"COMMAND",
	"SHOW",
	"KILL",
	"POOLS",
	"STATS",
	"LISTS",
	"SERVERS",
	"CLIENTS",
	"DATABASES",
	"SHUTDOWN",
	"LISTEN",
	"REGISTER",
	"UNREGISTER",
	"ROUTER",
	"ROUTE",
	"CREATE",
	"ADD",
	"DROP",
	"LOCK",
	"UNLOCK",
	"SPLIT",
	"MOVE",
	"COMPOSE",
	"SHARDING",
	"COLUMN",
	"TABLE",
	"HASH",
	"FUNCTION",
	"KEY",
	"RANGE",
	"DATASPACE",
	"SHARDS",
	"KEY_RANGES",
	"ROUTERS",
	"SHARD",
	"HOST",
	"SHARDING_RULES",
	"RULE",
	"COLUMNS",
	"BY",
	"FROM",
	"TO",
	"WITH",
	"UNITE",
	"ALL",
	"ADDRESS",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc/console/sql.y:483

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 141

var yyAct = [...]int{

	71, 110, 127, 86, 113, 60, 88, 76, 28, 29,
	40, 98, 87, 89, 77, 107, 31, 30, 35, 25,
	90, 137, 36, 23, 24, 37, 38, 32, 33, 123,
	106, 105, 49, 54, 101, 52, 51, 50, 130, 73,
	115, 70, 42, 104, 134, 65, 85, 41, 66, 84,
	34, 80, 79, 43, 116, 78, 74, 94, 92, 75,
	53, 55, 56, 69, 44, 57, 68, 67, 64, 45,
	63, 62, 132, 81, 126, 103, 48, 46, 93, 72,
	91, 136, 95, 96, 97, 111, 128, 87, 99, 100,
	61, 118, 109, 89, 83, 77, 39, 1, 135, 117,
	82, 108, 21, 20, 19, 18, 59, 121, 17, 120,
	119, 122, 15, 16, 26, 11, 27, 124, 12, 129,
	125, 114, 102, 112, 22, 131, 5, 4, 3, 8,
	7, 6, 133, 9, 10, 14, 13, 58, 138, 47,
	2,
}
var yyPact = [...]int{

	2, -1000, -41, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 14, 36, 59, -1000, -1000, 24, 24,
	86, -1000, 38, 37, 35, 27, 13, 34, 33, -1000,
	-1000, 29, -1, 75, -3, 22, 10, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 21, 18, 17, 91, 90, 15, 12, 83,
	89, -27, -1000, 9, 8, -1000, -1000, -1000, 83, 83,
	83, -39, -1000, -1000, 83, 83, -11, -1000, 45, -1000,
	3, -1000, -1000, -1000, -1000, -14, -16, -32, 88, -1000,
	-1000, 81, 11, 87, 86, 83, 75, 83, -1000, -1000,
	-17, -1000, 11, -1000, 43, 82, 82, -1000, -1000, -1000,
	-6, -1000, -1000, 81, -1000, -1000, 40, -1000, -1000, -1000,
	81, 25, 77, -1000, -25, -1000, -1000, 75, -1000,
}
var yyPgo = [...]int{

	0, 140, 139, 137, 136, 135, 134, 133, 131, 130,
	129, 128, 127, 126, 124, 123, 4, 122, 121, 120,
	118, 116, 115, 114, 113, 112, 108, 105, 104, 103,
	102, 76, 0, 5, 1, 3, 7, 101, 6, 100,
	2, 99, 98, 97, 96,
}
var yyR1 = [...]int{

	0, 43, 44, 44, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 31, 31, 31, 31, 31, 31,
	31, 31, 31, 2, 3, 4, 40, 41, 42, 38,
	39, 34, 35, 32, 33, 22, 12, 14, 13, 15,
	15, 16, 17, 17, 18, 18, 19, 19, 11, 20,
	6, 7, 23, 21, 26, 5, 27, 28, 25, 24,
	37, 36, 29, 8, 9, 10, 30,
}
var yyR2 = [...]int{

	0, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 11, 3, 6, 1,
	2, 2, 2, 0, 2, 2, 3, 0, 6, 1,
	4, 4, 4, 4, 8, 2, 6, 6, 2, 1,
	1, 1, 5, 4, 4, 3, 3,
}
var yyChk = [...]int{

	-1000, -43, -1, -11, -12, -13, -8, -9, -10, -7,
	-6, -22, -20, -4, -5, -25, -24, -26, -27, -28,
	-29, -30, -14, 21, 22, 17, -23, -21, 6, 7,
	15, 14, 25, 26, 48, 16, 20, 23, 24, -44,
	51, 33, 28, 39, 28, 33, 18, -2, -31, 8,
	13, 12, 11, 36, 9, 37, 38, 41, -3, -31,
	-33, 4, 33, 33, 33, 18, 35, 33, 33, 34,
	42, -32, 4, 42, 34, 49, -36, 4, 34, 34,
	34, -36, -39, 4, 34, 34, -35, 4, -38, 4,
	47, -38, 49, -35, 49, -35, -35, -35, 50, -35,
	-35, 45, -17, 30, 40, 45, 46, 47, -37, 4,
	-34, 4, -15, -16, -18, 29, 43, -41, 4, -33,
	-35, -32, -35, 46, -16, -19, 31, -40, 4, -40,
	44, -34, 32, -34, 19, -42, 4, 46, -32,
}
var yyDef = [...]int{

	0, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 0, 0, 0, 45, 59, 0, 0,
	0, 69, 0, 0, 0, 0, 0, 0, 0, 1,
	3, 0, 0, 0, 0, 0, 0, 35, 33, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 65, 34,
	68, 44, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 0, 0, 75, 76, 71, 0, 0,
	0, 0, 47, 40, 0, 0, 0, 42, 53, 39,
	0, 60, 74, 61, 73, 0, 0, 0, 0, 62,
	63, 0, 0, 0, 0, 0, 0, 0, 72, 70,
	0, 41, 48, 49, 57, 0, 0, 52, 37, 58,
	0, 66, 67, 0, 50, 51, 0, 54, 36, 55,
	0, 0, 0, 64, 0, 56, 38, 0, 46,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50,
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

	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/console/sql.y:113
		{
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:114
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:119
		{
			setParseTree(yylex, yyDollar[1].add)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:123
		{
			setParseTree(yylex, yyDollar[1].add)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:127
		{
			setParseTree(yylex, yyDollar[1].add)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:131
		{
			setParseTree(yylex, yyDollar[1].dropAll)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:135
		{
			setParseTree(yylex, yyDollar[1].dropAll)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:139
		{
			setParseTree(yylex, yyDollar[1].dropAll)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:143
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:147
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:151
		{
			setParseTree(yylex, yyDollar[1].lock)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:155
		{
			setParseTree(yylex, yyDollar[1].unlock)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:159
		{
			setParseTree(yylex, yyDollar[1].show)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:163
		{
			setParseTree(yylex, yyDollar[1].kill)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:167
		{
			setParseTree(yylex, yyDollar[1].listen)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:171
		{
			setParseTree(yylex, yyDollar[1].shutdown)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:175
		{
			setParseTree(yylex, yyDollar[1].split)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:179
		{
			setParseTree(yylex, yyDollar[1].move)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:183
		{
			setParseTree(yylex, yyDollar[1].unite)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:187
		{
			setParseTree(yylex, yyDollar[1].register_router)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:191
		{
			setParseTree(yylex, yyDollar[1].unregister_router)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:195
		{
			setParseTree(yylex, yyDollar[1].add)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:212
		{
			switch v := string(yyDollar[1].str); v {
			case ShowDatabasesStr, ShowRoutersStr, ShowPoolsStr, ShowShardsStr, ShowKeyRangesStr, ShowShardingRules:
				yyVAL.str = v
			default:
				yyVAL.str = ShowUnsupportedStr
			}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:223
		{
			switch v := string(yyDollar[1].str); v {
			case KillClientsStr:
				yyVAL.str = v
			default:
				yyVAL.str = "unsupp"
			}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:235
		{
			yyVAL.show = &Show{Cmd: yyDollar[2].str}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:241
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:247
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:253
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:259
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:265
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:271
		{
			yyVAL.bytes = []byte(yyDollar[1].str)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:277
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:284
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:290
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 46:
		yyDollar = yyS[yypt-11 : yypt+1]
//line yacc/console/sql.y:299
		{
			yyVAL.add = &Add{Element: &AddKeyRange{LowerBound: yyDollar[6].bytes, UpperBound: yyDollar[8].bytes, ShardID: yyDollar[11].str, KeyRangeID: yyDollar[4].str}}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:305
		{
			yyVAL.add = &Add{Element: &AddDataspace{ID: yyDollar[3].str}}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
//line yacc/console/sql.y:317
		{
			yyVAL.add = &Add{Element: &AddShardingRule{ID: yyDollar[4].str, TableName: yyDollar[5].str, Entries: yyDollar[6].entrieslist}}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:322
		{
			yyVAL.entrieslist = make([]ShardingRuleEntry, 0)
			yyVAL.entrieslist = append(yyVAL.entrieslist, yyDollar[1].shruleEntry)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:328
		{
			yyVAL.entrieslist = append(yyDollar[1].entrieslist, yyDollar[2].shruleEntry)
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:334
		{
			yyVAL.shruleEntry = ShardingRuleEntry{
				Column:       yyDollar[1].str,
				HashFunction: yyDollar[2].str,
			}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:343
		{
			yyVAL.str = yyDollar[2].str
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/console/sql.y:346
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:350
		{
			yyVAL.str = yyDollar[2].str
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:355
		{
			yyVAL.str = yyDollar[2].str
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:361
		{
			yyVAL.str = yyDollar[3].str
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/console/sql.y:364
		{
			yyVAL.str = ""
		}
	case 58:
		yyDollar = yyS[yypt-6 : yypt+1]
//line yacc/console/sql.y:368
		{
			yyVAL.add = &Add{Element: &AddShard{Id: yyDollar[3].str, Hosts: []string{yyDollar[6].str}}}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:377
		{
			yyVAL.drop = &Drop{Element: &DropShardingRule{ID: yyDollar[4].str}}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:383
		{
			yyVAL.drop = &Drop{Element: &DropKeyRange{KeyRangeID: yyDollar[4].str}}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:389
		{
			yyVAL.lock = &Lock{KeyRangeID: yyDollar[4].str}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:395
		{
			yyVAL.unlock = &Unlock{KeyRangeID: yyDollar[4].str}
		}
	case 64:
		yyDollar = yyS[yypt-8 : yypt+1]
//line yacc/console/sql.y:402
		{
			yyVAL.split = &SplitKeyRange{KeyRangeID: yyDollar[4].str, KeyRangeFromID: yyDollar[6].str, Border: yyDollar[8].bytes}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:408
		{
			yyVAL.kill = &Kill{Cmd: yyDollar[2].str}
		}
	case 66:
		yyDollar = yyS[yypt-6 : yypt+1]
//line yacc/console/sql.y:414
		{
			yyVAL.move = &MoveKeyRange{KeyRangeID: yyDollar[4].str, DestShardID: yyDollar[6].str}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
//line yacc/console/sql.y:420
		{
			yyVAL.unite = &UniteKeyRange{KeyRangeIDL: yyDollar[4].str, KeyRangeIDR: yyDollar[5].str}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:426
		{
			yyVAL.listen = &Listen{addr: yyDollar[2].str}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:432
		{
			yyVAL.shutdown = &Shutdown{}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:440
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:446
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
//line yacc/console/sql.y:452
		{
			yyVAL.register_router = &RegisterRouter{ID: yyDollar[3].str, Addr: yyDollar[5].str}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:458
		{
			yyVAL.dropAll = &DropAll{Entity: EntityKeyRanges}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:465
		{
			yyVAL.dropAll = &DropAll{Entity: EntityShardingRule}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:472
		{
			yyVAL.dropAll = &DropAll{Entity: EntityRouters}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:478
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: yyDollar[3].str}
		}
	}
	goto yystack /* stack new state and value */
}
