# Foolslide Proxy

A proxy that converts some manga reading websites to a fake Foolslide website made to be parsed by the Tachiyomi extension "FoolSlide Customizable"

This project was made because I was too lazy to learn how to make a Tachiyomi extension

## Advantages and disadvantages

Advantages

- Automatic updates
- Easier to add support for websites

Disadvantages

- You need to rely on an external server
- No filters or other custom Tachiyomi interface
- The origins could rate limit or block the service
- When searching for mangas the covers don't load without using a [custom version](https://github.com/Minettyx/tachiyomi-extensions) of the Tachiyomi extension

## Supported websites

Check them [here](https://github.com/Minettyx/FoolslideProxy/wiki/Available-sources)

## Usage

You can run the proxy on your own or use the publicly available instance at foolslideproxy.minettyx.com

- Install the "FoolSlide Customizable" extension in Tachiyomi or use my [custom version](https://github.com/Minettyx/tachiyomi-extensions) of the extension that adds covers when searching
- Go the the Extensions tab, find the extension and click on "Settings"
- Click on the gear icon, change the URL to your server and make sure no leading slash is present (or use my public instance "https://foolslideproxy.minettyx.com"), click OK and restart Tachiyomi (you may need to force stop it from the applications settings)
