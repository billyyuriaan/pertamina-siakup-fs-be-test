# Base Project — Announcement Feature (Fullstack)

Two folders: `backend/` (Go + Fiber) and `frontend/` (React + TypeScript +
Vite + TanStack Query). No database, Docker, or credentials are required —
the backend stores data in memory.

## Setup

**Backend** (in `backend/`):
```bash
go mod tidy      # downloads Fiber; requires internet access
go run main.go   # starts the API on :8080
```

**Frontend** (in `frontend/`, in a separate terminal):
```bash
npm install
npm run dev      # starts the app on :5173
```

Start the backend first — the frontend's dev server proxies nothing; it
calls `http://localhost:8080` directly (see `frontend/src/api/client.ts`).

## What to build

### Backend — `backend/module/announcement/usecase/usecase.go`

Implement `GetList` and `Create` per the contract documented on
`domain.AnnouncementUseCase` (in `domain/usecase.go`). The repository and
HTTP handler are already fully implemented and wired — once your usecase
compiles and behaves correctly, both endpoints work as-is:

- `GET /announcements?search=&target_role=&page=&size=`
- `POST /announcements` — body `{"title", "content", "target_role"}`

Validation (empty/too-long title, invalid `target_role`) belongs in the
usecase, not the handler or repository.

### Frontend — three files

- `frontend/src/api/announcements.ts` — the `useAnnouncements` query hook and
  `useCreateAnnouncement` mutation hook. Wire them to `apiGet`/`apiPost`
  (already implemented in `api/client.ts`).
- `frontend/src/components/AnnouncementList.tsx` — search + role filter,
  loading/error/empty states, and the rendered list.
- `frontend/src/components/AnnouncementForm.tsx` — the create form, with
  client-side validation before submitting.

`App.tsx` and `main.tsx` are already wired — you should not need to touch
them.

## What we're evaluating

- Correctness end-to-end: creating an announcement in the form actually shows
  up in the list without a manual page reload.
- Backend: validation lives in the usecase layer; sensible error handling and
  HTTP status codes.
- Frontend: loading/error/empty states are handled, not just the happy path;
  the query key includes the filters so search/filter actually refetch;
  the mutation invalidates the list on success.
- Idiomatic code on both sides — no over-engineering for this scope.
- Bonus (not required): a unit test for the backend usecase, or a component
  test for one of the two frontend components.

## Submitting

Push to a private repo and share access, or zip the project (excluding
`node_modules/`, `dist/`, and `go.sum`) and send it back within the 2-hour
window.
