import { useState, useEffect } from "react";
import { useAnnouncements } from "../api/announcements";
import type { TargetRole } from "../types";

// TODO(candidate): implement this component.
//
// Requirements:
//   - A text input bound to `search`, debounced or not (your choice), that
//     filters announcements by title (case-insensitive substring match — the
//     backend already does the matching, you just need to send the param).
//   - A <select> bound to `targetRole` with options: "" (all), "student",
//     "lecturer", "all".
//   - Call useAnnouncements({ search, targetRole }) and render:
//       - a loading state while the query is pending
//       - an error state if the query failed
//       - the list of announcements (title, content, target_role) once loaded
//       - an empty state ("No announcements found") when the list is empty
//   - Show the total count from the response's `meta.total` somewhere in the UI.
export function AnnouncementList() {
  const [search, setSearch] = useState("");
  const [targetRole, setTargetRole] = useState<TargetRole | "">("");
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);
  const [data, setData] = useState([]);

  // Wire this up once useAnnouncements is implemented.
  void search;
  void targetRole;
  void page;
  void size;

  useEffect(() => {
    setData(useAnnouncements({ search, targetRole, page, size }));

    return () => {
      setData([]);
    };
  }, [search, targetRole, page, size]);

  return <div>
    <input type="text" value={search} onChange={(e) => setSearch(e.target.value)} placeholder="Search by title" />
    <select value={targetRole} onChange={(e) => setTargetRole(e.target.value as TargetRole | "")}>
      <option value="">All</option>
      <option value="student">Student</option>
      <option value="lecturer">Lecturer</option>
      <option value="all">All</option>
    </select>
    <table>
      <thead>
        <tr>
          <th>Title</th>
          <th>Content</th>
          <th>Target Role</th>
        </tr>
      </thead>
      <tbody>
        {data.map((announcement) => (
          <tr key={announcement.id}>
            <td>{announcement.title}</td>
            <td>{announcement.content}</td>
            <td>{announcement.target_role}</td>
          </tr>
        ))}
      </tbody>
      <tfoot>
        <tr>
          <td colSpan={3}>Total: {data.length}</td>
          <td colSpan={3}>
            <button onClick={() => setPage((prev) => Math.max(prev - 1, 1))} disabled={page === 1}>
              Previous
            </button>
            <span>Page {page}</span>
            <button onClick={() => setPage((prev) => prev + 1)} disabled={data.length < size}>
              Next
            </button>
          </td>
          <tr>
            <span>Page size : {size}</span>
            <select value={size} onChange={(e) => setSize(Number(e.target.value))}>
              <option value="5">5</option>
              <option value="10">10</option>
              <option value="20">20</option>
            </select>
          </tr>
        </tr>
      </tfoot>
    </table>
  </div>;
}
