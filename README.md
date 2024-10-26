# Bedrock Chatbot CLI
This project is a simple (sample) command-line interface (CLI) tool for a chatbot using Amazon Bedrock.




https://github.com/user-attachments/assets/189f5fa5-21b0-4d5d-9497-ab31cd2be253




## Prerequisites
- Go 1.16 or higher
- AWS SDK for Go
- AWS credentials configured

## Installation

### as a CLI

```
% go install github.com/timtoronto634/bedrock-chat@latest
% AWS_PROFILE=your_profile bedrock-chat
```

### as a repository

1. Clone the repository:

```sh
git clone https://github.com/yourusername/bedrock-chat.git
cd bedrock-chat
```

2. Install dependencies:

```sh
go mod tidy
```

## Usage

Run the application:

```
AWS_PROFILE=your_profile go run .
```

Type `exit` when you want to end the conversation
