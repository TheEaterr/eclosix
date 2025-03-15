import PocketBase from 'pocketbase';
import type { WordsResponse, ProblemsResponse } from './generated/pocketBaseTypes';
import { cyrb128, shuffleArray } from './hash';

const API_URL =
	process.env.NODE_ENV === 'production' ? 'https://eclosix.maoune.fr' : 'http://127.0.0.1:8090';
const pb = new PocketBase(API_URL);

type TexpandWord = {
	word: WordsResponse;
};

type TexpandTopMatches = {
	top_matches: WordsResponse[];
};

export const testWord = async (word: string) => {
	const words = await pb.collection('words').getList<WordsResponse>(0, 1, {
		filter: `word = "${word}"`,
		skipTotal: true
	});
	if (words.items.length === 0) {
		return '';
	}
	return words.items[0].raw;
};

export type Problem = {
	id: string;
	availableLetters: string[];
	centerLetter: string;
	maxPoints: number;
};

export const getProblemFromId = async (id: string): Promise<Problem> => {
	const problem = await pb
		.collection('problems')
		.getOne<ProblemsResponse<string[], TexpandWord>>(id, {
			expand: 'word'
		});

	if (!problem.letters) {
		throw new Error('No available letters');
	}

	return {
		id,
		availableLetters: shuffleArray(problem.letters, problem.id),
		centerLetter: problem.center_letter,
		maxPoints: problem.max_points
	};
};

export const getRandomProblem = async (): Promise<Problem> => {
	const problem = await pb
		.collection('problems')
		.getFirstListItem<ProblemsResponse<string[], TexpandWord>>('', {
			sort: '@random',
			expand: 'word'
		});

	if (!problem.letters) {
		throw new Error('No available letters');
	}

	return {
		id: problem.id,
		availableLetters: shuffleArray(problem.letters, problem.id),
		centerLetter: problem.center_letter,
		maxPoints: problem.max_points
	};
};

export const getDailyProblem = async (): Promise<Problem> => {
	const oneDay = 24 * 60 * 60 * 1000;
	const initialDate = new Date(2001, 12, 10);
	const currentDate = new Date();
	initialDate.setHours(0, 0, 0);
	currentDate.setHours(0, 0, 0);
	const diffDays = Math.round(Math.abs((currentDate.getTime() - initialDate.getTime()) / oneDay));
	// Hash the date
	const hash = cyrb128(diffDays.toString());
	const problemsCount = (await pb.collection('problems').getList()).totalItems;
	const randomProblemIndex = Math.floor(hash[0] % problemsCount);
	const problem = (
		await pb
			.collection('problems')
			.getList<ProblemsResponse<string[], TexpandWord>>(randomProblemIndex, 1, {
				expand: 'word',
				skipTotal: true
			})
	).items[0];

	if (!problem.letters) {
		throw new Error('No available letters');
	}
	return {
		id: problem.id,
		availableLetters: shuffleArray(problem.letters, problem.id),
		centerLetter: problem.center_letter,
		maxPoints: problem.max_points
	};
};

export const getTopMatches = async (problem: Problem): Promise<string[]> => {
	const matches = await pb
		.collection('problems')
		.getOne<ProblemsResponse<string[], TexpandTopMatches>>(problem.id, {
			expand: 'top_matches'
		});
	return matches.expand?.top_matches.map((m) => m.raw) ?? [];
};

export default pb;
