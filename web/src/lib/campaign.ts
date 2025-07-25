import type { Errors } from "$lib";
import type { Character } from "$lib/character";

export type Campaign = {
	master: number;
	players: Character[];
};

export async function getCampaign(id: number): Promise<Campaign> {
	const res = await fetch(`/api/v1/campaigns/${id}`, {
		method: "GET",
		headers: {
			"Content-Type": "application/json",
		},
	});
	const { errors, campaigns }: { errors: Errors; campaigns: Campaign } =
		await res.json();
	if (Object.entries(errors).length) {
		throw new Error("Errore in getCampaigns()", { cause: errors });
	}
	return campaigns;
}
