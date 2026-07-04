import { createContext, useContext, useMemo, useState, type ReactNode } from 'react'

interface AuthState {
  token: string | null
  username: string | null
  login: (token: string, username: string) => void
  logout: () => void
}

const AuthContext = createContext<AuthState | null>(null)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(null)
  const [username, setUsername] = useState<string | null>(null)

  const value = useMemo<AuthState>(
    () => ({
      token,
      username,
      login: (nextToken: string, nextUsername: string) => {
        setToken(nextToken)
        setUsername(nextUsername)
      },
      logout: () => {
        setToken(null)
        setUsername(null)
      },
    }),
    [token, username],
  )

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error('useAuth must be used within AuthProvider')
  return ctx
}
