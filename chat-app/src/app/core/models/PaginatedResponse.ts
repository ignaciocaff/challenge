export type PaginationResult<T> = {
  data: T[];
  total: number;
  pageNumber: number;
  totalPages: number;
  pageSize: number;
};
