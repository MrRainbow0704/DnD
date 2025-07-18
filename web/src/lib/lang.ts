import {
	setLocale,
	getLocale,
	locales,
	type Locale,
} from "$lib/paraglide/runtime";
import { m } from "$lib/paraglide/messages";

export const langs = locales;

export type Lang = Locale;

export function setLang(l: Lang): void {
	setLocale(l);
}

export function getLang(): Lang {
	return getLocale();
}

export function t<K extends keyof typeof m>(
	k: K,
	...args: Parameters<(typeof m)[K]>
): string {
	try {
		return (m[k] as (...args: any[]) => string)(...args);
	} catch {
		return k
	}
}

export function langName(l: Lang): string {
	return new Intl.DisplayNames(l, {
		type: "language",
	}).of(l)!;
}

export function langFlag(l: Lang): string {
	if (l === "en") {
		l = "gb" as Lang;
	}
	let first = l.charCodeAt(0) + 127365;
	let second = l.charCodeAt(1) + 127365;
	let flag = `\\0${first.toString(16)}\\0${second.toString(16)}`;
	return flag;
}
