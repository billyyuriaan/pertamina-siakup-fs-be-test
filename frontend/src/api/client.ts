// Already implemented for you — you should not need to change this file.
// The backend runs on :8080 (see backend/main.go); run it separately with
// `go run main.go` from the backend/ directory before starting the frontend.

const BASE_URL = "http://localhost:8080";

export async function apiGet<T>(path: string): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`);
  if (!res.ok) {
    throw new Error(`GET ${path} failed with status ${res.status}`);
  }
  return res.json() as Promise<T>;
}

export async function apiPost<T>(path: string, body: unknown): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  const json = (await res.json()) as T;
  if (!res.ok) {
    const message =
      (json as { message?: string })?.message ?? `POST ${path} failed with status ${res.status}`;
    throw new Error(message);
  }
  return json;
}
