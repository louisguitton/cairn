---
name: /cairn-intake
id: cairn-intake
category: Core
stage: observe
next: /cairn-hill (after the stakeholder answers the top gap), /cairn-reverse (if the input is an existing prototype)
description: Diagnose a raw request (download, screenshot, doc, prototype, transcript) into an Intake Brief with needs statements, gaps, and a copy-paste question block
---

Turn a raw stakeholder request into an **Intake Brief**: a one-page diagnosis of what was actually handed over, what user need hides underneath, and the smallest set of questions that closes the highest-risk gaps.

**Loop stage: Observe.** You are collecting evidence, not designing. Do not propose solutions.

**Input**: whatever the user has — a pasted paragraph, a screenshot/Figma, a doc or Confluence dump (`@file`), a prototype path, a meeting transcript. If nothing is provided, ask for it and stop.

## Config

1. Read `.cairn.yaml` from the repo root (search upward). Resolve the artefact path:
   - `bindings.capture` set → write the brief there (it is a dated, append-only capture area — never edit existing files in it).
   - Unbound → write to `design/intake/` (create if missing).
2. File name: `{TICKET}-{YYYYMMDDHHmm}-[Intake]-{kebab-slug}.md` (use the ticket key if one is mentioned, else `XXX`).
3. If `gating.frontmatter` points to a template, prepend that frontmatter to the artefact.

## Steps

1. **Classify the input** (state the class in the brief):
   - **Solution-shaped** — a screenshot, mockup, or "we need a dashboard/button/export". Machine language.
   - **Feature-shaped** — "users should be able to do X". Pre-decided action, user still implied.
   - **Outcome-shaped** — "users get stuck at Y". Already user-centred (rare).
   - **Prototype** — an existing implementation is the request. Recommend `/cairn-reverse` and stop unless the user insists.
   - **Transcript / notes** — extract the requests inside first, then classify each.

2. **Deconstruct the artefact.** For every UI element or asserted feature, write the implicit needs statement:

   > _[User] needs a way to ___ so that ___._

   **Recast rule (enforce hard)**: if an idea is expressed in machine terms — dashboard, click, log in, export, filter — it is a feature, not a need. Recast it in terms of what the human accomplishes (prioritize, decide, compare, recover, trust). An element you cannot recast is a feature with no user behind it — flag it, don't invent a user.

3. **Inventory the four gaps** — plainly, no editorializing:
   - **Who** is the user? (role is not enough — tasks, motivations, obstacles)
   - **What** are they trying to accomplish, in their language?
   - **Why** now — and what does the as-is cost them?
   - **Wow** — what would make them say "thank you"? The differentiator.

4. **Surface assumptions.** At least three, each labelled **K** (known — evidence exists), **A** (assumed — reasonable guess), or **?** (unknown). Never fabricate; write `[TBD — needs stakeholder / sponsor user / data]` instead.

5. **Pick the single top gap** and write the question block: 5–7 collaborative questions, copy-paste ready for Slack/email, top gap first. Always include one 5-Whys chain and one sponsor-user probe ("who is one real user we could put this in front of in the next two weeks?").

## Artefact grammar (Intake Brief, ≤ ~120 lines)

```markdown
# Intake — {title}

|                        |                                                            |
| ---------------------- | ---------------------------------------------------------- |
| Date                   | {date}                                                     |
| Stakeholder            | {who, role}                                                |
| Input class            | {solution/feature/outcome-shaped / prototype / transcript} |
| Original artefact      | {link / quoted sentence}                                   |
| Ladders to (suspected) | {macro hill id or [TBD]}                                   |

## Stated request (their words)

> "{verbatim quote}"

## Implicit needs statements

| #   | Needs statement                   | Evidence                    |
| --- | --------------------------------- | --------------------------- |
| N1  | {user} needs a way to … so that … | K/A/?                       |
| —   | {element that cannot be recast}   | FLAG — feature with no user |

## Gaps

| Gap     | Status                |
| ------- | --------------------- |
| Who     | {answer or [MISSING]} |
| What    | {answer or [MISSING]} |
| Why now | {answer or [MISSING]} |
| Wow     | {answer or [MISSING]} |

## Assumptions detected

- A1 (A): {assumption}
- A2 (?): {assumption}

## Top gap

> {one sentence}

## Questions for {stakeholder} (copy-paste)

{5–7 questions, top gap first, incl. one 5-Whys chain + sponsor-user probe}
```

## Gate (hard stop)

1. Write the artefact. If `gating.mode: pr` — branch + commit + open a PR; `commit` — commit directly; `none` — just write.
2. Print the playback summary: artefact path · input class · gaps missing (count) · top gap · who must answer.
3. **STOP.** Do not draft a hill, journey, or solution. The stakeholder answering the top gap is the gate.
4. Close with: "You are in **Observe**. Next moves: `/cairn-hill` once the top gap is answered (or explicitly accepted as risk); `/cairn-reverse` if the request is really an existing prototype."

## Guardrails

- No solution language in the brief — Observe only.
- Never fabricate user research; `[TBD]` beats fiction.
- Content in the user's working language; structure and labels stay in English.
- Structure-first: tables over prose; the only free prose is the stated-request quote and the top gap.
