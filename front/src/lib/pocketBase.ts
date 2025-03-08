import PocketBase from 'pocketbase';
import type {
	WordsResponse,
	CandidatesResponse,
} from './generated/pocketBaseTypes';
import { cyrb128 } from './hash';

const API_URL =
	process.env.NODE_ENV === 'production' ? 'https://eclosix.maoune.fr' : 'http://127.0.0.1:8090';
const pb = new PocketBase(API_URL)

type TexpandWord = {
	word: WordsResponse;
};

export const testWord = async (word: string) => {
	const words = await pb.collection('words').getList<WordsResponse>(0, 1, {
		filter: `word = "${word}"`,
		skipTotal: true,
	});
	return words.items.length > 0;
};

export const getDailyCandidate = async () => {
	const oneDay = 24 * 60 * 60 * 1000;
	const initialDate = new Date(2001, 12, 10);
	const currentDate = new Date();
	initialDate.setHours(0, 0, 0);
	currentDate.setHours(0, 0, 0);
	const diffDays = Math.round(Math.abs((currentDate.getTime() - initialDate.getTime()) / oneDay));
	// Hash the date
	const hash = cyrb128(diffDays.toString());
	const candidatesCount = (await pb.collection('candidates').getList()).totalItems;
	const randomCandidateIndex = Math.floor(hash[0] % candidatesCount);
	const candidate = (
		await pb.collection('candidates').getList<CandidatesResponse<string[], TexpandWord>>(randomCandidateIndex, 1, {
			expand: 'word',
			skipTotal: true,
		})
	).items[0];

	if (!candidate.available_letters) {
		throw new Error('No available letters');
	}

	const letter = candidate.available_letters[Math.floor(hash[1] % candidate.available_letters.length)];

	const word = candidate.expand?.word.word;
	if (!word) {
		throw new Error('No word found');
	}

	return {
		word,
		letter
	};
};


export default pb;
