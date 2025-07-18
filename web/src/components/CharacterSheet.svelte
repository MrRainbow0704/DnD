<script lang="ts">
import type { Character } from "$lib/character";
import { t } from "$lib/lang";

let { c }: { c: Character } = $props();

let level: number = $state(0);
let hitdicestr: string = $state("");
let hitdice = { 6: 0, 8: 0, 10: 0, 12: 0 };
let classes: string[] = [];
c.misc.classlevel.forEach((e) => {
	classes = classes.concat(`${e.name} ${e.level}`);
	level += e.level;
	hitdice[e.hitdie as keyof typeof hitdice] += e.level;
});

Object.entries(hitdice).forEach((e) => {
	if (e[1] !== 0) {
		hitdicestr += `${e[1]}d${e[0]}`;
	}
});
let clsLevel: string = classes.join(" / ");
let features: string = c.traits.map((e) => `${e.name}: ${e.desc}`).join("\n");
</script>

<form class="charsheet">
	<header>
		<section class="charname">
			<label for="charname">{t("charactersheet.charname")}</label>
			<input name="charname" type="text" value={c.charname} />
		</section>
		<section class="misc">
			<ul>
				{#each Object.entries(c.misc) as [k, v]}
					<li>
						<label for={k}>
							{t(
								`charactersheet.misc.${k}` as Parameters<
									typeof t
								>[0]
							)}
						</label>
						{#if k === "classlevel"}
							<input name={k} type="text" value={clsLevel} />
						{:else}
							<input name={k} type="text" value={v} />
						{/if}
					</li>
				{/each}
			</ul>
		</section>
	</header>
	<main>
		<section>
			<section class="attributes">
				<div class="scores">
					<ul>
						{#each Object.entries(c.scores) as [k, v]}
							<li>
								<div class="score">
									<label for={k + "score"}>
										{t(
											`charactersheet.score.${k}` as Parameters<
												typeof t
											>[0]
										)}
									</label>
									<input
										name={k + "score"}
										type="text"
										value={v} />
								</div>
								<div class="modifier">
									<input
										name={k + "mod"}
										type="text"
										value={Math.floor((v - 10) / 2)} />
								</div>
							</li>
						{/each}
					</ul>
				</div>
				<div class="attr-applications">
					<div class="inspiration box">
						<div class="label-container">
							<label for="inspiration">
								{t("charactersheet.inspiration")}
							</label>
						</div>
						<div class="inspiration check-container">
							<input name="inspiration" type="checkbox" />
						</div>
					</div>
					<div class="proficiencybonus box">
						<div class="label-container">
							<label for="proficiencybonus">
								{t("charactersheet.proficiencybonus")}
							</label>
						</div>
						<input
							name="proficiencybonus"
							type="text"
							value={Math.floor(level / 4) + 2} />
					</div>
					<div class="saves list-section box">
						<ul>
							{#each Object.entries(c.saves) as [k, v]}
								<li>
									<label for={k + "-save"}>
										{t(
											`charactersheet.score.${k}` as Parameters<
												typeof t
											>[0]
										)}
									</label>
									<input
										name={k + "-save"}
										type="text"
										value={v
											? Math.floor(level / 4) +
												2 +
												Math.floor(
													(c.scores[
														k as keyof typeof c.scores
													] -
														10) /
														2
												)
											: Math.floor(
													(c.scores[
														k as keyof typeof c.scores
													] -
														10) /
														2
												)} />
									<input
										name={k + "-save-prof"}
										type="checkbox"
										checked={v} />
								</li>
							{/each}
						</ul>
						<div class="label">
							{t("charactersheet.savingthrows")}
						</div>
					</div>
					<div class="skills list-section box">
						<ul>
							{#each Object.entries(c.skills) as [k, v]}
								<li>
									<label for={k}>
										{t(
											`charactersheet.skill.${k}` as Parameters<
												typeof t
											>[0]
										)}
										<span class="skill"
											>({t(
												`charactersheet.score.${v.score}` as Parameters<
													typeof t
												>[0]
											).slice(0, 3)})</span>
									</label>
									<input
										name={k}
										type="text"
										value={v.proficency
											? v.expertise
												? (Math.floor(level / 4) + 2) *
														2 +
													Math.floor(
														(c.scores[
															v.score as keyof typeof c.scores
														] -
															10) /
															2
													)
												: Math.floor(level / 4) +
													2 +
													Math.floor(
														(c.scores[
															v.score as keyof typeof c.scores
														] -
															10) /
															2
													)
											: Math.floor(
													(c.scores[
														v.score as keyof typeof c.scores
													] -
														10) /
														2
												)} />
									<input
										name={k + "-prof"}
										type="checkbox"
										checked={v.proficency} />
									<input
										name={k + "-exper"}
										class="expertise"
										type="checkbox"
										checked={v.expertise} />
								</li>
							{/each}
						</ul>
						<div class="label">{t("charactersheet.skill.")}</div>
					</div>
				</div>
			</section>
			<div class="passive-perception box">
				<div class="label-container">
					<label for="passiveperception">
						{t("charactersheet.passiveperception")}
					</label>
				</div>
				<input
					name="passiveperception"
					value={c.skills.perception.proficency
						? c.skills.perception.expertise
							? (Math.floor(level / 4) + 2) * 2 +
								Math.floor(
									(c.scores[
										c.skills.perception
											.score as keyof typeof c.scores
									] -
										10) /
										2
								) +
								10
							: Math.floor(level / 4) +
								2 +
								Math.floor(
									(c.scores[
										c.skills.perception
											.score as keyof typeof c.scores
									] -
										10) /
										2
								) +
								10
						: Math.floor(
								(c.scores[
									c.skills.perception
										.score as keyof typeof c.scores
								] -
									10) /
									2
							) + 10} />
			</div>
			<div class="otherprofs box textblock">
				<label for="otherprofs">
					{t("charactersheet.otherprofs")}
				</label>
				<textarea name="otherprofs">{c.otherprofs}</textarea>
			</div>
		</section>
		<section>
			<section class="combat">
				<div class="armorclass">
					<div>
						<label for="ac">
							{t("charactersheet.combat.ac")}
						</label>
						<input name="ac" type="text" value={c.combat.ac} />
					</div>
				</div>
				<div class="initiative">
					<div>
						<label for="initiative">
							{t("charactersheet.combat.initiative")}
						</label>
						<input
							name="initiative"
							type="text"
							value={Math.floor((c.scores.dexterity - 10) / 2)} />
					</div>
				</div>
				<div class="speed">
					<div>
						<label for="speed">
							{t("charactersheet.combat.speed")}
						</label>
						<input
							name="speed"
							type="text"
							value={`${c.combat.speed}m`} />
					</div>
				</div>
				<div class="hp">
					<div class="regular">
						<div class="max">
							<label for="maxhp">
								{t("charactersheet.combat.maxhp")}
							</label>
							<input
								name="maxhp"
								type="text"
								value={c.combat.maxhp} />
						</div>
						<div class="current">
							<label for="currenthp">
								{t("charactersheet.combat.currenthp")}
							</label>
							<input
								name="currenthp"
								type="text"
								value={c.combat.currenthp} />
						</div>
					</div>
					<div class="temporary">
						<label for="temphp">
							{t("charactersheet.combat.temphp")}
						</label>
						<input
							name="temphp"
							type="text"
							value={c.combat.temphp} />
					</div>
				</div>
				<div class="hitdice">
					<div>
						<div class="total">
							<label for="totalhd">
								{t("charactersheet.combat.hitdice.total")}
							</label>
							<input
								name="totalhd"
								type="text"
								value={hitdicestr} />
						</div>
						<div class="remaining">
							<label for="remaininghd">
								{t("charactersheet.combat.hitdice.remaning")}
							</label>
							<input name="remaininghd" type="text" />
						</div>
					</div>
				</div>
				<div class="deathsaves">
					<div>
						<div class="label">
							<label>
								{t("charactersheet.combat.deathsaves.")}
							</label>
						</div>
						<div class="marks">
							<div class="deathsuccesses">
								<label>
									{t(
										"charactersheet.combat.deathsaves.success"
									)}
								</label>
								<div class="bubbles">
									{#each Array(3).keys() as i}
										<input
											name="deathsuccess1"
											type="checkbox"
											checked={i + 1 <=
											c.combat.deathsuccess
												? true
												: false} />
									{/each}
								</div>
							</div>
							<div class="deathfails">
								<label>
									{t(
										"charactersheet.combat.deathsaves.faliure"
									)}
								</label>
								<div class="bubbles">
									{#each Array(3).keys() as i}
										<input
											name="deathfail1"
											type="checkbox"
											checked={i + 1 <= c.combat.deathfail
												? true
												: false} />
									{/each}
								</div>
							</div>
						</div>
					</div>
				</div>
			</section>
			<section class="attacksandspellcasting">
				<div>
					<label>
						{t("charactersheet.attacksandspellcasting.")}
					</label>
					<table>
						<thead>
							<tr>
								<th>
									{t(
										"charactersheet.attacksandspellcasting.name"
									)}
								</th>
								<th>
									{t(
										"charactersheet.attacksandspellcasting.bonus"
									)}
								</th>
								<th>
									{t(
										"charactersheet.attacksandspellcasting.damage"
									)}
								</th>
							</tr>
						</thead>
						<tbody>
							{#each c.combat.attacks as a}
								<tr>
									<td>
										<input
											name="atkname1"
											type="text"
											value={a.name} />
									</td>
									<td>
										<input
											name="atkbonus1"
											type="text"
											value={(a.bonus < 0 ? "" : "+") +
												a.bonus} />
									</td>
									<td>
										<input
											name="atkdamage1"
											type="text"
											value={a.damage} />
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
					<textarea> </textarea>
				</div>
			</section>
			<section class="equipment">
				<div>
					<label>{t("charactersheet.equipment.")}</label>
					<div class="money">
						<ul>
							<li>
								<label for="cp">
									<abbr
										title={t(
											"charactersheet.equipment.coin.copper"
										)}>
										{t(
											"charactersheet.equipment.coin.abbr.copper"
										)}
									</abbr>
								</label>
								<input
									name="cp"
									value={c.equipment.money.copper} />
							</li>
							<li>
								<label for="sp">
									<abbr
										title={t(
											"charactersheet.equipment.coin.silver"
										)}>
										{t(
											"charactersheet.equipment.coin.abbr.silver"
										)}
									</abbr>
								</label>
								<input
									name="sp"
									value={c.equipment.money.silver} />
							</li>
							<li>
								<label for="ep">
									<abbr
										title={t(
											"charactersheet.equipment.coin.electrum"
										)}>
										{t(
											"charactersheet.equipment.coin.abbr.electrum"
										)}
									</abbr>
								</label>
								<input
									name="ep"
									value={c.equipment.money.electrum} />
							</li>
							<li>
								<label for="gp">
									<abbr
										title={t(
											"charactersheet.equipment.coin.gold"
										)}>
										{t(
											"charactersheet.equipment.coin.abbr.gold"
										)}
									</abbr>
								</label>
								<input
									name="gp"
									value={c.equipment.money.gold} />
							</li>
							<li>
								<label for="pp">
									<abbr
										title={t(
											"charactersheet.equipment.coin.platinum"
										)}>
										{t(
											"charactersheet.equipment.coin.abbr.platinum"
										)}
									</abbr>
								</label>
								<input
									name="pp"
									value={c.equipment.money.platinum} />
							</li>
						</ul>
					</div>
					<textarea placeholder="Equipment list here"
						>{c.equipment.other}</textarea>
				</div>
			</section>
		</section>
		<section>
			<section class="flavor">
				<div class="personality">
					<label for="personality">
						{t("charactersheet.flavor.personality")}
					</label>
					<textarea name="personality"
						>{c.flavor.personality}</textarea>
				</div>
				<div class="ideals">
					<label for="ideals">
						{t("charactersheet.flavor.ideals")}
					</label>
					<textarea name="ideals">{c.flavor.ideals}</textarea>
				</div>
				<div class="bonds">
					<label for="bonds">
						{t("charactersheet.flavor.bonds")}
					</label>
					<textarea name="bonds">{c.flavor.bonds}</textarea>
				</div>
				<div class="flaws">
					<label for="flaws">
						{t("charactersheet.flavor.flaws")}
					</label>
					<textarea name="flaws">{c.flavor.flaws}</textarea>
				</div>
			</section>
			<section class="features">
				<div>
					<label for="features">
						{t("charactersheet.features")}
					</label>
					<textarea name="features">{features}</textarea>
				</div>
			</section>
		</section>
	</main>
</form>

<style>
textarea {
	font-size: 12px;
	text-align: left;
	width: calc(100% - 20px - 2px);
	border-radius: 10px;
	padding: 10px;
	resize: none;
	overflow: hidden;
	height: 15em;
}

input[type="checkbox"] {
	cursor: pointer;
}

.box {
	margin-top: 10px;
}

form {
	width: 800px;
	right: 0;
	left: 0;
	margin-right: auto;
	margin-left: auto;
	margin-top: 10px;
}

.textblock {
	display: flex;
	flex-direction: column-reverse;
	width: 100%;
	align-items: center;
}

.textblock label {
	text-align: center;
	border: 1px solid black;
	border-top: 0;
	font-size: 10px;
	width: calc(100% - 20px - 2px);
	border-radius: 0 0 10px 10px;
	padding: 4px 0;
	font-weight: bold;
}

.textblock textarea {
	border: 1px solid black;
}

ul {
	margin: 0;
	padding: 0;
}

ul li {
	list-style-image: none;
	display: block;
}

label {
	text-transform: uppercase;
	font-size: 12px;
}

header {
	display: flex;
	align-items: stretch;
}

header section.charname {
	border: 1px solid black;
	border-right: 0;
	border-radius: 10px 0 0 10px;
	padding: 5px 0;
	background-color: #eee;
	width: 30%;
	display: flex;
	flex-direction: column-reverse;
	bottom: 0;
	top: 0;
	margin: auto;
}

header section.charname input {
	padding: 0.5em;
	margin: 5px;
	border: 0;
}

header section.charname label {
	padding-left: 1em;
}

header section.misc {
	width: 70%;
	border: 1px solid black;
	border-radius: 10px;
	padding-left: 1em;
	padding-right: 1em;
}

header section.misc ul {
	padding: 10px 0px 5px 0px;
	display: flex;
	flex-wrap: wrap;
}

header section.misc ul li {
	width: 33.33333%;
	display: flex;
	flex-direction: column-reverse;
}

header section.misc ul li label {
	margin-bottom: 5px;
}

header section.misc ul li input {
	border: 0;
	border-bottom: 1px solid #ddd;
}

main {
	display: flex;
	justify-content: space-between;
	margin-top: 20px;
}

main .label-container {
	position: relative;
	width: 100%;
	height: 18px;
	margin-top: 6px;
	border: 1px solid black;
	border-left: 0;
	text-align: center;
}

main .label-container > label {
	position: absolute;
	left: 0;
	top: 1px;
	transform: translate(0%, 50%);
	width: 100%;
	font-size: 8px;
}

main > section {
	width: 32%;
	display: flex;
	flex-direction: column;
}

main > section section.attributes {
	width: 100%;
	display: flex;
	flex-direction: row;
	justify-content: space-between;
}

main > section section.attributes .scores {
	width: 101px;
	background-color: #bbb;
	border-radius: 10px;
	padding-bottom: 5px;
}

main > section section.attributes .scores label {
	font-size: 8px;
	font-weight: bold;
}

main > section section.attributes .scores ul {
	display: flex;
	flex-direction: column;
	justify-content: space-around;
	align-items: center;
	height: 100%;
}

main > section section.attributes .scores ul li {
	height: 80px;
	width: 70px;
	background-color: white;
	border: 1px solid black;
	text-align: center;
	display: flex;
	flex-direction: column;
	border-radius: 10px;
}

main > section section.attributes .scores ul li input {
	width: 100%;
	padding: 0;
	border: 0;
}

main > section section.attributes .scores ul li .score input {
	text-align: center;
	font-size: 40px;
	padding: 2px 0px 0px 0px;
	background: white;
}

main > section section.attributes .scores ul li .modifier {
	margin-top: 3px;
}

main > section section.attributes .scores ul li .modifier input {
	background: white;
	width: 30px;
	height: 20px;
	border: 1px inset black;
	border-radius: 20px;
	margin: -1px;
	text-align: center;
}

main > section section.attributes .attr-applications .inspiration {
	display: flex;
	flex-direction: row-reverse;
	justify-content: flex-end;
}

main
	> section
	section.attributes
	.attr-applications
	.inspiration
	.check-container {
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	padding: 8px;
	border: 1px solid black;
	border-radius: 10px;
	order: 1;
}

main
	> section
	section.attributes
	.attr-applications
	.inspiration
	.check-container
	input {
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	width: 10px;
	height: 10px;
	border-radius: 10px;
	border: 1px solid black;
	order: 1;
}

main
	> section
	section.attributes
	.attr-applications
	.inspiration
	.check-container
	input:checked {
	background-color: black;
}

main > section section.attributes .attr-applications .proficiencybonus {
	display: flex;
	flex-direction: row-reverse;
	justify-content: flex-end;
}

main > section section.attributes .attr-applications .proficiencybonus input {
	width: 30px;
	height: 28px;
	border: 1px solid black;
	text-align: center;
	border-radius: 10px;
}

main > section section.attributes .attr-applications .list-section {
	border: 1px solid black;
	border-radius: 10px;
	padding: 10px 5px;
}

main > section section.attributes .attr-applications .list-section .label {
	margin-top: 10px;
	margin-bottom: 2.5px;
	text-align: center;
	text-transform: uppercase;
	font-size: 10px;
	font-weight: bold;
}

main > section section.attributes .attr-applications .list-section ul li {
	display: flex;
	align-items: center;
	position: relative;
}

main > section section.attributes .attr-applications .list-section ul li > * {
	margin-left: 5px;
}

main > section section.attributes .attr-applications .list-section ul li label {
	text-transform: none;
	font-size: 8px;
	text-align: left;
	order: 4;
}

main
	> section
	section.attributes
	.attr-applications
	.list-section
	ul
	li
	label
	span.skill {
	color: #bbb;
	font-size: 6px;
}

main
	> section
	section.attributes
	.attr-applications
	.list-section
	ul
	li
	input[type="text"] {
	width: 30px;
	font-size: 12px;
	text-align: center;
	border: 0;
	border-bottom: 1px solid black;
	order: 3;
}

main
	> section
	section.attributes
	.attr-applications
	.list-section
	ul
	li
	input[type="checkbox"] {
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	width: 10px;
	height: 10px;
	border: 1px solid black;
	border-radius: 10px;
	order: 1;
}

main
	> section
	section.attributes
	.attr-applications
	.list-section
	ul
	li
	input[type="checkbox"].expertise {
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	position: absolute;
	margin: 1px;
	width: 6px;
	height: 6px;
	border: 1px solid black;
	border-radius: 10px;
	order: 2;
	top: 0;
	left: 12px;
}

main
	> section
	section.attributes
	.attr-applications
	.list-section
	ul
	li
	input[type="checkbox"]:checked {
	background-color: black;
}

main > section .passive-perception {
	display: flex;
	flex-direction: row-reverse;
	justify-content: flex-end;
}

main > section .passive-perception input {
	width: 30px;
	height: 28px;
	text-align: center;
	border: 1px solid black;
	border-radius: 10px;
}

main > section .otherprofs textarea {
	height: 26em;
}

main section.combat {
	background-color: #eee;
	display: flex;
	flex-wrap: wrap;
	border-radius: 10px;
}

main section.combat > div {
	overflow: hidden;
}

main section.combat > .armorclass,
main section.combat > .initiative,
main section.combat > .speed {
	flex-basis: 33.333%;
}

main section.combat > .armorclass > div,
main section.combat > .initiative > div,
main section.combat > .speed > div {
	display: flex;
	flex-direction: column-reverse;
	align-items: center;
	margin-top: 10px;
}

main section.combat > .armorclass > div label,
main section.combat > .initiative > div label,
main section.combat > .speed > div label {
	font-size: 8px;
	width: 50px;
	border: 1px solid black;
	border-top: 0;
	background-color: white;
	text-align: center;
	padding-top: 5px;
	padding-bottom: 5px;
	border-radius: 0 0 10px 10px;
}

main section.combat > .armorclass > div input,
main section.combat > .initiative > div input,
main section.combat > .speed > div input {
	height: 70px;
	width: 70px;
	border-radius: 10px;
	border: 1px solid black;
	text-align: center;
	font-size: 30px;
}

main section.combat > .hp {
	flex-basis: 100%;
}

main section.combat > .hp > .regular {
	background-color: white;
	border: 1px solid black;
	margin: 10px;
	margin-bottom: 5px;
	border-radius: 10px 10px 0 0;
}

main section.combat > .hp > .regular > .max {
	display: flex;
	justify-content: space-around;
	align-items: baseline;
}

main section.combat > .hp > .regular > .max label {
	font-size: 10px;
	text-transform: none;
	color: #bbb;
}

main section.combat > .hp > .regular > .max input {
	width: 40%;
	border: 0;
	border-bottom: 1px solid #ddd;
	font-size: 12px;
	text-align: center;
}

main section.combat > .hp > .regular > .current {
	display: flex;
	flex-direction: column-reverse;
}

main section.combat > .hp > .regular > .current input {
	border: 0;
	width: 100%;
	padding: 1em 0;
	font-size: 20px;
	text-align: center;
}

main section.combat > .hp > .regular > .current label {
	font-size: 10px;
	padding-bottom: 5px;
	text-align: center;
	font-weight: bold;
}

main section.combat > .hp > .temporary {
	margin: 10px;
	margin-top: 0;
	border: 1px solid black;
	border-radius: 0 0 10px 10px;
	display: flex;
	flex-direction: column-reverse;
	background-color: white;
}

main section.combat > .hp > .temporary input {
	padding: 1em 0;
	font-size: 20px;
	border: 0;
	text-align: center;
}

main section.combat > .hp > .temporary label {
	font-size: 10px;
	padding-bottom: 5px;
	text-align: center;
	font-weight: bold;
}

main section.combat > .hitdice,
main section.combat > .deathsaves {
	flex: 1 50%;
	height: 100px;
}

main section.combat > .hitdice > div,
main section.combat > .deathsaves > div {
	height: 80px;
}

main section.combat > .hitdice > div {
	background-color: white;
	margin: 10px;
	border: 1px solid black;
	border-radius: 10px;
	display: flex;
	flex-direction: column;
}

main section.combat > .hitdice > div > .total {
	display: flex;
	align-items: baseline;
	justify-content: space-around;
	padding: 5px 0;
}

main section.combat > .hitdice > div > .total label {
	font-size: 10px;
	color: #bbb;
	margin: 0.25em;
	text-transform: none;
}

main section.combat > .hitdice > div > .total input {
	font-size: 12px;
	flex-grow: 1;
	border: 0;
	border-bottom: 1px solid #ddd;
	width: 50%;
	margin-right: 0.25em;
	padding: 0 0.25em;
	text-align: center;
}

main section.combat > .hitdice > div > .remaining {
	flex: 1;
	display: flex;
	flex-direction: column-reverse;
}

main section.combat > .hitdice > div > .remaining label {
	text-align: center;
	padding: 2px;
	font-size: 10px;
}

main section.combat > .hitdice > div > .remaining input {
	text-align: center;
	border: 0;
	flex: 1;
}

main section.combat > .deathsaves > div {
	margin: 10px;
	background: white;
	border: 1px solid black;
	border-radius: 10px;
	display: flex;
	flex-direction: column-reverse;
}

main section.combat > .deathsaves > div > .label {
	text-align: center;
}

main section.combat > .deathsaves > div > .label label {
	font-size: 10px;
}

main section.combat > .deathsaves > div > .marks {
	display: flex;
	flex: 1;
	flex-direction: column;
	justify-content: center;
}

main section.combat > .deathsaves > div > .marks .deathsuccesses,
main section.combat > .deathsaves > div > .marks .deathfails {
	display: grid;
	align-items: center;
	justify-content: center;
}

main section.combat > .deathsaves > div > .marks .deathsuccesses label,
main section.combat > .deathsaves > div > .marks .deathfails label {
	font-size: 8px;
	text-align: right;
}

main section.combat > .deathsaves > div .bubbles {
	margin-left: 5px;
}

main section.combat > .deathsaves > div .bubbles input[type="checkbox"] {
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	width: 10px;
	height: 10px;
	border: 1px solid black;
	border-radius: 10px;
	margin: 0;
}

main
	section.combat
	> .deathsaves
	> div
	.bubbles
	input[type="checkbox"]:checked {
	background-color: black;
}

main section.attacksandspellcasting {
	border: 1px solid black;
	border-radius: 10px;
	margin-top: 10px;
}

main section.attacksandspellcasting > div {
	margin: 10px;
	display: flex;
	flex-direction: column;
}

main section.attacksandspellcasting > div > label {
	order: 3;
	text-align: center;
}

main section.attacksandspellcasting > div > table {
	width: 100%;
}

main section.attacksandspellcasting > div > table th {
	font-size: 10px;
	color: #ddd;
}

main section.attacksandspellcasting > div > table input {
	width: calc(100% - 4px);
	border: 0;
	background-color: #eee;
	font-size: 10px;
	padding: 3px;
}

main section.attacksandspellcasting > div textarea {
	border: 0;
}

main section.equipment {
	border: 1px solid black;
	border-radius: 10px;
	margin-top: 10px;
}

main section.equipment > div {
	padding: 10px;
	display: flex;
	flex-direction: row;
	flex-wrap: wrap;
}

main section.equipment > div > .money ul {
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	height: 100%;
}

main section.equipment > div > .money ul > li {
	display: flex;
	align-items: center;
}

main section.equipment > div > .money ul > li label {
	border: 1px solid black;
	border-radius: 10px 0 0 10px;
	border-right: 0;
	width: 20px;
	font-size: 8px;
	text-align: center;
	padding: 3px 0;
}

main section.equipment > div > .money ul > li input {
	border: 1px solid black;
	border-radius: 10px;
	width: 40px;
	padding: 10px 3px;
	font-size: 12px;
	text-align: center;
}

main section.equipment > div > label {
	order: 3;
	text-align: center;
	flex: 100%;
}

main section.equipment > div > textarea {
	flex: 1;
	border: 0;
}

main section.flavor {
	padding: 10px;
	background: #bbb;
	border-radius: 10px;
}

main section.flavor div {
	background: white;
	display: flex;
	flex-direction: column-reverse;
	padding: 5px;
	border: 1px solid black;
}

main section.flavor div label {
	text-align: center;
	font-size: 10px;
	margin-top: 3px;
}

main section.flavor div textarea {
	border: 0;
	border-radius: 0;
	height: 4em;
}

main section.flavor :first-child {
	border-radius: 10px 10px 0 0;
}

main section.flavor :not(:first-child) {
	margin-top: 10px;
}

main section.flavor :last-child {
	border-radius: 0 0 10px 10px;
}

main section.features {
	padding: 10px;
}

main section.features div {
	padding: 10px;
	border: 1px solid black;
	border-radius: 10px;
	display: flex;
	flex-direction: column-reverse;
}

main section.features div label {
	text-align: center;
}

main section.features div textarea {
	border: 0;
	padding: 5px;
	height: 43em;
}
</style>
