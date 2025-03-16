import type { HandleFetch } from '@sveltejs/kit';

export const handleFetch: HandleFetch = async ({ request, fetch }) => {
	if (request.url.startsWith('https://eclosix.fr/')) {
		// clone the original request, but change the URL
		request = new Request(
			request.url.replace('https://eclosix.fr/', 'http://eclosix_back:8080/'),
			request
		);
	}

	return fetch(request);
};
