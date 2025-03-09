export type ChosenWord = {
	word: string;
	points: number;
};

export type GameContext = {
	points: number;
	chosenWords: ChosenWord[];
	problemId: string;
};
