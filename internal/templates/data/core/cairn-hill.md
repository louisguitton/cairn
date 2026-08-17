---
name: /cairn-hill
id: cairn-hill
category: Core
stage: reflect
next: /cairn-canvas (after the Hill playback agrees on Who/What/Wow)
description: Reframe an intake into a problem statement, needs statements, a micro hill laddering to a macro hill, as-is/to-be scenario maps, and an assumptions grid
---

Reframe what Observe collected into a **Hill artefact**: problem statement, needs statements, one **micro hill** (Who/What/Wow) laddering to a macro hill, paired as-is and to-be scenario maps, and an assumptions grid.

**Loop stage: Reflect.** You are converging on intent. You are not designing screens or flows.

**Input**: an Intake Brief (`@file`), or a direct description if the request is small enough to skip intake. If you skip intake, say so in the artefact: `entered at: hill`.

## Config

1. Read `.cairn.yaml`, then take the **work identity from the input artefact's `Work` header**: the home folder, ticket and slug. Reuse them exactly. Never re-derive a slug or invent a second folder for the same work.
   - Input has no `Work` header (you entered mid-pipeline)? Propose ticket, slug and home the way `/cairn-intake` does, then **ask the user to confirm before writing**.
2. Write to `<home>/2-hill.md`. The number is the review order: the hill is read before the canvas, always.
3. `bindings.hills` is the **macro hill register**. Read it. Never write it. If unbound, no macro hills exist yet. See step 2 below.
4. Prepend the `gating.frontmatter` template if one is declared.

## Steps

1. **Read the macro hills** from the register. Note each hill's id, its Who/What/Wow, and its owner.

2. **Pick the ladder.** Which macro hill does this work serve?

   - Exactly one fits. Record `ladders_to: {id}`.
   - None fits. **Do not invent one and do not skip.** Record `ladders_to: NONE` with one honest sentence explaining why, and flag it in the gate summary. Either this is the wrong work, or a macro-level conversation is due. That finding is a valid, gate-worthy output.
   - The register is unbound or empty. Draft the micro hill anyway, mark `ladders_to: [TBD, no macro register bound]`, and suggest creating one.

3. **Problem statement.** One sentence, user-centred, anchored in the as-is, with no solution in it:

   > _{Specific user} struggles to {do something} because {as-is reason}, which costs them {concrete cost}._

   If the cost is unknown, write `[TBD, needs interview or data]`. Never invent it.

4. **Needs statements.** One per user task. Enforce the recast rule: no click, view or export verbs. Use task verbs such as prioritise, decide, compare, recover, trust. If you have more than five, cluster them and write one statement that covers the cluster.

5. **Micro hill.** Same grammar as a macro hill, smaller blast radius:

   - **Who**: a specific persona. Link the personas file if it is bound.
   - **What**: the new capability in human terms. Name no technology.
   - **Wow**: measurable, ideally a slice of the macro Wow (a time, a count, a percentage, or a feeling with a probe attached).
   - **Timeframe**: day, week or sprint.

6. **As-is and to-be scenario maps.** Four rows (Phases, Doing, Thinking, Feeling), four to seven phases. Mark **every cell** as known, assumed or unknown, spelled out in words. The negative cells in the as-is Feeling row are where the Wow has to land. Say which ones.

   Leave a cell empty and mark it unknown rather than inventing a plausible feeling. An empty row is a finding: nobody has observed a user yet.

7. **Assumptions grid.** Place every assumption on certainty against risk. **Key each assumption by a short name**, two or three words, so a reviewer can say it out loud. Do not number them. The uncertain and high-risk group blocks the build, and each entry there needs a validation plan: a sponsor-user session, a data pull, or a spike.

## Artefact grammar (Hill artefact)

