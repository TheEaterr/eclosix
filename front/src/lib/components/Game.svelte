<script lang="ts">
	import { getContext } from 'svelte';
	import { type Writable } from 'svelte/store';
	import type { Problem } from '$lib/pocketBase';

	let { problem, isDaily }: { problem: Problem; isDaily: boolean } = $props();

	const chosenWords = getContext<Writable<string[]>>('chosenWords');
	const points = getContext<Writable<number>>('points');
	const gameWon = getContext<Writable<boolean>>('gameWon');
	gameWon.set(false);

	// const shareToClipboard = async () => {
	// 	if (goalTaxonData === undefined) return;
	// 	if (isDaily) {
	// 		await navigator.clipboard.writeText(
	// 			`I was able to reach the daily taxon ${goalTaxonData.scientific} in ${$numberSteps} steps 🏆 on Taxonomicle 🐻!\nTry it yourself at https://taxonomicle.com/game/daily 🎯 before it expires ⏰!`
	// 		);
	// 	} else {
	// 		await navigator.clipboard.writeText(
	// 			`I was able to reach the taxon ${goalTaxonData.scientific} in ${$numberSteps} steps 🏆 on Taxonomicle 🐻!\nTry it yourself at https://taxonomicle.com/game/custom/${goalTaxonData.id} 🎯!`
	// 		);
	// 	}
	// 	toggleTooltip(true);
	// };
</script>

<div class="flex flex-col justify-center gap-5">
	<div class="m-auto flex max-w-screen-md flex-wrap items-center justify-center gap-5">
		TODO chosenWords: {JSON.stringify($chosenWords)}
		points: {$points}
		availableLetters: {problem.availableLetters}
		centerLetter: {problem.centerLetter}
		id: {problem.id}
		isDaily: {isDaily}
		gameWon: {$gameWon}
	</div>
</div>

<style>
</style>
