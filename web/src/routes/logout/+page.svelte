<script lang="ts">
import { logout } from "$src/lib/auth";

logout().then(() => {
	window.location.href = "/";
});
</script>
