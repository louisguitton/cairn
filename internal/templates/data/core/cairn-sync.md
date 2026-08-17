---
name: /cairn-sync
id: cairn-sync
category: Core
stage: loop
next: /cairn-canvas (if the sync demands re-design), /cairn-handoff (next slice), /cairn-hill (if the hill itself broke)
description: Close the Loop by harvesting a git diff or prototype, a playback-meeting transcript, or loose notes into decision records and canvas updates
---

Close the EDT Loop. Read what reality produced, whether that is a built prototype, a playback meeting or a hallway decision, and sync it **back into the Design Canvas** as decision records and journey corrections. This is the command that stops the canvas becoming a stale oracle, so it has to stay the cheapest command to run.

**Loop stage: Make, returning to Observe.** Reality is the new observation.

**Input**, one or more, all first-class:

- **A diff or prototype**: a git range, a branch, or a prototype directory in the proto repo.
- **A playback transcript**: notes or an auto-transcript of a session with stakeholders, the client, or sponsor users (`@file` or pasted).
- **Loose notes**: a Slack thread, a hallway decision, notes from a grilling session.

If no input is given, ask which of the three exists and stop.

## Config

1. Read `.cairn.yaml`, then locate the work: the canvas at `<home>/3-canvas.md`, taken from the input artefact's `Work` header. Given only raw input such as a transcript or a diff, list the candidate work folders under `bindings.specs` and **ask which one this belongs to**. Never guess.
   - No canvas exists for this work yet? This command has nowhere to write. Say so, and route the user to `/cairn-hill` (if the Who and the outcome are still open) or `/cairn-canvas` (if they are settled and only the decisions need recording). Do not create a canvas here: a canvas assembled from meeting notes alone would skip the hill it is supposed to serve.
2. `proto_repo.path` is where you inspect a diff.
3. `bindings.capture` is where you file a dated copy of the raw transcript or notes. Cite that path in Sources.

## Steps

1. **Extract candidate decisions conservatively.** The bar, per input:

   - Transcript or notes: a decision needs an **explicit commitment** in the source, such as "we'll go with X", "agreed", or "der Kunde will X". Phrases like "let's revisit", "probably" and "I think" become an open question or an assumption. They are not decisions.
   - Diff or prototype: implemented behaviour that contradicts or extends the canvas journey is a **candidate** decision. Record it as built, with status `proposed`, unless a human confirms it was deliberate.
   - Every extracted decision carries a **verbatim provenance quote**, or a file and line, or a route, in the Sources line. No quote means you do not write it. A miss beats noise.

2. **Classify each decision**: driver (business, ux, feasibility or legal), owner (the person who committed to it, taken from the transcript speaker or the commit author, `[TBD]` if unknown), the date, and the hill it serves.

3. **Update the canvas**:

   - Add a new decision **block** for each one, in the block format the canvas uses. Never a table row. Give it the next number and a short title.
   - Supersede, never delete. A reversed decision keeps its block with status `superseded`, and the new block says which decision replaced it, by title as well as number.
   - Journey: correct the diverging nodes. Change `(assumed)` to `(known)` where the input is evidence. If the prototype diverges from the canvas and no decision anywhere covers it, that is drift, and it becomes an open question.
   - Assumptions that got answered move out of the grid. New ones move in, keyed by short name.
   - A decision that crosses systems or is hard to reverse gets `Promote to ADR` in its `Who owns it.` line, targeting `bindings.decisions`.

4. **Route the raw input.** If `bindings.capture` is bound and the transcript or notes are not stored yet, save them there first, dated, and cite that path in Sources.

## Output shape (an edit to the canvas, not a new artefact)

The canvas diff should show only this:

```markdown
## Decisions

### 7. {short title}

**Decided.** {what was decided}
**Why.** {the reason it won}
**Rejected.** {the option, and why it lost}
**What it costs.** {the honest downside}
**Who owns it.** {name}. Driver: {business, ux, feasibility or legal}. Status: decided.

### 8. {short title of a built-as finding}

**Decided.** Built as {behaviour}, found in the prototype rather than agreed in a playback.
**Why.** {the reason given in the source, or "no reason recorded"}
**Rejected.** {what the canvas said instead, and why it lost, or "nothing, this contradicts the canvas"}
**What it costs.** {the honest downside}
**Who owns it.** [TBD]. Driver: feasibility. Status: proposed.

## To-be journey

{corrected nodes, and assumed changed to known where there is now evidence}

## Assumptions

{answered ones out, new ones in, keyed by short name}

## Open questions

{undecided divergence lands here, under a criticality heading}

Sources: {capture path, the quotes you relied on, the commit range}
```

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- **One language per sentence, and one content language per artefact.** Name the content language in the artefact header. Never mix languages inside a sentence, in either direction.
  - The sentence patterns in this template are shown in English because the template is English. **Translate the pattern into the content language.** Do not keep the English connectives and drop content into the gaps.
  - Wrong: `Ein Beschaffer needs a way to seinen Bedarf zu begruenden so that die Sparsamkeit belegt ist`
  - Right: `Ein Beschaffer braucht eine Moeglichkeit, seinen Bedarf zu begruenden, damit die Sparsamkeit belegt ist`
  - Furniture stays English: section headings, table column labels, frontmatter keys, and the controlled vocabularies (known, assumed, unknown; proposed, decided, superseded; business, ux, feasibility, legal). These are short, fixed anchors that keep artefacts recognisable and searchable across teams. Give their translation once in the legend if that helps the reader, then use the English term consistently.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the canvas update to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: how many decisions you extracted and under which drivers, how many are proposed rather than decided, which assumptions became known, what drifted, and what you flagged for promotion to an ADR.
3. **STOP.**
4. Close with: "Loop closed: **Make, returning to Observe**. Next moves: `/cairn-handoff` for the next slice. `/cairn-canvas` if drift demands re-design. `/cairn-hill` if the Wow itself was refuted."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **Conservative extraction is the law.** An explicit commitment plus a verbatim quote, or it stays an assumption or a question.
- Never edit the macro hill register, never edit existing files under `bindings.capture`, and never touch anything outside the resolved canvas and capture paths.
- Never delete a decision. Supersede it.
- Garbled auto-transcripts, with mixed languages and broken names, get quoted as they are. Mark the owner `[TBD]`. Do not guess who spoke.
- **Protect the uncomfortable findings.** Drift with no decision behind it, a claim the prototype never measured, a contradiction between two sources: all of it stays visible. This command exists to surface them.
- If several canvases match, ask. Do not pick.
