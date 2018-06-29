package main

import (
	"fmt"
	testing "testing"

	"github.com/orcaman/concurrent-map"

	"github.com/json-iterator/go"
)

func TestConvertToMapped(t *testing.T) {
	boggleChars := BoggleChars{
		Lang: "en_US",
		Rows: []BoggleRows{
			{
				Cols: []BoggleCols{
					{
						Char: "a",
					}, {
						Char: "c",
					},
				},
			}, {
				Cols: []BoggleCols{
					{
						Char: "t",
					}, {
						Char: "m",
					},
				},
			},
		},
	}

	if len(boggleChars.Rows) != 2 {
		t.Errorf("format error")
	}
	if len(boggleChars.Rows[0].Cols) != 2 {
		t.Errorf("format error")
	}

	mapped := ConvertToMapped(boggleChars)

	if (*mapped)[0][0].Char != "a" {
		t.Error("ConvertToMapped maps chars (a)")
	}
	if (*mapped)[0][1].Char != "c" {
		t.Error("ConvertToMapped maps chars (c)")
	}
	if (*mapped)[1][0].Char != "t" {
		t.Error("ConvertToMapped maps chars (t)")
	}
	if (*mapped)[1][1].Char != "m" {
		t.Error("ConvertToMapped maps chars (m)")
	}

	if (*mapped)[0][0].East.Char != "c" {
		t.Error("ConvertToMapped maps East (c)")
	}
	if (*mapped)[0][0].South.Char != "t" {
		t.Error("ConvertToMapped maps South (c)")
	}
	if (*mapped)[0][0].SouthEast.Char != "m" {
		t.Error("ConvertToMapped maps SouthEast (m)")
	}

	if (*mapped)[0][1].West.Char != "a" {
		t.Error("ConvertToMapped maps West (a)")
	}
	if (*mapped)[0][1].SouthWest.Char != "t" {
		t.Error("ConvertToMapped maps West (a)")
	}
	if (*mapped)[0][1].South.Char != "m" {
		t.Error("ConvertToMapped maps West (a)")
	}

	if (*mapped)[1][0].East.Char != "m" {
		t.Error("ConvertToMapped maps West (m)")
	}
	if (*mapped)[1][0].North.Char != "a" {
		t.Error("ConvertToMapped maps North (a)")
	}
	if (*mapped)[1][0].NorthEast.Char != "c" {
		t.Error("ConvertToMapped maps NorthEast (c)")
	}

	if (*mapped)[1][1].West.Char != "t" {
		t.Error("ConvertToMapped maps West (t)")
	}
	if (*mapped)[1][1].North.Char != "c" {
		t.Error("ConvertToMapped maps North (c)")
	}
	if (*mapped)[1][1].NorthWest.Char != "a" {
		t.Error("ConvertToMapped maps NorthWest (a)")
	}
}

func TestGetAllPossibleWords(t *testing.T) {

	hunLangs, err := LoadAllLanguageFiles(10)
	boggleChars := BoggleChars{
		Lang: "en_US",
		Rows: []BoggleRows{
			{
				Cols: []BoggleCols{
					{
						Char: "a",
					}, {
						Char: "c",
					},
				},
			}, {
				Cols: []BoggleCols{
					{
						Char: "t",
					}, {
						Char: "m",
					},
				},
			},
		},
	}

	mapped := ConvertToMapped(boggleChars)
	allwords, err := GetAllValidWords(hunLangs[boggleChars.Lang], mapped, 10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, word := range allwords {
		fmt.Println(word)
	}
}

func TestGetAllValidWords(t *testing.T) {
	langMap, err := LoadAllLanguageFiles(10)

	boggleChars := BoggleChars{
		Lang: "en_US",
		Rows: []BoggleRows{
			{
				Cols: []BoggleCols{
					{
						Char: "a",
					}, {
						Char: "c",
					},
				},
			}, {
				Cols: []BoggleCols{
					{
						Char: "t",
					}, {
						Char: "m",
					},
				},
			},
		},
	}

	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 10)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}
}

func TestLargeBoard(t *testing.T) {

	//arrange
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	langMap, err := LoadAllLanguageFiles(10)
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"h"},{"char":"m"},{"char":"v"},{"char":"y"}]},{"cols":[{"char":"b"},{"char":"u"},{"char":"x"},{"char":"a"}]},{"cols":[{"char":"y"},{"char":"t"},{"char":"a"},{"char":"w"}]},{"cols":[{"char":"s"},{"char":"o"},{"char":"o"},{"char":"p"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)

	//act
	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 10)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}

	if len(validWords) != 41 {
		t.Errorf("ConvertToMapped did not find 41 words found: %s ", string(len(validWords)))
	}

	if (*mapped)[0][0].Char != "h" {
		t.Error("ConvertToMapped maps chars (a)")
	}
	if (*mapped)[0][0].South.Char != "b" {
		t.Error("ConvertToMapped maps chars (b")
	}
	if (*mapped)[0][0].East.Char != "m" {
		t.Error("ConvertToMapped maps chars (m)")
	}
	if (*mapped)[0][0].SouthEast.Char != "u" {
		t.Error("ConvertToMapped maps chars (m)")
	}
	// spells but
	if (*mapped)[1][0].Char != "b" {
		t.Error("Does NOT spell but")
	}
	if (*mapped)[1][0].East.Char != "u" {
		t.Error("Does NOT spell but")
	}
	if (*mapped)[1][0].SouthEast.Char != "t" {
		t.Error("Does NOT spell but")
	}
}

func Test5x5Board(t *testing.T) {
	//arrange
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	langMap, err := LoadAllLanguageFiles(7)
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"g"},{"char":"i"},{"char":"x"},{"char":"b"},{"char":"x"}]},{"cols":[{"char":"r"},{"char":"e"},{"char":"y"},{"char":"b"},{"char":"i"}]},{"cols":[{"char":"b"},{"char":"w"},{"char":"t"},{"char":"y"},{"char":"t"}]},{"cols":[{"char":"k"},{"char":"y"},{"char":"u"},{"char":"l"},{"char":"i"}]},{"cols":[{"char":"e"},{"char":"i"},{"char":"l"},{"char":"h"},{"char":"f"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)

	//act
	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 7)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}
}

