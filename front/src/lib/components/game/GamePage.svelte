<script lang="ts">
	import { getContext } from 'svelte';
	import { type Writable } from 'svelte/store';
	import Game from '$lib/components/game/Game.svelte';
	import Header from '$lib/components/header/Header.svelte';
	import type { Problem } from '$lib/pocketBase';
	import LoadingLogo from '../LoadingLogo.svelte';
	import type { ChosenWord } from '$lib/gameContext';

	const chosenWords = getContext<Writable<ChosenWord[]>>('chosenWords');
	const points = getContext<Writable<number>>('points');

	let { problem, isDaily, reset }: { problem?: Problem; isDaily: boolean; reset: () => void } =
		$props();

	chosenWords.set([]);
	points.set(0);
</script>

{#if problem}
	<Header {reset} showHowToPlay={true} />
	<Game {problem} {isDaily} />
{:else}
	<div class="flex h-screen w-screen items-center justify-center">
		<span class="animate-suspense text-neutral w-40 opacity-0">
			<LoadingLogo />
		</span>
	</div>
{/if}
