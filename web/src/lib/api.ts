export class ApiError extends Error {}

interface RequestOptions {
  method?: string
  body?: unknown
  token?: string
}

async function request<T>(path: string, options: RequestOptions = {}): Promise<T> {
  const headers: Record<string, string> = { 'Content-Type': 'application/json' }
  if (options.token) headers.Authorization = `Bearer ${options.token}`

  const res = await fetch(path, {
    method: options.method ?? 'GET',
    headers,
    body: options.body !== undefined ? JSON.stringify(options.body) : undefined,
  })

  const data = await res.json().catch(() => ({}))

  if (!res.ok) {
    throw new ApiError(typeof data?.error === 'string' ? data.error : 'Что-то пошло не так')
  }

  return data as T
}

interface LoginResponse {
  status: string
  access_token: string
}

interface RegistrationResponse {
  status: string
}

export function login(loginValue: string, password: string) {
  return request<LoginResponse>('/login', { method: 'POST', body: { login: loginValue, password } })
}

export function register(username: string, loginValue: string, password: string) {
  return request<RegistrationResponse>('/registration', {
    method: 'POST',
    body: { username, login: loginValue, password },
  })
}

export interface TransactionPayload {
  item: string
  price: number
  class: string
}

export interface TransactionDTO {
  id: string
  item: string
  price: number
  class: string
}

export function listTransactions(token: string) {
  return request<TransactionDTO[]>('/app/transactions', { token })
}

export function createTransaction(token: string, payload: TransactionPayload) {
  return request<TransactionDTO>('/app/transactions', { method: 'POST', body: payload, token })
}

export function deleteTransaction(token: string, id: string) {
  return request<{ message: string }>(`/app/transactions/${id}`, { method: 'DELETE', token })
}
