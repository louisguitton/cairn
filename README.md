# cairn

> Prompt-driven **Product Design Thinking** on the IBM EDT Loop — five gated stages, lean markdown artefacts, for teams that prototype with AI.

Cairns are the stacked stones that mark a path up a hill. This kit leaves one artefact per stage of the [IBM Enterprise Design Thinking](https://www.ibm.com/design/thinking/) Loop, so anyone can follow how a product decision was reached — and no decision gets lost in a prototype again.

## The problem

Product specification with AI-assisted prototyping tends to fail the same way everywhere:

- Specs arrive as **oracle documents** (Confluence pages nobody can interrogate).
- The **prototype becomes the spec** — journeys get reverse-engineered from clicks, the user pain point gets lost, and decisions (business? UX? feasibility?) dilute into code.
- There are **no gates** — no defined points where PO, design, and engineering pause, review one artefact, and reconvene.

## The method

Five agent commands mapped onto the EDT Loop (Observe → Reflect → Make). Each writes **one markdown artefact**, prints a playback summary, and **stops** — the pause is the point. Artefacts land in your working tree; nothing is committed, pushed, or opened as a PR unless you ask.

| Command          | Loop stage     | Artefact                                                           | Gate                         |
| ---------------- | -------------- | ------------------------------------------------------------------ | ---------------------------- |
| `/cairn-intake`  | Observe        | Intake Brief + copy-paste question block                           | stakeholder answers top gap  |
| `/cairn-hill`    | Reflect        | Problem statement, needs, **micro hill**, as-is/to-be, assumptions | Hill playback (Who/What/Wow) |
| `/cairn-canvas`  | Reflect → Make | **Design Canvas**: journey, decision records, MVE, signals         | Spec playback                |
| `/cairn-handoff` | Make           | Prototype brief · story set · SPDD input                           | build kickoff                |
| `/cairn-sync`    | Make → Observe | Decision rows + canvas updates from diffs/transcripts/notes        | iteration playback           |

Beta commands are opt-in — `generate --all` installs the five core stages only:

```bash
cairn list --beta                 # see what's available
cairn generate cairn-reverse      # install one by name
```

`/cairn-reverse` reverse-engineers an existing prototype into a draft canvas, every reconstructed claim labelled Assumed until a playback validates it.

Key mechanics:

- **Hill Cascade** — macro hills (≤3 per release, read-only register) vs micro hills (day/week iterations, mandatory `ladders_to:`). Same Who/What/Wow grammar, no conflicts by construction.
- **Decision Records** — one row per decision: driver (business / ux / feasibility / legal), owner, status, options considered, honest consequences.
- **Repo variables** — `.cairn.yaml` binds a docs repo and a prototype repo, plus artefact paths (`hills`, `personas`, `specs`, `capture`). Works in any setup; unbound types fall back to a `design/` tree.
- **One folder per piece of work, numbered in review order** — `/cairn-intake` proposes a ticket, slug and home folder, you confirm it once, and every later stage writes beside it. The folder listing is the review agenda:

  ```text
  1-intake.md   2-hill.md   3-canvas.md   4-brief.md   5-stories.md
  ```

  The hill is always read before the canvas it serves. A missing number is a finding: no `1-intake.md` means the work started from a prototype rather than a stakeholder request.

- **Playback transcripts** are first-class sync input, so the weekly sponsor-user hour lands in the canvas rather than in someone's notebook.
- **Written for a review meeting.** Artefacts are shaped for four people reading together for twenty minutes, most of them reading English as a second language: one block per decision rather than a wide table, confidence written as known / assumed / unknown rather than letter codes, no cross-references by bare identifier, open questions triaged by criticality, and plain declarative sentences.

## Install

```bash
# macOS
brew tap louisguitton/tools
brew trust louisguitton/tools     # Homebrew blocks casks from untrusted third-party taps
brew install --cask cairn

# any platform
go install github.com/louisguitton/cairn/cmd/cairn@latest
```

The tap ships a Homebrew **cask**, which is macOS-only. On Linux and Windows use `go install`, or grab a binary from [Releases](https://github.com/louisguitton/cairn/releases).

## Quick start

```bash
cd your-project          # most likely your prototype repo
cairn init               # declare repo roles, bind artefact paths → .cairn.yaml
cairn generate --all     # install the stage commands for your AI tool
cairn pathcheck          # review where artefacts will land
```

Then, in Claude Code / Cursor:

```text
/cairn-intake @stakeholder-request.md
# … answer the question block …
/cairn-hill @…/1-intake.md
# … Hill playback …
/cairn-canvas @…/2-hill.md
# … Spec playback (PR merge) …
/cairn-handoff @…/3-canvas.md
# build the prototype from the brief, then:
/cairn-sync @playback-transcript.md
```

Supports **Claude Code**, **Cursor** (slash commands), and **Claude** (`cairn generate --tool claude` emits an uploadable skill bundle).

## Companions

- [OpenSPDD](https://github.com/gszhangwei/open-spdd) — the development-side sibling. `/cairn-handoff` emits a ready `/spdd-analysis` input block; cairn covers design, SPDD covers implementation.
- Story pipelines — handoff stories use the As-a/I-want/so-that + Given/When/Then grammar, pipeable to any story skill or straight into Jira.

## License

MIT — see [LICENSE](LICENSE) and [NOTICE](NOTICE). CLI architecture derived from OpenSPDD (MIT). Method based on concepts from IBM Enterprise Design Thinking; not affiliated with or endorsed by IBM.
