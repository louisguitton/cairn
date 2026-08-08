# SPDD Analysis: Prompt-Driven Product Design Thinking Kit (EDT Loop)

## Original Business Requirement

I would like to design and build a structured, prompt-driven process for Product Design Thinking, similar to referneces/n-ai1st-kit and to references/open-spdd, which would focus on Design Thinking instead of Development, using elements from IBM EDT like in references/ibm-enterprise-design-thinking or @docs/synthesis_report_design_thinking.md

the main pain point right now is that Product Owners and Engineers doing product specification are not structured, are too slow, and don't have a clear process with defined intermediary steps or gates where we can pause and reconvene, nor with defined markdown artefacts that we can review or pass each other.

product specs come in the form of markdown docs or confluence pages a bit like an oracle. otherwise product people, designers and engineers prototype high-fidelity prototypes leveraging AI-assisted coding like in ~/workspace/procureai/bmds/version0, and the prototype becomes the spec which is very insatisfactiry for me because we don't know hwat to click on, we have to reverse engineer the user journeys, we lose track of the user pain point (or product hill) that we're fixing, and we dilute all the decisions that were taken, whether they were business driven, UX preferences driven, or technical feasibility driven.

admittedly the 15+ steps process of n-ai1st-kit is too much. the ~5 steps process of open-spdd is much more aligned with what I have in mind. but somehow I want these steps to represents steps in the IBM EDT loop, so that the agent can guide the human as to what could be the next steps, etc...

in terms of delivery model, open-spdd with brew and a couple of init commands is perfect.

### Clarifications (round 2, 2026-08-08)

codify the existing practice yes but also make it generic to other work places and work setups where version0 and procureai-docs are not available. if a prototype repo and a "second brain" or "documentation" repos are prerequisites that's also fine, but they need to be declared as variables so that we can use the toolkit in any setup

also the sync command needs to work with transcript of playback meetings with stakeholders or the client or sponsor users

also we should be compatible with Claude Code, Cursor and Claude

if the hills and personas already exist again, we should symlink or link them a bit like n-ai1st-kit does with projects, but keep in mind we want to use this more like open-spdd: install it in an existing project (e.g. version0 most likely)

also we need this to work at different levels of complexity. A project spanning from August 2026 to July 2027 might have 3 Hills — those are the 3 "macro hills". but I might be using the pdd toolkit for the iteration in the next day or the next week. so the "hills" in the context of the pdd toolkit need to be usable or compatible with the smaller level of complexity for the shorter timeframe, without overriding or conflicting with the macro hills

also diagramming with mermaid diagrams, UML sequence diagrams, potentially drawio should also be permitted and perhaps encouraged. Markdown tables as well to structure information. overall the markdown artefacts should not be heavy on prose. quality criteria: (1) minimal, cleanly cut scope; (2) one decision per record; (3) the decision and its rationale are the core — considered options documented with brief, objective pros/cons to capture the trade-off space without rambling; (4) honest consequences (both positive and negative), factual drivers and trade-offs stated directly without defensive or persuasive language; (5) timeless and self-sufficient, no links to gitignored files or living docs

definitely use English structure and content in user's language: the fact that our project for BMDS is in German is just specific to procureai-docs and version0 which would be the beta tester of this kit, but we need this to be generic

this process feeds Jira: once we have done product specifications and prototyping, we can more easily write jira stories (see /story skill)

product-prompt is the git repo itself, akin to open-spdd on github

portfolio-level prioritisation is not needed here and we can assume a request exists. it can be a follow up command like open-spdd has beta commands on top of the core ones.

POs, Designers and fullstack Engineers run the commands interchangeably. they collaborate on the prototype and second-brain repos, and are Claude Code or Cursor users.

the librarian is specific to procureai-docs and can be ignored. any overlap will be smoothed out later.

yes the reverse path would also be helpful for requests arriving as existing prototypes

sponsor user is always reachable; that's a fine assumption and in the case of BMDS that's the case — we have at least 1h per week + async rounds with the sponsor user

yes we need a way to sync decisions taken outside the process with a sync command

adoption cannot be a requirement: we don't want to force this onto everyone. but if truly useful and helpful, the people using the kit will be faster and have better input to colleagues, thus driving adoption organically.

the method duplication with the existing EDT skill is not a problem — the existing EDT skill can be viewed as my first shot at this problem. I now have gathered more experience in the last 2 months, and if this kit is successful, I can even retire the IBM design thinking skill

the template bloat in n-ai1st-kit is real, and we prefer the leaner templates from open-spdd (see ~/workspace/procureai/bmds/agora/spdd for reference)

---

## Domain Concept Identification

### Existing Concepts (from codebase)

**In this repository (`cairn`, then named `product-prompt`) — currently near-greenfield**

- **SPDD command templates (copied in)**: `.claude/commands/spdd-{analysis,reasons-canvas,generate,prompt-update,sync}.md` — the five OpenSPDD development commands already installed here. Evidence that this repo is the intended home of the new kit, and that the target AI tool is Claude Code (`.claude/commands/`).
- **EDT synthesis report** (external research, kept outside the repo under `references/`) — the doctrinal source: Principles (user outcomes / restless reinvention / diverse empowered teams), the Loop (Observe → Reflect → Make), the Keys (Hills, Playbacks, Sponsor Users), plus adoption obstacles (silos, buyer-vs-user confusion, solution-first trap) and ROI framing. This is the vocabulary the new process must speak.
- **No build system, no CLI, no `spdd/` artefacts, no commits yet** — the repo is a blank slate with three reference symlinks and one doctrine document.

**In `references/open-spdd` (the delivery model to copy)**

- **CLI binary (`openspdd`)**: Go + cobra + `huh`/lipgloss TUI. Commands: `init`, `generate`, `list`, `pathcheck`, `version`, `uninstall`. MIT licensed (Copyright 2026 gszhangwei) — a port or fork is legally straightforward with attribution.
- **Tool detector**: `internal/detector` maps an `AIToolType` (Cursor, Claude Code, Antigravity, Copilot, OpenCode, Codex) to a config dir (`.claude/commands`, `.cursor/commands`, …) and to signature files used for auto-detection.
- **Embedded template store**: `internal/templates/data/{core,optional,tools}` embedded via `go:embed`, so the binary ships the whole method. Templates are markdown files with YAML frontmatter (`name`, `id`, `category`, `description`).
- **Generation strategy registry**: per-tool strategies (flat markdown, Copilot instruction-merge, Codex skill-bundle) behind one interface — the extension point that lets one method target many agents.
- **Distribution**: goreleaser → Homebrew tap (`gszhangwei/homebrew-tools`) + `go install` + install script + a real `uninstall` that classifies the install method. This is precisely the "brew and a couple of init commands" model the requirement asks for.
- **REASONS Canvas**: the 7-dimension _design contract_ for development (Requirements, Entities, Approach, Structure, Operations, Norms, Safeguards), grouped R+E+A = decisions, S+O = path, N+S = guardrails.
- **Workflow and artefact convention**: `/spdd-analysis` → `/spdd-reasons-canvas` → `/spdd-generate` → `/spdd-sync`, writing to `spdd/analysis/` and `spdd/prompt/` with filenames `{JIRA}-{YYYYMMDDHHmm}-[Stage]-{kebab-slug}.md`. Optional commands exist for `spdd-reverse` (reverse-engineer code into a canvas), `spdd-story`, `spdd-api-test`, `spdd-code-review`.
- **Design philosophy** (`docs/design-philosophy.md`): the control-vs-capability argument — code records "what is", not "what should be"; reverse-engineering intent from a projection is lossy. This is the _same_ argument the requirement makes about prototypes, one layer up.

