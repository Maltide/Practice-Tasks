# Sprint 03 — Hints (Last Resort)

Use these only after a serious attempt.

1) **Intersection of Two Slices** — Build a set from the smaller slice; emit each common element once.  
2) **Remove Duplicates from Sorted Slice** — Two pointers: advance write on value change.  
3) **Rotate Slice Right by k** — Normalize `k := k % n`; consider triple-reverse or slicing.  
4) **Maximum Sum Subarray (fixed k)** — Sliding window: add new, subtract old.  
5) **Merge Two Sorted Slices** — Two indices; append the smaller; then drain remainder.  
6) **Longest Common Prefix** — Start from the shortest word; shrink prefix until all match.  
7) **First Non-Repeating Character** — Count first, then scan to find the first with count==1.  
8) **Run-Length Encoding** — Track current rune + count; flush on change/end.  
9) **Run-Length Decoding** — Parse multi-digit counts; watch for large repeats and invalid patterns.  
10) **Invert Map with Unique Values** — Values must be unique; swap key/value while iterating.  
11) **Two Sum** — Map value→index; check complement before insert to avoid pairing same index.