func Test20x20Board(t *testing.T) {
	//arrange
	validWordMap := cmap.New()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	langMap, err := LoadAllLanguageFiles(5)
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"p"},{"char":"w"},{"char":"l"},{"char":"z"},{"char":"z"},{"char":"p"},{"char":"x"},{"char":"k"},{"char":"j"},{"char":"z"},{"char":"k"},{"char":"z"},{"char":"m"},{"char":"d"},{"char":"q"},{"char":"g"},{"char":"f"},{"char":"a"},{"char":"x"},{"char":"f"}]},{"cols":[{"char":"j"},{"char":"e"},{"char":"v"},{"char":"s"},{"char":"n"},{"char":"v"},{"char":"v"},{"char":"h"},{"char":"o"},{"char":"x"},{"char":"y"},{"char":"p"},{"char":"m"},{"char":"s"},{"char":"q"},{"char":"e"},{"char":"b"},{"char":"k"},{"char":"z"},{"char":"m"}]},{"cols":[{"char":"r"},{"char":"d"},{"char":"x"},{"char":"r"},{"char":"r"},{"char":"o"},{"char":"s"},{"char":"t"},{"char":"x"},{"char":"y"},{"char":"c"},{"char":"r"},{"char":"s"},{"char":"b"},{"char":"b"},{"char":"u"},{"char":"b"},{"char":"r"},{"char":"m"},{"char":"o"}]},{"cols":[{"char":"t"},{"char":"s"},{"char":"f"},{"char":"d"},{"char":"r"},{"char":"l"},{"char":"o"},{"char":"z"},{"char":"h"},{"char":"f"},{"char":"w"},{"char":"v"},{"char":"l"},{"char":"h"},{"char":"l"},{"char":"b"},{"char":"n"},{"char":"w"},{"char":"w"},{"char":"m"}]},{"cols":[{"char":"t"},{"char":"o"},{"char":"m"},{"char":"u"},{"char":"k"},{"char":"y"},{"char":"i"},{"char":"o"},{"char":"c"},{"char":"p"},{"char":"w"},{"char":"a"},{"char":"w"},{"char":"h"},{"char":"l"},{"char":"q"},{"char":"g"},{"char":"g"},{"char":"r"},{"char":"z"}]},{"cols":[{"char":"v"},{"char":"b"},{"char":"z"},{"char":"w"},{"char":"l"},{"char":"a"},{"char":"r"},{"char":"j"},{"char":"a"},{"char":"x"},{"char":"r"},{"char":"r"},{"char":"p"},{"char":"e"},{"char":"j"},{"char":"l"},{"char":"n"},{"char":"s"},{"char":"f"},{"char":"z"}]},{"cols":[{"char":"s"},{"char":"c"},{"char":"t"},{"char":"r"},{"char":"q"},{"char":"m"},{"char":"s"},{"char":"w"},{"char":"l"},{"char":"n"},{"char":"p"},{"char":"i"},{"char":"d"},{"char":"z"},{"char":"o"},{"char":"w"},{"char":"z"},{"char":"e"},{"char":"e"},{"char":"h"}]},{"cols":[{"char":"u"},{"char":"e"},{"char":"k"},{"char":"w"},{"char":"a"},{"char":"w"},{"char":"x"},{"char":"i"},{"char":"b"},{"char":"t"},{"char":"w"},{"char":"f"},{"char":"u"},{"char":"t"},{"char":"r"},{"char":"i"},{"char":"l"},{"char":"e"},{"char":"u"},{"char":"y"}]},{"cols":[{"char":"u"},{"char":"q"},{"char":"v"},{"char":"a"},{"char":"l"},{"char":"o"},{"char":"s"},{"char":"h"},{"char":"h"},{"char":"l"},{"char":"u"},{"char":"m"},{"char":"x"},{"char":"m"},{"char":"q"},{"char":"u"},{"char":"i"},{"char":"e"},{"char":"t"},{"char":"r"}]},{"cols":[{"char":"f"},{"char":"w"},{"char":"b"},{"char":"t"},{"char":"z"},{"char":"g"},{"char":"o"},{"char":"j"},{"char":"j"},{"char":"n"},{"char":"h"},{"char":"l"},{"char":"i"},{"char":"i"},{"char":"n"},{"char":"t"},{"char":"x"},{"char":"z"},{"char":"w"},{"char":"g"}]},{"cols":[{"char":"n"},{"char":"g"},{"char":"w"},{"char":"b"},{"char":"e"},{"char":"n"},{"char":"q"},{"char":"k"},{"char":"k"},{"char":"y"},{"char":"b"},{"char":"f"},{"char":"s"},{"char":"f"},{"char":"o"},{"char":"e"},{"char":"q"},{"char":"v"},{"char":"i"},{"char":"b"}]},{"cols":[{"char":"h"},{"char":"h"},{"char":"s"},{"char":"c"},{"char":"g"},{"char":"n"},{"char":"n"},{"char":"i"},{"char":"v"},{"char":"h"},{"char":"g"},{"char":"n"},{"char":"i"},{"char":"g"},{"char":"m"},{"char":"a"},{"char":"g"},{"char":"l"},{"char":"r"},{"char":"x"}]},{"cols":[{"char":"r"},{"char":"y"},{"char":"b"},{"char":"p"},{"char":"g"},{"char":"u"},{"char":"p"},{"char":"p"},{"char":"m"},{"char":"e"},{"char":"f"},{"char":"i"},{"char":"z"},{"char":"h"},{"char":"s"},{"char":"u"},{"char":"u"},{"char":"q"},{"char":"f"},{"char":"b"}]},{"cols":[{"char":"g"},{"char":"v"},{"char":"h"},{"char":"i"},{"char":"i"},{"char":"m"},{"char":"u"},{"char":"e"},{"char":"i"},{"char":"e"},{"char":"p"},{"char":"w"},{"char":"v"},{"char":"h"},{"char":"s"},{"char":"l"},{"char":"s"},{"char":"u"},{"char":"e"},{"char":"i"}]},{"cols":[{"char":"d"},{"char":"g"},{"char":"i"},{"char":"b"},{"char":"r"},{"char":"m"},{"char":"x"},{"char":"f"},{"char":"o"},{"char":"h"},{"char":"k"},{"char":"g"},{"char":"e"},{"char":"x"},{"char":"t"},{"char":"l"},{"char":"l"},{"char":"e"},{"char":"h"},{"char":"x"}]},{"cols":[{"char":"f"},{"char":"p"},{"char":"s"},{"char":"f"},{"char":"l"},{"char":"t"},{"char":"o"},{"char":"n"},{"char":"b"},{"char":"a"},{"char":"l"},{"char":"n"},{"char":"o"},{"char":"k"},{"char":"a"},{"char":"t"},{"char":"k"},{"char":"n"},{"char":"z"},{"char":"j"}]},{"cols":[{"char":"c"},{"char":"s"},{"char":"w"},{"char":"j"},{"char":"v"},{"char":"b"},{"char":"a"},{"char":"t"},{"char":"i"},{"char":"q"},{"char":"k"},{"char":"o"},{"char":"t"},{"char":"r"},{"char":"u"},{"char":"j"},{"char":"a"},{"char":"w"},{"char":"u"},{"char":"j"}]},{"cols":[{"char":"w"},{"char":"y"},{"char":"p"},{"char":"s"},{"char":"t"},{"char":"x"},{"char":"b"},{"char":"o"},{"char":"y"},{"char":"y"},{"char":"p"},{"char":"s"},{"char":"r"},{"char":"p"},{"char":"i"},{"char":"r"},{"char":"q"},{"char":"e"},{"char":"u"},{"char":"y"}]},{"cols":[{"char":"v"},{"char":"d"},{"char":"t"},{"char":"b"},{"char":"z"},{"char":"q"},{"char":"k"},{"char":"d"},{"char":"p"},{"char":"w"},{"char":"x"},{"char":"p"},{"char":"r"},{"char":"q"},{"char":"c"},{"char":"h"},{"char":"a"},{"char":"c"},{"char":"g"},{"char":"d"}]},{"cols":[{"char":"h"},{"char":"z"},{"char":"h"},{"char":"q"},{"char":"u"},{"char":"t"},{"char":"i"},{"char":"o"},{"char":"r"},{"char":"h"},{"char":"c"},{"char":"c"},{"char":"b"},{"char":"q"},{"char":"v"},{"char":"t"},{"char":"v"},{"char":"r"},{"char":"f"},{"char":"c"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)

	//act
	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 6)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert

	fmt.Println("valid words...")
	for _, word := range validWords {
		validWordMap.Set(word, struct{}{})

		fmt.Println(word)
	}

	if !validWordMap.Has("abbe") ||
		!validWordMap.Has("abbot") ||
		!validWordMap.Has("ace") ||
		!validWordMap.Has("achoo") ||
		!validWordMap.Has("achy") ||
		!validWordMap.Has("aeon") ||
		!validWordMap.Has("age") ||
		!validWordMap.Has("agent") ||
		!validWordMap.Has("ahem") ||
		!validWordMap.Has("ail") ||
		!validWordMap.Has("air") ||
		!validWordMap.Has("airs") ||
		!validWordMap.Has("airy") ||
		!validWordMap.Has("ajar") ||
		!validWordMap.Has("alb") ||
		!validWordMap.Has("all") ||
		!validWordMap.Has("alms") ||
		!validWordMap.Has("alone") ||
		!validWordMap.Has("along") ||
		!validWordMap.Has("altar") ||
		!validWordMap.Has("alum") ||
		!validWordMap.Has("amen") ||
		!validWordMap.Has("amuse") ||
		!validWordMap.Has("ankle") ||
		!validWordMap.Has("ant") ||
		!validWordMap.Has("anti") ||
		!validWordMap.Has("ape") ||
		!validWordMap.Has("arc") ||
		!validWordMap.Has("arch") ||
		!validWordMap.Has("argue") ||
		!validWordMap.Has("arid") ||
		!validWordMap.Has("ark") ||
		!validWordMap.Has("arm") ||
		!validWordMap.Has("arms") ||
		!validWordMap.Has("art") ||
		!validWordMap.Has("arts") ||
		!validWordMap.Has("ash") ||
		!validWordMap.Has("ashen") ||
		!validWordMap.Has("ass") ||
		!validWordMap.Has("ate") ||
		!validWordMap.Has("atom") ||
		!validWordMap.Has("atop") ||
		!validWordMap.Has("attar") ||
		!validWordMap.Has("auk") ||
		!validWordMap.Has("aura") ||
		!validWordMap.Has("awake") ||
		!validWordMap.Has("awe") ||
		!validWordMap.Has("awed") ||
		!validWordMap.Has("awl") ||
		!validWordMap.Has("awn") ||
		!validWordMap.Has("awry") ||
		!validWordMap.Has("axe") ||
		!validWordMap.Has("axon") ||
		!validWordMap.Has("baa") ||
		!validWordMap.Has("bah") ||
		!validWordMap.Has("bait") ||
		!validWordMap.Has("bake") ||
		!validWordMap.Has("balk") ||
		!validWordMap.Has("balky") ||
		!validWordMap.Has("ban") ||
		!validWordMap.Has("bat") ||
		!validWordMap.Has("bate") ||
		!validWordMap.Has("baton") ||
		!validWordMap.Has("bawl") ||
		!validWordMap.Has("bee") ||
		!validWordMap.Has("been") ||
		!validWordMap.Has("beg") ||
		!validWordMap.Has("begun") ||
		!validWordMap.Has("bet") ||
		!validWordMap.Has("beta") ||
		!validWordMap.Has("bib") ||
		!validWordMap.Has("big") ||
		!validWordMap.Has("bilge") ||
		!validWordMap.Has("bin") ||
		!validWordMap.Has("bio") ||
		!validWordMap.Has("bison") ||
		!validWordMap.Has("bit") ||
		!validWordMap.Has("biz") ||
		!validWordMap.Has("bled") ||
		!validWordMap.Has("blew") ||
		!validWordMap.Has("blue") ||
		!validWordMap.Has("boa") ||
		!validWordMap.Has("boat") ||
		!validWordMap.Has("bod") ||
		!validWordMap.Has("body") ||
		!validWordMap.Has("bop") ||
		!validWordMap.Has("box") ||
		!validWordMap.Has("boy") ||
		!validWordMap.Has("brig") ||
		!validWordMap.Has("brim") ||
		!validWordMap.Has("brr") ||
		!validWordMap.Has("bub") ||
		!validWordMap.Has("bulb") ||
		!validWordMap.Has("bull") ||
		!validWordMap.Has("bun") ||
		!validWordMap.Has("bung") ||
		!validWordMap.Has("but") ||
		!validWordMap.Has("can") ||
		!validWordMap.Has("cant") ||
		!validWordMap.Has("cap") ||
		!validWordMap.Has("car") ||
		!validWordMap.Has("carat") ||
		!validWordMap.Has("cat") ||
		!validWordMap.Has("catch") ||
		!validWordMap.Has("caw") ||
		!validWordMap.Has("char") ||
		!validWordMap.Has("chat") ||
		!validWordMap.Has("chi") ||
		!validWordMap.Has("chip") ||
		!validWordMap.Has("chirp") ||
		!validWordMap.Has("choir") ||
		!validWordMap.Has("coal") ||
		!validWordMap.Has("coax") ||
		!validWordMap.Has("coil") ||
		!validWordMap.Has("coo") ||
		!validWordMap.Has("cool") ||
		!validWordMap.Has("coot") ||
		!validWordMap.Has("coral") ||
		!validWordMap.Has("corm") ||
		!validWordMap.Has("crane") ||
		!validWordMap.Has("crank") ||
		!validWordMap.Has("craw") ||
		!validWordMap.Has("cry") ||
		!validWordMap.Has("cud") ||
		!validWordMap.Has("cue") ||
		!validWordMap.Has("dell") ||
		!validWordMap.Has("dew") ||
		!validWordMap.Has("dip") ||
		!validWordMap.Has("ditz") ||
		!validWordMap.Has("dopy") ||
		!validWordMap.Has("dot") ||
		!validWordMap.Has("drape") ||
		!validWordMap.Has("draw") ||
		!validWordMap.Has("drawl") ||
		!validWordMap.Has("drew") ||
		!validWordMap.Has("drip") ||
		!validWordMap.Has("drool") ||
		!validWordMap.Has("drop") ||
		!validWordMap.Has("drum") ||
		!validWordMap.Has("dry") ||
		!validWordMap.Has("ducat") ||
		!validWordMap.Has("due") ||
		!validWordMap.Has("dug") ||
		!validWordMap.Has("duly") ||
		!validWordMap.Has("dumb") ||
		!validWordMap.Has("ear") ||
		!validWordMap.Has("eat") ||
		!validWordMap.Has("ebb") ||
		!validWordMap.Has("eek") ||
		!validWordMap.Has("eel") ||
		!validWordMap.Has("egg") ||
		!validWordMap.Has("ego") ||
		!validWordMap.Has("elite") ||
		!validWordMap.Has("elk") ||
		!validWordMap.Has("ell") ||
		!validWordMap.Has("emu") ||
		!validWordMap.Has("ennui") ||
		!validWordMap.Has("eon") ||
		!validWordMap.Has("erst") ||
		!validWordMap.Has("eta") ||
		!validWordMap.Has("exalt") ||
		!validWordMap.Has("exile") ||
		!validWordMap.Has("exit") ||
		!validWordMap.Has("fax") ||
		!validWordMap.Has("fee") ||
		!validWordMap.Has("feel") ||
		!validWordMap.Has("feet") ||
		!validWordMap.Has("fen") ||
		!validWordMap.Has("fey") ||
		!validWordMap.Has("fez") ||
		!validWordMap.Has("fib") ||
		!validWordMap.Has("fie") ||
		!validWordMap.Has("fief") ||
		!validWordMap.Has("fig") ||
		!validWordMap.Has("film") ||
		!validWordMap.Has("fin") ||
		!validWordMap.Has("fine") ||
		!validWordMap.Has("finis") ||
		!validWordMap.Has("fir") ||
		!validWordMap.Has("five") ||
		!validWordMap.Has("fix") ||
		!validWordMap.Has("flu") ||
		!validWordMap.Has("flunk") ||
		!validWordMap.Has("foal") ||
		!validWordMap.Has("foam") ||
		!validWordMap.Has("fob") ||
		!validWordMap.Has("foe") ||
		!validWordMap.Has("fog") ||
		!validWordMap.Has("font") ||
		!validWordMap.Has("fox") ||
		!validWordMap.Has("frat") ||
		!validWordMap.Has("fuel") ||
		!validWordMap.Has("full") ||
		!validWordMap.Has("fun") ||
		!validWordMap.Has("funk") ||
		!validWordMap.Has("funky") ||
		!validWordMap.Has("fur") ||
		!validWordMap.Has("furl") ||
		!validWordMap.Has("furor") ||
		!validWordMap.Has("fury") ||
		!validWordMap.Has("fuse") ||
		!validWordMap.Has("fusee") ||
		!validWordMap.Has("futz") ||
		!validWordMap.Has("fuze") ||
		!validWordMap.Has("game") ||
		!validWordMap.Has("gas") ||
		!validWordMap.Has("gash") ||
		!validWordMap.Has("gear") ||
		!validWordMap.Has("gee") ||
		!validWordMap.Has("geek") ||
		!validWordMap.Has("gem") ||
		!validWordMap.Has("genii") ||
		!validWordMap.Has("gent") ||
		!validWordMap.Has("get") ||
		!validWordMap.Has("gimp") ||
		!validWordMap.Has("gin") ||
		!validWordMap.Has("girl") ||
		!validWordMap.Has("girt") ||
		!validWordMap.Has("give") ||
		!validWordMap.Has("given") ||
		!validWordMap.Has("glaze") ||
		!validWordMap.Has("glib") ||
		!validWordMap.Has("glop") ||
		!validWordMap.Has("glow") ||
		!validWordMap.Has("gnu") ||
		!validWordMap.Has("goal") ||
		!validWordMap.Has("gone") ||
		!validWordMap.Has("gong") ||
		!validWordMap.Has("goo") ||
		!validWordMap.Has("gook") ||
		!validWordMap.Has("goon") ||
		!validWordMap.Has("goop") ||
		!validWordMap.Has("gorp") ||
		!validWordMap.Has("gosh") ||
		!validWordMap.Has("got") ||
		!validWordMap.Has("grace") ||
		!validWordMap.Has("gruel") ||
		!validWordMap.Has("gulf") ||
		!validWordMap.Has("gull") ||
		!validWordMap.Has("gum") ||
		!validWordMap.Has("gums") ||
		!validWordMap.Has("gun") ||
		!validWordMap.Has("gunk") ||
		!validWordMap.Has("gush") ||
		!validWordMap.Has("gust") ||
		!validWordMap.Has("guy") ||
		!validWordMap.Has("gyve") ||
		!validWordMap.Has("habit") ||
		!validWordMap.Has("hake") ||
		!validWordMap.Has("halo") ||
		!validWordMap.Has("hat") ||
		!validWordMap.Has("heel") ||
		!validWordMap.Has("heels") ||
		!validWordMap.Has("hell") ||
		!validWordMap.Has("hello") ||
		!validWordMap.Has("hem") ||
		!validWordMap.Has("heme") ||
		!validWordMap.Has("hemp") ||
		!validWordMap.Has("hen") ||
		!validWordMap.Has("hep") ||
		!validWordMap.Has("hew") ||
		!validWordMap.Has("hex") ||
		!validWordMap.Has("hey") ||
		!validWordMap.Has("hie") ||
		!validWordMap.Has("hilt") ||
		!validWordMap.Has("him") ||
		!validWordMap.Has("hinge") ||
		!validWordMap.Has("hip") ||
		!validWordMap.Has("his") ||
		!validWordMap.Has("hiss") ||
		!validWordMap.Has("hoax") ||
		!validWordMap.Has("hob") ||
		!validWordMap.Has("hoe") ||
		!validWordMap.Has("hog") ||
		!validWordMap.Has("hon") ||
		!validWordMap.Has("hone") ||
		!validWordMap.Has("hoot") ||
		!validWordMap.Has("hora") ||
		!validWordMap.Has("hot") ||
		!validWordMap.Has("hots") ||
		!validWordMap.Has("hue") ||
		!validWordMap.Has("hulk") ||
		!validWordMap.Has("hull") ||
		!validWordMap.Has("hum") ||
		!validWordMap.Has("hunk") ||
		!validWordMap.Has("hunky") ||
		!validWordMap.Has("hurt") ||
		!validWordMap.Has("hut") ||
		!validWordMap.Has("ibex") ||
		!validWordMap.Has("ibis") ||
		!validWordMap.Has("ikon") ||
		!validWordMap.Has("ilk") ||
		!validWordMap.Has("imp") ||
		!validWordMap.Has("info") ||
		!validWordMap.Has("ink") ||
		!validWordMap.Has("inn") ||
		!validWordMap.Has("into") ||
		!validWordMap.Has("ion") ||
		!validWordMap.Has("iota") ||
		!validWordMap.Has("irk") ||
		!validWordMap.Has("iron") ||
		!validWordMap.Has("ism") ||
		!validWordMap.Has("item") ||
		!validWordMap.Has("ivy") ||
		!validWordMap.Has("jar") ||
		!validWordMap.Has("jaw") ||
		!validWordMap.Has("jaws") ||
		!validWordMap.Has("jell") ||
		!validWordMap.Has("jello") ||
		!validWordMap.Has("jew") ||
		!validWordMap.Has("jog") ||
		!validWordMap.Has("josh") ||
		!validWordMap.Has("jot") ||
		!validWordMap.Has("jowl") ||
		!validWordMap.Has("joy") ||
		!validWordMap.Has("judge") ||
		!validWordMap.Has("jug") ||
		!validWordMap.Has("junk") ||
		!validWordMap.Has("jut") ||
		!validWordMap.Has("karat") ||
		!validWordMap.Has("karma") ||
		!validWordMap.Has("kart") ||
		!validWordMap.Has("kayo") ||
		!validWordMap.Has("kazoo") ||
		!validWordMap.Has("keep") ||
		!validWordMap.Has("keg") ||
		!validWordMap.Has("ken") ||
		!validWordMap.Has("keno") ||
		!validWordMap.Has("kepi") ||
		!validWordMap.Has("kid") ||
		!validWordMap.Has("kin") ||
		!validWordMap.Has("kink") ||
		!validWordMap.Has("kit") ||
		!validWordMap.Has("knee") ||
		!validWordMap.Has("knell") ||
		!validWordMap.Has("knelt") ||
		!validWordMap.Has("knew") ||
		!validWordMap.Has("knot") ||
		!validWordMap.Has("kola") ||
		!validWordMap.Has("kook") ||
		!validWordMap.Has("kooky") ||
		!validWordMap.Has("kraal") ||
		!validWordMap.Has("krone") ||
		!validWordMap.Has("lab") ||
		!validWordMap.Has("lac") ||
		!validWordMap.Has("lain") ||
		!validWordMap.Has("lair") ||
		!validWordMap.Has("laity") ||
		!validWordMap.Has("lake") ||
		!validWordMap.Has("lam") ||
		!validWordMap.Has("lama") ||
		!validWordMap.Has("lap") ||
		!validWordMap.Has("lapel") ||
		!validWordMap.Has("lard") ||
		!validWordMap.Has("lark") ||
		!validWordMap.Has("late") ||
		!validWordMap.Has("lath") ||
		!validWordMap.Has("lathe") ||
		!validWordMap.Has("lave") ||
		!validWordMap.Has("law") ||
		!validWordMap.Has("lax") ||
		!validWordMap.Has("lay") ||
		!validWordMap.Has("laze") ||
		!validWordMap.Has("led") ||
		!validWordMap.Has("lee") ||
		!validWordMap.Has("lees") ||
		!validWordMap.Has("lei") ||
		!validWordMap.Has("lens") ||
		!validWordMap.Has("let") ||
		!validWordMap.Has("liar") ||
		!validWordMap.Has("lib") ||
		!validWordMap.Has("lie") ||
		!validWordMap.Has("lieu") ||
		!validWordMap.Has("limn") ||
		!validWordMap.Has("lira") ||
		!validWordMap.Has("lit") ||
		!validWordMap.Has("lite") ||
		!validWordMap.Has("loam") ||
		!validWordMap.Has("log") ||
		!validWordMap.Has("loge") ||
		!validWordMap.Has("logo") ||
		!validWordMap.Has("lone") ||
		!validWordMap.Has("long") ||
		!validWordMap.Has("look") ||
		!validWordMap.Has("loon") ||
		!validWordMap.Has("loot") ||
		!validWordMap.Has("lop") ||
		!validWordMap.Has("lord") ||
		!validWordMap.Has("lorn") ||
		!validWordMap.Has("lost") ||
		!validWordMap.Has("lot") ||
		!validWordMap.Has("loth") ||
		!validWordMap.Has("lots") ||
		!validWordMap.Has("low") ||
		!validWordMap.Has("lox") ||
		!validWordMap.Has("loyal") ||
		!validWordMap.Has("luau") ||
		!validWordMap.Has("lube") ||
		!validWordMap.Has("lug") ||
		!validWordMap.Has("lull") ||
		!validWordMap.Has("lung") ||
		!validWordMap.Has("lurk") ||
		!validWordMap.Has("lush") ||
		!validWordMap.Has("lust") ||
		!validWordMap.Has("lute") ||
		!validWordMap.Has("mag") ||
		!validWordMap.Has("mail") ||
		!validWordMap.Has("malt") ||
		!validWordMap.Has("mar") ||
		!validWordMap.Has("mark") ||
		!validWordMap.Has("marl") ||
		!validWordMap.Has("mart") ||
		!validWordMap.Has("mash") ||
		!validWordMap.Has("mass") ||
		!validWordMap.Has("maul") ||
		!validWordMap.Has("maw") ||
		!validWordMap.Has("may") ||
		!validWordMap.Has("mayo") ||
		!validWordMap.Has("meek") ||
		!validWordMap.Has("men") ||
		!validWordMap.Has("menu") ||
		!validWordMap.Has("met") ||
		!validWordMap.Has("mil") ||
		!validWordMap.Has("mine") ||
		!validWordMap.Has("mink") ||
		!validWordMap.Has("mint") ||
		!validWordMap.Has("mix") ||
		!validWordMap.Has("moan") ||
		!validWordMap.Has("moat") ||
		!validWordMap.Has("mob") ||
		!validWordMap.Has("mom") ||
		!validWordMap.Has("mono") ||
		!validWordMap.Has("most") ||
		!validWordMap.Has("mot") ||
		!validWordMap.Has("mote") ||
		!validWordMap.Has("mow") ||
		!validWordMap.Has("mud") ||
		!validWordMap.Has("mug") ||
		!validWordMap.Has("mull") ||
		!validWordMap.Has("mum") ||
		!validWordMap.Has("murk") ||
		!validWordMap.Has("murky") ||
		!validWordMap.Has("muse") ||
		!validWordMap.Has("mush") ||
		!validWordMap.Has("muss") ||
		!validWordMap.Has("must") ||
		!validWordMap.Has("nab") ||
		!validWordMap.Has("nacho") ||
		!validWordMap.Has("nap") ||
		!validWordMap.Has("narc") ||
		!validWordMap.Has("natal") ||
		!validWordMap.Has("nee") ||
		!validWordMap.Has("net") ||
		!validWordMap.Has("nevi") ||
		!validWordMap.Has("new") ||
		!validWordMap.Has("next") ||
		!validWordMap.Has("nib") ||
		!validWordMap.Has("nigh") ||
		!validWordMap.Has("nip") ||
		!validWordMap.Has("nit") ||
		!validWordMap.Has("nix") ||
		!validWordMap.Has("nook") ||
		!validWordMap.Has("nor") ||
		!validWordMap.Has("nosh") ||
		!validWordMap.Has("not") ||
		!validWordMap.Has("note") ||
		!validWordMap.Has("nth") ||
		!validWordMap.Has("nub") ||
		!validWordMap.Has("null") ||
		!validWordMap.Has("nun") ||
		!validWordMap.Has("nut") ||
		!validWordMap.Has("oak") ||
		!validWordMap.Has("oar") ||
		!validWordMap.Has("oat") ||
		!validWordMap.Has("obi") ||
		!validWordMap.Has("obit") ||
		!validWordMap.Has("odor") ||
		!validWordMap.Has("oil") ||
		!validWordMap.Has("oils") ||
		!validWordMap.Has("oily") ||
		!validWordMap.Has("okra") ||
		!validWordMap.Has("omega") ||
		!validWordMap.Has("omen") ||
		!validWordMap.Has("one") ||
		!validWordMap.Has("onto") ||
		!validWordMap.Has("ooh") ||
		!validWordMap.Has("oops") ||
		!validWordMap.Has("ooze") ||
		!validWordMap.Has("oral") ||
		!validWordMap.Has("owl") ||
		!validWordMap.Has("owlet") ||
		!validWordMap.Has("own") ||
		!validWordMap.Has("oxen") ||
		!validWordMap.Has("pal") ||
		!validWordMap.Has("pan") ||
		!validWordMap.Has("pant") ||
		!validWordMap.Has("par") ||
		!validWordMap.Has("paw") ||
		!validWordMap.Has("pawl") ||
		!validWordMap.Has("pee") ||
		!validWordMap.Has("peek") ||
		!validWordMap.Has("peg") ||
		!validWordMap.Has("peon") ||
		!validWordMap.Has("pep") ||
		!validWordMap.Has("per") ||
		!validWordMap.Has("pert") ||
		!validWordMap.Has("pew") ||
		!validWordMap.Has("phew") ||
		!validWordMap.Has("phi") ||
		!validWordMap.Has("pic") ||
		!validWordMap.Has("pie") ||
		!validWordMap.Has("pig") ||
		!validWordMap.Has("piggy") ||
		!validWordMap.Has("pimp") ||
		!validWordMap.Has("pin") ||
		!validWordMap.Has("ping") ||
		!validWordMap.Has("pink") ||
		!validWordMap.Has("pinko") ||
		!validWordMap.Has("pinup") ||
		!validWordMap.Has("pip") ||
		!validWordMap.Has("pipe") ||
		!validWordMap.Has("pis") ||
		!validWordMap.Has("piss") ||
		!validWordMap.Has("pod") ||
		!validWordMap.Has("poi") ||
		!validWordMap.Has("point") ||
		!validWordMap.Has("poky") ||
		!validWordMap.Has("pol") ||
		!validWordMap.Has("polka") ||
		!validWordMap.Has("pone") ||
		!validWordMap.Has("poor") ||
		!validWordMap.Has("port") ||
		!validWordMap.Has("post") ||
		!validWordMap.Has("pot") ||
		!validWordMap.Has("pride") ||
		!validWordMap.Has("pro") ||
		!validWordMap.Has("prod") ||
		!validWordMap.Has("prone") ||
		!validWordMap.Has("prong") ||
		!validWordMap.Has("prop") ||
		!validWordMap.Has("pros") ||
		!validWordMap.Has("pry") ||
		!validWordMap.Has("psi") ||
		!validWordMap.Has("pug") ||
		!validWordMap.Has("puke") ||
		!validWordMap.Has("pump") ||
		!validWordMap.Has("pun") ||
		!validWordMap.Has("punk") ||
		!validWordMap.Has("pup") ||
		!validWordMap.Has("purr") ||
		!validWordMap.Has("put") ||
		!validWordMap.Has("putt") ||
		!validWordMap.Has("pyx") ||
		!validWordMap.Has("qua") ||
		!validWordMap.Has("quash") ||
		!validWordMap.Has("queen") ||
		!validWordMap.Has("quell") ||
		!validWordMap.Has("quiet") ||
		!validWordMap.Has("quirt") ||
		!validWordMap.Has("quit") ||
		!validWordMap.Has("quite") ||
		!validWordMap.Has("quiz") ||
		!validWordMap.Has("race") ||
		!validWordMap.Has("rah") ||
		!validWordMap.Has("rail") ||
		!validWordMap.Has("raja") ||
		!validWordMap.Has("rake") ||
		!validWordMap.Has("ram") ||
		!validWordMap.Has("ran") ||
		!validWordMap.Has("ranee") ||
		!validWordMap.Has("rank") ||
		!validWordMap.Has("rap") ||
		!validWordMap.Has("rape") ||
		!validWordMap.Has("rapid") ||
		!validWordMap.Has("rat") ||
		!validWordMap.Has("rats") ||
		!validWordMap.Has("raw") ||
		!validWordMap.Has("ray") ||
		!validWordMap.Has("red") ||
		!validWordMap.Has("rep") ||
		!validWordMap.Has("rev") ||
		!validWordMap.Has("rib") ||
		!validWordMap.Has("rich") ||
		!validWordMap.Has("rid") ||
		!validWordMap.Has("ride") ||
		!validWordMap.Has("rig") ||
		!validWordMap.Has("rile") ||
		!validWordMap.Has("rim") ||
		!validWordMap.Has("riot") ||
		!validWordMap.Has("rip") ||
		!validWordMap.Has("ripe") ||
		!validWordMap.Has("roach") ||
		!validWordMap.Has("roan") ||
		!validWordMap.Has("rod") ||
		!validWordMap.Has("roe") ||
		!validWordMap.Has("roil") ||
		!validWordMap.Has("roll") ||
		!validWordMap.Has("rook") ||
		!validWordMap.Has("roost") ||
		!validWordMap.Has("root") ||
		!validWordMap.Has("roots") ||
		!validWordMap.Has("rot") ||
		!validWordMap.Has("rotor") ||
		!validWordMap.Has("row") ||
		!validWordMap.Has("rue") ||
		!validWordMap.Has("rule") ||
		!validWordMap.Has("rum") ||
		!validWordMap.Has("run") ||
		!validWordMap.Has("rune") ||
		!validWordMap.Has("runt") ||
		!validWordMap.Has("rural") ||
		!validWordMap.Has("rut") ||
		!validWordMap.Has("rye") ||
		!validWordMap.Has("sag") ||
		!validWordMap.Has("sage") ||
		!validWordMap.Has("sail") ||
		!validWordMap.Has("same") ||
		!validWordMap.Has("sari") ||
		!validWordMap.Has("say") ||
		!validWordMap.Has("sec") ||
		!validWordMap.Has("sect") ||
		!validWordMap.Has("see") ||
		!validWordMap.Has("seen") ||
		!validWordMap.Has("sell") ||
		!validWordMap.Has("set") ||
		!validWordMap.Has("she") ||
		!validWordMap.Has("shew") ||
		!validWordMap.Has("shh") ||
		!validWordMap.Has("shin") ||
		!validWordMap.Has("shone") ||
		!validWordMap.Has("shoo") ||
		!validWordMap.Has("shot") ||
		!validWordMap.Has("shy") ||
		!validWordMap.Has("sigh") ||
		!validWordMap.Has("sigma") ||
		!validWordMap.Has("silt") ||
		!validWordMap.Has("sin") ||
		!validWordMap.Has("sine") ||
		!validWordMap.Has("sing") ||
		!validWordMap.Has("singe") ||
		!validWordMap.Has("sip") ||
		!validWordMap.Has("sis") ||
		!validWordMap.Has("sisal") ||
		!validWordMap.Has("six") ||
		!validWordMap.Has("ska") ||
		!validWordMap.Has("sky") ||
		!validWordMap.Has("slake") ||
		!validWordMap.Has("slap") ||
		!validWordMap.Has("slat") ||
		!validWordMap.Has("slaw") ||
		!validWordMap.Has("sled") ||
		!validWordMap.Has("slew") ||
		!validWordMap.Has("slim") ||
		!validWordMap.Has("slot") ||
		!validWordMap.Has("sloth") ||
		!validWordMap.Has("slue") ||
		!validWordMap.Has("slug") ||
		!validWordMap.Has("slum") ||
		!validWordMap.Has("slunk") ||
		!validWordMap.Has("slush") ||
		!validWordMap.Has("slut") ||
		!validWordMap.Has("sly") ||
		!validWordMap.Has("smart") ||
		!validWordMap.Has("smog") ||
		!validWordMap.Has("smote") ||
		!validWordMap.Has("smug") ||
		!validWordMap.Has("snip") ||
		!validWordMap.Has("snipe") ||
		!validWordMap.Has("snoot") ||
		!validWordMap.Has("soar") ||
		!validWordMap.Has("sob") ||
		!validWordMap.Has("soil") ||
		!validWordMap.Has("sol") ||
		!validWordMap.Has("solar") ||
		!validWordMap.Has("solo") ||
		!validWordMap.Has("son") ||
		!validWordMap.Has("song") ||
		!validWordMap.Has("soon") ||
		!validWordMap.Has("soot") ||
		!validWordMap.Has("sooth") ||
		!validWordMap.Has("sop") ||
		!validWordMap.Has("sorry") ||
		!validWordMap.Has("sort") ||
		!validWordMap.Has("sot") ||
		!validWordMap.Has("sow") ||
		!validWordMap.Has("sox") ||
		!validWordMap.Has("soy") ||
		!validWordMap.Has("spook") ||
		!validWordMap.Has("spoon") ||
		!validWordMap.Has("spoor") ||
		!validWordMap.Has("sport") ||
		!validWordMap.Has("spot") ||
		!validWordMap.Has("spry") ||
		!validWordMap.Has("spy") ||
		!validWordMap.Has("ssh") ||
		!validWordMap.Has("stake") ||
		!validWordMap.Has("stalk") ||
		!validWordMap.Has("stall") ||
		!validWordMap.Has("star") ||
		!validWordMap.Has("stark") ||
		!validWordMap.Has("start") ||
		!validWordMap.Has("stat") ||
		!validWordMap.Has("stoke") ||
		!validWordMap.Has("stone") ||
		!validWordMap.Has("stool") ||
		!validWordMap.Has("stoop") ||
		!validWordMap.Has("stop") ||
		!validWordMap.Has("stork") ||
		!validWordMap.Has("strep") ||
		!validWordMap.Has("strew") ||
		!validWordMap.Has("strip") ||
		!validWordMap.Has("strop") ||
		!validWordMap.Has("strut") ||
		!validWordMap.Has("sty") ||
		!validWordMap.Has("suck") ||
		!validWordMap.Has("sue") ||
		!validWordMap.Has("suet") ||
		!validWordMap.Has("sulk") ||
		!validWordMap.Has("sum") ||
		!validWordMap.Has("sumo") ||
		!validWordMap.Has("sushi") ||
		!validWordMap.Has("swam") ||
		!validWordMap.Has("swan") ||
		!validWordMap.Has("swap") ||
		!validWordMap.Has("swish") ||
		!validWordMap.Has("swoon") ||
		!validWordMap.Has("tab") ||
		!validWordMap.Has("take") ||
		!validWordMap.Has("taken") ||
		!validWordMap.Has("talk") ||
		!validWordMap.Has("tall") ||
		!validWordMap.Has("tan") ||
		!validWordMap.Has("tank") ||
		!validWordMap.Has("tar") ||
		!validWordMap.Has("taro") ||
		!validWordMap.Has("tarot") ||
		!validWordMap.Has("tarp") ||
		!validWordMap.Has("tart") ||
		!validWordMap.Has("tat") ||
		!validWordMap.Has("tau") ||
		!validWordMap.Has("taut") ||
		!validWordMap.Has("tax") ||
		!validWordMap.Has("tea") ||
		!validWordMap.Has("team") ||
		!validWordMap.Has("tee") ||
		!validWordMap.Has("teen") ||
		!validWordMap.Has("teens") ||
		!validWordMap.Has("telex") ||
		!validWordMap.Has("ten") ||
		!validWordMap.Has("text") ||
		!validWordMap.Has("the") ||
		!validWordMap.Has("then") ||
		!validWordMap.Has("thew") ||
		!validWordMap.Has("this") ||
		!validWordMap.Has("tho") ||
		!validWordMap.Has("thru") ||
		!validWordMap.Has("thy") ||
		!validWordMap.Has("tidy") ||
		!validWordMap.Has("tie") ||
		!validWordMap.Has("tile") ||
		!validWordMap.Has("tin") ||
		!validWordMap.Has("toe") ||
		!validWordMap.Has("toed") ||
		!validWordMap.Has("tofu") ||
		!validWordMap.Has("tog") ||
		!validWordMap.Has("togs") ||
		!validWordMap.Has("toil") ||
		!validWordMap.Has("toke") ||
		!validWordMap.Has("token") ||
		!validWordMap.Has("toll") ||
		!validWordMap.Has("tom") ||
		!validWordMap.Has("tomb") ||
		!validWordMap.Has("tome") ||
		!validWordMap.Has("ton") ||
		!validWordMap.Has("tone") ||
		!validWordMap.Has("tong") ||
		!validWordMap.Has("too") ||
		!validWordMap.Has("took") ||
		!validWordMap.Has("tool") ||
		!validWordMap.Has("top") ||
		!validWordMap.Has("tops") ||
		!validWordMap.Has("tor") ||
		!validWordMap.Has("tot") ||
		!validWordMap.Has("tow") ||
		!validWordMap.Has("town") ||
		!validWordMap.Has("toy") ||
		!validWordMap.Has("tram") ||
		!validWordMap.Has("trawl") ||
		!validWordMap.Has("trig") ||
		!validWordMap.Has("trim") ||
		!validWordMap.Has("trio") ||
		!validWordMap.Has("trip") ||
		!validWordMap.Has("troll") ||
		!validWordMap.Has("troop") ||
		!validWordMap.Has("trow") ||
		!validWordMap.Has("true") ||
		!validWordMap.Has("try") ||
		!validWordMap.Has("tub") ||
		!validWordMap.Has("tun") ||
		!validWordMap.Has("tune") ||
		!validWordMap.Has("tux") ||
		!validWordMap.Has("twig") ||
		!validWordMap.Has("twirl") ||
		!validWordMap.Has("typo") ||
		!validWordMap.Has("ulna") ||
		!validWordMap.Has("ulnae") ||
		!validWordMap.Has("ulnar") ||
		!validWordMap.Has("ump") ||
		!validWordMap.Has("unfix") ||
		!validWordMap.Has("unpin") ||
		!validWordMap.Has("untie") ||
		!validWordMap.Has("until") ||
		!validWordMap.Has("unto") ||
		!validWordMap.Has("uric") ||
		!validWordMap.Has("usage") ||
		!validWordMap.Has("use") ||
		!validWordMap.Has("vapid") ||
		!validWordMap.Has("vat") ||
		!validWordMap.Has("veep") ||
		!validWordMap.Has("veg") ||
		!validWordMap.Has("vent") ||
		!validWordMap.Has("verso") ||
		!validWordMap.Has("vet") ||
		!validWordMap.Has("vex") ||
		!validWordMap.Has("vim") ||
		!validWordMap.Has("wake") ||
		!validWordMap.Has("waken") ||
		!validWordMap.Has("waltz") ||
		!validWordMap.Has("wan") ||
		!validWordMap.Has("wane") ||
		!validWordMap.Has("want") ||
		!validWordMap.Has("war") ||
		!validWordMap.Has("ward") ||
		!validWordMap.Has("warn") ||
		!validWordMap.Has("warp") ||
		!validWordMap.Has("wart") ||
		!validWordMap.Has("watt") ||
		!validWordMap.Has("wave") ||
		!validWordMap.Has("wax") ||
		!validWordMap.Has("weak") ||
		!validWordMap.Has("wean") ||
		!validWordMap.Has("wear") ||
		!validWordMap.Has("wed") ||
		!validWordMap.Has("wee") ||
		!validWordMap.Has("well") ||
		!validWordMap.Has("wen") ||
		!validWordMap.Has("went") ||
		!validWordMap.Has("wet") ||
		!validWordMap.Has("why") ||
		!validWordMap.Has("wide") ||
		!validWordMap.Has("wife") ||
		!validWordMap.Has("wig") ||
		!validWordMap.Has("wile") ||
		!validWordMap.Has("wiles") ||
		!validWordMap.Has("wilt") ||
		!validWordMap.Has("win") ||
		!validWordMap.Has("wing") ||
		!validWordMap.Has("wipe") ||
		!validWordMap.Has("wish") ||
		!validWordMap.Has("wive") ||
		!validWordMap.Has("wiz") ||
		!validWordMap.Has("woe") ||
		!validWordMap.Has("woo") ||
		!validWordMap.Has("worm") ||
		!validWordMap.Has("wort") ||
		!validWordMap.Has("wrap") ||
		!validWordMap.Has("wry") ||
		!validWordMap.Has("yak") ||
		!validWordMap.Has("yam") ||
		!validWordMap.Has("yes") ||
		!validWordMap.Has("yin") ||
		!validWordMap.Has("yolk") ||
		!validWordMap.Has("yurt") ||
		!validWordMap.Has("zed") ||
		!validWordMap.Has("zeta") ||
		!validWordMap.Has("zing") ||
		!validWordMap.Has("zingy") ||
		!validWordMap.Has("zip") ||
		!validWordMap.Has("zit") ||
		!validWordMap.Has("zoo") {
		t.Error("Does NOT contain all six letter words")
	}

}