**In `references/n-ai1st-kit` (the anti-pattern for size, the pattern for depth)**

- **~30 numbered commands** across Phase 0–4 plus a legacy-modernization branch (`n-001`, `n-102-write-story`, `n-103-capture-ui`, `n-104-specify`, `n-105-clarify`, `n-106-plan`, `n-107-tasks`, `n-108-checklist`, `n-109-plan-confidence`, `n-110/111` Jira, `n-2xx` implement/test, `n-3xx` E2E, `n-4xx` verify, `n-6xx` legacy, `n-900` knowledge sync).
- **`.ai/2_templates/`**: ~24 document templates including `brd-template`, `srs-template`, `story-template`, `use-cases-overview-template`, `design-system-template`, `design-overview-template`, `adr-template`, `verify-ui-template`. Confirms that upstream (design/product) templates already exist in the organisation — but embedded in a 15+ step process.
- **`.ai_project_memory/` + `constitution.md`**: durable per-project context and rules, with a per-developer "active project" selection mechanism for the multi-repo BMDS workspace. Relevant precedent: the BMDS reality is multi-repo, and context selection is a real problem.

**In `references/ibm-enterprise-design-thinking` (the method content, already written)**

- **A Claude skill** (`SKILL.md` + 9 reference files, ~1,600 lines) implementing a **five-phase pipeline**: (1) Diagnose the input, (2) Deconstruct the artefact, (3) Interrogate the stakeholder, (4) Reframe into EDT artefacts, (5) Translate into engineering shape.
- **Artefact vocabulary already defined**: Intake Brief (12 sections), Problem Statement, Needs Statements ("[user] needs a way to {task} so that {benefit}"), Hill (Who/What/Wow), As-is / To-be Scenario Maps (Phases / Doing / Thinking / Feeling), Assumptions & Questions 2×2 grid, Hypothesis, Minimum Viable Experience, leading/lagging success metrics, architecture implications, explicit out-of-scope.
- **Two enforcement devices worth keeping verbatim**: the "people aren't machines" recast rule (machine language ⇒ it's a feature, not a need), and **K / A / ?** evidence labels (known / assumed / unknown) on every scenario cell, plus mandatory `[TBD — needs sponsor user / interview / data]` instead of fabrication.
- Note the framing: this skill's protagonist is the **engineer** translating a stakeholder download. The requirement's protagonist set is broader (POs _and_ engineers writing specs).

**In `~/workspace/procureai/bmds/version0` (the pain, in the flesh)**

- A React/Vite frontend-only prototype of the BMDS Marketplace (two portals, KERN design system, German UI, mock data beside each prototype).
- **`sitemap.md`**: a manually maintained route inventory where single table cells run to full paragraphs describing behaviour, components, wizard paths and toasts. This _is_ the reverse-engineered spec the requirement complains about — it documents "what to click", after the fact, at enormous cost, with no link to any user outcome.
- **`specs/*.md`** (15 files, e.g. `direktkauf-user-journey.md`, `webshop-storefront.md`, `roles_rights.md`): closer to real specs — they carry a Goal, a Scope boundary, a channel ranking, entry points, and even dated decision revisions ("Revised 2026-08-03 (see webshop-storefront.md, W8)", "Source: two grilling sessions 2026-07-31, 23 decisions"). Proof that the practice _wants_ decision provenance but has no structure for it: 23 decisions are referenced, none are individually recorded with driver, owner, or status.
- `CLAUDE.md`, `design-system.md`, `checklist.md`, `REFACTOR-NOTES.md` — the prototype repo is already agent-driven, so the new kit must be able to _attach_ to such a repo, not only to a fresh one.

**In `~/workspace/procureai/bmds/agora/spdd` (SPDD in production — the leanness benchmark)**

- Five REASONS canvases (165–426 lines each) and five analyses from real shipped work (`APP-5988` PII erasure, `APP-6142` observability, …). Density profile: bullet-dense Requirements with resolved-decisions and boundaries inline, mermaid `classDiagram` for Entities, tables and short blocks elsewhere — prose only where rationale demands it. This is the artefact quality bar the design-thinking kit's templates must match, and proof the team already reads/writes this density happily.

**In `~/workspace/procureai/bmds/procureai-docs` (the team second brain — target codebase #2, and the biggest finding of this analysis)**

The second brain already implements a substantial fraction of the target method **by hand**. The kit is therefore not introducing EDT to this team — it is codifying and operationalising a practice that exists but has no commands, no gates, and no repeatable pipeline.

- **Two-speed model** (`decisions/rfc-0001-team-second-brain.md`): **truth** (`product/`, `architecture/`, `guides/`, `decisions/`, `glossary.md` — always current, PR-gated) vs **stream** (`stream/` — dated, append-only, immutable provenance) vs `archive/`. Any artefact the kit writes must land on the correct side of this line.
- **Hills are already a first-class truth doc**: `product/hills.md` — five Hills in canonical Who/What/Wow form, explicit "at most three per release" budget, problem statement, per-persona needs statements, owner per Hill, epic codes, and references into version0 by branch and route ("Reference: version0 repo, `direktkauf` branch, `/procurement/requests/register`"). The Hill Register the first analysis proposed as a new concept **exists**.
- **Personas are already a truth doc**: `product/personas.md` — P1–P6 as roles-in-workflow, shared across both delivery streams, `[TBD]` validation labels, and a **designated sponsor user** (P2 = "designated UAT sponsor user"). The Sponsor User Register partially exists.
- **A complete manual EDT pass exists**: `product/specs/procurement-office/design-thinking.md` — problem statements, recast needs statements, one Hill per module, as-is/to-be, K/A/? legend, assumptions-instead-of-answers discipline ("Auto-mode caveat… every place we would have asked is recorded as an assumption"). This file is the single best specimen of what the kit's Reflect-stage output should look like — and proof the team will produce it when someone drives.
- **Spec layout convention**: `product/specs/{direct-purchase,procurement-office}/`, with per-epic folders (`APP-4946-hill1-companyportfolio/`, `APP-4947-hill2-direct-purchase-agent/` each holding `README.md` + `stories.md` + `qa-handoff/`). Artefact naming runs on **Jira epic keys + hill number**, not on timestamps. This collides with the `design/{intake,hill,canvas}/{TICKET}-{TIMESTAMP}` convention proposed in the first pass.
- **Frontmatter contract on every truth doc**: `owner` (a team, never a person), `status: draft|published|committed|deprecated`, `last_reviewed`, `decay`, `ttl_days`, `sources` (up-links into `stream/` or Jira), `supersedes`. Kit-emitted truth artefacts must carry this frontmatter or they break the second brain's decay machinery.
- **Agent rules of the road** (`AGENTS.md`): never auto-merge a semantic change to a truth doc (PR, human merges); cite provenance verbatim or don't write it; `stream/` is read-only; relative links only, no wikilinks; lychee CI fails on broken links; prettier + pre-commit; agents do not push. Every kit command that writes into this repo inherits these constraints.
- **Templates + skills already resident**: `_templates/{doc,adr,rfc,story}.md` (the story template mandates As-a/I-want/so-that + Given/When/Then, and — critically — says "No design mockups on this project — reference the **version0 prototype** instead, by route"), and `.claude/skills/` already contains `ibm-enterprise-design-thinking`, `story`, `board-triage`, and the planned `librarian` maintenance agent.
- **Decisions layer**: `decisions/` holds ADRs + RFCs, one file per decision, append-only. **Product-level lightweight decisions have no home** — hills.md embeds some, version0 spec files reference "23 decisions" from grilling sessions, `stream/` holds the raw notes, but no structure links a decision to its driver (business/UX/feasibility), owner, and Hill. The Decision Record gap from the first pass is confirmed, and `decisions/` is where its escape-hatch (promotion to ADR) already lives.
- **The prototype-as-spec loop is institutionalised, in both directions**: hills.md and the story template cite version0 routes as the design reference; stories are written against prototype behaviour. The kit does not need to _break_ this loop — it needs to make the missing upstream artefacts (journey, decisions, hill linkage) exist so the prototype reference is grounded instead of oracular.

### New Concepts Required

- **Repo Variables (generic two-repo operating model)**: the kit's central new mechanism. The method presumes two _roles_ — a **docs repo** (the team's documentation / second brain: truth artefacts, hills, personas, specs) and a **prototype repo** (where Make happens) — declared as **variables** in a committed config (`.pdd.yaml`) written by `pdd init`, so the toolkit works in any workplace setup. A role can be unbound (solo repo: both roles in one), and the two roles can point at each other cross-repo. procureai-docs + version0 are merely the **beta binding** of these variables, not assumptions baked into the kit. `init` runs open-spdd-style _inside an existing project_ (most likely the prototype repo) and records where the docs repo lives.
- **Hill Cascade (macro / micro hills)**: two linked levels of the same Who/What/Wow grammar. **Macro hills** — the ≤3 release/project-level statements of intent (e.g. a project spanning Aug 2026 → Jul 2027) — live in the docs repo's hill register and are _read-only_ for day-to-day kit runs. **Micro hills** — iteration-scale (next day / next week) statements — are created freely by `/pdd-hill`, must declare `ladders_to: <macro-hill-id>`, and never modify the macro register. Same artefact grammar, different blast radius: the cascade lets the kit serve a one-week iteration without touching (or conflicting with) the year-scale intent. Promotion of a micro hill to macro status is an explicit, human-gated act.
- **Hill/Persona Linking**: when a hill register or persona file already exists in the docs repo, `pdd init` does not copy it — it records the path (or symlink) in `.pdd.yaml`, the way n-ai1st-kit's `active-project.local.md` `@import`s per-project memory. Commands then _read_ macro hills and personas through the link and _write_ only micro-level and per-feature artefacts. One home per artefact type, machine-known location.
- **Playback Transcript as sync input**: `/pdd-sync` accepts a transcript of a playback meeting (with stakeholders, the client, or sponsor users) as a first-class input alongside git diffs and prototype directories — extracting decisions (with driver + owner), answered/new assumptions, and journey corrections from it. This is how decisions taken outside the process, and the weekly sponsor-user hour, get captured without anyone re-running the pipeline.
- **Design Canvas** (working name; the design-thinking counterpart of REASONS Canvas): the single reviewable contract artefact that holds Who, pain, Hill, to-be journey, decisions, boundaries, signals, unknowns. Relates to REASONS Canvas as upstream to downstream: the Design Canvas is the input contract that `/spdd-analysis` consumes. In the second brain it materialises as the per-epic spec folder's spine (the role `design-thinking.md` plays today for the procurement-office stream), not as a parallel `design/` tree.
- **Loop Stage**: an explicit tag (Observe / Reflect / Make) attached to every command and every artefact, so the agent can answer "where are we in the Loop, and what are the legitimate next moves?". This is the mechanism that makes the process _guiding_ rather than merely sequential.
- **Gate (Playback)**: a named, mandatory stop between stages where the artefact is complete enough to be reviewed by humans and the process pauses until someone confirms. Maps 1:1 to the EDT Key "Playbacks". Owns the "pause and reconvene" requirement. In the second brain profile, the gate has a natural physical form already mandated by AGENTS.md: **the artefact lands as a branch/PR, and the human merge _is_ the playback sign-off**.
- **Decision Record (lightweight)**: one line/block per decision with **driver classification — business / UX / technical-feasibility / legal** — plus owner, date, status, and the Hill it serves. Directly answers "we dilute all the decisions that were taken". Confirmed as the biggest structural gap in _both_ target repos: version0 references "23 decisions" from grilling sessions without recording them; procureai-docs has append-only ADRs/RFCs for architecture but no home for product-level decisions. Lives inside the canvas, with promotion to `decisions/` when architectural.
- **Journey / Scenario artefact as a first-class file**: as-is and to-be scenario maps with K/A/? labels, produced _before_ any prototype, so "what to click" is derived from a journey instead of reverse-engineered from a prototype. The K/A/? discipline is already practised in `specs/procurement-office/design-thinking.md` — the concept is new only as a _mandatory, per-feature_ artefact.
- **Prototype Brief / Prototype Reverse-Sync**: the two directions of the Make edge — (a) emit a brief that tells a prototyping agent (working in version0) what to build, which Hill it serves and what question it answers, (b) read an existing prototype back into a canvas, recording decisions that were only ever encoded in code. Given hills.md and the story template already cite version0 by branch + route, the brief must speak that same addressing scheme (branch, route) so the citation chain closes.
- **Design Kit CLI**: a new single binary (Go) that ships the method as embedded command templates and installs them into `.claude/commands/` (and the other tools' equivalents) via `init` + `generate`, distributed by Homebrew tap.
- **Handoff contract to SPDD / stories**: an explicit, named section of the Design Canvas that is valid input to `/spdd-analysis` — and, in the second brain profile, valid input to the resident `story` skill's template (As-a/I-want/so-that + Given/When/Then), so the kit feeds the team's existing story pipeline rather than inventing a second one.

**No longer new (exist in the beta docs repo, kit links instead of introduces):**

- **Hill Register** → `product/hills.md` in procureai-docs already is one (Who/What/Wow, ≤3 per release, owners, epic codes). Generic form: `.pdd.yaml` declares `hills:` as a path variable; when bound, `/pdd-hill` reads macro hills from it and writes only micro hills laddering to them; when unbound, the kit scaffolds a fresh register.
- **Sponsor User Register** → `product/personas.md` already designates sponsor users (P2 as UAT sponsor) and carries `[TBD]` validation labels. Same treatment: a `personas:` path variable, linked not copied. Sponsor-user reachability is now an explicit **method assumption** (≥1h/week + async in the beta) — the kit plans validation against them rather than treating their absence as the norm.

### Key Business Rules

- **Every artefact ladders to exactly one Hill; a release carries at most three macro Hills.** The ≤3 budget applies to macro hills only; micro hills are unbounded but each must declare `ladders_to:` a macro hill, and micro-level work never edits the macro register. Governs Hill Cascade, Design Canvas, Journey, Decision Record. (EDT: "Projects usually focus on no more than 3 Hills".)
- **No machine language in needs statements.** Any statement expressed as "dashboard / click / export / log in" is a feature and must be recast as a human task, or flagged as a feature with no user behind it. Governs Intake Brief, Needs Statements, Hill.
- **Every unknown is labelled, never fabricated.** K / A / ? on scenario and empathy cells; `[TBD — needs X]` elsewhere. Governs all Observe/Reflect artefacts. This is the rule that stops the "oracle document".
- **Every decision carries a driver classification and an owner.** business | UX | technical-feasibility (extendable: legal/regulatory, which matters for German public procurement). Governs Decision Record.
- **Each stage ends by writing a markdown artefact to disk before the gate.** No stage may complete "in conversation only" — the artefact is the unit of review and handoff. Governs all commands.
- **The prototype is an experiment under a Hill, never the specification.** A prototype may only be built once a Hill and a to-be journey exist; anything learned in the prototype must be synced back as decisions. Governs Prototype Brief and reverse-sync.
- **The process must degrade gracefully.** A small change may enter at a later stage; entering late must be an explicit, recorded choice rather than an undocumented skip. Governs the guidance logic of every command.
- **Scope is defined by both what is in and what is explicitly out.** Governs Design Canvas boundaries section (mirrors REASONS "Safeguards" — negative space).
- **The kit ships the method, not the content.** Templates are embedded in a binary and installed per-repo; the user's artefacts live in their product repo and are theirs. Governs CLI design (open-spdd's `uninstall` explicitly leaves generated templates alone — same rule).
- **Truth artefacts obey the host docs repo's constitution.** In the beta binding (procureai-docs) that means: frontmatter contract (`owner`/`status`/`last_reviewed`/`decay`/`ttl_days`/`sources`/`supersedes`), provenance via `sources:` up-links, relative links, branch/PR for semantic changes, `stream/` read-only. Generically: `.pdd.yaml` can declare a frontmatter template and a gating mode, and the kit honours them. Governs every docs-repo-profile command.
- **One method, one home per artefact type.** Hills, personas, and architectural decisions each have exactly one home, declared in `.pdd.yaml`; the kit updates or links these, and never creates a competing copy. Governs Repo Variables path mapping.
- **Artefacts are structure-first, not prose-first.** Mermaid diagrams (journeys, sequence/UML, state), markdown tables (decision records, assumptions grids, AC coverage), and lists are the default carriers of information; prose is for rationale only. The resident specimen of the target density is `agora/spdd/prompt/*` (165–426-line canvases dominated by mermaid class diagrams and bullet blocks). drawio (embedded XML) permitted where mermaid runs out. Governs every template.
- **Record quality criteria (the five rules).** Every kit artefact — decision records above all — must satisfy: (1) minimal, cleanly cut scope; (2) one decision per record; (3) decision + rationale as the core, considered options with brief objective pros/cons; (4) honest consequences, positive and negative, stated factually without defensive or persuasive language; (5) timeless and self-sufficient — no links to gitignored files or living docs that will rot the record. Governs Decision Record, Design Canvas, and the promotion path to ADRs.

---

## Strategic Approach

### Solution Direction

Build a **standalone, EDT-shaped sibling of OpenSPDD**: a single Go binary that embeds a small set of markdown command templates and installs them into whichever AI coding tool the target repo uses, plus a single contract artefact (the Design Canvas) that sits upstream of the REASONS Canvas.

The shape, end to end:

```
Observe            Reflect                        Make
  │                  │                             │
  ├─ /pdd-intake ──► ├─ /pdd-hill ──► /pdd-canvas ─┤─► /pdd-handoff ──► prototype
  │   Intake Brief   │   Problem/Needs   Design     │    Prototype brief   or
  │   gaps, red      │   Hill (W/W/W)    Canvas     │    + /spdd-analysis  stories
  │   flags, Qs      │   As-is/To-be     decisions  │      input
  │                  │   Assumptions     boundaries │
  │                  │   2×2             signals    │
  └──────────────────┴───────────────────┬──────────┘
              ▲                          │
              │                          ▼
              └──────── /pdd-sync ◄── reality (prototype, shipped code, user feedback)
                        reverse-sync decisions & journeys back into the canvas
```

Five commands, each tagged with its Loop stage, each ending in a written artefact and a **Playback gate**:

| #   | Command        | Loop stage                       | Artefact written                                                                                    | Gate (who reconvenes)                                                   |
| --- | -------------- | -------------------------------- | --------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| 1   | `/pdd-intake`  | Observe                          | Intake Brief (+ copy-paste question block)                                                          | Stakeholder answers the top gap                                         |
| 2   | `/pdd-hill`    | Reflect                          | Problem statement, needs statements, Hill, as-is/to-be scenarios, assumptions 2×2                   | Hill playback — PO + design + eng agree on Who/What/Wow                 |
| 3   | `/pdd-canvas`  | Reflect → Make                   | **Design Canvas** (the contract: journey, decisions w/ drivers, boundaries, MVE, signals, unknowns) | Spec playback — the review artefact that replaces the Confluence oracle |
| 4   | `/pdd-handoff` | Make                             | Prototype brief and/or story set and/or `/spdd-analysis` input                                      | Build kickoff                                                           |
| 5   | `/pdd-sync`    | Make → Observe (closes the Loop) | Updated Design Canvas + new decision records                                                        | Iteration playback                                                      |

Plus `pdd init` / `pdd generate` from the CLI (not agent commands) to install the templates — mirroring `openspdd init` / `openspdd generate --all`.

Mechanically this reuses the OpenSPDD architecture almost unchanged: embedded `data/core` + `data/optional` templates with YAML frontmatter, the tool detector, the per-tool generation strategies, goreleaser + Homebrew tap. What changes is the _content_ of the templates and two new structural elements: **a Loop-stage frontmatter field** (`stage: observe|reflect|make`) plus a `next:` hint (what lets each command close by telling the human where they are and what the legitimate next moves are), and **a repo profile** written by `init`.

**The two-repo operating model (generic, with the beta binding in brackets).** The method presumes two repo _roles_, declared as variables in `.pdd.yaml`; the Loop's stages split naturally across them:

```
DOCS_REPO (truth)  [beta: procureai-docs]       PROTO_REPO (experiment)  [beta: version0]
┌─────────────────────────────────────┐        ┌────────────────────────────────┐
│ Observe   /pdd-intake               │        │ Make    prototype built from   │
│           reads capture, tickets,   │        │         the prototype brief    │
│           stakeholder downloads     │        │         (branch + routes)      │
│ Reflect   /pdd-hill → micro hill,   │  brief │                                │
│           ladders to macro register │ ─────► │ Learn   decisions surface in   │
│           /pdd-canvas → spec folder │        │         playback meetings,     │
│ Make      /pdd-handoff → prototype  │ ◄───── │         code, transcripts      │
│           brief + stories           │  sync  │                                │
│           (→ /story skill → Jira,   │        │ /pdd-sync reads diff / routes  │
│            or /spdd-analysis input) │        │ / playback transcripts, emits  │
│                                     │        │ decision rows                  │
└─────────────────────────────────────┘        └────────────────────────────────┘
```

- **Docs-repo profile**: artefacts land inside the host repo's existing structure, mapped by `.pdd.yaml`, not a kit-imposed tree. Beta binding: `/pdd-hill` writes micro hills laddering to `product/hills.md`, `/pdd-canvas` writes the per-epic spec folder (`product/specs/<stream>/<epic>/`), `/pdd-handoff` writes `stories.md` next to it. Where the host repo PR-gates truth (as procureai-docs does), **the PR merge is the playback gate** — the gate mechanic comes free from rules the team already enforces. Intake briefs, being dated capture rather than truth, land in the capture area (`stream/` in the beta; a `capture:` path variable generically).
- **Prototype-repo profile**: the primary commands are `/pdd-sync` and the consumption side of the prototype brief. The kit does not make the prototype host truth artefacts; it makes the prototype _addressable_ (branch + route — the citation scheme the beta's hills.md already uses) and _harvestable_.
- **`/pdd-sync` input classes** (all first-class): (a) a git diff or prototype directory — harvest implemented behaviour into journey/decision updates; (b) a **playback-meeting transcript** (stakeholder, client, or sponsor-user session) — extract decisions with driver + owner, answered/new assumptions, journey corrections; (c) loose notes from decisions taken outside the process (Slack threads, hallway calls, grilling-session notes). Output is always the same: decision rows + canvas updates, PR-gated in the docs repo.
- **Fallback layout**: a repo with no prior docs structure gets the kit-native tree (`design/{intake,hill,canvas,handoff}/` with `{TICKET}-{TIMESTAMP}-[Stage]-{slug}` naming). A repo with existing structure keeps it — the config knows the paths.
- **Tool targets at launch**: **Claude Code, Cursor, and Claude** (claude.ai / Desktop). Claude Code + Cursor via the OpenSPDD command-directory strategies (`.claude/commands/`, `.cursor/commands/`). Claude has no command directory — delivery there is a **skill/project-instructions bundle** the CLI can emit (`pdd generate --tool claude` producing an uploadable skill folder or project instructions), reusing the Codex skill-bundle strategy precedent.

The IBM EDT skill in `references/ibm-enterprise-design-thinking` is the **method content predecessor** — the author's first pass at this problem, now superseded by two months of practice. The kit's templates mine it (five phases → commands 1–4; the recast rule; K/A/?; the questioning playbook) but need not stay compatible with it; if the kit succeeds, the skill is retired.

**Jira handoff**: the process explicitly feeds Jira. `/pdd-handoff` produces story-shaped output (As-a/I-want/so-that + Given/When/Then) designed to be piped into the resident `/story` skill (or copied into Jira directly) — specification and prototyping happen first, story-writing becomes the cheap last step. The kit itself does not talk to the Jira API in core; that stays a beta-command candidate (open-spdd's optional-template precedent).

### Key Design Decisions

- **Standalone binary vs. fork of `openspdd` vs. skills-only**: A skills-only delivery (just drop `.claude/skills/`) is the cheapest and needs no Go, but loses multi-tool installation, versioning, and the brew upgrade path the requirement explicitly praises. Forking `openspdd` and adding a parallel command family risks entangling two release cadences and confusing `list`/`generate` output. → **Recommend a new binary that ports OpenSPDD's architecture** (MIT licence permits this; keep attribution). Same UX, independent versioning, clean template namespace, and the two tools compose at the artefact level rather than the binary level.
- **Number and granularity of stages**: three stages (one per Loop phase) is elegant but collapses intake and reframing into one oversized step and gives only two gates. Seven+ recreates the n-ai1st-kit problem. → **Recommend five commands mapped onto the three Loop phases**, with the Reflect phase carrying two commands (hill, canvas) because that is where the requirement's pain concentrates — losing the Hill and diluting decisions both happen in Reflect.
- **One canvas vs. two artefacts in Reflect**: merging Hill and Canvas into one file is fewer files, but then the cheapest, most valuable gate (agree on Who/What/Wow before anyone designs a journey) disappears. → **Recommend keeping them separate**: the Hill artefact is short and is the thing you playback to a CPTO; the Canvas is long and is the thing you review as a team.
- **Gate enforcement — hard stop vs. advisory**: A hard stop (command refuses to continue until a human confirms) enforces the discipline but is annoying for a solo PO iterating. → **Recommend a hard stop at the end of each command with an explicit "resume" affordance**: the command always writes the artefact, prints the playback summary, and stops; continuing is a separate command invocation. The pause is the product.
- **Relationship to prototyping**: banning prototypes contradicts EDT's "Make" and the org's actual velocity. → **Recommend positioning the prototype as an experiment under a Hill**, produced from a prototype brief (`/pdd-handoff`) and read back with `/pdd-sync`. version0-style repos become downstream consumers, not the source of truth.
- **Decision provenance format**: full ADRs per product decision are too heavy for 23 decisions from one grilling session; a bare bullet list loses the driver. → **Recommend a compact table row per decision** (id, date, decision, driver ∈ {business, UX, tech-feasibility, legal}, owner, status, Hill) inside the Canvas, with an escape hatch to promote a row to a real ADR when it is architectural.
- **Primary user**: the existing EDT skill assumes an engineer receiving a stakeholder download; the requirement names POs and engineers writing specs. → **Recommend designing for the pair** — the artefacts are the shared language — but making the _tone_ of `/pdd-intake` work in both directions (receiving a download, and structuring your own half-formed idea). Concretely: the intake command must accept a screenshot/Figma, a Confluence dump, a prototype path, or a raw paragraph.
- **Language**: **decided — English structure, content in the user's language.** German content is a beta-binding specific (procureai-docs/version0), not a kit property. EDT terms stay English proper nouns (Hill, Playback); templates are English; a repo's `glossary.md`-style domain-term file is linkable via config.
- **Generic variables vs. hard-coded beta layout**: **decided — declare the setup as variables.** `.pdd.yaml` declares the two repo roles (`docs_repo`, `proto_repo` — possibly the same repo), and per-artefact path bindings (`hills:`, `personas:`, `specs:`, `capture:`, `decisions:`), each optional with a kit-native fallback. Existing hills/personas are **linked, not copied** (the n-ai1st-kit `@import` idea, minus its ceremony), while installation stays open-spdd-style: run `pdd init` inside an existing project — most likely the prototype repo — and point it at the docs repo. The method is portable; the paths are local.
- **Hill granularity across timeframes**: a year-scale project has ≤3 macro hills, but the kit will mostly be run for day/week iterations. Options: (a) one flat hill list — iteration hills crowd out and conflict with the macro three; (b) forbid hills below release scale — the kit loses its daily driver; (c) **a two-level cascade** — micro hills in the same Who/What/Wow grammar, mandatory `ladders_to:` a macro hill, macro register read-only from below. → **Recommend (c)**: the Wow of a micro hill is typically a measurable slice of its macro Wow; conflicts are structurally impossible because micro never writes macro; and `/pdd-hill` gains a cheap validity check ("which macro hill does this ladder to? none → either it's the wrong work or a macro-level conversation is due").
- **Two repos, one feature**: where does the canonical state of a feature live when Observe/Reflect happen in the docs repo and Make happens in the prototype? → **Recommend: the docs repo is always canonical**; the prototype is addressed (branch + route) and harvested (`/pdd-sync`), never authoritative. This matches the citation direction the beta's hills.md and story template already use — the "prototype becomes the spec" pathology is resolved by _direction of reference_, not by prohibition.
- **Artefact density**: prose-heavy templates (n-ai1st-kit trajectory) vs. structure-first. → **Decided — lean, diagram-first**, with `agora/spdd/prompt/*` as the density benchmark (165–426-line canvases: mermaid class/sequence diagrams, tables, bullets, prose only for rationale). The five record-quality criteria (minimal scope, one decision per record, decision+rationale core with brief option trade-offs, honest consequences, timeless/self-sufficient) are embedded in every template as norms, not suggestions.
- **Tool support**: six tools (OpenSPDD parity) vs. the three the team actually uses. → **Decided — Claude Code, Cursor, Claude.** The first two are flat-markdown command strategies already proven in OpenSPDD; Claude needs a new emit-a-skill-bundle strategy (Codex strategy precedent). Other tools can join later behind the same strategy interface.

### Alternatives Considered

- **Extend n-ai1st-kit with design-thinking commands (`n-09x`)**: rejected — inherits the 15+ step surface the requirement explicitly calls too much, couples the design method to one org's dev kit, and the numbered-command namespace already signals "big process".
- **Contribute the templates upstream as OpenSPDD "optional" templates**: rejected as the primary route — design-thinking stages need their own artefact directory, their own Loop-stage frontmatter, and their own gates; shipping them as optional dev-workflow extras would bury them and make the two methods' vocabularies collide. Worth revisiting as a _secondary_ distribution once the method stabilises.
- **Confluence-native process (templates + page blueprints)**: rejected — loses the agent, loses git history and reviewability by diff, and Confluence is precisely where the "oracle document" pathology lives today. Markdown-first with a paste-friendly output is the compromise.
- **A single mega-command that produces the whole spec in one pass**: rejected — it reproduces the oracle document with better formatting. The requirement's core ask is _intermediary steps and gates_, which is exactly what a single pass removes.
- **Adopting the Double Diamond instead of the EDT Loop**: rejected — the requirement explicitly asks for the EDT Loop, and the Loop's non-linearity plus the Keys (Hills, Playbacks, Sponsor Users) is what supplies the gate and traceability mechanics the Double Diamond leaves implicit.

---

## Risk & Gap Analysis

### Requirement Ambiguities

**Resolved by the second-pass exploration of the two target codebases:**

- ~~Whose repo do the artefacts live in?~~ → **procureai-docs is canonical** (truth in `product/`, capture in `stream/`); version0 is addressed and harvested. Confirmed by the citation direction already in use (hills.md → version0 routes).
- ~~What is the unit of work a Hill attaches to?~~ → **Hills map to epics** in existing practice: `product/specs/direct-purchase/APP-4946-hill1-companyportfolio/` names both, and hills.md lists epic codes per Hill. The kit adopts this mapping.
- ~~Language of artefacts~~ → English truth docs, German domain terms via `glossary.md` — settled by existing practice.
- ~~Relationship to the EDT skill~~ → the skill is already installed in procureai-docs' `.claude/skills/`; the kit's templates should invoke/cite it rather than restate it (see Technical Risks).

**Resolved by the round-2 clarifications:**

- ~~Feed Jira?~~ → **Feeds Jira.** `/pdd-handoff` emits story-shaped output for the `/story` skill; direct Jira API integration stays a beta-command candidate.
- ~~Which AI tools at launch?~~ → **Claude Code, Cursor, Claude.**
- ~~Is `product-prompt` the kit repo?~~ → **Yes** — akin to open-spdd on GitHub.
- ~~Portfolio-level prioritisation in scope?~~ → **No.** Assume a request exists; upstream discovery can arrive later as a beta command (open-spdd optional-template precedent).
- ~~Who runs the commands?~~ → **POs, Designers, and fullstack Engineers interchangeably** — all Claude Code or Cursor users collaborating on both repos. Templates must not assume an engineering-only reader.
- ~~Kit vs. librarian boundary~~ → **librarian is procureai-docs-specific; ignore.** Overlap smoothed out later.
- ~~Generic vs. procureai-specific~~ → **generic kit, variables for the setup; procureai-docs + version0 are the beta testers.**

**Still open:**

- **`.pdd.yaml` shape**: which keys are required vs. optional, single-repo vs. cross-repo reference format (relative path? git URL?), and where it sits in a docs repo that separates tooling from knowledge. Needs one design pass before templates hard-code reads of it.
- **Claude (claude.ai) delivery mechanics**: a skill bundle can be generated, but installation is manual (upload/paste) and has no update path comparable to re-running `pdd generate`. Acceptable friction or does Claude get a reduced surface (canvas + sync only)?
- **drawio feasibility**: mermaid is native in GitLab/GitHub rendering; drawio needs the XML-in-markdown convention and an editor plugin. Encourage mermaid, permit drawio, or mermaid-only at launch?
- **Micro-hill lifecycle**: after the iteration ships, does the micro hill get archived, folded into the canvas's decision history, or accumulate in a register? Unbounded accumulation would recreate the noise the ≤3 rule exists to prevent.

### Edge Cases

- **The request arrives as an existing prototype** (the version0 case). Confirmed as wanted: a reverse path (OpenSPDD `spdd-reverse` precedent) that reads implemented routes/behaviour into a canvas + micro hill, flagging every reverse-engineered journey cell as A (assumed) until a playback validates it.
- **Sponsor user temporarily unavailable.** The method now _assumes_ reachability (≥1h/week + async in the beta) — the edge case shrinks to a scheduling gap, handled by K/A/? labels persisting until the next session, not by a separate mode.
- **Tiny changes.** A copy fix or a filter addition does not deserve five stages. Without an explicit lite path, users will bypass the process entirely — the classic failure mode of heavyweight kits. The hill cascade helps (a micro hill + canvas delta can be minutes of work), but an explicit "enter at `/pdd-canvas` with a one-row decision" fast lane should be documented.
- **Mid-process pivot.** A Playback kills the Hill after the canvas exists. Artefact lifecycle (supersede? archive? new timestamp?) is undefined.
- **Multiple Hills in one feature**, or a feature spanning two Hills — the "ladder to exactly one Hill" rule needs an escape hatch.
- **Decisions taken outside the process** (a hallway conversation, a Slack thread, a "grilling session" as in version0). If the only way to record a decision is to re-run a command, they will not be recorded. Note the second brain already has the right primitive for this: drop the raw notes in `stream/` (append-only capture), then `/pdd-sync` extracts decision rows from them with a `sources:` up-link — the two-speed model was designed for exactly this shape.
- **A Hill or spec artefact goes stale.** The frontmatter contract gives every truth doc a `ttl_days`; the (planned) librarian flags overdue docs for re-verification. Kit-emitted artefacts inherit this decay machinery — but only if the kit fills `decay`/`ttl_days` sensibly per artefact type (a canvas decays faster than a persona).
- **Concurrent editors.** Two people run `/pdd-canvas` on the same feature. In the second-brain profile this resolves to a git branch/PR conflict — visible and mergeable. The kit needs an "update existing" path (OpenSPDD's `/spdd-prompt-update` precedent), not a new-timestamped-file-per-run habit.

### Technical Risks

- **Adoption — reframed per round-2: not a requirement, organic by design.** The kit is not forced on anyone; the bet is that kit users are faster and hand colleagues better input, pulling adoption. Design consequence: every artefact must be valuable to a _non-user_ receiving it (a pasteable question block, a reviewable canvas PR, a ready story) — the kit's outputs are its own marketing. What remains a real risk: if outputs aren't visibly better than the status quo, the kit dies quietly, which is the acceptable failure mode.
- **Method duplication with the existing EDT skill — reframed per round-2: not a risk.** The skill is the author's first pass; the kit is the second, informed by two months of practice. Templates mine the skill freely and need not stay compatible; success path explicitly includes retiring the skill. Residual care: don't leave the stale skill installed in the beta repos alongside the kit long-term.
- **Divergence from hand-grown practice.** A docs repo's conventions (frontmatter, spec layout, naming, house style) predate the kit and keep evolving — and _every other workplace's docs repo differs_. Mitigation is now structural: the `.pdd.yaml` variables keep paths/frontmatter declarative and per-repo; the binary knows the method, the config knows the repo. Generic-first design makes the beta binding just one config file.
- **Transcript quality variance.** `/pdd-sync` on playback transcripts must cope with raw meeting notes, auto-transcripts (names garbled, German/English mixed), and half-decisions ("let's revisit"). Mitigation: extract conservatively — a decision row requires an explicit commitment in the source; everything else lands as an assumption or open question with a provenance quote. (Same "a miss beats noise" rule the beta's AGENTS.md already states.)
- **PR-gating friction.** Making every Reflect artefact a PR gives the gate teeth but adds latency for a solo PO drafting. Mitigation: `draft` status in the frontmatter contract already models this — the command can commit `status: draft` artefacts to a branch freely and only the `draft → published` flip demands the playback PR review.
- **Canvas/reality drift** — the same problem OpenSPDD solves with `/spdd-sync`, and the same one that made `version0/sitemap.md` a maintenance burden. Without a cheap sync path and a habit of running it, the canvas becomes another stale oracle. Mitigation: `/pdd-sync` must be the _easiest_ command to run, and should be able to work from a git diff or a prototype directory.
- **Template bloat.** Every gate wants "one more section"; the n-ai1st-kit `.ai/2_templates/` directory (24 templates) shows the trajectory. Confirmed preference for open-spdd leanness with `agora/spdd` as the living benchmark. Mitigation: a hard cap — five core commands, one canvas, structure-first density rules embedded in the templates, and a stated rule that new sections replace rather than add.
- **Per-tool strategy cost.** Three launch targets: Claude Code and Cursor are proven flat-markdown strategies; **Claude (claude.ai) is the new build** — a skill-bundle emitter with a manual install/update path. Mitigation: reuse OpenSPDD's Codex skill-bundle strategy as the pattern; accept reduced update ergonomics on Claude at launch.
- **IBM trademark and attribution.** "Enterprise Design Thinking", "Hills", "Playbacks" are IBM's framework and marks. A publicly distributed kit should attribute rather than imply endorsement, and the OpenSPDD MIT notice must be preserved in any ported code.
- **Homebrew tap operational cost**: goreleaser needs a tap repo and a `HOMEBREW_TAP_GITHUB_TOKEN`; whether this ships under a personal or a procureai org account is an open operational question with security implications (token custody).
- **Go toolchain**: nothing in this repo is Go today, and the surrounding org work is TypeScript/.NET/Python. A Go binary is the right call for single-file distribution, but is a maintenance island. Mitigation: keep the Go surface thin — it only copies embedded markdown — so the method (markdown) stays editable by anyone.

### Acceptance Criteria Coverage

The requirement contains no formal acceptance criteria. The table below assesses the implicit criteria derived verbatim from the requirement's own statements; each is grounded in a sentence of the original text, and none has been invented.

| AC# | Description (derived from the requirement)                                                                          | Addressable? | Gaps/Notes                                                                                                                                                                                                                                                        |
| --- | ------------------------------------------------------------------------------------------------------------------- | ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | A structured, prompt-driven process for Product Design Thinking exists, in the shape of `n-ai1st-kit` / `open-spdd` | Yes          | Five markdown command templates + CLI, ported from the OpenSPDD architecture                                                                                                                                                                                      |
| 2   | Process focuses on Design Thinking, not development                                                                 | Yes          | Artefacts are Intake/Hill/Journey/Canvas; no code is generated; handoff stops at `/spdd-analysis` input                                                                                                                                                           |
| 3   | Uses elements from IBM EDT (per the skill and the synthesis report)                                                 | Yes          | Hills (Who/What/Wow), Playbacks as gates, Sponsor Users register, Loop as stage model, Principles as review checklist                                                                                                                                             |
| 4   | Defined intermediary steps                                                                                          | Yes          | Five stages, each with one input, one artefact, one gate                                                                                                                                                                                                          |
| 5   | Defined gates where the team can pause and reconvene                                                                | Yes          | Every command hard-stops after writing its artefact and prints a playback summary; resuming is a separate invocation                                                                                                                                              |
| 6   | Defined markdown artefacts that can be reviewed or passed between people                                            | Yes          | Artefact types with fixed grammars; paths per `.pdd.yaml` bindings (kit-native `design/…` tree as fallback), reviewable by diff                                                                                                                                   |
| 7   | Faster than today's unstructured specification                                                                      | Partial      | Plausible but unmeasured. No baseline exists for "how long a spec takes today"; needs a pilot measurement, otherwise this AC is untestable                                                                                                                        |
| 8   | Specs stop being oracle-style markdown/Confluence documents                                                         | Partial      | The staged artefacts and mandatory `[TBD]`/K-A-? labelling attack the pathology, but nothing prevents a team from pasting a canvas into Confluence and treating it as an oracle. Cultural, not technical                                                          |
| 9   | Prototype no longer _is_ the spec; user journeys are not reverse-engineered                                         | Yes          | Resolved by direction of reference: the second brain is canonical, version0 is addressed (branch + route) and harvested via `/pdd-sync` — the citation direction hills.md and the story template already use. Reverse path = `/pdd-sync` in the prototype profile |
| 10  | The user pain point / product Hill stays visible throughout                                                         | Yes          | Hill Register + "every artefact ladders to exactly one Hill" rule + Hill echoed in every downstream artefact header                                                                                                                                               |
| 11  | Decisions are not diluted; business / UX / technical-feasibility drivers are preserved                              | Yes          | Decision Record table with a mandatory driver classification, owner, date, status, and Hill link                                                                                                                                                                  |
| 12  | Fewer steps than n-ai1st-kit's 15+; close to open-spdd's ~5                                                         | Yes          | Five agent commands + two CLI commands (`init`, `generate`)                                                                                                                                                                                                       |
| 13  | Steps map to the IBM EDT Loop so the agent can guide the human on next steps                                        | Yes          | `stage:` and `next:` frontmatter on every template; each command closes with "you are in {stage}; legitimate next moves are …"                                                                                                                                    |
| 14  | Delivery model: Homebrew + a couple of init commands, like open-spdd                                                | Yes          | goreleaser → Homebrew tap, `pdd init` (detect tool, create config dir), `pdd generate --all` (install templates); operational gap: which GitHub org owns the tap                                                                                                  |
| 15  | Works against the real environment (AI-assisted prototyping repos such as version0)                                 | Yes          | Prototype profile: `/pdd-sync` harvests routes/diffs/grilling notes; prototype briefs address version0 by branch + route                                                                                                                                          |
| 16  | Works in both target codebases — version0 (prototype) and procureai-docs (second brain)                             | Yes          | Repo Variables mechanism: same binary and commands, per-repo artefact paths via `.pdd.yaml`; docs repo hosts truth, prototype hosts experiments; beta binding obeys the frontmatter contract and AGENTS.md PR-gating                                              |
| 17  | Generic beyond procureai: prototype/docs repos declared as variables, usable in any setup                           | Yes          | `.pdd.yaml` declares repo roles + per-artefact path bindings, all optional with kit-native fallbacks; procureai-docs/version0 are only the beta binding                                                                                                           |
| 18  | `/pdd-sync` works with playback-meeting transcripts (stakeholders, client, sponsor users)                           | Yes          | Transcript is a first-class sync input alongside diffs/prototype dirs/loose notes; conservative extraction — decision rows only on explicit commitment, else assumptions with provenance quotes                                                                   |
| 19  | Compatible with Claude Code, Cursor and Claude                                                                      | Yes          | Claude Code + Cursor via flat-markdown command strategies (OpenSPDD-proven); Claude via a new skill-bundle emit strategy — manual install/update accepted at launch                                                                                               |
| 20  | Existing hills/personas linked, not copied; install open-spdd-style into an existing project                        | Yes          | `pdd init` in the target project records `hills:`/`personas:` path links in `.pdd.yaml` (n-ai1st-kit `@import` idea without the ceremony); commands read macro registers through the link, write only micro/per-feature artefacts                                 |
| 21  | Hills work at two complexity levels — macro (year-scale, ≤3) and micro (day/week iterations) without conflict       | Yes          | Hill Cascade: same Who/What/Wow grammar, micro hills carry mandatory `ladders_to:`, macro register read-only from below; promotion is explicit and human-gated. Open: micro-hill archival lifecycle                                                               |
| 22  | Artefacts are diagram/table-first, low prose, meeting the five record-quality criteria                              | Yes          | Mermaid (journey/sequence/state) + tables as default carriers, prose for rationale only; `agora/spdd/prompt/*` as density benchmark; five criteria embedded as template norms. Open: drawio at launch or later                                                    |
| 23  | Feeds Jira: specs + prototyping first, then cheap story-writing (via `/story` skill)                                | Yes          | `/pdd-handoff` emits story-grammar output consumable by the `/story` skill; direct Jira API stays a beta-command candidate                                                                                                                                        |
