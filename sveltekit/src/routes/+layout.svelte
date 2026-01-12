<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';

	import { Toaster } from '$lib/components/ui/sonner';

	import './layout.css';

	const { children } = $props();
</script>

{@render children?.()}

<Toaster />
<ModeWatcher />
