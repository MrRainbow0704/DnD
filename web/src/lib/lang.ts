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
	const s =  new Intl.DisplayNames(l, {
		type: "language",
	}).of(l)!;
	return String(s[0]).toUpperCase() + String(s).slice(1);
}
