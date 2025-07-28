## Regenerating Ignored Files

Some files are ignored by Git (see `.gitignore`) but are required for development or running the project.  
Follow these steps to regenerate them:

### 1. Generate Swagger Documentation

If you use [swaggo/swag](https://github.com/swaggo/swag):

```sh
swag init -g main.go -o docs
```
This will generate `docs/docs.go`, `swagger.json`, and `swagger.yaml` in the `docs/` folder.

### 2. Build Go Binaries

To generate build artifacts (e.g., in `bin/`):

```sh
go build -o bin/sarc
```

### 3. Install Go Dependencies

If you deleted `go.sum`, run:

```sh
go mod tidy
```

### 4. Create a `.env` File

Copy the example and edit as needed:

```sh
cp .env.example .env
```

### 5. Other Generated Files

- Test files: Re-run your tests to regenerate coverage or test artifacts.
- Config files: Copy or recreate as needed.

---

**Note:**  
These files are ignored in version control, so each developer must