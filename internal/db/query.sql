-- name: NewUser :one
INSERT INTO users (name, passwd, salt)
VALUES (
        sqlc.arg(name),
        sqlc.arg(passwd),
        sqlc.arg(salt)
    )
RETURNING *;

-- name: DelUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);

-- name: GetUser :one
SELECT *
FROM users
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: GetUserFromName :one
SELECT *
FROM users
WHERE name = sqlc.arg(name)
LIMIT 1;

-- name: SetUserRole :one
UPDATE users
SET role = sqlc.arg(role)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetUserName :one
UPDATE users
SET name = sqlc.arg(name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetUserPassword :one
UPDATE users
SET passwd = sqlc.arg(passwd),
    salt = sqlc.arg(salt)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: NewCharacter :one
INSERT INTO characters (owner, name)
VALUES (
        sqlc.arg(owner),
        sqlc.arg(name)
    )
RETURNING *;

-- name: DelCharacter :exec
DELETE FROM characters
WHERE id = sqlc.arg(id);

-- name: GetCharacter :one
SELECT *
FROM characters
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: SetCharacterName :one
UPDATE characters
SET name = sqlc.arg(name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterRace :one
UPDATE characters
SET race = sqlc.arg(race)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterBackground :one
UPDATE characters
SET background = sqlc.arg(background)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterAlignment :one
UPDATE characters
SET alignment = sqlc.arg(alignment)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterExperience :one
UPDATE characters
SET experience = sqlc.arg(experience)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterStrenghtProf :one
UPDATE characters
SET strenght_prof = sqlc.arg(strenght_prof)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterDexterityProf :one
UPDATE characters
SET dexterity_prof = sqlc.arg(dexterity_prof)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterConstitutionProf :one
UPDATE characters
SET constitution_prof = sqlc.arg(constitution_prof)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterIntelligenceProf :one
UPDATE characters
SET intelligence_prof = sqlc.arg(intelligence_prof)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterWisdomProf :one
UPDATE characters
SET wisdom_prof = sqlc.arg(wisdom_prof)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCharismaProf :one
UPDATE characters
SET charisma_prof = sqlc.arg(charisma_prof)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterArmorClass :one
UPDATE characters
SET armor_class = sqlc.arg(armor_class)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterMaxHp :one
UPDATE characters
SET max_hp = sqlc.arg(max_hp)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCurrentHp :one
UPDATE characters
SET current_hp = sqlc.arg(current_hp)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterTempHp :one
UPDATE characters
SET temp_hp = sqlc.arg(temp_hp)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterDeathSuccess :one
UPDATE characters
SET death_success = sqlc.arg(death_success)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterDeathFail :one
UPDATE characters
SET death_fail = sqlc.arg(death_fail)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterProficencies :one
UPDATE characters
SET proficencies = sqlc.arg(proficencies)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterEquipment :one
UPDATE characters
SET equipment = sqlc.arg(equipment)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCoinsCopper :one
UPDATE characters
SET coins_copper = sqlc.arg(coins_copper)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCoinsSilver :one
UPDATE characters
SET coins_silver = sqlc.arg(coins_silver)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCoinsElectrum :one
UPDATE characters
SET coins_electrum = sqlc.arg(coins_electrum)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCoinsGold :one
UPDATE characters
SET coins_gold = sqlc.arg(coins_gold)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterCoinsPlatinum :one
UPDATE characters
SET coins_platinum = sqlc.arg(coins_platinum)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterPersonality :one
UPDATE characters
SET personality = sqlc.arg(personality)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterIdeals :one
UPDATE characters
SET ideals = sqlc.arg(ideals)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterBonds :one
UPDATE characters
SET bonds = sqlc.arg(bonds)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCharacterFlaws :one
UPDATE characters
SET flaws = sqlc.arg(flaws)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: NewCampaign :one
INSERT INTO campaigns (name, master)
VALUES (
        sqlc.arg(name),
        sqlc.arg(master)
    )
RETURNING *;

-- name: DelCampaign :exec
DELETE FROM campaigns
WHERE id = sqlc.arg(id);

-- name: GetCampaign :one
SELECT *
FROM campaigns
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: SetCampaignName :one
UPDATE campaigns
SET name = sqlc.arg(name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SetCampaignDesc :one
UPDATE campaigns
SET desc = sqlc.arg(desc)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: AddCharacterToCampaign :one
INSERT INTO campaigns_characters (character, campaign)
VALUES (sqlc.arg(character), sqlc.arg(campaign))
RETURNING *;

-- name: RemoveCharacterFromCampaign :exec
DELETE FROM campaigns_characters
WHERE character = sqlc.arg(character)
    AND campaign = sqlc.arg(campaign);

-- name: NewCharacterScore :one
INSERT INTO character_scores (score, name, character, value, operation)
VALUES (
        sqlc.arg(score),
        sqlc.arg(name),
        sqlc.arg(character),
        sqlc.arg(value),
        sqlc.arg(operation)
    )
RETURNING *;

-- name: DelCharacterScore :exec
DELETE FROM character_scores
WHERE id = sqlc.arg(id);

-- name: UpdateCharacterScore :one
UPDATE character_scores
SET value = sqlc.arg(value),
    operation = sqlc.arg(operation)
RETURNING *;

-- name: NewCharacterSpeed :one
INSERT INTO character_speeds (name, character, type, value, operation)
VALUES (
        sqlc.arg(name),
        sqlc.arg(character),
        sqlc.arg(type),
        sqlc.arg(value),
        sqlc.arg(operation)
    )
RETURNING *;

-- name: DelCharacterSpeed :exec
DELETE FROM character_speeds
WHERE id = sqlc.arg(id);

-- name: UpdateCharacterSpeed :one
UPDATE character_speeds
SET value = sqlc.arg(value),
    operation = sqlc.arg(operation)
RETURNING *;