# Browser Cache/History Test

Not really an automated test, but a manual one. One day when the stars are aligned, it will be automated.

What stays in browser history buffers and what evaporates varies a lot between browsers. This can help you with some sanity.

## Current Conclusion
Use `Cache-Control: no-store` along with  `window.onunload` and `window.onbeforeunload`. Also, add in HTTPS if you can. 