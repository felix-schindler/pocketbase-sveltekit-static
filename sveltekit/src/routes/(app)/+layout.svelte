<script lang="ts">
	import { Logout, user } from '$lib/auth';
	import { Button } from '$lib/components/ui/button';

	let { children } = $props();
</script>

<nav class="flex justify-end items-center py-2 px-4 bg-secondary text-secondary-foreground shadow">
	<div class="mr-4">{$user.email}</div>
	<Button type="button" onclick={Logout}>Logout</Button>
</nav>

{@render children?.()}
