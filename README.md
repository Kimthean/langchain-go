# Go Lang Langchain Rag Agent

This project uses [Swago](https://github.com/swaggo/swag) to automatically generate API documentation from annotations in Go code.

## Generating API Documentation

To generate the API documentation, run the following command:

```bash
swag init --parseDependency --parseInternal --generalInfo ./cmd/main.go


This command tells Swag to parse all dependencies and internal packages, and to use ./cmd/main.go as the general information file.

After running this command, Swag will generate a docs directory with the API documentation.

Running the Development Server
To run the development server, use the make watch command:

This command will start the development server and automatically reload it whenever you make changes to the code.
```
