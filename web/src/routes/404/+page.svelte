<script lang="ts">
import { browser } from "$app/environment";

let p: string = ""
if (browser) {
	p = window.location.pathname;
}
</script>

<h1>Errore 404, Pagina non trovata</h1>
<p>
	Non esiste una pagina a {p}. Controlla di aver digitato l'URL correttamente.
</p>
