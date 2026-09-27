# Review PR — additional instructions

Appended to the default review-pr prompt. `/review review this PR carefully
and leave comment` is the core instruction; add repository conventions here.

- Flag missing tests for new code paths as `[Warning]`.
- Check that error messages are lowercase and wrapped with `%w`.
- Prefer `errors.Is` over direct error comparison.
- Reviewer tone: precise, kind, and concrete. Suggest the fix, not just the
  problem.