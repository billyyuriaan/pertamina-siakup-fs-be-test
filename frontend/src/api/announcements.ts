import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { apiGet, apiPost } from "./client";
import type {
  Announcement,
  AnnouncementFilters,
  ApiResponse,
  CreateAnnouncementInput,
} from "../types";

// TODO(candidate): build the query string from `filters` (search, targetRole,
// page, size — all optional, omit params that are empty/undefined) and call
// apiGet<ApiResponse<Announcement[]>>(`/announcements?${queryString}`).
//
// Then use it inside a React Query `useQuery` hook below. The query key MUST
// include `filters` so the list refetches whenever search/targetRole/page
// change.
export function useAnnouncements(filters: AnnouncementFilters) {
  return useQuery<ApiResponse<Announcement[]>>({
    queryKey: ["announcements", filters],
    queryFn: async () => {
      const data = await apiGet<ApiResponse<Announcement[]>>(`/announcements?${new URLSearchParams(filters as Record<string, string>).toString()}`);
      return data;
    },
  });
}

// TODO(candidate): call apiPost<ApiResponse<Announcement>>("/announcements", input)
// inside the mutationFn below, and invalidate the ["announcements"] query key
// on success so the list refetches and shows the newly created announcement.
export function useCreateAnnouncement() {
  const queryClient = useQueryClient();

  return useMutation<ApiResponse<Announcement>, Error, CreateAnnouncementInput>({
    mutationFn: async (_input) => {
      return apiPost<ApiResponse<Announcement>>("/announcements", _input);
    },
    onSuccess: () => {
      void queryClient;
      // TODO(candidate): invalidate the announcements list query here.
      queryClient.invalidateQueries(["announcements"]);
      return Promise.resolve({ success: true });
    },
  });
}
