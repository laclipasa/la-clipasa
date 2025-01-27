import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI la-clipasa
 * la-clipasa
 * OpenAPI spec version: 2.0.0
 */
import type { CacheDemoWorkItemResponse } from './cacheDemoWorkItemResponse';
import type { PaginationPage } from './paginationPage';

export interface PaginatedDemoWorkItemsResponse {
  /** @nullable */
  items: CacheDemoWorkItemResponse[] | null;
  page: PaginationPage;
}
