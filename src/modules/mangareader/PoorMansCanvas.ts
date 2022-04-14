// https://github.com/pandeynmn/nmns-extensions/blob/main/src/MangaReaderTo/PoorMansCanvas.ts

/* eslint-disable @typescript-eslint/ban-ts-comment */
// @ts-ignore
import decode from 'image-decode'
// @ts-ignore
// eslint-disable-next-line @typescript-eslint/no-var-requires
const encode = require('image-encode')

class Canvas {
    width = 0;
    height = 0;
    data: Buffer | undefined;

    setSize(width: number, height: number) {
        // console.log(`setting size to ${width}, ${height}`)
        this.width = width
        this.height = height
        this.data = Buffer.alloc(width * height * 4)
    }

    drawImage(
        image: ImageData,
        sx: number,
        sy: number,
        sw: number,
        sh: number,
        dx: number,
        dy: number
    ) {
        if (this.data == undefined) return

        for (let y = 0; y < sh; y++) {
            image.data.copy(
                this.data,
                ((dy + y) * this.width + dx) << 2,
                ((sy + y) * image.width + sx) << 2,
                ((sy + y) * image.width + sx + sw) << 2
            )
        }
    }

    encode(format: 'png' | 'jpg'): ArrayBuffer {
        try {
            return encode(this.data, format, [this.width, this.height])
        } catch (e) {
            // @ts-ignore
            console.log(e.message)
            // @ts-ignore
            console.log(e.stack)
            return new ArrayBuffer(0)
        }
    }
}

class ImageData {
    constructor(
        public width: number,
        public height: number,
        public data: Buffer
    ) {}

    static decode(imgData: Uint8Array): ImageData {
        const { data, width, height } = decode(imgData)
        return new ImageData(width, height, Buffer.from(data))
    }
}

export function createCanvas(): Canvas {
    return new Canvas()
}

export function createImage(data: Uint8Array): ImageData {
    return ImageData.decode(data)
}