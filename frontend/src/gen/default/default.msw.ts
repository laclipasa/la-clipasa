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

export const getPingResponseMock = (): string => (faker.word.sample())

export const getOpenapiYamlGetResponseMock = (): Blob => (new Blob(faker.helpers.arrayElements(faker.word.words(10).split(' '))))


export const getPingMockHandler = () => {
  return http.get('*/ping', async () => {await delay(200);
    return new HttpResponse(getPingResponseMock(),
      {
        status: 200,
        headers: {
          'Content-Type': 'text/plain',
        }
      }
    )
  })
}

export const getOpenapiYamlGetMockHandler = (overrideResponse?: Blob | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<Blob> | Blob)) => {
  return http.get('*/openapi.yaml', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse)
            : getOpenapiYamlGetResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}
export const getDefaultMock = () => [
  getPingMockHandler(),
  getOpenapiYamlGetMockHandler()
]