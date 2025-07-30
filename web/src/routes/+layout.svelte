<script lang="ts">
import Navbar from "$components/Navbar.svelte";
import Footer from "$components/Footer.svelte";
import type { LayoutProps } from "./$types";
import "$src/app.css";

let { data, children }: LayoutProps = $props();
</script>

<Navbar user={data.user} />
<main>
	{@render children()}
</main>
<Footer />

<style>
main {
	display: flex;
	flex-direction: column;
	align-items: center;
	padding: 0.5rem;
}
</style>
