The goal of this project is to create a multi-threaded web crawler, maybe breaking some stuff along the way but just see how that goes

## Web Crawlers


## Process

### URL queue

You start with a small set of seed URLs.
These go into a shared frontier (usually a thread-safe queue or priority queue).

Think of this as:

“URLs we still need to visit”

### Worker Threads (fixed pool)

Instead of spawning infinite threads, you create a fixed-size pool of workers.

Each worker repeatedly:

Takes a URL from the frontier
Fetches the page
Processes it
Extracts new links
Pushes new valid links back into the frontier

So threads are reused, not constantly created/destroyed.

### Fetching (I/O-bound step)

Each worker:

Sends HTTP request
Waits for response

This is the slow part (network latency), which is why multithreading helps.

### Parsing + Link Extraction

Once a page is fetched:

Parse HTML
Extract links
Normalize URLs (avoid duplicates like /page vs /page/)

### Deduplication (critical)

Before adding new links to the frontier:

Check against a visited set (thread-safe)

Without this, your crawler loops forever.

### Politeness + Rate Limiting

Real crawlers don’t hammer servers:

Respect robots.txt
Add delays per domain
Limit concurrent requests per host