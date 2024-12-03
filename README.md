# TODO CLI

A simple TODO CLI app which stores it's data in a sqlite database.

## Usage

```bash
# show the help
todo help

# list all todo items
todo list

# create a new todo item
todo new --title "Title of the new todo item"
# with a description it would look like this
todo new --title "Title of the new todo item" --description "Description of the new todo item"
# short version
todo new -t "Title of the new todo item" -d "Description of the new todo item"

# show the todo item with the id 7
todo show 7

# mark the todo with the id 12 as resolved
todo resolve 12

# delete the todo with the id 2
todo delete 2

# generate autocomplete scripts, e.g. for bash (other options: fish, powershell, zsh)
todo completion bash
```

## Development

Development is done using the Nix package manager.

```bash
# enter dev shell
nix develop

# apply code format
nix fmt

# build the binary
nix build .#

# run the todo cli directly using nix
nix run .#
```

In case you don't have the nix package manager installed, following commands should be helpful:

```bash
# build the binary
go build -o todo .

# run the tests
go test ./...

# generate a test coverage html report
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html
```

## Project structure

- `model` package: Contains the DTOs
- `persistence` package:
  - Contains the `DbHandler` and the `TransactionHandler` interfaces. These interfaces are used for persistence.
  - The package also contains a sqlite persistence implementation. Thanks to the interfaces, new storage/database types can be added easily.
  - Inside the `persistence` directory there's also the `dbinit.sql` file which contains the SQL statements to create the required database tables.
- `output` package: Contains some functions to build the strings used to print the todo items.
- `cmd` package: Contains the CLI logic. The CLI implementation is done using the [cobra](https://github.com/spf13/cobra) lib.
- `testutil` package: Contains some util methods used inside the tests to set up sqlite databases for the tests.

## Contributing

We welcome contributions from the community! Whether it's a bug report, feature suggestion, code improvement, or documentation update, your input is greatly appreciated.

### Bugreports and feature requests

In case you find any bugs or want to suggest a feature, please report it by [opening an issue](https://github.com/rubenhoenle/todo-cli/issues).
Please fill all fields provided by the issue template.

### Contributing to the codebase

1. Fork the repository on GitHub
2. Clone the fork:
   ```bash
   git clone https://github.com/your-username/todo-cli.git
   cd todo-cli
   ```
3. Create a new branch for your changes: `git checkout -b feature-or-bugfix-name`
4. Make your changes: Follow the guidelines of the project (see below) and ensure your changes include relevant tests.
5. Before committing your changes:
   - Run tests: `go test ./...`
   - Apply code format: `nix fmt`
6. Commit your changes with a clear commit message:
   ```bash
   git add .
   git commit -m "Add feature or fix description"
   git push origin feature-or-bugfix-name
   ```
7. Open a Pull Request: Go to the original repository and open a pull request. Don't forget to provide a clear description of your changes and reference any relevant issues.

**Guidelines**

- Ensure your code adheres to the existing style and structure.
- Write meaningful commit messages.
- Include tests for any new functionality or bug fixes.
- Update documentation if your changes affect usage.

## Further TODOs and ideas

- [ ] Add missing unit tests
  - [ ] `persistence` package (only tested via `cmd` package currently)
  - [ ] `cmd` package: tests for `delete` command
  - [ ] improve test coverage across the whole project with the help of the test coverage report
- [ ] GitHub Pull request template
- [ ] GitHub Issue template
- [ ] machine readable output (e.g. JSON)
- [ ] configuration options (e.g. where sqlite database file is located)
- [ ] CI Pipeline
- [ ] enable dependabot if it makes sense
- [ ] Automate release creation
