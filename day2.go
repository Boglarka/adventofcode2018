package main

import (
	"fmt"
	"strings"
)

var input = `mphcuiszrnjzxwkbgdzqeoyxfa
mihcuisgrnjzxwkbgdtqeoylia
mphauisvrnjgxwkbgdtqeiylfa
mphcuisnrnjzxwkbgdgqeoylua
mphcuisurnjzxwkbgdtqeoilfi
mkhcuisvrnjzowkbgdteeoylfa
mphcoicvrnjzxwksgdtqeoylfa
mxhcuisvrndzxwkbgdtqeeylfa
dphcuisijnjzxwkbgdtqeoylfa
mihvuisvrqjzxwkbgdtqeoylfa
mphcuisrrnvzxwkbgdtqeodlfa
mphtuisdrnjzxskbgdtqeoylfa
mphcutmvsnjzxwkbgdtqeoylfa
mphcunsvrnjzswkggdtqeoylfa
mphcuisvrwjzxwkbpdtqeoylfr
mphcujsdrnjzxwkbgdtqeovlfa
mpfcuisvrdjzxwkbgdtteoylfa
mppcuisvrpjzxwkbgdtqeoywfa
mphcuisvrnjzxwkbfptqroylfa
mphcuisvrnjzxwkbgstoeoysfa
mphcufsvrnjzcwkbgdeqeoylfa
mphcuissrnjzxwkbgdkquoylfa
sphcuxsvrnjzxwkbgdtqioylfa
mphcuiivrhjzxwkbgdtqevylfa
echcuisvrnjzxwkbgltqeoylfa
mphcuisvrljexwkbvdtqeoylfa
mpjcuisvrnjzxwkhidtqeoylfa
mphcuisvrfjzmwkbgdtqeoylfl
mwhcuisvrnjzxwkbgdtqeoytfm
mphcuisvrsjzxwkbgdaqeoylfh
mohcuisvrnjzxwkbgdtqtoymfa
maycuisvrnjzxwkbgdtqboylfa
pphcuisvqnjzxwkbgdtqeoylfd
mprcuisvrnjtxwmbgdtqeoylfa
mfhcuisgrnjzxckbgdtqeoylfa
mphiubsvrnjzxwkbgdtqeoyufa
dphctisvrnjzxwkbgdtqeoylfk
mphcuisvrnjznwksgdtqeoyzfa
mpwcuisvrnjziwkbgdtqaoylfa
mphduzsvrnjznwkbgdtqeoylfa
mphccisvrnjzxwebgdtqeoylqa
xphcuisvrnjzxwkfvdtqeoylfa
mphcupsvrnjzxwkbgdtfeoylpa
mphcuisvrtjzjwkbgdtqeoylfe
mpbcuisvrnjzxwkbgdmieoylfa
mphcuisvrnjzxwkbgjtqetylaa
mphcuisvrnjzxwpbgdtgdoylfa
ophcufsvrqjzxwkbgdtqeoylfa
iphcuhsvrnjzxwkbgetqeoylfa
mphcuisvunjzxwwbgdtqeoylqa
mphcpisvrnjzowkbgdtveoylfa
mphcuisvrnjzxhkbgdtqeotlla
mphcuisvrnjzxwkbodtgeoylha
mphcuisvrjjzxwkbwdtqtoylfa
mphcwisvrnjnxwkbgjtqeoylfa
mplcuicqrnjzxwkbgdtqeoylfa
mphcuisvrnjzxydbgdtqeoylfn
ophckisvrnjzxwkbgdtqeozlfa
mphcuisvrkjzxwkbgdtteoblfa
yphcuisvrnjcxwkbggtqeoylfa
mphcuisvrnazxwfbqdtqeoylfa
mphcuisvrmjzxwkbgdtlwoylfa
mphctksvrnjzxwibgdtqeoylfa
mphcuisprnjzxlebgdtqeoylfa
mphcuisnrnjzxakbgdtueoylfa
mphcuiavrnjoxwtbgdtqeoylfa
nphcuisvrnjzxwkbgdtqzoylfk
mphcuisrrnjmxwkbgdtqdoylfa
mphcuisvrujzxwkvgdtqehylfa
mphcuisvrnfzxwkogdtqebylfa
mphcuisvrnjwdwkbgdtqeoyxfa
mphcuisvrntzxwkrgxtqeoylfa
mpzcuisvrnjzxwebgdtqeoylsa
aphcuikvrnjzxwwbgdtqeoylfa
mphcqisvrnjzxwkpgdtqeoelfa
mphcuusvrnjzxwkbgdtjeodlfa
mphcuisvrnjzewkbgdtteoylza
mphcuisvanjzxwkbgdtheoylfc
mphcjishrnjzxwkbgltqeoylfa
mpxcuislrnjzxwkbgdtqeoynfa
mphcuisvrnjjxwkbgdtmeoxlfa
mphcimsvrnjzxwkbsdtqeoylfa
mphcxisvcnjzxwjbgdtqeoylfa
mphcuisbrvjzxwkbgdtqeoymfa
mplcuisvrnjzxwkbgdtaenylfa
mphcuihvrnjzxwkygytqeoylfa
mphcbisvrnjzxhkbgdtqezylfa
mphcuisarnjzxwkbgatqeoylfv
mphcumsvrnjzxwkbgdrqebylfa
mlhcuisvrnwzxwkbgdtqeoylfx
mpkcuisvrkjzxwkbgdtqeoylfo
mphcuissrnjzxwkbgdtqmoylfc
mphcuiwvrnjuxwkfgdtqeoylfa
mphcuicvlnjzxwkbgdvqeoylfa
mphcuisvrvvzxwkbfdtqeoylfa
myhcuisvrnjpxwkbgntqeoylfa
mpocuisvrnjzxwtbgitqeoylfa
mphcuisvrnjzxwkbgdtwewyqfa
mphcuisvtnjzxwwbgdtqeoolfa
mphcuisvrnjzxgkbgdyqeoyyfa
mphcuisvrdjzxwkbgpyqeoylfa
bphcuisvrnjzxwkbgxtqefylfa
sphcuisvrdjzxwktgdtqeoylfa
mphcuvsvrnjmxwobgdtqeoylfa
mphcuisvrnjzxwkbsdtqeuylfb
mnhcmisvynjzxwkbgdtqeoylfa
mphckisvrnjzxwkhgdkqeoylfa
mpacuisvrnjzxwkbgdtqeoolaa
mpgcuisvrnjzxwkbzdtqeoynfa
mphcuisvrojzxwkbzdtqeoylga
mphcuisvknjfxwkbydtqeoylfa
mphcuistrnjzxwkbgdqqeuylfa
bpvcuiszrnjzxwkbgdtqeoylfa
mphcuxsvrnjzswkbgdtqeoelfa
mphcuisvbnjzxwlbgdtqeoylla
mphcuisvonczxwkbgktqeoylfa
mphcuisvrnkzxwvbgdtquoylfa
mphcuisvrnjzxokfgdtqeoylia
tphcuisvrnjzxwkbjdwqeoylfa
mihcuisvrnjzpwibgdtqeoylfa
mphcuisvrejzxwkbgdtqjuylfa
mprcuisvrnjixwkxgdtqeoylfa
mpqcuiszrnjzxwkbgdtqeodlfa
mphcuasvrnjzzakbgdtqeoylva
mphcuisvrnjzmwkbtdtqeoycfa
mphcuisvrnjzxwkbcdtqioylxa
mphckisvrnjzxwkbcdtqeoylfm
mphcuisvrnjuxwbogdtqeoylfa
mphcuisdrnjzxwkbldtqeoylfx
mphcuisvrnjoxwkbgdtqeyyyfa
mphcuicvqnjzxwkbgdtqeoylna
mpmcuisvrnjzxwkbgdtqktylfa
mphcuisvrnqzxwkggdtqeoykfa
mphcuisvryjzxwkbydtqejylfa
mphcugsvrnjzxwkbghtqeeylfa
rphcuusvrnjzxwkwgdtqeoylfa
zphwuiyvrnjzxwkbgdtqeoylfa
cphcuivvrnjzxwkbgdtqenylfa
mphcuisvrnjzxwkagotqevylfa
mprcuisvrcjzxwkbgdtqeoytfa
mphjugsvrnezxwkbgdtqeoylfa
mphcuisvryjzxwkbgltqeoylaa
mphcursvrnjzxfkbgdtqeoydfa
mphcuisvrcuzxwkbgdtqeoylfw
mphcuisvrijzxwkbgdtqeoelfh
xphcuisvenjzxjkbgdtqeoylfa
mphcuisvrnazxwkbgdeqeoylaa
mphcuisbrsjzxwkbgdtqeoygfa
mlhvuisvrnjzxwkbgdtqeoylfh
mphcuisvrnjzxukbgdtqeoyhfy
mpzcuilvrnjzawkbgdtqeoylfa
hphcuisjfnjzxwkbgdtqeoylfa
mahcuisvrnjzxwkegdtqeoylfi
mphcuixvrnjzcwkbgdtqetylfa
mphcuisvrnjzxwkdgdtqeoklfj
mlhcuisvrnjzxwkbgdteeoylka
mphcuifvrnjbxwkrgdtqeoylfa
mphcuasvrnjzzwkbgdtqeoylva
mphcuisvrnjzxwkboutqeoylba
mbhcuisvcnjzxwklgdtqeoylfa
mpbcuisvrnjzxgkbgdtqesylfa
mphcuisvrnjfswkbgdtqeoylfd
mphcuisvrnjzxwkbgdoweoysfa
uphcuisvrnjzrwkbgdtqelylfa
mphcuisvrnjzxwkbgdtqyoylsi
mpqcuiqvxnjzxwkbgdtqeoylfa
mphcuisorfjzxwkbgatqeoylfa
mphcuisvrntfxwkbzdtqeoylfa
mphcuisvrnrzxwkbgdtueoylfl
mphcuisvrnjzewkagdtyeoylfa
mpocuisdrnjzxwkbgdtqeozlfa
mphcuisvrnjjxwkbgdtoeoylfm
mphcuisvenjzxwkbgdtqwoylza
mpmcuisvrnjzxwkbgdtqeoxlfr
mphcuisvgnjhxwkbgdtqeoplfa
mphcuisvrnjzowkdgdtqeoyyfa
mphcuisqynjzxwkbgdtqeoylda
hphcuisvgnjzxwkbgdtbeoylfa
iphcuipvrnuzxwkbgdtqeoylfa
mphcuisvrnjzsikbpdtqeoylfa
mpwcuhsvrnjzxbkbgdtqeoylfa
mnhjuisvcnjzxwkbgdtqeoylfa
mphcudsvrnjzxwkbgdtqloilfa
mpncuiwvrwjzxwkbgdtqeoylfa
mphcuisvrnjgawkbgdtqeoylya
mphcuisvrnjzxwkbggtteoslfa
mphcuisvrnjzxwkbgdvqeoylpe
mphcuisvrnczxfkbgktqeoylfa
mphcuifvrnjzxwkbgdbmeoylfa
mphcuisvrnjytwkbgdtqeoylla
mphcuisvrnjzxwkbgdtjeoxlfn
mphjuisvrnjzxwkbghtqeoyffa
mphcuisvrnjzxkrbgdtqeoylaa
mphcbisvrnjzxwkbgttqeoylfs
mphkuksvbnjzxwkbgdtqeoylfa
nphcuidvrnjzxwhbgdtqeoylfa
mphguzsvrnjzxwkbgdaqeoylfa
mihcuisfrnjzxwkbgdtqhoylfa
mphcuisvrnrzxwpbgdtqesylfa
zphcuisvrnjzxwkbddtqeoylaa
mphcuigvmnjzxwkbgdtqeoylba
mjhcuisvrnjzxjkbgdtqeoylha
mphnuisvrnjznwkbgdtqnoylfa
mkhcuisvrnjcxwkbgdqqeoylfa
mphcuisvenjzxwbbqdtqeoylfa
qphcuisnrnjzawkbgdtqeoylfa
mphcuisvrdjzxwkbgdtqeoywca
mphcuzsvvnjzxwfbgdtqeoylfa
pphcuxsvrnjzxwkbgdtmeoylfa
mphiuvsvrnjzxlkbgdtqeoylfa
mphlqisvrnjzxkkbgdtqeoylfa
mmhcuisvrnjzxwkbgatqeoylea
mphduisrrnjoxwkbgdtqeoylfa
mphcuisvrnjnxwkvgdyqeoylfa
mphcuvsvrnjzxgkbgdtqeoylfz
mphcuisvryjzxwkbggtqkoylfa
iphcuisvrdjzxwkbgotqeoylfa
mphcuisvrnjzxwhbgdtqwoyofa
mphcorbvrnjzxwkbgdtqeoylfa
mghcuisvrnpzxykbgdtqeoylfa
mphauisvrnjnxwkbzdtqeoylfa
mphcgisvrnjzxwkwgdtqeoygfa
mphcuisvrnjzxwkggotqeoylba
mphcuesvrnjzxwkbgdwqebylfa
yphcuisvrnjzxwkbgdxqeoylja
ephyuisvrnjzywkbgdtqeoylfa
mfhcuisqrnjzxwkbgdlqeoylfa
mphkuisvrnjzxwkbertqeoylfa
mphcuusgrnjzxwkbggtqeoylfa
mphcuildrnjvxwkbgdtqeoylfa
mphcuiuvrnjzlwkbgwtqeoylfa
mppcuisvrljzxwkbgdtqeoylfw
mphcwiwvrnjzxwsbgdtqeoylfa
mphcubivrnjzxwkqgdtqeoylfa
mphcuisvrnjpxwkngdtqeoylpa
pchcuisvrgjzxwkbgdtqeoylfa
mphcuisvlnjzxwkbgdtmeoylfw
mphcuisvrnjzywkbgdvqeoylfj
mpzcuisvrnezxwktgdtqeoylfa
mphcuisvrnjbxwkbgzrqeoylfa
mphcuisvrnjzxwktgdtqeodtfa
jphcuiavrnjzxwkbgdtqeoylfv
mphcuisvrnjzxwkbddppeoylfa
mphcuissrkjzxwkbgxtqeoylfa
mphcuisvrhjzxwxbgdtqeoylxa
mphcvisvgnjjxwkbgdtqeoylfa
mphcuisprnjwxwtbgdtqeoylfa
mphcuissrnjzxqkbgdtqeoymfa
mphcuiabrnjzxokbgdtqeoylfa
mphcuisvrnczxwkbgmtpeoylfa`

