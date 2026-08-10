---
name: /cairn-intake
id: cairn-intake
category: Core
stage: observe
next: /cairn-hill (after the stakeholder answers the top gap), /cairn-reverse (if the input is an existing prototype)
description: Diagnose a raw request (a download, screenshot, document, prototype or transcript) into an Intake Brief with needs statements, gaps and a copy-paste question block
---

Turn a raw stakeholder request into an **Intake Brief**: a one-page diagnosis of what was actually handed over, what user need hides underneath, and the smallest set of questions that closes the highest-risk gaps.

**Loop stage: Observe.** You are collecting evidence, not designing. Do not propose solutions.

**Input**: whatever the user has. A pasted paragraph, a screenshot or Figma frame, a document or Confluence dump (`@file`), a prototype path, a meeting transcript. If nothing is provided, ask for it and stop.

## Config

1. Read `.cairn.yaml` from the repo root, searching upward.
2. **Establish the work identity before writing anything.** This command names the work. Every later stage inherits that name and must never re-derive it.
   - **Ticket**: take the key from the request, for example `APP-4947`. If none is mentioned, ask for one, and only fall back to `XXX` if the user says there is no ticket.
   - **Slug**: three to six kebab-case words naming the outcome, not the mechanism.
   - **Home**: `{specs}/<existing stream>/<existing epic folder>/<slug>/` when `bindings.specs` is set. List the real subfolders and place the work under the matching one. If unbound, use `design/<slug>/`.
   - Show the full proposed path plus the ticket and slug, then **ask the user to confirm or edit it**. Do not write until they answer.
3. Write the brief to `<home>/intake.md`. All stage artefacts for this work live in that one folder.
4. `bindings.capture` is for **raw source material** only, such as transcripts and stakeholder dumps. If the user supplied raw input worth keeping, file a dated copy there and cite it. Never put the brief itself there.
5. If `gating.frontmatter` points to a template, prepend that frontmatter to the artefact.

## Steps

1. **Classify the input** and state the class in the brief:

   - **Solution-shaped**: a screenshot, a mockup, or "we need a dashboard, a button, an export". This is machine language.
   - **Feature-shaped**: "users should be able to do X". The action is pre-decided, the user is still implied.
   - **Outcome-shaped**: "users get stuck at Y". Already user-centred, and rare.
   - **Prototype**: an existing implementation is the request. Recommend `/cairn-reverse` and stop, unless the user insists.
   - **Transcript or notes**: extract the requests inside first, then classify each one.

2. **Deconstruct the artefact.** For every interface element or asserted feature, write the needs statement it presupposes:

   > _{user} needs a way to {do a task} so that {they benefit}._

   **Enforce the recast rule.** If an idea is expressed in machine terms such as dashboard, click, log in, export or filter, it is a feature rather than a need. Recast it as what the human accomplishes: prioritise, decide, compare, recover, trust. An element you cannot recast is a feature with no user behind it. Flag it. Do not invent a user.

3. **Inventory the four gaps**, plainly, with no editorialising:

   - **Who** is the user? A role is not enough. What are their tasks, motivations and obstacles?
   - **What** are they trying to accomplish, in their language?
   - **Why** now, and what does the current situation cost them?
   - **Wow**: what would make them say thank you? The differentiator.

4. **Surface assumptions.** At least three. Mark each as **known** (evidence exists), **assumed** (a reasonable guess) or **unknown** (nobody knows yet), spelled out in words. Give each one a short name of two or three words so a reviewer can say it out loud. Never fabricate. Write `[TBD, needs stakeholder, sponsor user or data]` instead.

5. **Pick the single top gap**, then write the question block: five to seven collaborative questions, ready to paste into Slack or email, top gap first. Always include one five-whys chain and one sponsor-user probe, such as "who is one real user we could put this in front of in the next two weeks?".

## Artefact grammar (Intake Brief, up to about 120 lines)

```markdown
# Intake: {title}

|                        |                                                                            |
| ---------------------- | -------------------------------------------------------------------------- |
| Work                   | {ticket} · `{home folder}` · slug `{slug}`                                 |
| Date                   | {date}                                                                     |
| Stakeholder            | {who, role}                                                                |
| Input class            | {solution-shaped, feature-shaped, outcome-shaped, prototype or transcript} |
| Original artefact      | {path, or the sentence quoted}                                             |
| Ladders to (suspected) | {macro hill id, or [TBD]}                                                  |

Confidence in this document is written as **known** (there is evidence),
**assumed** (a reasonable guess) or **unknown** (nobody knows yet).

## Stated request, in their words

> "{verbatim quote}"

## Needs statements the request presupposes

| Needs statement                   | Confidence                                              |
| --------------------------------- | ------------------------------------------------------- |
| {user} needs a way to … so that … | assumed                                                 |
| {element that cannot be recast}   | This is a feature with no user behind it. Ask about it. |

## Gaps

| Gap     | Status               |
| ------- | -------------------- |
| Who     | {answer, or MISSING} |
| What    | {answer, or MISSING} |
| Why now | {answer, or MISSING} |
| Wow     | {answer, or MISSING} |

## Assumptions

| Assumption          | What it claims | Confidence |
| ------------------- | -------------- | ---------- |
| **{two-word name}** | {one sentence} | assumed    |

## Top gap

> {one sentence}

## Questions for {stakeholder}, ready to paste

- {question}

{five to seven questions, top gap first, including one five-whys chain and one sponsor-user probe}
```

The `Work` row is the identity every later stage inherits: same ticket, same folder, same slug. Never restate it differently.

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the artefact to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: artefact path, input class, how many gaps are missing, the top gap, and who has to answer it.
3. **STOP.** Do not draft a hill, a journey or a solution. The stakeholder answering the top gap is the gate.
4. Close with: "You are in **Observe**. Next moves: `/cairn-hill` once the top gap is answered, or explicitly accepted as a risk. `/cairn-reverse` if the request is really an existing prototype."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- No solution language in the brief. This is Observe only.
- Never fabricate user research. `[TBD]` beats fiction.
- **Protect the uncomfortable findings.** An element with no user behind it, a gap nobody can fill, a stated need that is really a policy requirement: record them as they are. Do not dress them up.
- Content in the user's working language. Structure and labels stay in English.
- Structure first: tables over prose. The only free prose is the quoted request and the top gap.
