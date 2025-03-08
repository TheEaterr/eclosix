<script lang="ts">
	import { getContext } from 'svelte';
	import { type Writable } from 'svelte/store';
	import { testWord, type Problem } from '$lib/pocketBase';
	import LetterButton from './LetterButton.svelte';
	import { IconTrophy } from '@tabler/icons-svelte';
	import { cyrb128 } from '$lib/hash';

	let { problem, isDaily }: { problem: Problem; isDaily: boolean } = $props();

	const chosenWords = getContext<Writable<string[]>>('chosenWords');
	const points = getContext<Writable<number>>('points');
	const gameWon = getContext<Writable<boolean>>('gameWon');
	gameWon.set(false);

	const setBonusLetter = () => {
		// hash the id
		const hash = cyrb128(problem.id + $chosenWords.length);
		const nonCenterLetters = problem.availableLetters.filter(
			(letter) => letter !== problem.centerLetter
		);
		const index = hash[0] % nonCenterLetters.length;
		bonusLetter = nonCenterLetters[index];
	};

	let currentWord = $state('');
	let bonusLetter = $state('');
	setBonusLetter();

	const onLetterButtonClick = (letter: string) => {
		currentWord += letter.toUpperCase();
	};

	const getNumberOfPoints = (word: string): number => {
		let wordPoints = 0;
		if (word.length == 4) {
			wordPoints = 2;
		} else if (word.length == 5) {
			wordPoints = 4;
		} else if (word.length == 6) {
			wordPoints = 6;
		} else if (word.length == 7) {
			wordPoints = 12;
		} else {
			wordPoints = 12 + (word.length - 7) * 3;
		}

		if (word.includes(bonusLetter)) {
			wordPoints = wordPoints + 5;
		}

		let hasAllLetters = true;
		for (const letter of problem.availableLetters) {
			if (!word.includes(letter)) {
				hasAllLetters = false;
				break;
			}
		}
		if (hasAllLetters) {
			wordPoints = wordPoints + 7;
		}

		return wordPoints;
	};

	const submitWord = async () => {
		if (currentWord.length < 3) {
			return;
		}

		const isValid = await testWord(currentWord.toLowerCase());
		if (!isValid) {
			return;
		}

		chosenWords.update((words) => {
			if (words.includes(currentWord)) {
				return words;
			}
			return [...words, currentWord];
		});
		points.update((points) => points + getNumberOfPoints(currentWord.toLowerCase()));
		currentWord = '';
		setBonusLetter();
	};

	$effect(() => {
		if ($chosenWords.length == 12) {
			gameWon.set(true);
		}
	});

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
	<div class="m-auto flex max-w-screen-md flex-col flex-wrap items-center justify-center gap-5">
		isDaily: {isDaily}
		id: {problem.id}
		<div class="stats stats-vertical bg-base-200 lg:stats-horizontal shadow">
			<div class="stat">
				<div class="stat-figure text-primary">
					<IconTrophy size={30} />
				</div>
				<div class="stat-title text-absolute font-semibold">Points</div>
				<div class="small-title stat-value text-primary">{$points}</div>
			</div>
		</div>
		<div class="flex flex-row flex-wrap gap-2">
			{#each $chosenWords as word (word)}
				<div class="badge badge-primary">{word}</div>
			{/each}
		</div>
		<div class="join">
			<input
				type="text"
				value={currentWord}
				placeholder="Écrivez ou utilisez les boutons..."
				class="input input-primary input-xl w-100 rounded-l-xl"
				oninput={(event) => {
					const target = event.target as HTMLInputElement;
					const value = target.value;
					// keep only available letters
					const availableLetters = problem.availableLetters;
					const letters = value
						.split('')
						.filter((letter) => availableLetters.includes(letter.toLowerCase()));
					target.value = letters.join('').toUpperCase();
					currentWord = target.value;
				}}
				onkeyup={(event) => {
					if (event.key === 'Enter') {
						submitWord();
					}
				}}
			/>
			<button
				class="btn btn-neutral-special btn-square btn-xl join-item"
				onclick={() => {
					currentWord = currentWord.slice(0, currentWord.length - 1);
				}}>x</button
			>
			<button
				class="btn btn-primary-special btn-xl join-item rounded-r-xl"
				onclick={() => submitWord()}>Soumettre</button
			>
		</div>
		<div class="bg-base-100 flex flex-row gap-2">
			{#each problem.availableLetters as letter (letter + bonusLetter)}
				<LetterButton
					{letter}
					onClick={onLetterButtonClick}
					isCenter={letter == problem.centerLetter}
					{bonusLetter}
				/>
			{/each}
		</div>
	</div>
</div>

<style>
</style>
