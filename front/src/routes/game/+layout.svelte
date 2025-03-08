<script lang="ts">
	import { beforeNavigate } from '$app/navigation';
	import { getContext, setContext } from 'svelte';
	import { type Snapshot } from '@sveltejs/kit';
	import type { GameContext } from '$lib/gameContext';
	import { writable, type Writable } from 'svelte/store';

	let { children } = $props();

	const chosenWords = writable<string[]>([]);
	const points = writable<number>(0);
	const problemId = writable<string>('');
	const gameStarted = writable<boolean>(false);
	const gameWon = getContext<Writable<boolean>>('gameWon');
	setContext('chosenWords', chosenWords);
	setContext('points', points);
	setContext('problemId', problemId);
	setContext('gameStarted', gameStarted);

	beforeNavigate(() => {
		// since gameWon is set before, we have to take care
		// to leave it in a coherent state
		gameWon.set(false);
	});

	const snapshot: Snapshot<GameContext | undefined> = {
		capture: () => {
			if (!$gameWon) {
				return { chosenWords: $chosenWords, points: $points, problemId: $problemId };
			}
		},
		restore: async (value) => {
			if (!value || !value.chosenWords) {
				value = {
					chosenWords: [],
					points: 0,
					problemId: ''
				};
			}
			chosenWords.set(value.chosenWords);
			points.set(value.points);
		}
	};

	const reset = async () => {
		chosenWords.set([]);
		points.set(0);
		problemId.set('');
		gameWon.set(false);
	};

	setContext('fixedReset', reset);
	setContext('fixedSnapshot', snapshot);
</script>

<div class="flex flex-col content-center items-center">
	{@render children?.()}
</div>
