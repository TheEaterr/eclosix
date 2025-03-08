import { getDailyCandidate } from '$lib/pocketBase';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const candidate = await getDailyCandidate();
	return candidate;
};
