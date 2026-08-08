---
name: /cairn-canvas
id: cairn-canvas
category: Core
stage: reflect
next: /cairn-handoff (after the Spec playback merges the canvas)
description: Assemble the Design Canvas — the reviewable contract holding hill echo, to-be journey, decision records with drivers, MVE, signals, and boundaries
---

Assemble the **Design Canvas**: the single reviewable contract for a feature. It replaces the oracle spec — everything downstream (prototype brief, stories, `/spdd-analysis`) derives from it.

**Loop stage: Reflect → Make.** Design decisions get made and recorded here; implementation still doesn't.

**Input**: a Hill artefact (`@file`). For a tiny change (copy fix, small filter), you may enter directly with a one-row decision — record `entered at: canvas (fast lane)` and keep every section minimal.

## Config

1. Read `.cairn.yaml`. Resolve:
   - `bindings.specs` set → the canvas lives at `{specs}/{epic-or-slug}/canvas.md` (one canvas per feature — **update it in place** on re-runs, never fork a second one).
   - Unbound → `design/canvas/{TICKET}-{YYYYMMDDHHmm}-[Canvas]-{slug}.md`; on update, edit the existing file.
2. Prepend `gating.frontmatter` if declared. Read `bindings.glossary` for domain terms if bound.

## Steps

1. **Echo the hill.** Copy the micro hill (Who/What/Wow + `ladders_to`) verbatim into the canvas header. Every later section must serve it; anything that doesn't belongs in Out of scope.

2. **To-be journey** — as a mermaid diagram, not prose. Pick the fitting form:

   - user journey / flow: `flowchart LR` or `journey`
   - system interaction: `sequenceDiagram`
   - object lifecycle: `stateDiagram-v2`
     Carry the K/A/? labels from the Hill artefact into node/edge annotations where evidence is thin.

3. **Decision records** — the core. One table row per decision, satisfying all five quality criteria (see Norms). Drivers: `business | ux | feasibility | legal`. Include decisions inherited from intake/hill playbacks. A decision that is architectural (crosses systems, hard to reverse) gets flagged `→ promote to ADR` targeting `bindings.decisions`.

4. **Minimum Viable Experience** — the smallest end-to-end slice that lets a real user accomplish the hill's What and produces a learnable signal. Smallest _journey_, not smallest feature set. 3–6 bullets.

5. **Signals** — leading (behaviour, time-to-task) and lagging (retention, outcome) indicators, each tied to the Wow. Unknown baseline → `[TBD — needs data]`.

6. **Boundaries** — explicitly out of scope, and constraints (negative space). Killing scope is as important as picking it.

7. **Open questions** — everything unresolved, each with what would resolve it (sponsor user, data, spike).

## Artefact grammar (Design Canvas, ≤ ~400 lines)

````markdown
# Canvas — {title}

|            |                                          |
| ---------- | ---------------------------------------- |
| Hill       | **Who** {…} · **What** {…} · **Wow** {…} |
| Ladders to | {macro id}                               |
| Status     | draft                                    |
| Sources    | {hill artefact, playback dates}          |

## To-be journey

```mermaid
{flowchart / sequenceDiagram / stateDiagram-v2}
```

## Decisions

| ID  | Date   | Decision       | Driver | Owner  | Status  | Options considered (pros/cons, brief) | Consequences (+/−) |
| --- | ------ | -------------- | ------ | ------ | ------- | ------------------------------------- | ------------------ |
| D1  | {date} | {one decision} | ux     | {name} | decided | {opt A: +…/−… · opt B: +…/−…}         | {+… / −…}          |

## Minimum Viable Experience

- {slice bullet}

## Signals

| Signal | Type    | Target | Baseline     |
| ------ | ------- | ------ | ------------ |
| {…}    | leading | {…}    | {… or [TBD]} |

## Out of scope

- {explicitly not doing, and why in ≤1 clause}

## Open questions

- Q1: {question} — resolved by {sponsor user / data / spike}
````

## Gate (hard stop)

1. Write the artefact to the working tree. **Stop there — no git writes.** Never run `git commit`, `git push`, or open a PR/MR: the gate is a human reading the markdown locally, and publishing is their call, not yours. Keep `status: draft` until a playback agrees it.
2. Print the playback summary: canvas path · decision count by driver · MVE in one line · open questions count · what the playback must settle.
3. **STOP.** Do not emit briefs, stories, or code.
4. Close with: "You are in **Reflect → Make**. Next moves: `/cairn-handoff` once the Spec playback signs off; `/cairn-hill` if the playback breaks the hill."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr` — branch, commit, push, open the PR/MR as a draft.
- `commit` — commit on the current branch, no push.
- `none` — nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Norms — the five record-quality criteria (apply to every decision row)

1. Minimal, cleanly cut scope.
2. One decision per record — split compound decisions.
3. Decision + rationale are the core; considered options get brief, objective pros/cons.
4. Honest consequences, positive and negative, stated factually — no defensive or persuasive language.
5. Timeless and self-sufficient — no links to gitignored files or living docs; quote what you need.

## Guardrails

- Sections are capped by the grammar above — new content **replaces** stale content; never append a second journey or a "misc" section.
- Structure-first: mermaid + tables; prose only for rationale. drawio (embedded XML) only where mermaid genuinely cannot express it.
- Every canvas serves exactly one hill. A second hill = a second canvas.
- Content in the user's working language; keys and labels in English.
