// https://github.com/pandeynmn/nmns-extensions/blob/main/src/MangaReaderTo/interceptors/ImageInterceptor.ts

import { createImage, createCanvas } from './PoorMansCanvas';
import pkg from 'shuffle-seed';
const { unshuffle } = pkg;

interface Slice {
	x: number;
	y: number;
	width: number;
	height: number;
}

export class ImageUnscrambler {
	async interceptResponse(byteArray: Uint8Array): Promise<Buffer> {
		// console.log(`Request`)

		// console.log('response caught')
		const sliceSize = 200;
		const seed = 'stay';

		const image = createImage(byteArray);

		const canvas = createCanvas();
		canvas.setSize(image.width, image.height);

		const totalParts = Math.ceil(image.width / sliceSize) * Math.ceil(image.height / sliceSize);
		// console.log(`totalParts: ${totalParts}`)
		// console.log(`image.width: ${image.width}`)
		// console.log(`image.height: ${image.height}`)

		const noOfHoriSeg = Math.ceil(image.width / sliceSize);
		const someArray: Slice[][] = [];

		// console.log('Begin someArray loop')
		// for loop to get all slides
		for (let i = 0; i < totalParts; i++) {
			const row = Math.floor(i / noOfHoriSeg);
			const slice = {
				x: (i - row * noOfHoriSeg) * sliceSize,
				y: row * sliceSize,
				width: 0,
				height: 0
			};

			slice.width =
				sliceSize - (slice.x + sliceSize <= image.width ? 0 : slice.x + sliceSize - image.width);

			slice.height =
				sliceSize - (slice.y + sliceSize <= image.height ? 0 : slice.y + sliceSize - image.height);

			if (!someArray[slice.width - slice.height]) {
				someArray[slice.width - slice.height] = [];
			}

			someArray[slice.width - slice.height].push(slice);
		}
		// console.log('finished someArray loop')
		// console.log(JSON.stringify(someArray))
		// console.log(`Some Array Length: ${someArray[0]!.length}`)

		// console.log('Begin drawing loop')
		for (const property in someArray) {
			const baseRangeArray = this.baseRange(0, someArray[property]?.length, 1, false);
			const shuffleInd = unshuffle(baseRangeArray, seed);
			// console.log(JSON.stringify(shuffleInd))
			const groups = this.getGroup(someArray[property]);
			// console.log(JSON.stringify(groups))

			for (const [key, slice] of someArray[property].entries()) {
				const s = shuffleInd[key];

				const row = Math.floor(s / groups.cols);
				const col = s - row * groups.cols;
				const x = col * slice.width;
				const y = row * slice.height;

				canvas.drawImage(
					image,
					groups.x + x,
					groups.y + y,
					slice.width,
					slice.height,
					slice.x,
					slice.y
				);
			}
		}
		// console.log('finished drawing loop')

		const encodedImg = canvas.encode('jpg');

		return Buffer.from(encodedImg);
	}

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	private getColsInGroup(slices: any) {
		if (slices.length == 1) return 1;
		let t = 'init';
		// eslint-disable-next-line no-var
		for (var i = 0; i < slices.length; i++) {
			if (t == 'init') t = slices[i].y;
			if (t != slices[i].y) {
				return i;
			}
		}
		return i;
	}

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	private getGroup(slices: string | any[]) {
		const group = {
			slices: slices.length,
			cols: this.getColsInGroup(slices),
			rows: 0,
			x: slices[0].x,
			y: slices[0].y
		};
		group.rows = slices.length / group.cols;

		return group;
	}

	// https://github.com/lodash/lodash
	private baseRange(start: number, end: number, step: number, fromRight: boolean) {
		let index = -1;
		let length = Math.max(Math.ceil((end - start) / (step || 1)), 0);
		const result = new Array(length);

		while (length--) {
			result[fromRight ? length : ++index] = start;
			start += step;
		}
		return result;
	}
}
