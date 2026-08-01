<!-- What are the steps I followed -->

<!-- init UV -->
The package manager i used for this tutorial is 'uv'

```
uv init
uv add "fastapi[standard]"
uv add ruff
```

Added ruff.toml for highlevel ruff configurations

#### Testing
```
    ruff check .
    ruff check --fix

    ruff format --check #This runs a dry run (won't update the files but shows the formatting issues)
    ruff format . #This will update the files and do the formattings
```