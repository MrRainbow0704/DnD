<script lang="ts">
import Icon from "@iconify/svelte";
import type { User } from "$lib/user";
import { getLang, setLang, langs, langName, type Lang } from "$lib/lang";
import { onMount } from "svelte";

const { user }: { user: User } = $props();
const isLoggedIn: boolean = user && user.id ? true : false;

let root: HTMLBodyElement;
let langDialog: HTMLDialogElement;
let darkmode: boolean = $state(false);
let newLang: Lang = $state(getLang());

const toggleDarkmode = (old: string) => {
	if (old === "true") {
		root.classList.remove("dark");
		darkmode = false;
		localStorage.setItem("dark", "false");
	} else {
		root.classList.add("dark");
		darkmode = true;
		localStorage.setItem("dark", "true");
	}
};

onMount(() => {
	darkmode = localStorage.getItem("dark") === "true";
	if (localStorage.getItem("dark") === "true") {
		root.classList.add("dark");
	} else {
		root.classList.remove("dark");
	}
	window
		.matchMedia("(prefers-color-scheme: dark)")
		.addEventListener("change", (e) => {
			toggleDarkmode(`${!e.matches}`);
		});

	return () => {
		window
			.matchMedia("(prefers-color-scheme: dark)")
			.removeEventListener("change", () => {});
	};
});
</script>

<svelte:body bind:this={root} />
<nav>
	<span>
		<a href="/"><Icon icon="fa6-solid:house" />&nbsp;Home</a>
		<a href="/campaigns"><Icon icon="fa6-solid:map" />&nbsp;Campagne</a>
		<a href="/characters"
			><Icon icon="fa6-solid:person" />&nbsp;Personaggi</a>
		{#if isLoggedIn}
			<a href="/profile"
				><Icon icon="fa6-solid:circle-user" />&nbsp;Profilo</a>
		{/if}
	</span>
	<span>
		{#if isLoggedIn}
			<a href="/logout"
				><Icon icon="fa6-solid:user-slash" />&nbsp;Logout</a>
		{:else}
			<a href="/register">
				<Icon icon="fa6-solid:user-plus" />&nbsp;Registrati
			</a>
			<a href="/login"><Icon icon="fa6-solid:user-check" />&nbsp;Login</a>
		{/if}
		<button
			class="link"
			type="button"
			onclick={() => langDialog.showModal()}>
			<Icon icon={`circle-flags:${getLang()}`} />&nbsp;Seleziona Lingua
		</button>
		<dialog bind:this={langDialog}>
			<select
				bind:value={newLang}
				onchange={() => {
					setLang(newLang);
				}}>
				{#each langs as l}
					<option value={l} selected={l === getLang() ? true : false}>
						<Icon icon={`circle-flags:${l}`} />
						{langName(l)}
					</option>
				{/each}
			</select>
			<button type="button" onclick={() => langDialog.close()}>
				<Icon icon="fa6-solid:x" />
			</button>
		</dialog>
		<button
			class="link"
			type="button"
			onclick={() => {
				toggleDarkmode(localStorage.getItem("dark")!);
			}}>
			{#if darkmode}
				<Icon icon="fa6-solid:sun" />&nbsp;Modalità Chiara
			{:else}
				<Icon icon="fa6-solid:moon" />&nbsp;Modalità Scura
			{/if}
		</button>
	</span>
</nav>

<style>
nav {
	display: flex;
	flex-direction: row;
	flex-wrap: nowrap;
	align-items: center;
	justify-content: space-between;
	padding: 1rem;
	background-color: var(--backgorund-accent);
}

a,
button.link {
	display: inline-block;
	padding: 0 0.5rem 0 0.5rem;
	margin: none;
}

button.link {
	background: none;
	border: none;
}
</style>
