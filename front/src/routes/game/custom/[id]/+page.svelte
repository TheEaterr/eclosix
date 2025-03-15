<script lang="ts">
	import GamePage from '$lib/components/game/GamePage.svelte';
	import { getContext } from 'svelte';
	import { type Snapshot } from '@sveltejs/kit';
	import type { GameContext } from '$lib/gameContext';
	import type { Writable } from 'svelte/store';

	let { data } = $props();
	const problemId = getContext<Writable<string>>('problemId');
	problemId.set(data.id);

	const fixedReset = getContext<() => void>('fixedReset');
	export const snapshot = getContext<Snapshot<GameContext | undefined>>('fixedSnapshot');
</script>

<svelte:head>
	<title>Personnalisée - Éclosix</title>
	<meta name="description" content="Trouver les mots les plus longs possibles avec 7 lettres." />
</svelte:head>

<GamePage gameType="custom" problem={data} reset={fixedReset} />
