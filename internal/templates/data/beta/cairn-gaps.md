---
name: /cairn-gaps
id: cairn-gaps
category: Beta
stage: observe
writes: none
next: /cairn-intake (to record answers once you have them), /cairn-hill (once the top gap is closed or accepted as a risk)
description: Explain the gaps and blocking assumptions in an intake or hill in plain language, and lay out the concrete ways to close each one
---

Take the gaps an artefact recorded and explain them to a human who was not in the room. What is actually missing, why it matters for this particular piece of work, and what the realistic options are for closing it.

**Loop stage: Observe.** This command exists because "Who: MISSING" tells you a fact without telling you what to do about it.

**This command writes nothing.** It produces no artefact and has no gate, because it makes no decisions and records nothing new. It is a reading aid. Once you have answers, `/cairn-intake` records them, and `/cairn-hill` moves on.

**Input**: an Intake Brief (`1-intake.md`) or a Hill artefact (`2-hill.md`). Without one, ask for it and stop.

## Config

Read `.cairn.yaml`, then take the work identity from the input artefact's `Work` header, so you can name the work you are talking about. Read `bindings.personas` if it is bound: a gap about Who is easier to close if you can see which personas already exist.

## Steps

1. **Collect what is open.** From the input artefact, gather:

   - every gap marked MISSING (Who, What, Why now, Wow)
   - every assumption marked assumed or unknown, and any in the blocking group
   - every cell marked unknown in a scenario map
   - every element flagged as a feature with no user behind it
   - the top gap the artefact already nominated

2. **Explain each one to somebody new.** Four short parts per gap, in this order:

   - **What is missing.** Say the actual missing thing, not its category. "We do not know who signs this off" beats "the Who is missing".
   - **Why it matters here.** The consequence for this specific work, not a general principle. If getting it wrong would waste a week of building, say so. If it would only change a label, say that too: not every gap is expensive.
   - **How to close it.** Two to four concrete options. For each: who does it, roughly how long it takes, and what you get at the end. Prefer the cheapest option that would actually settle the question.
   - **Or live with it.** What it costs to accept the gap as a known risk and carry on. Some gaps are worth accepting, and saying so is more useful than implying everything must be resolved.

3. **Rank them.** Say which one to close first and why. Usually it is the gap that would invalidate the most work if the answer surprises you, not the gap that is easiest to close.

4. **Offer the ask.** For any gap that a person could simply answer, draft the message to send them: short, specific, and answerable in one reply. The intake brief already carries a question block for the stakeholder, so do not repeat it. Cover the people it does not: a sponsor user, a domain expert, whoever owns the data.

## How to explain

Assume the reader is intelligent, busy, and new to this work. They do not know the method's vocabulary and should not have to learn it to understand the answer.

- No method jargon without a definition. If you need the words hill, Wow, sponsor user, playback or ladders to, define each in one short clause the first time it appears.
- Explain the thing, not the label. A reader who has understood you can restate the gap in their own words without using any cairn terminology.
- Use a concrete example wherever it is quicker than a definition. "For example, if the buyer turns out to be the department head rather than the assistant, the whole approval step changes" teaches more than a paragraph about personas.
- Say what nobody knows, and say it plainly. If the honest answer is that nobody in the team has ever spoken to this user, that is the finding, and it should be the first sentence rather than the last.
- Never invent the answer while explaining the question. A gap explained is still a gap.

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- **One language per sentence, and one content language throughout.** Never mix languages inside a sentence, in either direction. If the artefact you are reading is written in German, answer in German, including the sentence patterns. Translate the pattern, never keep English connectives around localised content.
  - Wrong: `Ein Beschaffer needs a way to seinen Bedarf zu begruenden so that die Sparsamkeit belegt ist`
  - Right: `Ein Beschaffer braucht eine Moeglichkeit, seinen Bedarf zu begruenden, damit die Sparsamkeit belegt ist`
  - Furniture stays English: section headings, table column labels, and the controlled vocabularies (known, assumed, unknown). These are short, fixed anchors.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. Ticket and commit ids are addresses rather than references: keep them, with a description attached, as in "the price visibility bug, BUG-2578".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Close

1. **Write nothing.** No file, no commit, no branch. Explaining is the whole job.
2. Print the ranked list, then the single sentence a reader should walk away with: which question to answer first, and who can answer it.
3. Close with: "You are in **Observe**. Next moves: `/cairn-intake` to record the answers once you have them. `/cairn-hill` once the top gap is closed, or explicitly accepted as a risk."

## Guardrails

- **This command never writes an artefact.** If the user wants the explanation kept, tell them to paste it where they want it. Do not invent a sixth numbered file in the work folder: the numbering is the review agenda, and a reading aid is not part of it.
- Never close a gap by guessing. Proposing how to find the answer is the job. Producing the answer is not.
- **Protect the uncomfortable findings.** If the honest explanation is that the work rests on an assumption nobody has tested, say it in the plainest sentence you can write. Making a gap sound smaller than it is defeats the purpose of having recorded it.
- Do not restate the intake's stakeholder question block. Add the asks it does not already cover.
