<script lang="ts">
	import '../app.css';
	import { navigating } from '$app/state';
	import { writable } from 'svelte/store';
	import { setContext } from 'svelte';
	import DecoratedBackground from '../decorations/DecoratedBackground.svelte';
	import LoadingLogo from '$lib/components/LoadingLogo.svelte';

	let { children } = $props();

	// put this in global layout to pass to background
	const gameWon = writable<boolean>(false);
	setContext('gameWon', gameWon);
</script>

<DecoratedBackground />

{#if navigating.to}
	<div class="flex h-screen w-screen items-center justify-center">
		<span class="animate-suspense text-neutral w-40 opacity-0">
			<LoadingLogo />
		</span>
	</div>
{:else}
	{@render children?.()}
{/if}

<style>
</style>
