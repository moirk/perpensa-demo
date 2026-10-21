# perpensa-demo

**Intentionally vulnerable demo project — do not use.**

This repository exists only to demonstrate dependency scanning with [Perpensa](https://stackpulse.mserdar.net). It pins known-vulnerable versions on purpose:

| Ecosystem | Package | Pinned | Advisory | Fixed in |
|---|---|---|---|---|
| Go | github.com/containers/podman/v5 | v5.8.7 | CVE-2026-94603 / GHSA-2cvf-wqm6-wr9g | 5.8.8 |
| npm | @tanstack/start-server-core | 1.169.38 | CVE-2026-102989 / GHSA-qx66-fv34-fjm8 | 1.169.39 |

Do not deploy or reuse this code.
