<script lang="ts">
import type { Errors } from "$lib";
import { login } from "$lib/auth";
import type { PageProps } from "../$types";

const { data }: PageProps = $props();
if (data.user && data.user.id) {
	window.location.href = "/";
}

let username: string = $state("");
let password: string = $state("");
let errors: Errors = $state({} as Errors);

function handleSubmit(e: Event): void {
	e.preventDefault();

	login(username, password)
		.then(() => {
			window.location.href = "/";
		})
		.catch((e: Error) => {
			errors = e.cause as Errors;
		});
}
</script>

<h1>Accedi</h1>
<div id="errors">
	{#if Object.entries(errors)}
		{#each Object.entries(errors) as [k, v]}
			<p><strong>{k}:</strong> {v}</p>
		{/each}
	{/if}
</div>
<form onsubmit={handleSubmit}>
	<label for="username">Username</label>
	<input
		type="text"
		id="username"
		name="username"
		required
		bind:value={username} />
	<label for="password">Password</label>
	<input
		type="text"
		id="password"
		name="password"
		required
		bind:value={password} />
	<button type="submit">Login</button>
</form>

<style>
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
