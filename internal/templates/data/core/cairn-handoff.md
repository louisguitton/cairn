---
name: /cairn-handoff
id: cairn-handoff
category: Core
stage: make
next: build the prototype from the brief, or run /spdd-analysis on the canvas; /cairn-sync when reality reports back
description: Emit the Make-stage handoffs from a signed-off canvas — prototype brief, story set (As-a/I-want + Given/When/Then), and/or an /spdd-analysis input block
---

Convert a signed-off Design Canvas into **handoff artefacts**: a prototype brief for the prototyping agent, a story set for the backlog, and/or an input block for `/spdd-analysis`. Specification happened upstream — this stage makes it executable by others.

**Loop stage: Make.** Still no code: the outputs _commission_ making.

**Input**: a Design Canvas (`@file`) whose Spec playback is done (status no longer draft — if it is still draft, warn and ask to confirm before proceeding). Ask which handoffs are wanted if not stated: `prototype` | `stories` | `spdd` (any combination).

## Config

1. Read `.cairn.yaml`. Resolve:
   - Handoffs live next to the canvas: `{specs}/{epic}/brief.md` and `{specs}/{epic}/stories.md`; fallback `design/handoff/{TICKET}-{YYYYMMDDHHmm}-[Handoff]-{slug}.md`.
   - `proto_repo.path` → the target the brief addresses (branch + route scheme).
2. Prepend `gating.frontmatter` if declared.

## Steps

1. **Verify the chain.** The canvas must carry a hill echo and ≥1 decided decision. Missing → stop and route back to `/cairn-canvas`.

2. **Prototype brief** (if wanted) — the experiment commission:

   - **Hill id + Wow** (verbatim echo — the prototype exists to test this).
   - **Question this experiment answers** — one sentence. A prototype without a question is scope creep.
   - **Journey slice** — which part of the to-be journey to build; copy the relevant mermaid subgraph.
   - **Addressing** — target repo (`proto_repo`), branch name to use, route(s) to mount. Branch + route is how every citation back to the prototype will be made.
   - **Fidelity + constraints** — mock data OK? design system? what NOT to build (from Out of scope).
   - **Report-back** — tell the builder: decisions made while building are recorded via `/cairn-sync`, not lost in the code.

3. **Story set** (if wanted) — one story per journey slice the team will build:

   - Grammar (non-negotiable): **As a** {persona} / **I want to** {action} / **so that** {goal}; acceptance criteria as **Given / When / Then** scenarios (happy path + the edge cases that matter).
   - Reference the prototype **by branch + route**, never by screenshot.
   - Each story links its hill id and the canvas decisions it implements (D-ids).
   - If the host has a story skill/template (e.g. `/story`), emit in that shape for direct piping.

4. **SPDD input block** (if wanted) — a self-contained requirement block for `/spdd-analysis`: hill echo, the journey slice, the decided decisions that constrain implementation, boundaries, and the signals to instrument. Quote — don't link — anything SPDD will need.

## Artefact grammar (Prototype brief, ≤ ~80 lines)

````markdown
# Brief — {title}

|          |                                                     |
| -------- | --------------------------------------------------- |
| Hill     | {id}: **Wow** {…}                                   |
| Question | {what this experiment answers}                      |
| Target   | {proto repo} · branch `{name}` · route(s) `{/path}` |
| Fidelity | {mock data / design system / …}                     |

## Journey slice

```mermaid
{subgraph of the canvas journey}
```

## Build

- {concrete slice bullet}

## Do not build

- {from Out of scope}

## Report back

- Record decisions taken while building via /cairn-sync (transcript, notes, or diff).
````

## Gate

1. Write the artefact to the working tree. **Stop there — no git writes.** Never run `git commit`, `git push`, or open a PR/MR: the gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: artefacts written · hill id · question under test · story count.
3. **STOP.** The gate is **build kickoff** — a human takes the brief/stories and starts.
4. Close with: "You are in **Make**. Next moves: build the prototype from the brief; pipe stories to the backlog (e.g. `/story`); run `/spdd-analysis` with the SPDD block; `/cairn-sync` as soon as reality produces decisions."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr` — branch, commit, push, open the PR/MR as a draft.
- `commit` — commit on the current branch, no push.
- `none` — nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **No code generation.** Briefs, stories, analysis input — nothing else.
- Every handoff artefact echoes the hill id; an artefact that can't name its hill doesn't ship.
- Stories without Given/When/Then scenarios are not stories — refuse to emit them bare.
- The SPDD block must be self-sufficient (quote, don't link) — it will be read in another repo.
