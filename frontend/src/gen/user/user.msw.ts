import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI la-clipasa
 * la-clipasa
 * OpenAPI spec version: 2.0.0
 */
import {
  faker
} from '@faker-js/faker'
import {
  HttpResponse,
  delay,
  http
} from 'msw'
import {
  ProjectName,
  Role,
  Scope
} from '.././model'
import type {
  PaginatedUsersResponse,
  UserResponse
} from '.././model'

export const getGetPaginatedUsersResponseMock = (overrideResponse: Partial< PaginatedUsersResponse > = {}): PaginatedUsersResponse => ({items: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), apiKey: {apiKey: faker.word.sample(), expiresOn: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID}, createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), projects: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({boardConfig: {fields: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({isEditable: faker.datatype.boolean(), isVisible: faker.datatype.boolean(), name: faker.word.sample(), path: faker.word.sample(), showCollapsed: faker.datatype.boolean()})), header: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => (faker.word.sample())), visualization: {}}, createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.helpers.arrayElement(Object.values(ProjectName)), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, updatedAt: (() => faker.date.past())()})), role: faker.helpers.arrayElement(Object.values(Role)), scopes: faker.helpers.arrayElements(Object.values(Scope)), teams: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, updatedAt: (() => faker.date.past())()})), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()})), page: {nextCursor: faker.word.sample()}, ...overrideResponse})

export const getGetCurrentUserResponseMock = (overrideResponse: Partial< UserResponse > = {}): UserResponse => ({age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), apiKey: {apiKey: faker.word.sample(), expiresOn: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID}, createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: (() => faker.internet.email())(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), projects: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({boardConfig: {fields: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({isEditable: faker.datatype.boolean(), isVisible: faker.datatype.boolean(), name: faker.word.sample(), path: faker.word.sample(), showCollapsed: faker.datatype.boolean()})), header: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => (faker.word.sample())), visualization: {}}, createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.helpers.arrayElement(Object.values(ProjectName)), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, updatedAt: (() => faker.date.past())()})), role: faker.helpers.arrayElement(Object.values(Role)), scopes: faker.helpers.arrayElements(Object.values(Scope)), teams: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, updatedAt: (() => faker.date.past())()})), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample(), ...overrideResponse})

export const getUpdateUserResponseMock = (overrideResponse: Partial< UserResponse > = {}): UserResponse => ({age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), apiKey: {apiKey: faker.word.sample(), expiresOn: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID}, createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: (() => faker.internet.email())(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), projects: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({boardConfig: {fields: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({isEditable: faker.datatype.boolean(), isVisible: faker.datatype.boolean(), name: faker.word.sample(), path: faker.word.sample(), showCollapsed: faker.datatype.boolean()})), header: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => (faker.word.sample())), visualization: {}}, createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.helpers.arrayElement(Object.values(ProjectName)), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, updatedAt: (() => faker.date.past())()})), role: faker.helpers.arrayElement(Object.values(Role)), scopes: faker.helpers.arrayElements(Object.values(Scope)), teams: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, updatedAt: (() => faker.date.past())()})), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample(), ...overrideResponse})


export const getGetPaginatedUsersMockHandler = (overrideResponse?: PaginatedUsersResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<PaginatedUsersResponse> | PaginatedUsersResponse)) => {
  return http.get('*/user/page', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getGetPaginatedUsersResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetCurrentUserMockHandler = (overrideResponse?: UserResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<UserResponse> | UserResponse)) => {
  return http.get('*/user/me', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getGetCurrentUserResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateUserAuthorizationMockHandler = () => {
  return http.patch('*/user/:id/authorization', async () => {await delay(200);
    return new HttpResponse(null,
      {
        status: 204,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteUserMockHandler = () => {
  return http.delete('*/user/:id', async () => {await delay(200);
    return new HttpResponse(null,
      {
        status: 204,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateUserMockHandler = (overrideResponse?: UserResponse | ((info: Parameters<Parameters<typeof http.patch>[1]>[0]) => Promise<UserResponse> | UserResponse)) => {
  return http.patch('*/user/:id', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getUpdateUserResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}
export const getUserMock = () => [
  getGetPaginatedUsersMockHandler(),
  getGetCurrentUserMockHandler(),
  getUpdateUserAuthorizationMockHandler(),
  getDeleteUserMockHandler(),
  getUpdateUserMockHandler()
]