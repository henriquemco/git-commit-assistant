![cli image](./assets/screenshot.png)
<h2 align="center">Git Commit Assistant</h2>

A CLI assistant that automatically generates commit messages based on the diff,
the type of change, and the patch size. Additionally, the tool relies on a brief
description of the commit's purpose provided by the user. I developed it for personal use, but decided to publish it on GitHub.

### Features

- Analysis of diffs, patches, and change types
- Interpretation of user messages
- Automatic generation of commit messages
- Automatic commit (with user approval)

### Project Structure
```bash
git_assistant/
├── cmd/
    └── main.go # Entry point
├── internal/
    ├── git/     # Git functionalities
    ├── parser/  # Message parser 
    ├── handler/ # LLM request
    ├── ui/      # Interface elements
    ├── auth/    # LLM credentials
    └── model/   # Data model 
├── go.mod
└── README.md
```


### Requirements
Before installing and running the project, make sure you have:

- Go (version 1.20 or higher recommended)
- An OpenRouter account
    👉 [OpenRouter account](https://openrouter.ai/) 
- An OpenRouter API key


### Installation 

 > When you use the tool for the first time, it will ask for your API key and the LLM model you want to use.

#### Release

#### Build source

- **Clone the repository:**

```bash 
git clone https://github.com/your-username/git-commit-assistant.git
cd git-commit-assistant
```

- **Build the project:**

```bash 
go build -o git-commit-assistant ./cmd
 ```

 - **Test the application:**
 > I advise you to define it as an executable binary and place it in your bin folder.

```bash
 ./git-commit-assistant
 ```

