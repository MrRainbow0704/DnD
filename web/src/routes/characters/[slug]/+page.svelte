<script lang="ts">
import { page } from "$app/state";
import { getCharacter, type Character } from "$src/lib/character";
import CharacterSheet from "$components/CharacterSheet.svelte";
import type { Errors } from "$src/lib";

let character: Character = {
	charname: "Lúmë",
	otherprofs:
		"Lingue: Comune, Celestiale, Elfico e Primordiale.\nArmi: Daghe, Dardi, Fionde, Staffe e Balestre Leggere",
	misc: {
		classlevel: [
			{
				name: "Mago",
				level: 6,
				hitdie: 6,
			},
		],
		background: "Vagabondo Astrale",
		playername: "Marco",
		race: "Aasimar (caduto)",
		alignment: "Legale/Neutrale",
		experience: 0,
	},
	scores: {
		strenght: 10,
		dexterity: 12,
		constitution: 16,
		intelligence: 18,
		wisdom: 15,
		charisma: 15,
	},
	saves: {
		strenght: false,
		dexterity: false,
		constitution: false,
		intelligence: true,
		wisdom: true,
		charisma: false,
	},
	skills: {
		athletics: {
			proficency: false,
			expertise: false,
			score: "strenght",
		},
		acrobatics: {
			proficency: false,
			expertise: false,
			score: "dexterity",
		},
		sleightofhand: {
			proficency: false,
			expertise: false,
			score: "dexterity",
		},
		stealth: {
			proficency: false,
			expertise: false,
			score: "dexterity",
		},
		arcana: {
			proficency: true,
			expertise: false,
			score: "intelligence",
		},
		history: {
			proficency: true,
			expertise: false,
			score: "intelligence",
		},
		investigation: {
			proficency: false,
			expertise: false,
			score: "intelligence",
		},
		nature: {
			proficency: false,
			expertise: false,
			score: "intelligence",
		},
		religion: {
			proficency: true,
			expertise: false,
			score: "intelligence",
		},
		animalhandling: {
			proficency: false,
			expertise: false,
			score: "wisdom",
		},
		insight: { proficency: true, expertise: false, score: "wisdom" },
		medicine: { proficency: false, expertise: false, score: "wisdom" },
		perception: {
			proficency: false,
			expertise: false,
			score: "wisdom",
		},
		survival: { proficency: false, expertise: false, score: "wisdom" },
		deception: {
			proficency: false,
			expertise: false,
			score: "charisma",
		},
		intimidation: {
			proficency: false,
			expertise: false,
			score: "charisma",
		},
		performance: {
			proficency: false,
			expertise: false,
			score: "charisma",
		},
		persuasion: {
			proficency: false,
			expertise: false,
			score: "charisma",
		},
	},
	combat: {
		ac: 11,
		maxhp: 38,
		currenthp: 38,
		temphp: 0,
		deathfail: 1,
		deathsuccess: 2,
		speeds: [
			{
				type: "walk",
				value: 9,
			},
		],
		attacks: [
			{
				name: "Staffa",
				bonus: 4,
				damage: "1d6 Contundente",
			},
		],
	},
	equipment: {
		other: "Staffa, Pugnale, Libro di incantesimi, Orologio",
		money: {
			copper: 8,
			silver: 4,
			electrum: 0,
			gold: 46,
			platinum: 1,
		},
	},
	flavor: {
		personality: "",
		ideals: "",
		bonds: "",
		flaws: "",
	},
	traits: [
		{ name: "Incantesimi", desc: "Incantesimi da mago" },
		{ name: "Scurovisione", desc: "Vedo al buio per 18m" },
	],
};
// let character: Character = {} as Character;
let errors: Errors = {};
const id = parseInt(page.params.slug);
getCharacter(id)
	.then((c) => (character = c))
	.catch((e: Error) => {
		errors = e.cause as Errors;
	});
</script>

<div>
	{#if Object.entries(errors)}
		{#each Object.entries(errors) as [k, v]}
			<p><strong>{k}:</strong> {v}</p>
		{/each}
	{/if}
</div>
<CharacterSheet c={character} />

<style>
div {
	width: 20rem;
	max-width: 80%;
}
p {
	width: 100%;
	text-align: center;
	padding: 0.5rem;
	margin: 0.1rem;
	background-color: red;
}
</style>
