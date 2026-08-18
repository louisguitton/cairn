---
name: /cairn-coverage
id: cairn-coverage
category: Beta
stage: reflect
next: /cairn-canvas (to add what fell out, or to rule it out explicitly), then /cairn-handoff
description: Check a canvas back against the original source material and report what is covered, what was ruled out, and what fell out silently
---

Check a Design Canvas back against the material it came from. Every earlier stage compresses: a request becomes needs statements, needs become one Wow, a journey becomes decisions. Compression loses things, and nothing in the pipeline looks back at the source to see what. This command does that one job.

**Loop stage: Reflect.** Run it at the canvas gate, before the Spec playback signs the canvas off. That is the last point where adding something back is a cheap edit rather than a reopened agreement.

**Input**: the canvas (`3-canvas.md`) **and** the original source material: the requirements catalogue, the wiki export, the meeting notes, the ticket, whatever the work was actually derived from. Without the source there is nothing to check against, so ask for it and stop.

## Config

1. Read `.cairn.yaml`, then take the work identity from the input artefact's `Work` header. Reuse it exactly.
2. Write to `<home>/coverage.md`. **Unnumbered on purpose.** The numbers 1 to 5 are the review agenda for the team. This artefact is addressed to one person, the one who knows the source material best, so it sits beside the agenda rather than inside it.
3. Read `bindings.capture` for any raw source already filed there.

## Steps

1. **Enumerate the source, and say how.** List the claims, requirements or asks the source contains. Then state plainly how you enumerated them and what you could not: which sections you read, which you skipped, how long the document is, whether anything was an image, a table you could not parse, or a language you handled less well.

   **This step decides whether the whole report is worth anything.** A coverage report that implies completeness it does not have is worse than no report, because it converts an open worry into false comfort.

2. **Classify every enumerated item** into exactly one of three outcomes:

   - **Covered.** It appears in the canvas. Say where: which decision, which journey step, which signal.
   - **Ruled out.** The canvas explicitly excludes it. Quote the line from `Out of scope`.
   - **Fell out.** It is in the source and appears nowhere in the canvas, neither included nor excluded. This is the finding. Everything else is bookkeeping.

3. **Say what you suspect.** Beyond what you enumerated, name up to three things you think may have been missed and cannot confirm: a section you could not parse, a theme that appears in the source but nowhere in the canvas, a term used in the source that the canvas never mentions. Mark each as a suspicion, not a finding.

   Note the inversion. `/cairn-sync` extracts conservatively, because there a miss beats noise. Here a miss is precisely what you are hunting, so a conservative check is useless. Report the suspicion and label it as one.

4. **Ask three questions.** Three questions for the source owner, each aimed at something you could not settle from the documents alone. Make each answerable in one reply. Good questions sound like "the catalogue mentions X in section 4 but the canvas never does, is that deliberate?" rather than "is anything missing?".

5. **State the count.** How many items you enumerated, how many of each outcome, and your honest estimate of the denominator. "Enumerated 40 items. Covered 31, ruled out 5, fell out 4. The catalogue is about 40 pages and probably holds nearer 120 items, so treat this as a sample rather than a full audit."

## Artefact grammar (Coverage report, up to about 200 lines)

```markdown
# Coverage: {title}

|          |                                                             |
| -------- | ----------------------------------------------------------- |
| Work     | {ticket} · `{home folder}` · slug `{slug}`                  |
| Canvas   | `3-canvas.md`, as of {date}                                 |
| Sources  | {each source, with what it is}                              |
| Language | {the content language of this document, for example German} |
| For      | {the person who knows the source material}                  |

## How far this check reaches

{what you read, what you could not, and the honest denominator. First, because
every number below depends on it.}

## Fell out

The findings. In the source, absent from the canvas, neither included nor excluded.

| From the source      | Where it says so   | Why it may matter             |
| -------------------- | ------------------ | ----------------------------- |
| {the item, in words} | {section or quote} | {consequence if it stays out} |

## Ruled out

| From the source | The canvas line that excludes it |
| --------------- | -------------------------------- |

## Covered

| From the source | Where the canvas holds it |
| --------------- | ------------------------- |

## Suspicions

- {something you think was missed and cannot confirm, marked as a suspicion}

## Three questions for {name}

- {question answerable in one reply}

## Count

{enumerated, covered, ruled out, fell out, and the estimated denominator}
```

## Writing rules

Your reader is the one person who knows the source material best. They are busy, they will read this once, and they are looking for one thing: what did we drop.

- **The title is one line of plain words that stands on its own.** Someone who has never opened this folder should understand what the work is from the title alone. State the outcome for the user, not the mechanism that delivers it. No jargon, no feature or product codes, no internal labels.
- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- **One language per sentence, and one content language throughout.** Never mix languages inside a sentence, in either direction. If the source and the canvas are in German, write this report in German, including the sentence patterns. Translate the pattern, never keep English connectives around localised content.
  - Wrong: `Ein Beschaffer needs a way to seinen Bedarf zu begruenden so that die Sparsamkeit belegt ist`
  - Right: `Ein Beschaffer braucht eine Moeglichkeit, seinen Bedarf zu begruenden, damit die Sparsamkeit belegt ist`
  - Furniture stays English: section headings, table column labels, and the controlled vocabularies (known, assumed, unknown). These are short, fixed anchors.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what a source section says, not just its number. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the artefact to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the summary: how many items fell out, the reach of the check in one line, and the three questions.
3. **STOP.** Do not edit the canvas. Deciding what to do with something that fell out is the human's call, and it has exactly two legitimate destinations: into the canvas, or into `Out of scope` with a reason.
4. Close with: "You are in **Reflect**. Next moves: `/cairn-canvas` to add what fell out or to rule it out explicitly. Then the Spec playback, then `/cairn-handoff`."

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- **Never imply completeness you do not have.** The reach section comes first for that reason. "I checked everything" is almost always false on a real catalogue.
- Never resolve a finding yourself. Something that fell out may have been dropped deliberately by someone who knew more than the document says.
- Never edit the canvas from this command. One artefact, one author.
- **Protect the uncomfortable findings.** An item that fell out is uncomfortable by definition: it means the compression lost something. Report it plainly, at the top, before the bookkeeping. If the honest count is that a quarter of the source is unaccounted for, lead with that.
- Do not restate the canvas. This report is a difference, not a summary.
