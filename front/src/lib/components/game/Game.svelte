<script lang="ts">
	import { getContext } from 'svelte';
	import { type Writable } from 'svelte/store';
	import { getTopMatches, testWord, type Problem } from '$lib/pocketBase';
	import LetterButton from './LetterButton.svelte';
	import {
		IconTrophy,
		IconTrash,
		IconArrowLeft,
		IconExclamationCircle,
		IconShare
	} from '@tabler/icons-svelte';
	import { getNumberOfPoints } from './utils';
	import WordBadge from './WordBadge.svelte';
	import { goto } from '$app/navigation';
	import type { ChosenWord } from '$lib/gameContext';
	import { shuffleArray } from '$lib/hash';

	let { problem, isDaily }: { problem: Problem; isDaily: boolean } = $props();

	const chosenWords = getContext<Writable<ChosenWord[]>>('chosenWords');
	const points = getContext<Writable<number>>('points');
	const gameWon = getContext<Writable<boolean>>('gameWon');
	gameWon.set(false);

	let inputField: HTMLInputElement | null = $state(null);

	const setBonusLetter = () => {
		bonusLetter = bonusLetters[$chosenWords.length % bonusLetters.length];
	};
	const sideLetters = $derived(
		problem.availableLetters.filter((letter) => letter !== problem.centerLetter)
	);
	const bonusLetters = $derived(shuffleArray([...sideLetters], problem.id));
	$effect(() => {
		setBonusLetter();
	});

	let top_matches: string[] | null = $state(null);
	let currentWord = $state('');
	let bonusLetter = $state('');
	let showAlert = $state(false);
	let alertMessage = $state('');
	let showShared = $state(false);
	let gameWonMessage = $state('');
	let gameWonClass = $state('');

	const onLetterButtonClick = (letter: string) => {
		showAlert = false;
		currentWord += letter.toLowerCase();
	};

	const submitWord = async () => {
		if (currentWord.length <= 3) {
			alertMessage = 'Le mot doit contenir au moins 4 lettres.';
			showAlert = true;
			return;
		}

		if (!currentWord.includes(problem.centerLetter)) {
			alertMessage = 'Le mot doit contenir la lettre centrale.';
			showAlert = true;
			return;
		}

		if ($chosenWords.map((word) => word.word).includes(currentWord)) {
			alertMessage = 'Ce mot a déjà été soumis.';
			showAlert = true;
			return;
		}

		const rawWord = await testWord(currentWord.toLowerCase());
		if (!rawWord) {
			alertMessage = "Ce mot n'existe pas.";
			showAlert = true;
			return;
		}

		const newPoints = getNumberOfPoints(currentWord.toLowerCase(), problem, bonusLetter);
		chosenWords.update((words) => {
			const newChosenWord = { word: currentWord, raw: rawWord, points: newPoints };
			return [...words, newChosenWord];
		});
		points.update((points) => points + newPoints);
		currentWord = '';
		setBonusLetter();
	};

	$effect(() => {
		if ($chosenWords.length == 12) {
			gameWon.set(true);
			if ($points >= 350) {
				gameWonMessage = `CLASSE EXCEPTIONNELLE`;
				gameWonClass = 'gold';
			} else if ($points >= 300) {
				gameWonMessage = `PREMIÈRE CLASSE`;
				gameWonClass = 'silver';
			} else if ($points >= 250) {
				gameWonMessage = `DEUXIÈME CLASSE`;
				gameWonClass = 'bronze';
			}
			getTopMatches(problem).then((matches) => {
				top_matches = matches;
			});
		}
	});

	const shareToClipboard = async () => {
		if (isDaily) {
			await navigator.clipboard.writeText(
				`J'ai réussi à atteindre un score de ${$points}🏆 sur Éclosix 🌸 !\nEssaie de me battre sur https://eclosix.maoune.fr/game/daily 🎯 avant qu'il expire ⏰ !`
			);
		} else {
			await navigator.clipboard.writeText(
				`J'ai réussi à atteindre un score de ${$points}🏆 sur Éclosix 🌸 !\nEssaie de me battre sur https://eclosix.maoune.fr/game/custom/${problem.id} 🎯 !`
			);
		}
		showShared = true;
	};

	const onKeyDown = (event: KeyboardEvent) => {
		// check if event is one of the letters
		const letter = event.key.toLowerCase();
		if (inputField && problem.availableLetters.includes(letter)) {
			inputField.focus();
		}
	};

	const getWordFromMatch = (match: string) => {
		return $chosenWords.filter((word) => word.word == match)[0];
	};
</script>

