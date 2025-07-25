CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    passwd BLOB NOT NULL,
    salt BLOB NOT NULL,
    role TEXT NOT NULL DEFAULT "user"
);

CREATE TABLE IF NOT EXISTS campaigns (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    desc TEXT NOT NULL DEFAULT "",
    master INTEGER NOT NULL,
    FOREIGN KEY (master) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS campaigns_characters (
    character INTEGER NOT NULL,
    campaign INTEGER NOT NULL,
    PRIMARY KEY (character, campaign),
    FOREIGN KEY (character) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (campaign) REFERENCES campaings(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS characters (
    id INTEGER PRIMARY KEY,
    owner INTEGER NOT NULL,
    name TEXT NOT NULL,
    race TEXT NOT NULL DEFAULT "custom",
    background TEXT NOT NULL DEFAULT "custom",
    alignment TEXT NOT NULL DEFAULT "NN" CHECK(
        alignment IN (
            "LB",
            "LN",
            "LC",
            "NB",
            "NN",
            "NC",
            "CB",
            "CN",
            "CC"
        )
    ),
    experience INTEGER NOT NULL DEFAULT 0 CHECK(experience >= 0),
    strenght_prof BOOLEAN NOT NULL DEFAULT 0 CHECK(strenght_prof IN (0, 1)),
    dexterity_prof BOOLEAN NOT NULL DEFAULT 0 CHECK(dexterity_prof IN (0, 1)),
    constitution_prof BOOLEAN NOT NULL DEFAULT 0 CHECK(constitution_prof IN (0, 1)),
    intelligence_prof BOOLEAN NOT NULL DEFAULT 0 CHECK(intelligence_prof IN (0, 1)),
    wisdom_prof BOOLEAN NOT NULL DEFAULT 0 CHECK(wisdom_prof IN (0, 1)),
    charisma_prof BOOLEAN NOT NULL DEFAULT 0 CHECK(charisma_prof IN (0, 1)),
    armor_class INTEGER NOT NULL DEFAULT 0 CHECK(armor_class >= 0),
    max_hp INTEGER NOT NULL DEFAULT 0 CHECK(max_hp >= 0),
    current_hp INTEGER NOT NULL DEFAULT 0 CHECK(
        current_hp >= - max_hp
        AND current_hp <= max_hp + temp_hp
    ),
    temp_hp INTEGER NOT NULL DEFAULT 0 CHECK(temp_hp >= 0),
    death_success INTEGER NOT NULL DEFAULT 0 CHECK(
        death_success >= 0
        AND death_success <= 3
    ),
    death_fail INTEGER NOT NULL DEFAULT 0 CHECK(
        death_fail >= 0
        AND death_fail <= 3
    ),
    proficencies TEXT NOT NULL DEFAULT "",
    equipment TEXT NOT NULL DEFAULT "",
    coins_copper INTEGER NOT NULL DEFAULT 0 CHECK(coins_copper >= 0),
    coins_silver INTEGER NOT NULL DEFAULT 0 CHECK(coins_silver >= 0),
    coins_electrum INTEGER NOT NULL DEFAULT 0 CHECK(coins_electrum >= 0),
    coins_gold INTEGER NOT NULL DEFAULT 0 CHECK(coins_gold >= 0),
    coins_platinum INTEGER NOT NULL DEFAULT 0 CHECK(coins_platinum >= 0),
    personality TEXT NOT NULL DEFAULT "",
    ideals TEXT NOT NULL DEFAULT "",
    bonds TEXT NOT NULL DEFAULT "",
    flaws TEXT NOT NULL DEFAULT "",
    FOREIGN KEY (owner) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (background) REFERENCES backgrounds(name) ON DELETE CASCADE,
    FOREIGN KEY (race) REFERENCES races(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS backgrounds (
    name TEXT PRIMARY KEY,
    desc TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS races (
    name TEXT PRIMARY KEY,
    desc TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS scores (name TEXT PRIMARY KEY);

CREATE TABLE IF NOT EXISTS character_scores (
    id INTEGER PRIMARY KEY,
    score TEXT NOT NULL,
    name TEXT NOT NULL,
    character INTEGER NOT NULL,
    value INTEGER NOT NULL,
    operation TEXT NOT NULL DEFAULT "add",
    FOREIGN KEY (score) REFERENCES scores(name) ON DELETE CASCADE,
    FOREIGN KEY (character) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (operation) REFERENCES operations(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS character_speeds (
    id INTEGER PRIMARY KEY,
    character INTEGER NOT NULL,
    value INTEGER NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL DEFAULT "walk" CHECK(
        type IN ("walk", "fly", "swim", "climb", "burrow")
    ),
    operation TEXT NOT NULL DEFAULT "add",
    FOREIGN KEY (character) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (operation) REFERENCES operations(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS classes (name TEXT PRIMARY KEY);

CREATE TABLE IF NOT EXISTS character_classes (
    character INTEGER NOT NULL,
    class INTEGER NOT NULL,
    subclass TEXT,
    level INTEGER NOT NULL DEFAULT 1 CHECK(
        level >= 1
        AND level <= 20
    ),
    PRIMARY KEY (character, class),
    FOREIGN KEY (character) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (class) REFERENCES classes(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS skills (
    name TEXT PRIMARY KEY,
    score TEXT NOT NULL,
    FOREIGN KEY (score) REFERENCES scores(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS character_skills (
    character INTEGER NOT NULL,
    skill TEXT NOT NULL,
    prof BOOLEAN NOT NULL DEFAULT 0 CHECK(prof IN (0, 1)),
    expert BOOLEAN NOT NULL DEFAULT 0 CHECK(prof IN (0, 1)),
    PRIMARY KEY (character, skill),
    FOREIGN KEY (character) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (skill) REFERENCES skills(name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS operations (name TEXT PRIMARY KEY);

INSERT
    OR IGNORE INTO operations (name)
VALUES ("add"),
    ("sub"),
    ("set"),
    ("set-if-less"),
    ("set-if-more"),
    ("multiply");

INSERT
    OR IGNORE INTO scores (name)
VALUES ("strenght"),
    ("dexterity"),
    ("constitution"),
    ("intelligence"),
    ("wisdom"),
    ("charisma");

INSERT
    OR IGNORE INTO skills (name, score)
VALUES ("athletics", "strenght"),
    ("acrobatics", "dexterity"),
    ("sleightofhand", "dexterity"),
    ("stealth", "dexterity"),
    ("arcana", "intelligence"),
    ("history", "intelligence"),
    ("investigation", "intelligence"),
    ("nature", "intelligence"),
    ("religion", "intelligence"),
    ("animalhandling", "wisdom"),
    ("insight", "wisdom"),
    ("medicine", "wisdom"),
    ("perception", "wisdom"),
    ("survival", "wisdom"),
    ("deception", "charisma"),
    ("intimidation", "charisma"),
    ("performance", "charisma"),
    ("persuasion", "charisma");

INSERT
    OR IGNORE INTO classes (name)
VALUES ("artificer"),
    ("barbarian"),
    ("bard"),
    ("cleric"),
    ("druid"),
    ("fighter"),
    ("monk"),
    ("paladin"),
    ("ranger"),
    ("rogue"),
    ("sorcerer"),
    ("warlock"),
    ("wizard");

INSERT
    OR IGNORE INTO backgrounds (name, desc)
VALUES ("custom", "Custom background");

INSERT
    OR IGNORE INTO races (name, desc)
VALUES ("custom", "Custom race"),
    ("arakocra", ""),
    ("aasimar", ""),
    ("bugbear", ""),
    ("centaur", ""),
    ("changeling", ""),
    ("elf", ""),
    ("halfling", ""),
    ("dwarf", ""),
    ("human", ""),
    ("dragonborn", ""),
    ("dhampir", ""),
    ("duergar", ""),
    ("fairy", ""),
    ("genasi", ""),
    ("goblin", ""),
    ("tabaxi", ""),
    ("goliath", ""),
    ("gnome", ""),
    ("half-elf", ""),
    ("half-orc", ""),
    ("tiefling", "");