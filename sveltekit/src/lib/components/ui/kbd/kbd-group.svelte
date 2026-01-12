<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	import { type WithElementRef, cn } from '$lib/components/utils.js';

	let {
		ref = $bindable(null),
		class: className,
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLElement>> = $props();
</script>

<kbd
	bind:this={ref}
	data-slot="kbd-group"
	class={cn('inline-flex items-center gap-1', className)}
	{...restProps}
>
	{@render children?.()}
</kbd>
