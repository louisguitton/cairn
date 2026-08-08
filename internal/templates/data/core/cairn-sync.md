---
name: /cairn-sync
id: cairn-sync
category: Core
stage: loop
next: /cairn-canvas (if the sync demands re-design), /cairn-handoff (next slice), /cairn-hill (if the hill itself broke)
description: Close the Loop — harvest a git diff/prototype, a playback-meeting transcript, or loose notes into decision records and canvas updates
---

Close the EDT Loop: read what reality produced — built prototype, playback meeting, hallway decision — and sync it **back into the Design Canvas** as decision records and journey corrections. This is the command that keeps the canvas from becoming a stale oracle. It must stay the cheapest command to run.

**Loop stage: Make → Observe.** Reality is the new observation.

**Input** (one or more, all first-class):

- **(a) Diff / prototype**: a git range, branch, or prototype directory in the proto repo.
- **(b) Playback transcript**: notes or auto-transcript of a session with stakeholders, the client, or sponsor users (`@file` or pasted).
- **(c) Loose notes**: Slack thread, hallway decision, grilling-session notes.

If no input given, ask which of the three exists and stop.

## Config

1. Read `.cairn.yaml`, then locate the work: the canvas at `<home>/canvas.md`, taken from the input artefact's `Work` header. Given only raw input (a transcript, a diff), list the candidate work folders under `bindings.specs` and **ask which one this belongs to** — never guess.
2. `proto_repo.path` → where to run diff inspection for input (a).
3. `bindings.capture` → file a dated copy of the raw transcript or notes there, and cite that path in Sources.

## Steps

1. **Extract candidate decisions — conservatively.** The bar, per input:

   - Transcript/notes: a decision row requires an **explicit commitment** in the source ("we'll go with X", "agreed", "der Kunde will X"). "Let's revisit", "probably", "I think" → open question or assumption, NOT a decision.
   - Diff/prototype: implemented behaviour that contradicts or extends the canvas journey is a **candidate** decision ("built as X") — mark `status: proposed` unless a human confirms it was deliberate.
   - Every extracted row carries a **verbatim provenance quote** (or file:line / route for code) in a Sources footnote. No quote → don't write it. A miss beats noise.

2. **Classify each decision row**: driver (`business | ux | feasibility | legal`), owner (who committed it — from the transcript speaker or commit author; unknown → `[TBD]`), date, hill id.

3. **Update the canvas**:

   - Append new rows to the Decisions table (next D-id).
   - Supersede, never delete: a reversed decision gets `status: superseded` and the new row links it.
   - Journey: correct diverging nodes; flip cell labels **A → K** (or **? → K/A**) where the input is evidence; flag prototype-vs-canvas divergence that was NOT decided anywhere as an open question — that's drift, surface it.
   - Answered assumptions move out of the 2×2; new ones move in.
   - Architectural decisions get `→ promote to ADR` flags targeting `bindings.decisions`.

4. **Route raw input**: if `bindings.capture` is bound and the transcript/notes aren't stored yet, save them there first (dated, append-only) and cite that path in Sources.

## Output shape (canvas edit, not a new artefact)

The canvas diff should show only:

```markdown
## Decisions

| D7 | {date} | {decision} | ux | {owner} | decided | {options, brief} | {+/−} |
| D8 | {date} | {built-as finding} | feasibility | [TBD] | proposed | … | … |

## To-be journey ← corrected nodes, A→K flips

## Assumptions ← answered rows out, new rows in

## Open questions ← undecided divergence lands here

Sources: {capture path / transcript quote refs / commit range}
```

## Gate

1. Write the artefact to the working tree. **Stop there — no git writes.** Never run `git commit`, `git push`, or open a PR/MR: the gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: decisions extracted (by driver) · proposed vs decided · A→K flips · drift findings · promotion flags.
3. **STOP.**
4. Close with: "Loop closed: **Make → Observe**. Next moves: `/cairn-handoff` for the next slice; `/cairn-canvas` if drift demands re-design; `/cairn-hill` if the Wow itself was refuted."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr` — branch, commit, push, open the PR/MR as a draft.
- `commit` — commit on the current branch, no push.
- `none` — nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **Conservative extraction is the law**: explicit commitment + verbatim quote, or it's an assumption/question.
- Never edit the macro hill register, `bindings.capture` existing files, or anything outside the resolved canvas + capture paths.
- Never delete decision rows — supersede.
- Garbled auto-transcripts (mixed languages, broken names): quote as-is, mark owner `[TBD]`, don't guess speakers.
- Multiple canvases matched → ask, don't pick.
