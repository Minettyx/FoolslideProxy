import { modules } from '$lib/server/modules';
import { fromhex } from '$lib/utils';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ params }) => {
	const modid = params.mod;
	const mangaid = fromhex(params.manga);
	const id = fromhex(params.chapter);

	for (const mod of modules) {
		if (mod.id === modid) {
			const images = await mod.chapter(mangaid, id);

			const data: { url: string }[] = images.map((v) => ({ url: v }));

			return new Response('<script>var pages = ' + JSON.stringify(data) + ';</script>', {
				headers: { 'content-type': 'text/html' }
			});
		}
	}

	return new Response('', {
		status: 404
	});
};
