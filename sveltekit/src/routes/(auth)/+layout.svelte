<script lang="ts">
	import * as Card from '$lib/components/ui/card';

	const { children } = $props();
</script>

<div class="min-h-screen flex items-center justify-center">
	<Card.Root class="max-w-md w-full">
		<Card.Header>
			<Card.Title>Login</Card.Title>
		</Card.Header>
		<Card.Content>
			{@render children?.()}
		</Card.Content>
	</Card.Root>
</div>
