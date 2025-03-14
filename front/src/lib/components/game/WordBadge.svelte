<script lang="ts">
	import type { ChosenWord } from '$lib/gameContext';
	import { IconLaurelWreath1, IconLaurelWreath2, IconLaurelWreath3 } from '@tabler/icons-svelte';

	let { word, fixedSize }: { word: ChosenWord, fixedSize: boolean } = $props();

	const getTier = (points: number) => {
		if (points >= 30) {
			return 'gold';
		}
		if (points >= 25) {
			return 'silver';
		}
		if (points >= 20) {
			return 'bronze';
		}
		return 'regular';
	};

	const tier = getTier(word.points);

	const getTierColor = (tier: string) => {
		switch (tier) {
			case 'bronze':
				return 'var(--color-bronze)';
			case 'silver':
				return 'var(--color-silver)';
			case 'gold':
				return 'var(--color-gold)';
			default:
				return 'var(--color-info)';
		}
	};

	const getTierFontSize = (tier: string) => {
		if (fixedSize)
			return 'text-lg';
		switch (tier) {
			case 'bronze':
				return 'text-xl';
			case 'silver':
				return 'text-2xl';
			case 'gold':
				return 'text-3xl';
			default:
				return 'text-lg';
		}
	};

	const getTierFontWeight = (tier: string) => {
		if (fixedSize)
			return 'font-normal';
		switch (tier) {
			case 'bronze':
				return 'font-bold';
			case 'silver':
				return 'font-bold';
			case 'gold':
				return 'font-bold';
			default:
				return 'font-normal';
		}
	};

	const getTierBadgeSize = (tier: string) => {
		if (fixedSize)
			return '';
		switch (tier) {
			case 'bronze':
				return 'badge-lg';
			case 'silver':
				return 'badge-xl';
			case 'gold':
				return 'badge-xl h-10';
			default:
				return '';
		}
	};
</script>

<div
	class="badge badge-info {getTierBadgeSize(tier)} border-0 {getTierFontWeight(
		tier
	)} {getTierFontSize(tier)}"
	style="background-color: {getTierColor(tier)}"
>
	{word.raw}
	<div class="ml-4">{word.points}</div>
	{#if tier === 'bronze'}
		<IconLaurelWreath3 />
	{:else if tier === 'silver'}
		<IconLaurelWreath2 />
	{:else if tier === 'gold'}
		<IconLaurelWreath1 />
	{/if}
</div>

<style>
</style>
