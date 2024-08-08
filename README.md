# GMD | Downloader

## Why this bot? 

I noticed that several bots on telegram add ads every time they send the video, I found it very annoying, I looked for the source code of those bots and I didn't find it so well, so you know....

## Installation

```sh
sudo apt update
sudo apt install -y libssl-dev zlib1g-dev gperf g++
sudo apt install -y software-properties-common
sudo add-apt-repository -y ppa:kitware/release
sudo apt install -y cmake
```

Then [compile](https://github.com/tdlib/telegram-bot-api?tab=readme-ov-file#installation) the local [telegram-bot-api](https://core.telegram.org/bots/api) this is used to avoid restrictions when uploading files to telegram, using the api of the nearest DC is limited to 50mb.

Using the local API we can:
* Download files without a size limit.
* Upload files up to 2000 MB.
* Upload files using their local path and [the file URI scheme](https://en.wikipedia.org/wiki/File_URI_scheme).
* Use an HTTP URL for the webhook.
* Use any local IP address for the webhook.
* Use any port for the webhook.
* Set *max_webhook_connections* up to 100000.
* Receive the absolute local path as a value of the *file_path* field without the need to download the file after a *getFile* request.