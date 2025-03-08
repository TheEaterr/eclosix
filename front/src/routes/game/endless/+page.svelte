<script lang="ts">
	import { getProblemFromId, getRandomProblem, type Problem } from '$lib/pocketBase';
	import { getContext } from 'svelte';
	import type { Snapshot } from './$types';
	import { type Writable } from 'svelte/store';
	import type { GameContext } from '$lib/gameContext';
	import GamePage from '$lib/components/GamePage.svelte';

	let { data } = $props();

	const chosenWords = getContext<Writable<string[]>>('chosenWords');
	const points = getContext<Writable<number>>('points');
	const problemId = getContext<Writable<string>>('problemId');
	const gameWon = getContext<Writable<boolean>>('gameWon');

	let problem: Problem | undefined = $state(data);
	points.set(0);
	chosenWords.set([]);
	problemId.set(data.id);

	export const snapshot: Snapshot<GameContext | undefined> = {
		capture: () => {
			if (!$gameWon) {
				return { points: $points, chosenWords: $chosenWords, problemId: $problemId };
			}
		},
		restore: async (value) => {
			if (!value || !value.chosenWords || !value.problemId) {
				value = {
					chosenWords: [],
					points: 0,
					problemId: ''
				};
			} else {
				problem = await getProblemFromId(value.problemId);
			}
			chosenWords.set(value.chosenWords);
			points.set(value.points);
			problemId.set(value.problemId);
		}
	};

	const reset = async () => {
		points.set(0);
		gameWon.set(false);
		chosenWords.set([]);
		// set to undefined before awaiting to get loading indicator
		problem = undefined;
		problem = await getRandomProblem();
		console.log(problem);
		// set after awaiting the goal taxon so the query is made
		// with goal taxon data.
		problemId.set(problem.id);
	};
</script>

<svelte:head>
	<title>Infini - Éclosix</title>
	<meta name="description" content="Trouver les mots les plus longs possibles avec 7 lettres." />
</svelte:head>

<GamePage isDaily={false} {problem} {reset} />
