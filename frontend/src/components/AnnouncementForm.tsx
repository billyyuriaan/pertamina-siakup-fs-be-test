import { useState, useEffect, React } from "react";
import { useCreateAnnouncement } from "../api/announcements";
import type { TargetRole } from "../types";

// TODO(candidate): implement this component.
//
// Requirements:
//   - Fields: title (text), content (textarea), target_role (select: student
//     / lecturer / all).
//   - Client-side validation before submitting: title is required and must be
//     200 characters or fewer. Show the validation error inline — do not rely
//     on the backend error alone.
//   - On submit, call useCreateAnnouncement()'s mutate function.
//   - While the mutation is pending, disable the submit button.
//   - On success, clear the form. On error (e.g. backend validation failure),
//     show the error message returned by the API.
export function AnnouncementForm() {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [targetRole, setTargetRole] = useState<TargetRole>("all");
  const handleSubmit = (e: React.FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    useCreateAnnouncement().mutate({
      title,
      content,
      target_role: targetRole,
    });

    setTitle("");
    setContent("");
    setTargetRole("all");
  }

  void title;
  void content;
  void targetRole;

  useEffect(() => {
    
  }, []);

  return <div>
    <form>
      <textarea value={content} onChange={(e) => setContent(e.target.value)} />
      <select value={targetRole} onChange={(e) => setTargetRole(e.target.value as TargetRole)}>
        <option value="student">Student</option>
        <option value="lecturer">Lecturer</option>
        <option value="all">All</option>
      </select>
      <button type="submit" onClick={handleSubmit}>
        Submit
      </button>
    </form>
  </div>;
    </form>
  </div>;
}
