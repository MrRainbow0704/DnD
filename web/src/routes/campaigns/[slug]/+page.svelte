<script lang="ts">
import { page } from "$app/state";
import { getCampaign, type Campaign } from "$src/lib/campaign";
import type { Errors } from "$src/lib";

let campaign: Campaign = $state({} as Campaign);
let errors: Errors = $state({} as Errors);
const id = parseInt(page.params.slug);
getCampaign(id)
	.then((c) => (campaign = c))
	.catch((e: Error) => {
		errors = e.cause as Errors;
	});
</script>

<div id="errors">
	{#if Object.entries(errors)}
		{#each Object.entries(errors) as [k, v]}
			<p><strong>{k}:</strong> {v}</p>
		{/each}
	{/if}
</div>
{campaign}

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
