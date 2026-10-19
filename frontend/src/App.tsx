import { AnnouncementForm } from "./components/AnnouncementForm";
import { AnnouncementList } from "./components/AnnouncementList";

// Already wired for you — you should not need to change this file.
export function App() {
  return (
    <main style={{ maxWidth: 640, margin: "2rem auto", fontFamily: "sans-serif" }}>
      <h1>Announcements</h1>
      <AnnouncementForm />
      <hr style={{ margin: "1.5rem 0" }} />
      <AnnouncementList />
    </main>
  );
}
