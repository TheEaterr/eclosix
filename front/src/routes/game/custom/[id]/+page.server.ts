import { getProblemFromId } from '$lib/pocketBase';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	return await getProblemFromId(params.id);
};
