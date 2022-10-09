import type { Handle } from '@sveltejs/kit';
import status from 'http-status';

export const handle: Handle = async ({ event, resolve }) => {
	const res = await resolve(event);
	console.log(event.url.href)

	switch (status[`${res.status}_CLASS`]) {
		case status.classes.CLIENT_ERROR:
		case status.classes.SERVER_ERROR: {
			const msg = status[res.status].toString();
			return new Response(msg, {
				status: res.status
			});
		}
		default:
			return res;
	}
};