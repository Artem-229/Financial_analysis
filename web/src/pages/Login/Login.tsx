import { useState, type FormEvent } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { AuthLayout } from '../../components/AuthLayout/AuthLayout'
import { Input } from '../../components/Input/Input'
import { Button } from '../../components/Button/Button'
import { useAuth } from '../../auth/AuthContext'
import { login as loginRequest, ApiError } from '../../lib/api'
import styles from './Login.module.css'

export function LoginPage() {
  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const { login: setAuth } = useAuth()
  const navigate = useNavigate()

  async function handleSubmit(e: FormEvent) {
    e.preventDefault()
    setError(null)
    setLoading(true)
    try {
      const res = await loginRequest(login, password)
      setAuth(res.access_token, login)
      navigate('/dashboard')
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'Не удалось войти. Попробуйте ещё раз.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <AuthLayout title="С возвращением" subtitle="Войдите, чтобы увидеть свои траты.">
      <form className={styles.form} onSubmit={handleSubmit}>
        <Input
          label="Логин"
          name="login"
          autoComplete="username"
          value={login}
          onChange={(e) => setLogin(e.target.value)}
          placeholder="Ваш логин"
          required
        />
        <Input
          label="Пароль"
          name="password"
          type="password"
          autoComplete="current-password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="••••••••"
          required
        />
        {error && <div className={styles.error}>{error}</div>}
        <Button type="submit" fullWidth disabled={loading}>
          {loading ? 'Входим…' : 'Войти'}
        </Button>
      </form>
      <p className={styles.footer}>
        Нет аккаунта?{' '}
        <Link to="/register" className={styles.link}>
          Зарегистрироваться
        </Link>
      </p>
    </AuthLayout>
  )
}
