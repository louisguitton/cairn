---
name: /cairn-hill
id: cairn-hill
category: Core
stage: reflect
next: /cairn-canvas (after the Hill playback agrees on Who/What/Wow)
description: Reframe an intake into a problem statement, needs statements, a micro hill laddering to a macro hill, as-is/to-be scenario maps, and an assumptions 2×2
---

Reframe what Observe collected into a **Hill artefact**: problem statement, needs statements, one **micro hill** (Who/What/Wow) laddering to a macro hill, paired as-is/to-be scenario maps, and an assumptions grid.

**Loop stage: Reflect.** You are converging on intent, not designing screens or flows.

**Input**: an Intake Brief (`@file`), or a direct description if the request is small enough to skip intake (say so in the artefact: `entered at: hill`).

## Config

1. Read `.cairn.yaml`, then take the **work identity from the input artefact's `Work` header** — the home folder, ticket and slug. Reuse them exactly; never re-derive a slug or invent a second folder for the same work.
   - Input has no `Work` header (entered mid-pipeline) → propose ticket, slug and home the way `/cairn-intake` does, and **ask the user to confirm before writing**.
2. Write to `<home>/hill.md`.
3. `bindings.hills` → the **macro hill register**. Read it; NEVER write it. Unbound → no macro hills exist yet; see step 2 of Steps.
4. Prepend `gating.frontmatter` template if declared.

## Steps

1. **Read the macro hills** from the register. Note each hill's id, Who/What/Wow, and owner.

2. **Pick the ladder.** Which macro hill does this work serve?

   - Exactly one fits → record `ladders_to: {id}`.
   - None fits → **do not invent one and do not skip**. Record `ladders_to: NONE — {one honest sentence why}` in the artefact and flag it in the gate summary: either this is the wrong work, or a macro-level conversation is due. That finding is a valid, gate-worthy output.
   - Register unbound/empty → draft the micro hill anyway, mark `ladders_to: [TBD — no macro register bound]`, and suggest creating one.

3. **Problem statement** — one sentence, user-centred, anchored in the as-is, no solution:

   > _{Specific user} struggles to {do something} because {as-is reason}, which costs them {concrete cost}._

   Cost unknown → `[TBD — needs interview / data]`, never invented.

4. **Needs statements** — one per user task, recast rule enforced (no click/view/export verbs; use prioritize/decide/compare/recover/trust). Cluster if >5 and write one über-statement.

5. **Micro hill** — same grammar as macro, smaller blast radius:

   - **Who**: specific persona (link the personas file if bound).
   - **What**: the new capability in human terms — no technology named.
   - **Wow**: measurable, ideally a slice of the macro Wow (time, count, %, feeling with a probe).
   - **Timeframe**: day | week | sprint.

6. **As-is / To-be scenario maps** — 4 rows (Phases / Doing / Thinking / Feeling), 4–7 phases, **every cell labelled K, A, or ?**. The negative cells in the as-is Feeling row are where the Wow must land; say which.

7. **Assumptions 2×2** — every assumption placed on certainty × risk. The uncertain/high-risk quadrant blocks the build; each entry there gets a validation plan (sponsor-user session, data pull, spike).

## Artefact grammar (Hill artefact)

```markdown
# Hill — {title}

|            |                                              |
| ---------- | -------------------------------------------- |
| Work       | {ticket} · `{home folder}` · slug `{slug}`   |
| Ladders to | {macro id + one-line quote of its What}      |
| Timeframe  | {day/week/sprint}                            |
| Entered at | {intake → hill / hill (skipped intake, why)} |
| Source     | {intake brief link / request quote}          |

## Problem statement

> {one sentence}

## Needs statements

- N1: {user} needs a way to … so that … (K/A/?)

## Micro hill

- **Who:** {persona}
- **What:** {capability, human terms}
- **Wow:** {measurable differentiator}

## As-is

|          | {phase 1} | {phase 2} | …   |
| -------- | --------- | --------- | --- |
| Doing    | (K) …     | (A) …     |     |
| Thinking | (A) …     | (?) …     |     |
| Feeling  | (A) …     | (?) …     |     |

## To-be

{same 4-row shape; mark the cells the Wow changes}

## Assumptions (certainty × risk)

| #   | Assumption | Certainty | Risk | Validation plan (if uncertain+high) |
| --- | ---------- | --------- | ---- | ----------------------------------- |
```

The `Work` row is the identity every later stage inherits — same ticket, same folder, same slug. Never restate it differently.

## Gate (hard stop)

1. Write the artefact to the working tree. **Stop there — no git writes.** Never run `git commit`, `git push`, or open a PR/MR: the gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: artefact path · micro hill (Who/What/Wow in one line) · ladders_to (or the NONE finding) · count of ?-cells and blocking assumptions.
3. **STOP.** The gate is the **Hill playback**: PO + design + engineering agree on Who/What/Wow before anyone designs a journey. Suggest the 3 sentences to open that playback with.
4. Close with: "You are in **Reflect**. Next moves: `/cairn-canvas` after the Hill playback; back to `/cairn-intake` if the playback exposes a missing Who/Why."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr` — branch, commit, push, open the PR/MR as a draft.
- `commit` — commit on the current branch, no push.
- `none` — nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **Never write the macro hill register.** Read-only, always. Promotion of a micro hill to macro is a human act outside this command.
- No solution or screen language in Who/What/Wow.
- Every scenario cell carries K/A/?; unknown cost/metric is `[TBD — needs X]`.
- Structure-first; prose only in the problem statement and rationale lines.
