<script lang="ts">
import Icon from "@iconify/svelte";
import { register } from "$lib/auth";
import type { PageProps } from "../$types";
import type { Errors } from "$lib";

const { data }: PageProps = $props();
if (data.user && data.user.id) {
	window.location.href = "/";
}

let username: string = $state("");
let password: string = $state("");
let passwordRepeat: string = $state("");
let error: string = $state("");
let showPassword: boolean = $state(false);
let showPassword2: boolean = $state(false);
let errors: Errors = $state({} as Errors);

function handleSubmit(e: Event): void {
	e.preventDefault();
	if (password !== passwordRepeat) {
		error = "Le password non corrispondono!";
		return;
	}

	register(username, password)
		.then(() => {
			window.location.href = "/";
		})
		.catch((e: Error) => {
			errors = e.cause as Errors;
		});
};
</script>

<h1>Registrati</h1>
<div id="errors">
	{#if Object.entries(errors)}
		{#each Object.entries(errors) as [k, v]}
			<p><strong>{k}:</strong> {v}</p>
		{/each}
	{/if}
</div>
<form onsubmit={handleSubmit}>
	<div>
		<label for="username">Username</label>
		<input
			type="text"
			id="username"
			name="username"
			autocomplete="nickname"
			required
			bind:value={username} />
		<label for="password">Password</label>
		<span>
			<input
				type={showPassword ? "text" : "password"}
				id="password"
				name="password"
				autocomplete="new-password"
				required
				bind:value={password} />
			<button
				class="togglePassword"
				type="button"
				onclick={() => {
					showPassword = !showPassword;
				}}>
				{#if showPassword}
					<Icon icon="fa6-solid:eye" />
				{:else}
					<Icon icon="fa6-solid:eye-slash" />
				{/if}
			</button>
		</span>
		<label for="password_repeat">Ripeti Password</label>
		<span>
			<input
				type={showPassword2 ? "text" : "password"}
				id="password_repeat"
				name="password_repeat"
				required
				bind:value={passwordRepeat} />
			<button
				class="togglePassword"
				type="button"
				onclick={() => {
					showPassword2 = !showPassword2;
				}}>
				{#if showPassword2}
					<Icon icon="fa6-solid:eye" />
				{:else}
					<Icon icon="fa6-solid:eye-slash" />
				{/if}
			</button>
		</span>
	</div>
	<span class="buttonWrapper">
		<button type="submit">Registrati</button>
	</span>
</form>
<p>Hai già un account? <a href="/login">accedi</a></p>

<style>
form {
	background-color: var(--backgorund-secondary);
	padding: 1rem;
	border-radius: 8px;
	div {
		display: grid;
		grid-template-columns: auto auto;
		input {
			margin: 0.5rem 0 0.5rem 0;
			border: solid 1px black;
			border-radius: 8px;
			padding: 4px;
		}

		label {
			padding: 4px 0 4px 0;
			margin: 0.5rem 1rem 0.5rem 0;
		}

		span {
			position: relative;
		}

		button {
			position: absolute;
			top: 0;
			right: 0;
			margin: 0.5rem 0 0.5rem 0;
			border: solid 1px black;
			border-radius: 8px;
			padding: 4px;
			height: 1.51rem;
		}
	}
}

.buttonWrapper {
	width: 100%;
	display: flex;
	justify-content: center;
	align-items: center;
	padding-top: 2rem;
	button {
		width: 75%;
		border-radius: 8px;
		border: solid 1px darkgreen;
		background-color: green;
		padding: 4px;
	}
}

div#errors {
	width: 20rem;
	max-width: 80%;
	p {
		width: 100%;
		text-align: center;
		padding: 0.5rem;
		margin: 0.1rem;
		background-color: red;
	}
}
</style>
