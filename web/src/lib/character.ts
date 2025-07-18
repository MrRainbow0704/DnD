export type Class = {
	name: string;
	level: number;
	hitdie: number;
};

export type Character = {
	charname: string;
	otherprofs: string;
	misc: {
		classlevel: Class[];
		background: string;
		playername: string;
		race: string;
		alignment: `${"Caotico" | "Neutrale" | "Legale"}/${
			| "Buono"
			| "Neutrale"
			| "Cattivo"}`;
		experience: number;
	};
	scores: {
		strenght: number;
		dexterity: number;
		constitution: number;
		intelligence: number;
		wisdom: number;
		charisma: number;
	};
	saves: {
		strenght: boolean;
		dexterity: boolean;
		constitution: boolean;
		intelligence: boolean;
		wisdom: boolean;
		charisma: boolean;
	};
	skills: {
		athletics: {
			proficency: boolean;
			expertise: boolean;
			score: "strenght";
		};
		acrobatics: {
			proficency: boolean;
			expertise: boolean;
			score: "dexterity";
		};
		sleightofhand: {
			proficency: boolean;
			expertise: boolean;
			score: "dexterity";
		};
		stealth: {
			proficency: boolean;
			expertise: boolean;
			score: "dexterity";
		};
		arcana: {
			proficency: boolean;
			expertise: boolean;
			score: "intelligence";
		};
		history: {
			proficency: boolean;
			expertise: boolean;
			score: "intelligence";
		};
		investigation: {
			proficency: boolean;
			expertise: boolean;
			score: "intelligence";
		};
		nature: {
			proficency: boolean;
			expertise: boolean;
			score: "intelligence";
		};
		religion: {
			proficency: boolean;
			expertise: boolean;
			score: "intelligence";
		};
		animalhandling: {
			proficency: boolean;
			expertise: boolean;
			score: "wisdom";
		};
		insight: { proficency: boolean; expertise: boolean; score: "wisdom" };
		medicine: { proficency: boolean; expertise: boolean; score: "wisdom" };
		perception: {
			proficency: boolean;
			expertise: boolean;
			score: "wisdom";
		};
		survival: { proficency: boolean; expertise: boolean; score: "wisdom" };
		deception: {
			proficency: boolean;
			expertise: boolean;
			score: "charisma";
		};
		intimidation: {
			proficency: boolean;
			expertise: boolean;
			score: "charisma";
		};
		performance: {
			proficency: boolean;
			expertise: boolean;
			score: "charisma";
		};
		persuasion: {
			proficency: boolean;
			expertise: boolean;
			score: "charisma";
		};
	};
	combat: {
		ac: number;
		maxhp: number;
		currenthp: number;
		temphp: number;
		deathsuccess: number;
		deathfail: number;
		speeds: {
			type: "walk" | "swim" | "fly" | "climb" | "burrow";
			value: number;
		}[];
		attacks: {
			name: string;
			bonus: number;
			damage: string;
		}[];
	};
	equipment: {
		other: string;
		money: {
			copper: number;
			silver: number;
			electrum: number;
			gold: number;
			platinum: number;
		};
	};
	flavor: {
		personality: string;
		ideals: string;
		bonds: string;
		flaws: string;
	};
	traits: {
		name: string;
		desc: string;
	}[];
};
