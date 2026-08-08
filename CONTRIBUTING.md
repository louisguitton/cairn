# Contributing

## Setup

```bash
go build ./...
go test ./...
pre-commit install --install-hooks -t pre-commit -t commit-msg
```

## Commits

[Conventional Commits](https://www.conventionalcommits.org/) are enforced by a `commit-msg` hook, and releases are derived from them:

| Prefix                                              | Release |
| --------------------------------------------------- | ------- |
| `fix:`                                              | patch   |
| `feat:`                                             | minor   |
| `BREAKING CHANGE:`                                  | major   |
| `docs:` `test:` `ci:` `chore:` `build:` `refactor:` | none    |

`main` is squash-merge only, and the PR title becomes the commit subject — so **the PR title must be conventional**. Merging to `main` triggers semantic-release, which tags and hands off to goreleaser (GitHub release + Homebrew cask).

## Changing the method

The five stage templates live in [`internal/templates/data/core/`](internal/templates/data/core/) (`cairn-reverse` and future experiments in `data/beta/`). They are embedded in the binary at build time, so editing one and rebuilding is the whole loop.

Every template carries `stage:` and `next:` frontmatter — that is what lets a command tell the user where they are in the EDT Loop and what may legitimately come next.

Two rules keep the kit from turning into the bloated process it exists to replace:

- **Sections are capped.** New content replaces stale content. A template that grows a section should lose one.
- **Structure over prose.** Mermaid diagrams and tables carry the information; prose is for rationale only.

## Artefact quality criteria

Decision records — the artefacts that carry the most weight — must be:

1. Minimal, cleanly cut in scope.
2. One decision per record.
3. Decision and rationale first, with brief objective pros/cons for options considered.
4. Honest about consequences, positive and negative, stated factually.
5. Timeless and self-sufficient — no links to gitignored files or living docs.