<div class="flex flex-col justify-center gap-5">
	<div class="m-auto flex max-w-screen-sm flex-col flex-wrap items-center justify-center gap-5">
		<h3 class="small-title text-primary text-3xl font-bold">Éclosix</h3>
		{#if !$gameWon}
			<fieldset class="fieldset">
				<legend class="fieldset-legend text-sm">Écrivez ou utilisez les boutons</legend>
				<input
					bind:this={inputField}
					type="text"
					value={currentWord}
					placeholder="_"
					class="input input-primary input-xl w-full uppercase"
					oninput={(event) => {
						showAlert = false;
						const target = event.target as HTMLInputElement;
						const value = target.value;
						// keep only available letters
						const availableLetters = problem.availableLetters;
						const letters = value
							.split('')
							.filter((letter) => availableLetters.includes(letter.toLowerCase()));
						target.value = letters.join('').toLowerCase();
						currentWord = target.value;
					}}
					onkeyup={(event) => {
						if (event.key === 'Enter') {
							submitWord();
						}
					}}
				/>
			</fieldset>
			{#if showAlert}
				<div role="alert" class="alert alert-error m-[-0.625rem]">
					<IconExclamationCircle size={30} class="stroke-current" />
					<span class="font-semibold">{alertMessage}</span>
				</div>
			{/if}
			<div class="flex flex-row flex-wrap items-center justify-center gap-2">
				<button
					class="btn btn-neutral-special btn-square btn-xl"
					onclick={() => {
						showAlert = false;
						currentWord = currentWord.slice(0, currentWord.length - 1);
					}}
				>
					<IconArrowLeft size={30} />
				</button>
				<button
					class="btn btn-error-special btn-square btn-xl"
					onclick={() => {
						showAlert = false;
						currentWord = '';
					}}
				>
					<IconTrash size={30} />
				</button>
				<button class="btn btn-primary-special btn-xl" onclick={() => submitWord()}
					>Soumettre</button
				>
			</div>
			<div class="flex flex-row">
				<div class="mt-10 flex flex-col">
					<LetterButton
						letter={sideLetters[0]}
						onClick={onLetterButtonClick}
						isCenter={false}
						{bonusLetter}
					/>
					<LetterButton
						letter={sideLetters[1]}
						onClick={onLetterButtonClick}
						isCenter={false}
						{bonusLetter}
					/>
				</div>
				<div class="mr-[-0.625rem] ml-[-0.625rem] flex flex-col">
					<LetterButton
						letter={sideLetters[2]}
						onClick={onLetterButtonClick}
						isCenter={false}
						{bonusLetter}
					/>
					<LetterButton
						letter={problem.centerLetter}
						onClick={onLetterButtonClick}
						isCenter={true}
						{bonusLetter}
					/>
					<LetterButton
						letter={sideLetters[3]}
						onClick={onLetterButtonClick}
						isCenter={false}
						{bonusLetter}
					/>
				</div>
				<div class="mt-10 flex flex-col">
					<LetterButton
						letter={sideLetters[4]}
						onClick={onLetterButtonClick}
						isCenter={false}
						{bonusLetter}
					/>
					<LetterButton
						letter={sideLetters[5]}
						onClick={onLetterButtonClick}
						isCenter={false}
						{bonusLetter}
					/>
				</div>
			</div>
		{/if}
		<div class="stats stats-vertical bg-base-200 m-5 shadow">
			<div class="stat">
				<div class="stat-figure text-primary">
					<IconTrophy size={30} />
				</div>
				<div class="stat-title mb-2 text-lg font-semibold">Points</div>
				<div class="small-title stat-value text-primary">
					{#if $gameWon}
						{$points} / {problem.maxPoints} ({Math.round(($points / problem.maxPoints) * 100)}%)
					{:else}
						{$points}
					{/if}
				</div>
			</div>
			{#if $gameWon}
				{#if gameWonMessage}
					<div class="stat">
						<div class="stat-title mb-2 text-lg font-semibold">Félicitations vous êtes de</div>
						<div class="stat-value justify-self-center">
							<h3
								class="small-title text-center text-3xl font-bold text-wrap"
								style="color: var(--color-{gameWonClass})"
							>
								{gameWonMessage}
							</h3>
						</div>
					</div>
				{/if}
				<div class="stat">
					<div class="stat-value justify-self-center">
						<button
							onclick={shareToClipboard}
							onfocusout={() => (showShared = false)}
							class="btn-secondary-special btn text-lg"><IconShare />Partage ton score</button
						>
					</div>
					{#if showShared}
						<div class="goal-stat-desc stat-desc" id="description-tooltip-share">
							Copié dans le presse-papier.
						</div>
					{/if}
				</div>
				<div class="stat">
					<div class="stat-value justify-self-center">
						<button
							onclick={() => goto('/game/endless')}
							class="btn-primary-special btn m-auto mb-0 w-fit text-lg">Rejouer</button
						>
					</div>
				</div>
				<div class="stat">
					<div class="stat-title mb-2 text-lg font-semibold">Meilleurs mots</div>
					{#if top_matches}
						<div class="text-neutral flex flex-wrap justify-evenly gap-1 text-sm">
							{#each top_matches as match (match)}
								{#if $chosenWords.map((word) => word.word).includes(match)}
									<span>
										<WordBadge word={getWordFromMatch(match)} fixedSize={true} />
									</span>
								{:else}
									<span>
										{match}
									</span>
								{/if}
							{/each}
						</div>
					{:else}
						<div class="stat-value justify-self-center">
							<span class="loading loading-dots loading-lg text-neutral"></span>
						</div>
					{/if}
				</div>
			{/if}
			{#if $chosenWords.length > 0}
				<div class="stat">
					<div class="stat-title mb-2 text-lg font-semibold">
						Mots soumis ({$chosenWords.length}/12)
					</div>
					<div class="flex flex-col flex-wrap items-center gap-2">
						{#each $chosenWords.toReversed() as word (word)}
							<WordBadge {word} fixedSize={false} />
						{/each}
					</div>
				</div>
			{/if}
		</div>
	</div>
</div>

<svelte:window on:keydown={onKeyDown} />

<style>
</style>
