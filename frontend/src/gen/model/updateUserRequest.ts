import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI la-clipasa
 * la-clipasa
 * OpenAPI spec version: 2.0.0
 */

/**
 * represents User data to update
 */
export interface UpdateUserRequest {
  /**
   * originally from auth server but updatable
   * @pattern ^[a-zA-Z '-]+$
   */
  firstName?: string;
  /**
   * originally from auth server but updatable
   * @pattern ^[a-zA-Z '-]+$
   */
  lastName?: string;
}