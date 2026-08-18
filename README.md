# cairn

> Prompt-driven **Product Design Thinking** on the IBM EDT Loop. Five gated stages, lean markdown artefacts, for teams that prototype with AI.

Cairns are the stacked stones that mark a path up a hill. This kit leaves one artefact per stage of the [IBM Enterprise Design Thinking](https://www.ibm.com/design/thinking/) Loop, so anyone can follow how a product decision was reached, and no decision gets lost inside a prototype.

## The problem

Product specification with AI-assisted prototyping tends to fail the same way everywhere:

- Specs arrive as **oracle documents**: wiki pages nobody can interrogate.
- The **prototype becomes the spec**. Journeys get reverse-engineered from clicks, the user pain point gets lost, and decisions dilute into code. Was that a business call, a UX preference, or a technical constraint? Nobody can tell six weeks later.
- There are **no gates**. No defined point where a product owner, a designer and an engineer pause over one artefact and reconvene.

## How it works

Five commands, one per step of the Loop, each writing exactly one markdown file and then stopping.

```text
Observe             Reflect                           Make
  │                   │                                │
  ├─ /cairn-intake ─► ├─ /cairn-hill ─► /cairn-canvas ─►┤─► /cairn-handoff ─► prototype
  │   1-intake.md     │   2-hill.md      3-canvas.md    │    4-brief.md         or
  │   gaps and        │   Who/What/Wow   journey,       │    5-stories.md       stories
  │   questions for   │   as-is/to-be    decisions,     │    (or an
  │   the stakeholder │   assumptions    boundaries     │    /spdd-analysis
  │                   │                  signals        │    input block)
  └───────────────────┴─────────────────────────┬───────┘
              ▲                                 │
              │                                 ▼
              └───────── /cairn-sync ◄────── reality: a diff, a playback
                         decisions and             transcript, loose notes
                         journey corrections,
                         written back into
                         3-canvas.md
```

**The stop is the feature.** Each command writes its artefact, prints a playback summary, and ends. Nothing is committed, pushed or opened as a pull request unless you ask for it. What happens between two commands is a human reading a file, and usually a short conversation. That is the gate.

| Command          | Loop stage     | Writes                                                  | Then you                                           |
| ---------------- | -------------- | ------------------------------------------------------- | -------------------------------------------------- |
| `/cairn-intake`  | Observe        | `1-intake.md`, plus questions ready to paste            | send the questions to the stakeholder              |
| `/cairn-hill`    | Reflect        | `2-hill.md`: Who/What/Wow, as-is and to-be, assumptions | hold the Hill playback: agree the Who/What/Wow     |
| `/cairn-canvas`  | Reflect → Make | `3-canvas.md`: journey, decisions, boundaries, signals  | hold the Spec playback: review it as a team        |
| `/cairn-handoff` | Make           | `4-brief.md` and `5-stories.md`                         | build the prototype, push the stories to a backlog |
| `/cairn-sync`    | Make → Observe | decisions and corrections back into `3-canvas.md`       | hold the iteration playback, then go round again   |

## Where artefacts live

The method assumes two repo roles, which you declare once in `.cairn.yaml`. They can be the same repo.

```text
DOCS_REPO (truth)                              PROTO_REPO (experiment)
┌──────────────────────────────────────┐       ┌────────────────────────────────┐
│ Observe  /cairn-intake               │       │ Make    the prototype is built │
│          reads the request, the      │       │         from 4-brief.md, on    │
│          tickets, the transcripts    │       │         the named branch and   │
│                                      │ brief │         routes                 │
│ Reflect  /cairn-hill                 │ ────► │                                │
│          2-hill.md, laddering to the │       │ Learn   decisions surface in   │
│          macro hill register         │       │         playback meetings, in  │
│          /cairn-canvas → 3-canvas.md │       │         code, in transcripts   │
│                                      │ ◄──── │                                │
│ Make     /cairn-handoff              │ sync  │ /cairn-sync reads the diff,    │
│          4-brief.md + 5-stories.md   │       │ the routes, or a transcript,   │
│          (stories go on to a         │       │ and writes the decisions back  │
│          backlog, or /spdd-analysis) │       │ into the canvas                │
└──────────────────────────────────────┘       └────────────────────────────────┘
```

The docs repo is always canonical. The prototype is addressed by branch and route, and harvested by `/cairn-sync`. That direction of reference is what stops the prototype quietly becoming the spec.

If you have no docs repo, leave the role unbound and everything falls back to a `design/` tree in the repo you are standing in.

## Install

```bash
# macOS
brew tap louisguitton/tools
brew trust louisguitton/tools     # Homebrew blocks casks from untrusted third-party taps
brew install --cask cairn

# any platform
go install github.com/louisguitton/cairn/cmd/cairn@latest
```

The tap ships a Homebrew **cask**, which is macOS-only. On Linux and Windows use `go install`, or take a binary from [Releases](https://github.com/louisguitton/cairn/releases).

### Upgrade

```bash
brew update && brew upgrade --cask cairn && cairn --version
```

On other platforms, re-run `go install github.com/louisguitton/cairn/cmd/cairn@latest`.

**Then refresh the templates in each repo that uses cairn:**

```bash
cd your-project
cairn generate --all --force     # plus any beta commands you installed
```

Upgrading the binary does not touch the stage commands already written into
`.claude/commands/` or `.cursor/commands/`. Those are files in your repo, yours
to edit, so cairn never overwrites them without `--force`. Until you regenerate,
you are running the previous version's method.

## Set up a repo

### Which repo do I run `init` in?

**Either of the two. Pick the prototype if you have no reason to prefer otherwise**, because that is the repo most people already have open. The choice is not a commitment about where artefacts land: that is what the path bindings decide, and artefacts normally land in the docs repo either way.

The choice is sticky in one respect. `init` writes `.cairn.yaml` into the repo you run it in, and cairn finds that file by searching **upward from the working directory**. So every cairn command has to run from that repo, or from a directory inside it. Run one from the other repo and cairn finds no config, silently falls back to a `design/` tree, and writes your artefact in the wrong place.

`cairn pathcheck` is the check for this. If it prints `bound` rows, you are in the right repo. If rows say `fallback` when you expect bindings, your working directory has drifted.

### Run it

Run this in the repo you chose.

```bash
cd your-project
cairn init               # asks the repo's role, where the docs repo is, which paths to use
cairn generate --all     # installs the five stage commands for your AI tool
cairn pathcheck          # prints the table of where each artefact will be written
```

`init` asks before it binds anything. If your docs repo already has a hills file, a personas file and a specs folder, it offers those paths and you confirm them. Existing files are linked, never copied or rewritten. `pathcheck` is the command to run whenever you are unsure where something will land.

### Open both repos side by side

Two repos means switching between them constantly: the canvas is in one, the code is in the other. A VS Code multi-root workspace solves that, and pins the working directory so the config is always found. Save this next to your prototype as `prototype.code-workspace`:

```jsonc
{
  // The repo holding .cairn.yaml is listed FIRST on purpose. cairn resolves
  // .cairn.yaml by searching upward from the working directory, so a terminal
  // rooted in the docs repo would miss it and fall back to a design/ tree.
  // terminal.integrated.cwd below keeps every terminal in the right repo.
  "folders": [
    { "name": "prototype", "path": "." },
    { "name": "team-docs", "path": "../team-docs" },
  ],
  "settings": {
    "terminal.integrated.cwd": "${workspaceFolder:prototype}",
    "search.exclude": { "**/node_modules": true, "**/dist": true },
  },
}
```

Open it with `code prototype.code-workspace`. Both repos appear with their own source control, and `cairn pathcheck` in any integrated terminal reports bound paths rather than fallbacks.

The second folder is a relative path, so it assumes the two repos are siblings on disk. If a colleague clones them elsewhere, that folder shows as missing and they need to adjust it.

Supports **Claude Code**, **Cursor** (slash commands), and **Claude** (`cairn generate --tool claude` emits a skill bundle you upload).

## Your first feature, end to end

Say a stakeholder sends you a screenshot and one sentence.

```text
/cairn-intake @the-screenshot-and-what-they-said.md
```

The command proposes a home for this work and waits for you to confirm it:

```text
Proposed home:  product/specs/checkout/buyers-see-real-prices/
Ticket: APP-1234   Slug: buyers-see-real-prices
Confirm, or give me a different path?
```

Then it writes `1-intake.md` and stops. Inside you get the needs statements the screenshot presupposes, which of Who / What / Why / Wow are missing, and five to seven questions ready to paste into Slack. **Send those questions.** That is the first gate, and the whole point of it is that you have not designed anything yet.

When answers come back:

```text
/cairn-hill @product/specs/checkout/buyers-see-real-prices/1-intake.md
```

You get `2-hill.md`: a problem statement, needs statements, one micro hill in Who/What/Wow form that must ladder to one of your macro hills, an as-is and to-be journey with every cell marked known, assumed or unknown, and an assumptions grid. **Hold a twenty-minute playback and agree the Who/What/Wow.** Cheapest gate in the process, because everything downstream inherits it.

```text
/cairn-canvas @…/2-hill.md
```

You get `3-canvas.md`, the contract: the to-be journey as a diagram, one block per decision with its reason, the option that lost and why, the honest cost, the minimum viable experience, the signals, and the open questions triaged by criticality. **Review this one as a team.** It replaces the oracle document.

```text
/cairn-handoff @…/3-canvas.md
```

You get `4-brief.md`, which commissions the prototype and names the one question the experiment answers, and `5-stories.md` in As a / I want / so that form with Given / When / Then criteria. Build the prototype, move the stories onto a backlog.

Then reality happens: you build, you meet the sponsor user, someone decides something in a corridor.

```text
/cairn-sync @the-meeting-transcript.md
```

Decisions come back into the canvas with their driver and owner, assumed cells become known where there is now evidence, and anything the prototype does that no decision covers is flagged as drift. The Loop is closed. Go round again.

## You do not have to start at the beginning

| Your situation                                                        | Start here                                                                      |
| --------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| A stakeholder sent a request, screenshot or doc                       | `/cairn-intake`                                                                 |
| You already know the user and the problem                             | `/cairn-hill`                                                                   |
| Small change, one decision to record                                  | `/cairn-canvas`, and note `entered at: canvas (fast lane)`                      |
| A prototype exists and there is no spec                               | `/cairn-reverse` (beta, see below)                                              |
| The intake lists gaps and you are unsure what they mean or what to do | `/cairn-gaps` (beta, see below)                                                 |
| The source was long and you want to know what fell out of the canvas  | `/cairn-coverage` (beta, see below)                                             |
| A meeting produced decisions, and this work already has a canvas      | `/cairn-sync`                                                                   |
| A meeting produced decisions, but nothing is written down yet         | `/cairn-hill`, then `/cairn-canvas`. `/cairn-sync` needs a canvas to write into |

Entering late is fine, as long as it is recorded rather than silent. Each command asks for what it is missing.

Beta commands are opt-in, because `generate --all` installs the five core stages only:

```bash
cairn list --beta                 # see what is available
cairn generate cairn-reverse      # install one by name
```

`/cairn-gaps` reads an intake or a hill and explains its gaps in plain language: what is actually missing, why it matters for this specific work, two to four concrete ways to close each one, and what it costs to accept it as a risk instead. It writes nothing, because explaining is the whole job, and it will not guess an answer to make a gap go away.

`/cairn-coverage` checks a canvas back against the source material it came from. Every stage compresses, and nothing else in the pipeline ever looks back at the input, so things fall out silently. It sorts each source item into covered, ruled out, or fell out, and leads with how far the check actually reaches, because a coverage report that implies completeness it does not have is worse than none. Run it at the canvas gate, where adding something back is still a cheap edit. It reports and never resolves: what to do with a finding is yours.

`/cairn-reverse` reconstructs a draft hill and canvas from an existing prototype. Every reconstructed claim is marked assumed until a playback validates it, and it deliberately writes no `1-intake.md`, because the missing number records that no stakeholder intake ever happened.

## Design choices worth knowing

- **Hill Cascade.** Macro hills, at most three per release, live in a register the kit only ever reads. Micro hills are day or week scale, must declare which macro hill they ladder to, and can never edit the register. Same Who/What/Wow grammar at both sizes, so a one-week iteration cannot conflict with a year-scale intent.
- **One folder per piece of work, numbered in review order.** The folder listing is the review agenda, and the hill is always read before the canvas that serves it. A missing number is information.
- **Decision records, one block each.** Decided, why, what was rejected and why it lost, the honest cost, who owns it. A rejected option with no reason recorded is not a decision, it is a fragment.
- **Written for a review meeting.** Four people, one screen, twenty minutes, most of them reading English as a second language. So: one block per decision rather than a wide table, confidence spelled out as known, assumed or unknown rather than letter codes, no cross-references by bare identifier, open questions triaged by criticality, one idea per sentence.
- **One canvas, one outcome.** Intake asks whether the request fits a single canvas: one Who, one Wow, one journey a team can review in twenty minutes. It never estimates size in days, because that guess would be wrong and would be believed. If the answer is no, it names the pieces and you choose. A request that is three outcomes wearing one name is the most expensive thing to discover at the canvas stage.
- **Fog is tracked separately from questions.** An open question is one you can phrase. Fog, in the canvas's `Not yet specified`, is what you can tell is coming but cannot phrase sharply yet, and the test between them is whether you can state the question precisely now. Fog graduates into a question or a decision as work proceeds, or turns out to be out of scope, which is terminal.
- **One language per sentence.** Artefacts are written in one content language, declared in the header. The sentence patterns in the templates are English because the templates are English, and they get translated rather than filled in: `Ein Beschaffer braucht eine Moeglichkeit, … damit …`, never `Ein Beschaffer needs a way to … so that …`. Headings, column labels and the controlled vocabularies stay English, so artefacts stay recognisable and searchable across teams.
- **Nothing is fabricated.** An unknown is written as unknown. An empty scenario row means nobody has observed a user yet, and that is a finding rather than a gap to fill in.

## Companions

- [OpenSPDD](https://github.com/gszhangwei/open-spdd), the development-side sibling. `/cairn-handoff` emits a ready `/spdd-analysis` input block. cairn covers design, SPDD covers implementation.
- Story pipelines. Handoff stories use As a / I want / so that with Given / When / Then criteria, so they pipe into any story tool or straight into a backlog.

## License

MIT. See [LICENSE](LICENSE) and [NOTICE](NOTICE). CLI architecture derived from OpenSPDD (MIT). Method based on concepts from IBM Enterprise Design Thinking, and not affiliated with or endorsed by IBM.