```markdown
# Hill: {title}

|            |                                                             |
| ---------- | ----------------------------------------------------------- |
| Work       | {ticket} · `{home folder}` · slug `{slug}`                  |
| Ladders to | {macro id, plus a one-line quote of its What}               |
| Timeframe  | {day, week or sprint}                                       |
| Entered at | {intake then hill, or hill directly with the reason}        |
| Source     | {intake brief path, or the request quoted}                  |
| Language   | {the content language of this document, for example German} |

Confidence in this document is written as **known** (there is evidence),
**assumed** (a reasonable guess) or **unknown** (nobody knows yet).

## Problem statement

> {one sentence}

## Needs statements

- {user} needs a way to … so that … ({known, assumed or unknown})

## Micro hill

- **Who:** {persona}
- **What:** {capability, in human terms}
- **Wow:** {measurable differentiator}

## As-is

|          | {phase 1}     | {phase 2}     |
| -------- | ------------- | ------------- |
| Doing    | {…} (known)   | {…} (assumed) |
| Thinking | {…} (assumed) | {…} (unknown) |
| Feeling  | {…} (assumed) | (unknown)     |

## To-be

{the same four-row shape. Mark the cells the Wow changes.}

## Assumptions

| Assumption          | What it claims | Certainty | Risk | Validation plan                            |
| ------------------- | -------------- | --------- | ---- | ------------------------------------------ |
| **{two-word name}** | {one sentence} | assumed   | high | {sponsor-user session, data pull or spike} |

Blocking group (uncertain and high risk): {name them in words, not codes}.
```

The `Work` row is the identity every later stage inherits: same ticket, same folder, same slug. Never restate it differently.

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- **The title is one line of plain words that stands on its own.** Someone who has never opened this folder should understand what the work is from the title alone. State the outcome for the user, not the mechanism that delivers it. No jargon, no feature or product codes, no internal labels such as "Hill 2", no words that only mean something inside your team. Write it lowercase after the colon, as a phrase.
  - Good: `answer a described need with products the buyer can actually order`
  - Good: `one list of products and services, and any file can fill it`
  - Weak: `agentic catalogue in chat` (jargon, and it names the mechanism)
  - Weak: `portfolio upload and facets` ("facets" is an implementation word)
- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- **One language per sentence, and one content language per artefact.** Name the content language in the artefact header. Never mix languages inside a sentence, in either direction.
  - The sentence patterns in this template are shown in English because the template is English. **Translate the pattern into the content language.** Do not keep the English connectives and drop content into the gaps.
  - Wrong: `Ein Beschaffer needs a way to seinen Bedarf zu begruenden so that die Sparsamkeit belegt ist`
  - Right: `Ein Beschaffer braucht eine Moeglichkeit, seinen Bedarf zu begruenden, damit die Sparsamkeit belegt ist`
  - Furniture stays English: section headings, table column labels, frontmatter keys, and the controlled vocabularies (known, assumed, unknown; proposed, decided, superseded; business, ux, feasibility, legal). These are short, fixed anchors that keep artefacts recognisable and searchable across teams. Give their translation once in the legend if that helps the reader, then use the English term consistently.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose. Write the cost as a sentence.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?. Include the two-line legend shown above.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the artefact to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: artefact path, the micro hill in one line, what it ladders to (or the NONE finding), how many cells are unknown, and which assumptions block the build.
3. **STOP.** The gate is the **Hill playback**: product owner, design and engineering agree on Who/What/Wow before anyone designs a journey. Offer the three sentences to open that playback with.
4. Close with: "You are in **Reflect**. Next moves: `/cairn-canvas` after the Hill playback. Back to `/cairn-intake` if the playback exposes a missing Who or Why."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **Never write the macro hill register.** Read-only, always. Promoting a micro hill to macro is a human act outside this command.
- No solution or screen language in Who, What or Wow.
- Every scenario cell says known, assumed or unknown. An unknown cost or metric is `[TBD, needs X]`.
- **Protect the uncomfortable findings.** If the Wow describes a mechanism instead of an outcome, say so rather than quietly improving it. If a stated user need is really a policy requirement served through the user's screen, record it that way. If two sources contradict each other, surface the contradiction instead of picking the convenient one. These findings are the point of the artefact.
- Structure first. Prose belongs in the problem statement and the rationale lines.
