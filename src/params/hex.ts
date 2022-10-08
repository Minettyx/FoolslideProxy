import type { ParamMatcher } from '@sveltejs/kit';

export const match: ParamMatcher = (param) => {
	return /[0-9a-fA-F]+/.test(param);
};
