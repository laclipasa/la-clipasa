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
import type {
  WorkItemTypeResponse
} from '.././model'

export const getCreateWorkItemTypeResponseMock = (overrideResponse: Partial< WorkItemTypeResponse > = {}): WorkItemTypeResponse => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID, ...overrideResponse})

export const getGetWorkItemTypeResponseMock = (overrideResponse: Partial< WorkItemTypeResponse > = {}): WorkItemTypeResponse => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID, ...overrideResponse})

export const getUpdateWorkItemTypeResponseMock = (overrideResponse: Partial< WorkItemTypeResponse > = {}): WorkItemTypeResponse => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID, ...overrideResponse})


export const getCreateWorkItemTypeMockHandler = (overrideResponse?: WorkItemTypeResponse | ((info: Parameters<Parameters<typeof http.post>[1]>[0]) => Promise<WorkItemTypeResponse> | WorkItemTypeResponse)) => {
  return http.post('*/project/:projectName/work-item-type/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getCreateWorkItemTypeResponseMock()),
      {
        status: 201,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetWorkItemTypeMockHandler = (overrideResponse?: WorkItemTypeResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<WorkItemTypeResponse> | WorkItemTypeResponse)) => {
  return http.get('*/work-item-type/:workItemTypeID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getGetWorkItemTypeResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateWorkItemTypeMockHandler = (overrideResponse?: WorkItemTypeResponse | ((info: Parameters<Parameters<typeof http.patch>[1]>[0]) => Promise<WorkItemTypeResponse> | WorkItemTypeResponse)) => {
  return http.patch('*/work-item-type/:workItemTypeID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getUpdateWorkItemTypeResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteWorkItemTypeMockHandler = () => {
  return http.delete('*/work-item-type/:workItemTypeID', async () => {await delay(200);
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
export const getWorkItemTypeMock = () => [
  getCreateWorkItemTypeMockHandler(),
  getGetWorkItemTypeMockHandler(),
  getUpdateWorkItemTypeMockHandler(),
  getDeleteWorkItemTypeMockHandler()
]