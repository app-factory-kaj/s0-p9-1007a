# greeter — PRD

## Problem Statement

Teams building and testing services on this platform often need a tiny, dependable HTTP endpoint to use as a smoke-test target, a demo backend, or a reference example — standing one up today means writing boilerplate for something trivial. S0 marker s0-p9-1007a.

## Solution

Greeter is a small Go HTTP service that returns a JSON greeting for a name supplied on the query string, giving teams a minimal, reliable endpoint to call, test against, or use as a reference implementation.

## Actors

- **API Caller** — any client (person or system) that sends HTTP requests to the service; anonymous, with no account or role.

## User Stories

1. As an API Caller, I want to GET /hello?name=X, so that I receive a JSON greeting addressed to X.
2. As an API Caller, I want to GET /hello without a name, so that I still receive a friendly default JSON greeting instead of an error.

## Product Decisions

- Authentication: the service is fully public — no sign-in, no token required to call it. *assumed*
- Missing/empty `name`: the service falls back to a generic default (e.g. "World") rather than returning an error. *assumed*
- Scope: GET /hello is the entire service — no additional endpoints (e.g. health checks) are included. *assumed*

## Out of Scope

- Authentication, authorization, or per-caller rate limiting.
- Persistence of any kind (no database, no request history).
- A health-check or readiness endpoint.
- Multiple languages/localized greetings.
- A user interface — this is an API-only service.

## Open Questions

None.

## Further Notes

None.