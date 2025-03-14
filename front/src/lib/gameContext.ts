export type ChosenWord = {
	word: string;
	raw: string;
	points: number;
};

export type GameContext = {
	points: number;
	chosenWords: ChosenWord[];
	problemId: string;
};
