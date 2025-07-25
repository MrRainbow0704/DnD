import type { Errors } from "$lib";
import type { User } from "$lib/user";

export async function register(
	username: string,
	password: string
): Promise<void> {
	const res = await fetch("/api/v1/user", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ username: username, password: password }),
	});
	const { errors, user }: { errors: Errors; user: User } = await res.json();
	if (!Object.entries(errors) && res.ok && user) {
		return login(username, password);
	}
	throw new Error("Errore in register()", { cause: errors });
}

export async function login(username: string, password: string): Promise<void> {
	const res = await fetch("/api/v1/login", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ username: username, password: password }),
	});
	const {
		errors,
		token,
		user,
	}: { errors: Errors; token: string; user: User } = await res.json();
	if (!Object.entries(errors) && res.ok && token && user) {
		localStorage.setItem("user", JSON.stringify(user));
		return;
	}
	throw new Error("Errore in login()", { cause: errors });
}

export async function logout(): Promise<void> {
	await fetch("/api/v1/logout", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
	});
	localStorage.removeItem("user");
}
