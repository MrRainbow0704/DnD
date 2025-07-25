import type { User } from "$src/lib/user";
import { browser } from "$app/environment";

export const prerender = true;
export const ssr = true;

export async function load() {
	if (!browser) {
		return {
			user: {} as User,
		};
	}
	return {
		user: JSON.parse(localStorage.getItem("user") || "{}") as User,
	};
}
