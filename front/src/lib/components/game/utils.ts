import type { Problem } from '$lib/pocketBase';

export const getNumberOfPoints = (word: string, problem: Problem, bonusLetter: string): number => {
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
