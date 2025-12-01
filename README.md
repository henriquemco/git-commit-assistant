# Git Commit Assistant

A CLI assistant that automatically generates commit messages based on the diff,
the type of change, and the patch size. Additionally, the tool relies on a brief
description of the commit's purpose provided by the user.

### Features

- Analysis of diffs, patches, and change types
- Interpretation of user messages
- Automatic generation of commit messages
- Automatic commit (with user approval)

### Project Structure
```bash
git_assistant/
├── main.go # Entry point
├── internal/
├── go.mod
└── README.md
```
