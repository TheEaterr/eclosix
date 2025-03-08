<script lang="ts">
	import '../app.css';
	import { browser } from '$app/environment';
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	import { SvelteQueryDevtools } from '@tanstack/svelte-query-devtools';
	import { navigating } from '$app/state';
	import { writable } from 'svelte/store';
	import { setContext } from 'svelte';

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser
			}
		}
	});

	// put this in global layout to pass to background
	const gameWon = writable<boolean>(false);
	setContext('gameWon', gameWon);
</script>

<QueryClientProvider client={queryClient}>
	{#if navigating.to}
		<div class="flex h-screen w-screen items-center justify-center">
			<span class="animate-suspense w-40 text-neutral opacity-0">
				<span class="loading loading-dots text-neutral-content"></span>
			</span>
		</div>
	{:else}
		<slot />
	{/if}
	{#if import.meta.env.MODE === 'development'}
		<SvelteQueryDevtools />
	{/if}
</QueryClientProvider>

<style lang="postcss">
	@keyframes fade-in {
		0% {
			opacity: 0;
		}
		100% {
			opacity: 1;
		}
	}
	@keyframes suspense {
		0% {
			opacity: 1;
		}
		100% {
			opacity: 1;
		}
	}

	.animate-suspense {
		animation: suspense 5s cubic-bezier(0.4, 0, 0.6, 1) infinite;
		animation-delay: 300ms;
	}
</style>
