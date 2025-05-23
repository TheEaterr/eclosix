export const GET = async () => {
	return new Response(
		`
		<?xml version="1.0" encoding="UTF-8" ?>
		<urlset
			xmlns="https://www.sitemaps.org/schemas/sitemap/0.9"
			xmlns:xhtml="https://www.w3.org/1999/xhtml"
			xmlns:mobile="https://www.google.com/schemas/sitemap-mobile/1.0"
			xmlns:news="https://www.google.com/schemas/sitemap-news/0.9"
			xmlns:image="https://www.google.com/schemas/sitemap-image/1.1"
			xmlns:video="https://www.google.com/schemas/sitemap-video/1.1"
		>
            <url>
                <loc>https://eclosix.fr</loc>
                <lastmod>2025-03-16</lastmod>
            </url>
            <url>
                <loc>https://eclosix.fr/game/daily</loc>
                <lastmod>2025-03-16</lastmod>
            </url>
            <url>
                <loc>https://eclosix.fr/game/endless</loc>
                <lastmod>2025-03-16</lastmod>
            </url>
            <url>
                <loc>https://eclosix.fr/about</loc>
                <lastmod>2025-03-16</lastmod>
            </url>
		</urlset>`.trim(),
		{
			headers: {
				'Content-Type': 'application/xml'
			}
		}
	);
};
