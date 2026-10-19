export type TargetRole = "student" | "lecturer" | "all";

export interface Announcement {
  id: string;
  title: string;
  content: string;
  target_role: TargetRole;
  created_at: string;
  updated_at: string;
}

export interface ListMeta {
  total: number;
  total_pages: number;
  page: number;
}

export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  meta?: ListMeta;
}

export interface AnnouncementFilters {
  search?: string;
  targetRole?: TargetRole | "";
  page?: number;
  size?: number;
}

export interface CreateAnnouncementInput {
  title: string;
  content: string;
  target_role: TargetRole;
}
