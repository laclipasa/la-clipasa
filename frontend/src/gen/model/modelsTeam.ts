import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI la-clipasa
 * la-clipasa
 * OpenAPI spec version: 2.0.0
 */

export interface ModelsTeam {
  createdAt: Date;
  description: string;
  name: string;
  projectID: EntityIDs.ProjectID;
  teamID: EntityIDs.TeamID;
  updatedAt: Date;
}