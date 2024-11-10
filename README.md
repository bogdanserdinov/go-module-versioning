# Awesome package versioning

### Versioning structure:
- Major version change (e.g., v1.0.0 → v2.0.0): Indicates breaking changes.
- Minor version change (e.g., v1.0.0 → v1.1.0): Adds functionality in a backward-compatible manner.
- Patch version change (e.g., v1.0.0 → v1.0.1): Bug fixes or small changes.

#### To add versions described above to the golang library/package, you can use the following steps:
- For major version changes:
  - Create a new branch for the major version (e.g., v2).
  - Modify the Go module path in go.mod to reflect the new versioning (e.g., github.com/username/repo/v2).
- For minor and patch version changes:
  - Use Git tags to create the version, like so:
```bash
git tag v1.0.1
git push origin v1.0.1
```

To restrict or deprecate a version in the Go module system, you can use the retract directive in the go.mod file.
This will signal that a particular version should not be used.

Example of retracting a version:
```go
module github.com/username/repo

retract v1.0.0 // RESTRICTED due to breaking changes
```

This will retract the version, and users trying to fetch it will see a warning like this:
```bash
go: warning: github.com/username/repo@v1.0.0: retracted by module author: RESTRICTED due to breaking changes
````
