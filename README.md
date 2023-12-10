# Otto Cron Chats

Otto is a bot designed to help you stay up-to-date with the latest news by monitoring RSS feeds and sending notifications via Telegram.

This part of the code only contains the Cron job that is fetching new information for each chats.
[Other repositories linked to otto](https://github.com/Vico1993?tab=repositories&q=otto&type=&language=&sort=)

## Table of Contents

-   [Getting Started](#getting-started)
-   [Usage](#usage)
-   [Contributing](#contributing)
-   [License](#license)

## Getting Started

To get started with Otto, clone the repository to your local machine:

```sh
git clone https://github.com/Vico1993/otto-cron-chats.git
cd otto-cron-chats
```

## Prerequisites

Make sure you have the following tools installed on your machine:

-   Go (at least version 1.20)
-   A valid Telegram bot:
    -   look at the great [bot father](https://core.telegram.org/bots)
-   Setup an `.env` file

```sh
# Bot token given by the @BotFather
TELEGRAM_BOT_TOKEN=<TELEGRAM_BOT>
# Optional value, to receive notification when deployed
TELEGRAM_ADMIN_CHAT_ID=<TELEGRAM_ADMIN_CHAT_ID>

# Api url for the Otto api
OTTO_API_URL=<OTTO_API_URL>
```

## Installing

To install Otto, run the following command:

```sh
make ensure_deps
```

## Running Tests

To run tests, use the following command:

```sh
make test
```

## TIPS

To make sure it's easy to build, I use: gow. Once install:

```sh
make watch
```

## Usage

To use Otto, run the following command:

```sh
make build && ./bin/bot
```

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) file for more information.

## License

This project is licensed under the [LICENSE](./LICENSE) file in the root directory of this repository.
