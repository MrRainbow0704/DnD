<script lang="ts">
import { login } from "$src/lib/auth";
import type { PageProps } from "../$types";

const { data }: PageProps = $props();
if (data.user && data.user.id) {
	window.location.href = "/";
}

let username: string = $state("");
let password: string = $state("");
let error: string = $state("");

const handleSubmit = (e: Event) => {
	e.preventDefault();

	login(username, password).then((errors) => {
		if (Object.keys(errors).length) {
			let errorMessage = "";
			Object.entries(errors).forEach((e) => {
				errorMessage += `${e[1]}\n`;
			});
			error = errorMessage;
		} else {
			window.location.href = "/";
		}
	});
};
</script>

<h1>Accedi</h1>
{#if error.length > 0}
	<div class="error">
		<p>{error}</p>
	</div>
{/if}
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
