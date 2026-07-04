import { useState, type FormEvent } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { AuthLayout } from '../../components/AuthLayout/AuthLayout'
import { Input } from '../../components/Input/Input'
import { Button } from '../../components/Button/Button'
import { useAuth } from '../../auth/AuthContext'
import { register as registerRequest, login as loginRequest, ApiError } from '../../lib/api'
import styles from './Register.module.css'

export function RegisterPage() {
  const [username, setUsername] = useState('')
  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')
  const [confirm, setConfirm] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const { login: setAuth } = useAuth()
  const navigate = useNavigate()

  async function handleSubmit(e: FormEvent) {
    e.preventDefault()
    setError(null)

    if (password !== confirm) {
      setError('Пароли не совпадают')
      return
    }
    if (password.length < 6) {
      setError('Пароль должен быть не короче 6 символов')
      return
    }

    setLoading(true)
    try {
      await registerRequest(username, login, password)
      const res = await loginRequest(login, password)
      setAuth(res.access_token, login)
      navigate('/dashboard')
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'Не удалось зарегистрироваться. Попробуйте ещё раз.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <AuthLayout title="Регистрация" subtitle="Пара минут — и вы начнёте видеть свои траты по полочкам.">
      <form className={styles.form} onSubmit={handleSubmit}>
        <Input
          label="Имя"
          name="username"
          autoComplete="name"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Как вас называть"
          required
        />
        <Input
          label="Логин"
          name="login"
          autoComplete="username"
          value={login}
          onChange={(e) => setLogin(e.target.value)}
          placeholder="Придумайте логин"
          required
        />
        <Input
          label="Пароль"
          name="password"
          type="password"
          autoComplete="new-password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Не короче 6 символов"
          required
        />
        <Input
          label="Повторите пароль"
          name="confirm"
          type="password"
          autoComplete="new-password"
          value={confirm}
          onChange={(e) => setConfirm(e.target.value)}
          placeholder="••••••••"
          required
        />
        {error && <div className={styles.error}>{error}</div>}
        <Button type="submit" fullWidth disabled={loading}>
          {loading ? 'Создаём аккаунт…' : 'Зарегистрироваться'}
        </Button>
      </form>
      <p className={styles.footer}>
        Уже есть аккаунт?{' '}
        <Link to="/login" className={styles.link}>
          Войти
        </Link>
      </p>
    </AuthLayout>
  )
}
