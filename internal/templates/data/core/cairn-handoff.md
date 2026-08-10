---
name: /cairn-handoff
id: cairn-handoff
category: Core
stage: make
next: build the prototype from the brief, or run /spdd-analysis on the canvas; /cairn-sync when reality reports back
description: Emit the Make-stage handoffs from a signed-off canvas: a prototype brief, a story set (As a, I want, so that, plus Given/When/Then), and an /spdd-analysis input block
---

Convert a signed-off Design Canvas into **handoff artefacts**: a prototype brief for the prototyping agent, a story set for the backlog, and an input block for `/spdd-analysis`. Specification happened upstream. This stage makes it executable by other people.

**Loop stage: Make.** Still no code. The outputs _commission_ the making.

**Input**: a Design Canvas (`@file`) whose Spec playback is done, so its status is no longer draft. If it is still draft, say so and ask to confirm before proceeding. If the user did not say which handoffs they want, ask: prototype, stories, spdd, or any combination.

## Config

1. Read `.cairn.yaml`, then take the **work identity from the input artefact's `Work` header**: the home folder, ticket and slug. Reuse them exactly. Never re-derive a slug or invent a second folder for the same work.
   - Input has no `Work` header (you entered mid-pipeline)? Propose ticket, slug and home the way `/cairn-intake` does, then **ask the user to confirm before writing**.
2. Write to `<home>/brief.md` and `<home>/stories.md`, next to the canvas they derive from.
3. `proto_repo.path` is the target the brief addresses, using its branch and route scheme.
4. Prepend `gating.frontmatter` if declared.

## Steps

1. **Verify the chain.** The canvas has to carry the hill and at least one decided decision. If either is missing, stop and route back to `/cairn-canvas`.

2. **Prototype brief**, if wanted. This is the commission for an experiment:

   - **The hill and its Wow**, quoted. The prototype exists to test this.
   - **The question this experiment answers**, in one sentence. A prototype with no question is scope creep.
   - **The journey slice**: which part of the to-be journey to build. Copy the relevant mermaid subgraph.
   - **Addressing**: the target repo, the branch name to use, the routes to mount. Branch and route is how every later citation points back at the prototype.
   - **Fidelity and constraints**: is mock data acceptable, which design system, and what must not be built, taken from Out of scope.
   - **Report back**: tell the builder that decisions taken while building are recorded through `/cairn-sync` rather than left in the code.

3. **Story set**, if wanted. One story per journey slice the team will build:

   - The grammar is not negotiable. **As a** {persona}, **I want to** {action}, **so that** {goal}. Acceptance criteria as **Given, When, Then** scenarios, covering the happy path and the edge cases that matter.
   - Reference the prototype **by branch and route**, never by screenshot.
   - Each story names its hill and the canvas decisions it implements, **by title as well as number**, so a reader never has to look up a bare code.
   - If the host repo has a story skill or template, emit in that shape so it can be piped straight in.

4. **SPDD input block**, if wanted. A self-contained requirement block for `/spdd-analysis`: the hill, the journey slice, the decided decisions that constrain implementation, the boundaries, and the signals to instrument. Quote what SPDD will need. Do not link to it, because it will be read in another repo.

## Artefact grammar (Prototype brief, up to about 80 lines)

````markdown
# Brief: {title}

|          |                                                   |
| -------- | ------------------------------------------------- |
| Work     | {ticket} · `{home folder}` · slug `{slug}`        |
| Hill     | {id}: **Wow** {…}                                 |
| Question | {what this experiment answers}                    |
| Target   | {proto repo} · branch `{name}` · routes `{/path}` |
| Fidelity | {mock data, design system, …}                     |

## Journey slice

```mermaid
{the relevant subgraph of the canvas journey}
```

## Build

- {concrete slice bullet}

## Do not build

- {taken from Out of scope}

## Report back

- Record decisions taken while building through /cairn-sync, using a transcript, notes or the diff.
````

The `Work` row is the identity every later stage inherits: same ticket, same folder, same slug. Never restate it differently.

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. A story that implements a canvas decision names the decision, not just its number. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the artefacts to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the playback summary: which artefacts you wrote, the hill, the question under test, and how many stories there are.
3. **STOP.** The gate is **build kickoff**: a human picks up the brief or the stories and starts.
4. Close with: "You are in **Make**. Next moves: build the prototype from the brief. Pipe the stories to the backlog. Run `/spdd-analysis` with the SPDD block. Run `/cairn-sync` as soon as reality produces decisions."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **No code generation.** Briefs, stories and analysis input. Nothing else.
- Every handoff artefact names its hill. An artefact that cannot name its hill does not ship.
- A story without Given, When, Then scenarios is not a story. Refuse to emit one bare.
- The SPDD block has to be self-sufficient. Quote rather than link, because it will be read in another repo.
- **Protect the uncomfortable findings.** If the canvas left something unmeasured or contradictory, the brief and the stories carry that forward rather than quietly resolving it.
