# cicd-linter
Linting CLI tool for Github Actions. Pass your YAML file into the linter and you will get a score and report.

## Usage
#### Setup
```bash
curl -L -o gh-actions-linter https://github.com/ktierney15/gh-actions-linter/releases/download/[version]/gh-actions-linter-[os]
chmod +x gh-actions-linter
./gh-actions-linter

# If you want to install globally
sudo mv gh-actions-linter /usr/local/bin/gh-actions-linter

```
#### Start using
```bash
gh-actions-linter [file name]
```