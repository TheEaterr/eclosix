<script lang="ts">
	import { getContext } from 'svelte';
	import { type Writable } from 'svelte/store';
	import Game, { type GameTypes } from '$lib/components/game/Game.svelte';
	import Header from '$lib/components/header/Header.svelte';
	import type { Problem } from '$lib/pocketBase';
	import LoadingLogo from '../LoadingLogo.svelte';
	import type { ChosenWord } from '$lib/gameContext';

	const chosenWords = getContext<Writable<ChosenWord[]>>('chosenWords');
	const points = getContext<Writable<number>>('points');

	let { problem, gameType, reset }: { problem?: Problem; gameType: GameTypes; reset: () => void } =
		$props();

	chosenWords.set([]);
	points.set(0);
</script>

{#if problem}
	<Header {reset} showHowToPlay={true} />
	<Game {problem} {reset} {gameType} />
{:else}
	<div class="flex h-screen w-screen items-center justify-center">
		<span class="animate-suspense text-neutral w-40 opacity-0">
			<LoadingLogo />
		</span>
	</div>
{/if}