func task1() int {
	two, three, hasTwo, hasThree, readChars, check, count := 0, 0, false, false, []string{}, true, 1
	rows := strings.Split(input, "\n")
	for _, word := range rows {
		hasTwo, hasThree, readChars = false, false, []string{}
		chars := strings.Split(word, "")
		for i, char := range chars {
			count, check = 1, true
			for _, readChar := range readChars {
				if char == readChar {
					check = false
					break
				}
			}
			if check {
				for j := i + 1; j < len(chars); j++ {
					if char == chars[j] {
						count++
						if count == 4 {
							break
						}
					}
				}
				readChars = append(readChars, char)
				if !hasTwo && count == 2 {
					two++
					hasTwo = true
				} else if !hasThree && count == 3 {
					three++
					hasThree = true
				}
				if hasTwo && hasThree {
					break
				}
			}
		}
	}
	return two * three
}

func task2() string {
	common, diff := "", 0
	rows := strings.Split(input, "\n")
	for i, word := range rows {
		for j := i + 1; j < len(rows); j++ {
			diff = 0
			fir := strings.Split(word, "")
			sec := strings.Split(rows[j], "")
			for k, char := range fir {
				if char != sec[k] {
					if diff != 0 {
						diff = 0
						break
					}
					diff = k
				}
			}
			if diff != 0 {
				fir[diff] = ""
				common = strings.Join(fir, "")
				break
			}
		}
		if diff != 0 {
			break
		}
	}
	return common
}

func main() {
	//fmt.Println(task1())
	fmt.Println(task2())
}
