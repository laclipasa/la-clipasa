import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI la-clipasa
 * la-clipasa
 * OpenAPI spec version: 2.0.0
 */

export interface CreateWorkItemTypeRequest {
  /** @pattern ^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$ */
  color: string;
  description: string;
  name: string;
}