export function fromhex(hex: string) {
	return Buffer.from(hex, 'hex').toString();
}

export function tohex(string: string) {
	return Buffer.from(string).toString('hex');
}

export function authorartist(author: string, artist: string): string {
	const names = [];
	if (author !== '') names.push(author);
	if (artist !== '') names.push(artist);

	if (names.length > 1) {
		if (names[0] === names[1]) {
			return names[0];
		} else {
			return names[0] + ', ' + names[1];
		}
	} else if (names.length === 1) {
		return names[0];
	}

	return '';
}
