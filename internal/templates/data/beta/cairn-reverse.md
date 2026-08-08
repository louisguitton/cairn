---
name: /cairn-reverse
id: cairn-reverse
category: Beta
stage: observe
next: /cairn-hill (validate the reconstructed hill), then a playback to convert A-cells to K
description: Reverse-engineer an existing prototype into a draft canvas + micro hill — every reconstructed claim labelled Assumed until a playback validates it
---

Read an **existing prototype** (the request arrived as an implementation, not a spec) and reconstruct the missing upstream artefacts: a draft micro hill and a draft Design Canvas. Reverse-engineering intent from code is lossy — everything reconstructed is **Assumed** until a human validates it.

**Loop stage: Observe** (the prototype is the observation).

**Input**: a prototype directory or repo path, optionally a sitemap/spec file and the branch + routes to focus on. Large prototypes: ask which journey/routes to scope to — never sweep everything.

## Config

Read `.cairn.yaml`: canvas target under `bindings.specs` (fallback `design/canvas/`); macro hills via `bindings.hills` (read-only) for the ladder guess.

## Steps

1. **Map the scoped journey**: routes, entry points, exits; what the user can click, in what order. Emit as a mermaid `flowchart LR`.
2. **Reconstruct the needs**: for each screen/action, write the needs statement it presupposes (recast rule). Elements with no plausible user task → flag, don't rationalize.
3. **Draft the micro hill**: infer Who from the UI's addressee, What from the journey, Wow honestly — usually `[TBD — Wow was never stated]`. Guess `ladders_to` from the macro register; mark the guess (A).
4. **Harvest embedded decisions**: visible choices (layout, flow order, data shown, wording) become decision rows with `status: proposed`, driver guessed, owner `[TBD]`, source = route/file reference.
5. **Assemble the draft canvas** per the `/cairn-canvas` grammar, with a banner at the top:

   > **⚠ Reverse-engineered from {repo}@{branch} on {date}. Every unlabelled claim is (A) Assumed. Validate at the next playback before building on this.**

   Every journey cell and decision defaults to **A**; only externally evidenced facts (a linked spec, a commit message stating intent) may carry K.

## Gate

1. Write the artefact to the working tree. **Stop there — no git writes.** Never run `git commit`, `git push`, or open a PR/MR: the gate is a human reading the markdown locally, and publishing is their call, not yours.
2. Summary: routes covered · needs reconstructed · proposed decisions · the three biggest unknowns.
3. **STOP.** Next: `/cairn-hill` to pressure-test the reconstructed hill, then a validation playback.

## Publishing (only on explicit request)

`gating.mode` in `.cairn.yaml` describes how this repo expects artefacts to be shared. It never authorises you to act. When the human has read the artefact and asks you to publish it:

- `pr` — branch, commit, push, open the PR/MR as a draft.
- `commit` — commit on the current branch, no push.
- `none` — nothing to do.

Until they ask, the artefact stays an uncommitted file in the working tree.

## Guardrails

- Never present a reconstruction as fact — the banner and A-labels are mandatory.
- Never write the macro hill register.
- Scope before sweeping; a 100-route prototype is many canvases, not one.
