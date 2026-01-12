<script lang="ts">
	import { enhance } from '$app/forms';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';

	import { RegisterAction } from '$lib/auth';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
</script>

<form use:enhance={RegisterAction} method="POST" class="space-y-6">
	<input type="hidden" name="next" value={page.url.searchParams.get('next')} />

	<Input placeholder="Email" name="email" type="text" />
	<Input placeholder="Password" name="password" type="password" />
	<Input placeholder="Confirm password" name="passwordConfirm" type="password" />

	{#if page.form?.response?.message}
		<p class="text-red-500">{page.form.response.message}</p>
	{/if}

	<Button type="submit">Register</Button>
	<Button href={resolve('/login')} variant="link">Already have an account</Button>
</form>
