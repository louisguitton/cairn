---
name: /cairn-reverse
id: cairn-reverse
category: Beta
stage: observe
next: /cairn-hill (validate the reconstructed hill), then a playback that turns assumed claims into known ones
description: Reverse-engineer an existing prototype into a draft canvas and micro hill, with every reconstructed claim marked assumed until a playback validates it
---

Read an **existing prototype**, where the request arrived as an implementation rather than a specification, and reconstruct the missing upstream artefacts: a draft micro hill and a draft Design Canvas. Recovering intent from code is lossy. Everything you reconstruct is **assumed** until a human validates it.

**Loop stage: Observe.** The prototype is the observation.

**Input**: a prototype directory or repo path. Optionally a sitemap or spec file, and the branch and routes to focus on. For a large prototype, ask which journey or routes to scope to. Never sweep everything: a hundred-route prototype is many canvases, not one.

## Config

Read `.cairn.yaml`. Establish the work identity the way `/cairn-intake` does: propose a ticket, a slug and a home folder under `bindings.specs`, then **ask the user to confirm before writing**. Artefacts go to `<home>/2-hill.md` and `<home>/3-canvas.md`. The numbers are the review order. There is deliberately no `1-intake.md`: the missing number records that no stakeholder intake happened, because the prototype was the request. Read macro hills from `bindings.hills`, which is read-only, to guess what this work ladders to.

## Steps

1. **Map the scoped journey.** Routes, entry points, exits. What the user can click, and in what order. Emit it as a mermaid `flowchart LR`.

2. **Reconstruct the needs.** For each screen and action, write the needs statement it presupposes, applying the recast rule. An element with no plausible user task behind it gets flagged. Do not rationalise it.

3. **Draft the micro hill.** Infer Who from who the interface addresses, and What from the journey. Be honest about the Wow: usually it is `[TBD, the Wow was never stated]`. Guess what it ladders to from the macro register and mark that guess as assumed.

   If the Wow you find describes a mechanism rather than an outcome for the user, say so. Do not quietly upgrade it into a better Wow than the prototype actually has.

4. **Harvest the embedded decisions.** Visible choices such as layout, flow order, what data is shown and the wording all became decisions at some point. Record each one as a block, never a table row, because a decision gets read aloud one at a time:

   ```markdown
   ### 1. {short title}

   **Decided.** {the choice the prototype makes, and where it is visible}
   **Why.** {the reason, if a source states one. Otherwise: no reason recorded.}
   **Rejected.** Unknown. The prototype records only the choice that shipped.
   **What it costs.** {the honest downside of the choice, if you can see one}
   **Who owns it.** [TBD]. Driver: {guessed}. Status: proposed.
   ```

   Never invent a rejected option to fill the line. The prototype cannot tell you what lost.

5. **Assemble the draft canvas** following the `/cairn-canvas` grammar, with this banner at the top:

   > **Reverse-engineered from {repo} on branch {branch}, {date}. Every claim in this document is assumed unless marked otherwise. Validate at the next playback before building on it.**

   Every journey cell and every decision defaults to **assumed**, spelled out in words. Only an externally evidenced fact, such as a linked specification or a commit message that states intent, may be marked known.

   Leave a scenario row empty and mark it unknown when nobody has observed a user. An empty row is a finding, not a gap to fill with plausible feelings.

## Writing rules

Your reader is a product owner or an engineer whose first language is probably German or French, reading English, in a meeting, with three other people and twenty minutes.

- **The title is one line of plain words that stands on its own.** Someone who has never opened this folder should understand what the work is from the title alone. State the outcome for the user, not the mechanism that delivers it. No jargon, no feature or product codes, no internal labels such as "Hill 2", no words that only mean something inside your team. Write it lowercase after the colon, as a phrase.
  - Good: `answer a described need with products the buyer can actually order`
  - Good: `one list of products and services, and any file can fill it`
  - Weak: `agentic catalogue in chat` (jargon, and it names the mechanism)
  - Weak: `portfolio upload and facets` ("facets" is an implementation word)
- One idea per sentence. Prefer a full stop to a semicolon, and a sentence to a subordinate clause.
- State facts. Do not argue for them. No rhetorical flourishes, no dramatic framing, no sentence whose job is to sound conclusive.
- **Never use an em dash.** Use a full stop, a comma, a colon, or brackets.
- Never use `+` or `-` as connectives in prose.
- **Never refer to anything by an identifier alone**, in this document or in a source. Write what it says. A route is an address, so keep it with a description attached, as in "the demand registration screen, /procurement/requests/register".
- Write confidence as words: known, assumed, unknown. Never as K, A or ?, including inside mermaid labels.
- Write figures in full, in the reader's convention: 50.000 EUR, not €50k.
- Prefer the plain word to the sophisticated one when they mean the same thing.
- **Simplify the sentence, never the claim.** A hedge, a confidence marker, or a note that something is unmeasured is content. If a situation is genuinely uncertain and hard to write plainly, write two plain sentences. Never delete the uncertainty.

## Gate (hard stop)

1. Write the artefacts to the working tree. **Stop there. No git writes.** Never run `git commit`, `git push`, or open a PR or MR. The gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Print the summary: which routes you covered, how many needs you reconstructed, how many decisions are proposed, and the three biggest unknowns.
3. **STOP.** Next: `/cairn-hill` to pressure-test the reconstructed hill, then a validation playback.

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr`: branch, commit, push, open the PR or MR as a draft.
- `commit`: commit on the current branch, no push.
- `none`: nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- Never present a reconstruction as fact. The banner and the assumed markers are mandatory.
- Never write the macro hill register.
- Scope before sweeping.
- **Protect the uncomfortable findings.** This is the most valuable thing this command produces. A Wow that describes a mechanism, a scenario row nobody has observed, a stated user need that is really a policy requirement served through the user's screen, two sources that contradict each other: record all of them as they are. Never smooth them over, and never let a later instruction to make the document more readable delete them.
