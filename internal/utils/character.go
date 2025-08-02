package utils

import (
	"context"
	"math"

	"github.com/MrRainbow0704/DnD/internal/db/sqlc"
	t "github.com/MrRainbow0704/DnD/internal/types"
)

type calculableValues interface {
	GetValue() int64
	GetOperation() string
}


func calculateValues[T calculableValues](v []T) int64 {
	var val int64
	var max int64 = math.MinInt64
	var min int64 = math.MaxInt64

	for _, s := range v {
		switch s.GetOperation() {
		case "sub":
			val -= s.GetValue()
		case "multiply":
			val *= s.GetValue()
		case "set":
			val = s.GetValue()
			return val
		case "set-if-less":
			if max < s.GetValue() {
				max = s.GetValue()
			}
		case "set-if-more":
			if min > s.GetValue() {
				min = s.GetValue()
			}
		default:
			val += s.GetValue()
		}
	}

	if val > min {
		val = min
	}
	if val < max {
		val = max
	}
	return val
}

func getScore(ctx context.Context, c sqlc.Character, name string) int64 {
	scores, err := db.GetCharacterScores(ctx, c.ID, name)
	if err != nil {
		return -1
	}
	return calculateValues(scores)
}

func getSkill(ctx context.Context, c sqlc.Character, name string) t.Skill {
	skill, err := db.GetCharacterSkills(ctx, c.ID, name)
	if err != nil {
		return t.Skill{}
	}
	s, err := db.GetSkillScore(ctx, skill.Skill)
	if err != nil {
		return t.Skill{}
	}
	return t.Skill{Score: s.Score, Proficiency: skill.Prof, Expertise: skill.Expert}
}

func getSpeeds(ctx context.Context, c sqlc.Character) []t.Speed {
	speeds := make([]t.Speed, 4)
	for i, tp := range []string{"walk", "swim", "fly", "climb"} {
		speed, err := db.GetCharacterSpeeds(ctx, c.ID, tp)
		if err != nil {
			return []t.Speed{}
		}
		speeds[i] = t.Speed{
			Type:  tp,
			Value: calculateValues(speed),
		}
	}
	return speeds
}

func PrepareCharacter(ctx context.Context, c sqlc.Character) (t.Character, error) {
	clsLevel := make([]t.ClassLevel, 0)
	u, err := db.GetUser(ctx, c.Owner)
	if err != nil {
		return t.Character{}, err
	}

	n := t.Character{
		Owner:      c.Owner,
		CharName:   c.Name,
		OtherProfs: c.Proficencies.(string),
		Misc: t.Misc{
			ClassLevel: clsLevel,
			Background: c.Background.(string),
			PlayerName: u.Name,
			Race:       c.Race.(string),
			Alignment:  c.Alignment.(string),
			Experience: c.Experience,
		},
		Scores: t.Scores{
			Strength:     getScore(ctx, c, "strength"),
			Dexterity:    getScore(ctx, c, "dexterity"),
			Constitution: getScore(ctx, c, "constitution"),
			Intelligence: getScore(ctx, c, "intelligence"),
			Wisdom:       getScore(ctx, c, "wisdom"),
			Charisma:     getScore(ctx, c, "charisma"),
		},
		Saves: t.Saves{
			Strength:     c.StrenghtProf,
			Dexterity:    c.DexterityProf,
			Constitution: c.ConstitutionProf,
			Intelligence: c.IntelligenceProf,
			Wisdom:       c.WisdomProf,
			Charisma:     c.CharismaProf,
		},
		Skills: t.Skills{
			Athletics:      getSkill(ctx, c, "athletics"),
			Acrobatics:     getSkill(ctx, c, "acrobatics"),
			SleightOfHand:  getSkill(ctx, c, "sleightofhand"),
			Stealth:        getSkill(ctx, c, "stealth"),
			Arcana:         getSkill(ctx, c, "arcana"),
			Historical:     getSkill(ctx, c, "historical"),
			Investigation:  getSkill(ctx, c, "investigation"),
			Nature:         getSkill(ctx, c, "nature"),
			Religion:       getSkill(ctx, c, "religion"),
			AnimalHandling: getSkill(ctx, c, "animalhandling"),
			Insight:        getSkill(ctx, c, "insight"),
			Medicine:       getSkill(ctx, c, "medicine"),
			Perception:     getSkill(ctx, c, "perception"),
			Survival:       getSkill(ctx, c, "survival"),
			Deception:      getSkill(ctx, c, "deception"),
			Intimidation:   getSkill(ctx, c, "intimidation"),
			Performance:    getSkill(ctx, c, "performance"),
			Persuasion:     getSkill(ctx, c, "persuasion"),
		},
		Combat: t.Combat{
			AC:           c.ArmorClass,
			MaxHP:        c.MaxHp,
			CurrentHP:    c.CurrentHp,
			TempHP:       c.TempHp,
			DeathSuccess: c.DeathSuccess,
			DeathFail:    c.DeathFail,
			Speeds:       getSpeeds(ctx, c),
			Attacks:      []t.Attack{}, // TODO
		},
		Equipment: t.Equipment{
			Other: c.Equipment.(string),
			Money: t.Money{
				Copper:   c.CoinsCopper,
				Silver:   c.CoinsSilver,
				Electrum: c.CoinsElectrum,
				Gold:     c.CoinsGold,
				Platinum: c.CoinsPlatinum,
			},
		},
		Flavor: t.Flavor{
			Personality: c.Personality.(string),
			Ideals:      c.Ideals.(string),
			Bonds:       c.Bonds.(string),
			Flaws:       c.Flaws.(string),
		},
		Traits: []t.Trait{}, // TODO
	}
	return n, nil
}
