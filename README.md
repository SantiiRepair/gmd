# GMD | Downloader

## Installation

```sh
sudo apt update
sudo apt install -y libssl-dev zlib1g-dev g++
sudo apt install -y software-properties-common
sudo add-apt-repository -y ppa:kitware/release
sudo apt install -y cmake
```

Then [compile](https://github.com/tdlib/telegram-bot-api?tab=readme-ov-file#installation) the local [telegram-bot-api](https://core.telegram.org/bots/api) this is used to avoid restrictions when uploading files to telegram, using the api of the nearest DC is limited to 50mb.
