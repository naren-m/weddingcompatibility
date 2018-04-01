package main

import (
	"fmt"
)

type Star struct {
	Num     int
	Name    string
	Meaning string
	Lord    string
	Symbol  string
	Deity   string
	Rasi    string
}

var Ashvini, Bharani, Krittika, Rohini, Mrigashirsha, Ardra, Punarvasu, Pushya,
	Ashlesha, Magha, PurvaPhalguni, UttaraPhalguni, Hasta, Chitra, Swati,
	Vishakha, Anuradha, Jyeshtha, Mula, PurvaAshadha, UttaraAshadha, Abhijit,
	Sravana, Dhanishta, Shatabhisha, PurvaBhadrapada, UttaraBhadrapada,
	Revati Star

func init() {

	Ashvini = Star{
		Num:     1,
		Name:    "Ashvini",
		Meaning: "Wife of the Ashvins",
		Lord:    "Ketu (South lunar node)",
		Symbol:  "Horses head",
		Deity:   "Ashvins, the horse-headed twins who are physicians to  gods",
		Rasi:    "Mesha",
	}

	Bharani = Star{
		Num:     2,
		Name:    "Bharani",
		Meaning: "The bearer",
		Lord:    "Shukra  (Venus)",
		Symbol:  "Yoni, the female organ of reproduction",
		Deity:   "Yama, god of death or Dharma",
		Rasi:    "Mesha",
	}

	Krittika = Star{
		Num:     3,
		Name:    "Krittika",
		Meaning: "an old name of the Pleiades; personified as the nurses of Kārttikeya, a son of Shiva.",
		Lord:    "Surya",
		Symbol:  "Knife or spear",
		Deity:   "Agni",
		Rasi:    "Mesha, Vrishabha",
	}

	Rohini = Star{
		Num:     4,
		Name:    "Rohini",
		Meaning: "the red one, a name of Aldebaran. Also known as brāhmī",
		Lord:    "Chandra",
		Symbol:  "Cart or chariot, temple, banyan tree",
		Deity:   "Bramha",
		Rasi:    "Vrishabha",
	}

	Mrigashirsha = Star{
		Num:     5,
		Name:    "Mrigashirsha",
		Meaning: "the deer's head. Also known as āgrahāyaṇī",
		Lord:    "Mangala (Mars)",
		Symbol:  "Deer's head",
		Deity:   "Soma, Chandra, the Moon god",
		Rasi:    "Vrishabha, Mithuna",
	}

	Ardra = Star{
		Num:     6,
		Name:    "Ardra",
		Meaning: "the storm god",
		Lord:    "Rahu (North lunar node)",
		Symbol:  "Teardrop, diamond, a human head",
		Deity:   "Rudra, the storm god",
		Rasi:    "Mithuna",
	}

	Punarvasu = Star{
		Num:     7,
		Name:    "Punarvasu ",
		Meaning: "'the two restorers of goods', also known as yamakau 'the two chariots'",
		Lord:    "Guru (Jupiter)",
		Symbol:  "Bow and quiver",
		Deity:   " Aditi, mother of the gods",
		Rasi:    "Mithuna, Karka",
	}

	Pushya = Star{
		Num:     8,
		Name:    "Pushya",
		Meaning: "'the nourisher', also known as sidhya or tiṣya",
		Lord:    "Shani (Saturn)",
		Symbol:  "Cow's udder, lotus, arrow and circle",
		Deity:   "Bṛhaspati, priest of the gods",
		Rasi:    "Karka",
	}

	Ashlesha = Star{
		Num:     9,
		Name:    "Ashlesha",
		Meaning: "the embrace",
		Lord:    "Budh (Mercury)",
		Symbol:  "Serpent",
		Deity:   "Sarpas or Nagas, deified snakes",
		Rasi:    "Karka",
	}

	Magha = Star{
		Num:     10,
		Name:    "Magha",
		Meaning: "the bountiful",
		Lord:    "Ketu (south lunar node)",
		Symbol:  "Royal Throne",
		Deity:   "Pitrs, 'The Fathers', family ancestors",
		Rasi:    "Simha",
	}

	PurvaPhalguni = Star{
		Num:     11,
		Name:    "PurvaPhalguni",
		Meaning: "first reddish one",
		Lord:    "Shukra (Venus)",
		Symbol:  "Front legs of bed, hammock, fig tree",
		Deity:   "Bhaga, god of marital bliss and prosperity",
		Rasi:    "Simha",
	}

	UttaraPhalguni = Star{
		Num:     12,
		Name:    "UttaraPhalguni",
		Meaning: "second reddish one",
		Lord:    "Surya (Sun)",
		Symbol:  " Four legs of bed, hammock",
		Deity:   "Aryaman, god of patronage and favours",
		Rasi:    "Simha, Kanya",
	}

	Hasta = Star{
		Num:     13,
		Name:    "Hasta",
		Meaning: "the hand",
		Lord:    "Chandra (Moon)",
		Symbol:  "Hand or fist",
		Deity:   "Saviti or Surya, the Sun god",
		Rasi:    "Kanya",
	}

	Chitra = Star{
		Num:     14,
		Name:    "Chitra",
		Meaning: "'the bright one', a name of Spica",
		Lord:    "Mangala (Mars)",
		Symbol:  "Bright jewel or pearl",
		Deity:   "Tvastar or Vishvakarman, the celestial architect",
		Rasi:    " Kanya, Tula",
	}

	Swati = Star{
		Num:     15,
		Name:    "Swati",
		Meaning: "'Su-Ati (sanskrit) Very good' name of Arcturus",
		Lord:    "Rahu (north lunar node)",
		Symbol:  "Shoot of plant, coral",
		Deity:   "Vayu, the Wind god",
		Rasi:    "Tula",
	}

	Vishakha = Star{
		Num:     16,
		Name:    "Vishakha",
		Meaning: "'forked, having branches'; also known as rādhā 'the gift'",
		Lord:    " Guru (Jupiter)",
		Symbol:  "Triumphal arch, potter's wheel",
		Deity:   " Indra, chief of the gods; Agni, god of Fire",
		Rasi:    "Tula  Vrishchika",
	}

	Anuradha = Star{
		Num:     17,
		Name:    "Anuradha",
		Meaning: "following rādhā",
		Lord:    "Shani (Saturn)",
		Symbol:  "Triumphal archway, lotus",
		Deity:   "Mitra, one of Adityas of friendship and partnership",
		Rasi:    "Vrishchika",
	}

	Jyeshtha = Star{
		Num:     18,
		Name:    "Jyeshtha",
		Meaning: "the eldest, most excellent",
		Lord:    "Budh (Mercury)",
		Symbol:  "circular amulet, umbrella, earring",
		Deity:   "Indra, chief of the gods",
		Rasi:    "Vrishchika",
	}

	Mula = Star{
		Num:     19,
		Name:    "Mula",
		Meaning: "the root",
		Lord:    "Ketu (south lunar node)",
		Symbol:  "Bunch of roots tied together, elephant goad",
		Deity:   "Nirrti, goddess of dissolution and destruction",
		Rasi:    "Dhanus",
	}

	PurvaAshadha = Star{
		Num:     20,
		Name:    "PurvaAshadha",
		Meaning: "'first of the aṣāḍhā', aṣāḍhā 'the invincible one' being the name of a constellation",
		Lord:    "Shukra (Venus)",
		Symbol:  "Elephant tusk, fan, winnowing basket",
		Deity:   "Apah, god of Water",
		Rasi:    "Dhanus",
	}

	UttaraAshadha = Star{
		Num:     21,
		Name:    "UttaraAshadha",
		Meaning: "second of the aṣāḍhā",
		Lord:    "Surya (Sun)",
		Symbol:  "Elephant tusk, small bed",
		Deity:   "Visvedevas, universal gods",
		Rasi:    "Dhanus, Makara",
	}

	Abhijit = Star{
		Num:     22,
		Name:    "Abhijit",
		Meaning: "victorious",
		Lord:    "Bramha",
		Symbol:  "",
		Deity:   "",
		Rasi:    "Makara",
	}

	Sravana = Star{
		Num:     23,
		Name:    "Sravana",
		Meaning: "",
		Lord:    "Chandra (Moon)",
		Symbol:  "Ear or Three Footprints",
		Deity:   "Vishnu, preserver of universe",
		Rasi:    "Makara",
	}

	Dhanishta = Star{
		Num:     24,
		Name:    "Dhanishta",
		Meaning: "'most famous', also Shravishthā 'swiftest'",
		Lord:    "Mangala (Mars)",
		Symbol:  "Drum or flute",
		Deity:   "Eight vasus, deities of earthly abundance",
		Rasi:    "Makara, Kumbha",
	}

	Shatabhisha = Star{
		Num:     24,
		Name:    "Shatabhisha",
		Meaning: "requiring a hundred physicians",
		Lord:    "Rahu (north lunar node)",
		Symbol:  "Empty circle, 1,000 flowers or stars",
		Deity:   "Varuna, god of cosmic waters, sky and earth",
		Rasi:    "Kumbha",
	}

	PurvaBhadrapada = Star{
		Num:     25,
		Name:    "PurvaBhadrapada",
		Meaning: "the first of the blessed feet",
		Lord:    "Guru (Jupiter)",
		Symbol:  "Swords or two front legs of funeral cot, man with two faces",
		Deity:   "Ajikapada, an ancient fire dragon",
		Rasi:    "Kumbha, Meena",
	}

	UttaraBhadrapada = Star{
		Num:     26,
		Name:    "UttaraBhadrapada",
		Meaning: "the second of the blessed feet",
		Lord:    "Shani (Saturn)",
		Symbol:  "Twins, back legs of funeral cot, snake in the water",
		Deity:   "Ahir Budhyana, serpent or dragon of the deep",
		Rasi:    "Meena",
	}

	Revati = Star{
		Num:     27,
		Name:    "Revati",
		Meaning: "prosperous",
		Lord:    "Budh (Mercury)",
		Symbol:  "Fish or a pair of fish, drum",
		Deity:   "Pushan, nourisher, the protective deity",
		Rasi:    "Meena",
	}

}

func main() {
	fmt.Println(Revati)
}
