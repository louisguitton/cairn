---
name: /cairn-canvas
id: cairn-canvas
category: Core
stage: reflect
next: /cairn-handoff (after the Spec playback signs off the canvas)
description: Assemble the Design Canvas, the reviewable contract holding the hill, the to-be journey, one block per decision, the minimum viable experience, signals and boundaries
---

Assemble the **Design Canvas**: the single reviewable contract for a feature. It replaces the oracle spec. Everything downstream (the prototype brief, the stories, `/spdd-analysis`) derives from it.

**Loop stage: Reflect, moving into Make.** Design decisions get made and recorded here. Implementation still does not happen.

**Input**: a Hill artefact (`@file`). For a tiny change such as a copy fix or a small filter, you may enter here directly with a single decision. Record `entered at: canvas (fast lane)` and keep every section minimal.

## Config

1. Read `.cairn.yaml`, then take the **work identity from the input artefact's `Work` header**: the home folder, ticket and slug. Reuse them exactly. Never re-derive a slug or invent a second folder for the same work.
   - Input has no `Work` header (you entered mid-pipeline)? Propose ticket, slug and home the way `/cairn-intake` does, then **ask the user to confirm before writing**.
2. Write to `<home>/canvas.md`. One canvas per work, **updated in place** on re-runs. Never fork a second canvas.
3. Prepend `gating.frontmatter` if declared. Read `bindings.glossary` for domain terms if it is bound.

## Steps

1. **Echo the hill.** Copy the micro hill (Who, What, Wow, and what it ladders to) into the canvas header. Every later section has to serve it. Anything that does not belongs in Out of scope.

2. **To-be journey.** A mermaid diagram, not prose. Pick the form that fits:

   - a user journey or flow: `flowchart LR`
   - a system interaction: `sequenceDiagram`
   - an object lifecycle: `stateDiagram-v2`

   Carry the confidence markers from the Hill artefact into the node and edge labels where evidence is thin, spelled out as `(assumed)` or `(unknown)`. Never abbreviate them to letters inside a diagram.

3. **Decisions.** The core of the canvas. **One block per decision, never a table row.** A decision is read aloud in a meeting, so it needs room to be a few short sentences. Use this shape:

   ```markdown
   ### 3. Selecting offerings invites their companies to one direct purchase

   **Decided.** Picking two or three offerings invites those companies to a single direct purchase. The chosen offering supplies the requirement text.
   **Why.** This is already the comparison evidence the procurement record wants. Selecting candidates and producing the evidence become the same act.
   **Rejected.** A basket of line items. It is familiar from ordinary shopping, but the direct purchase model has no container for line items, and it contradicts the rule of one direct purchase per flow.
   **What it costs.** Shopping-basket language must not leak into the interface. It would promise a container the model cannot hold.
   **Who owns it.** Product owner, to confirm. Driver: legal and usability. Status: decided.
   ```

   Rules for the block:

   - The numbered heading is an **address**, so a meeting can point at it. Give it a short title as well as a number.
   - **`Rejected.` must say why the option lost**, not just name it. A rejected option with no reason is not a record of a decision.
   - `What it costs.` states the honest downside. Every real decision has one. If you cannot name it, you have not finished thinking.
   - One decision per block. Split anything compound.
   - Drivers: business, ux, feasibility, legal. A decision that crosses systems or is hard to reverse gets `Promote to ADR` in the `Who owns it.` line, targeting `bindings.decisions`.

4. **Minimum viable experience.** The smallest end-to-end slice that lets a real user accomplish the hill's What and produces a signal you can learn from. The smallest _journey_, not the smallest feature set. Three to six bullets.

5. **Signals.** Leading indicators (behaviour, time to task) and lagging ones (retention, outcome), each tied to the Wow. An unknown baseline is `[TBD, needs data]`.

6. **Boundaries.** What is explicitly out of scope, and the constraints. Cutting scope matters as much as choosing it.

7. **Open questions.** Group them under criticality headings: Critical, High, Medium, Low. **Critical means it can invalidate the work or block downstream work.** Give each question the thing that would resolve it.

   Use bullets, never a numbered list. Numbered questions turn into numbered references, which is the habit these rules exist to break.

   **If you have more than about eight open questions, you are listing rather than triaging.** Merge questions that the same person resolves in the same conversation. Drop anything that is really a task.

## Artefact grammar (Design Canvas, up to about 400 lines)

````markdown
# Canvas: {title}

|            |                                            |
| ---------- | ------------------------------------------ |
| Work       | {ticket} · `{home folder}` · slug `{slug}` |
| Hill       | **Who** {…} · **What** {…} · **Wow** {…}   |
| Ladders to | {macro id}                                 |
| Status     | draft                                      |
| Sources    | {hill artefact path, playback dates}       |

Confidence in this document is written as **known** (there is evidence),
**assumed** (a reasonable guess) or **unknown** (nobody knows yet).

## To-be journey

```mermaid
{flowchart, sequenceDiagram or stateDiagram-v2}
```

## Decisions

### 1. {short title}

**Decided.** {what was decided}
**Why.** {the reason it won}
**Rejected.** {the option, and why it lost}
**What it costs.** {the honest downside}
**Who owns it.** {name or role}. Driver: {business, ux, feasibility or legal}. Status: {proposed, decided or superseded}.

### 2. {short title}

{same five lines}

## Minimum viable experience

- {slice bullet}

## Signals

| Signal | Type    | Target | Baseline     |
| ------ | ------- | ------ | ------------ |
| {…}    | leading | {…}    | {… or [TBD]} |

## Out of scope

- {what you are not doing, and why in one short clause}

## Open questions

### Critical

- **{the question}** {why it can invalidate or block. Resolved by: sponsor user, data or spike.}

### High

- **{the question}** {resolved by …}

### Medium

### Low
````

The `Work` row is the identity every later stage inherits: same ticket, same folder, same slug. Never restate it differently.

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose. `What it costs.` is a sentence, not a minus sign.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?, including inside mermaid labels.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the artefact to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours. Keep `status: draft` until a playback agrees it.
2. Print the playback summary: canvas path, decision count by driver, the minimum viable experience in one line, how many critical open questions there are, and what the playback has to settle.
3. **STOP.** Do not emit briefs, stories or code.
4. Close with: "You are in **Reflect, moving into Make**. Next moves: `/cairn-handoff` once the Spec playback signs off. `/cairn-hill` if the playback breaks the hill."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Norms, the five record-quality criteria (apply to every decision block)

1. Minimal, cleanly cut scope.
2. One decision per record. Split anything compound.
3. The decision and its reason are the core. The rejected option gets a brief, objective reason for losing.
4. Honest consequences, good and bad, stated factually. No defensive or persuasive language.
5. Timeless and self-sufficient. No links to gitignored files or living docs. Quote what you need.

## Guardrails

- Sections are capped by the grammar above. New content **replaces** stale content. Never append a second journey or a "misc" section.
- Structure first: mermaid and tables. Prose is for reasoning. Use drawio (embedded XML) only where mermaid genuinely cannot express the thing.
- Every canvas serves exactly one hill. A second hill means a second canvas.
- **Protect the uncomfortable findings.** Contradictions between sources, unmeasured claims, and needs that are really policy requirements all stay visible. A "make it more readable" instruction is the first thing that will erode them, so do not let readability delete a caveat.
- Content in the user's working language. Keys and labels in English.
