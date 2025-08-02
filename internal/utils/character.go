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

func getScore(ctx context.Context, c sqlc.Character, name string) (int64, error) {
	scores, err := db.GetCharacterScores(ctx, c.ID, name)
	if err != nil {
		return 0, err
	}
	return calculateValues(scores), nil
}

func getSkill(ctx context.Context, c sqlc.Character, name string) (t.Skill, error) {
	skill, err := db.GetCharacterSkills(ctx, c.ID, name)
	if err != nil {
		return t.Skill{}, err
	}
	s, err := db.GetSkillScore(ctx, skill.Skill)
	if err != nil {
		return t.Skill{}, err
	}
	return t.Skill{Score: s.Score, Proficiency: skill.Prof, Expertise: skill.Expert}, nil
}

func getSpeeds(ctx context.Context, c sqlc.Character) ([]t.Speed, error) {
	speeds := make([]t.Speed, 4)
	for i, tp := range []string{"walk", "swim", "fly", "climb"} {
		speed, err := db.GetCharacterSpeeds(ctx, c.ID, tp)
		if err != nil {
			return []t.Speed{}, err
		}
		speeds[i] = t.Speed{
			Type:  tp,
			Value: calculateValues(speed),
		}
	}
	return speeds, nil
}

func PrepareCharacter(ctx context.Context, c sqlc.Character) (t.Character, error) {
	clsLevel := make([]t.ClassLevel, 0)
	u, err := db.GetUser(ctx, c.Owner)
	if err != nil {
		return t.Character{}, err
	}

	scores := make([]int64, 6)
	for i, stat := range []string{"strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"} {
		s, err := getScore(ctx, c, stat)
		if err != nil {
			return t.Character{}, err
		}
		scores[i] = s
	}

	skills := make([]t.Skill, 18)
	for i, skill := range []string{
		"athletics",
		"acrobatics",
		"sleightofhand",
		"stealth",
		"arcana",
		"historical",
		"investigation",
		"nature",
		"religion",
		"animalhandling",
		"insight",
		"medicine",
		"perception",
		"survival",
		"deception",
		"intimidation",
		"performance",
		"persuasion",
	} {
		s, err := getSkill(ctx, c, skill)
		if err != nil {
			return t.Character{}, err
		}
		skills[i] = s
	}

	speeds, err := getSpeeds(ctx, c)
	if err != nil {
		return t.Character{}, err
	}

	return t.Character{
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
			Strength:     scores[0],
			Dexterity:    scores[1],
			Constitution: scores[2],
			Intelligence: scores[3],
			Wisdom:       scores[4],
			Charisma:     scores[5],
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
			Athletics:      skills[0],
			Acrobatics:     skills[1],
			SleightOfHand:  skills[2],
			Stealth:        skills[3],
			Arcana:         skills[4],
			Historical:     skills[5],
			Investigation:  skills[6],
			Nature:         skills[7],
			Religion:       skills[8],
			AnimalHandling: skills[9],
			Insight:        skills[10],
			Medicine:       skills[11],
			Perception:     skills[12],
			Survival:       skills[13],
			Deception:      skills[14],
			Intimidation:   skills[15],
			Performance:    skills[16],
			Persuasion:     skills[17],
		},
		Combat: t.Combat{
			AC:           c.ArmorClass,
			MaxHP:        c.MaxHp,
			CurrentHP:    c.CurrentHp,
			TempHP:       c.TempHp,
			DeathSuccess: c.DeathSuccess,
			DeathFail:    c.DeathFail,
			Speeds:       speeds,
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
	}, nil
}
