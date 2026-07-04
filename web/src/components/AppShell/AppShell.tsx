import { NavLink, Outlet } from 'react-router-dom'
import { Logo } from '../Logo/Logo'
import { useAuth } from '../../auth/AuthContext'
import { cn } from '../../lib/cn'
import styles from './AppShell.module.css'

const NAV = [
  { to: '/dashboard', label: 'Дашборд' },
  { to: '/history', label: 'История' },
]

export function AppShell() {
  const { username, logout } = useAuth()

  return (
    <div className={styles.shell}>
      <aside className={styles.sidebar}>
        <Logo size="sm" />
        <nav className={styles.nav}>
          {NAV.map((item) => (
            <NavLink
              key={item.to}
              to={item.to}
              className={({ isActive }) => cn(styles.navItem, isActive && styles.navItemActive)}
            >
              {item.label}
            </NavLink>
          ))}
        </nav>
        <div className={styles.sidebarFooter}>
          <div className={styles.userChip}>
            <span className={styles.avatar}>{(username ?? '?').slice(0, 1).toUpperCase()}</span>
            <span className={styles.userName}>{username}</span>
          </div>
          <button className={styles.logout} onClick={logout}>
            Выйти
          </button>
        </div>
      </aside>

      <div className={styles.mobileBar}>
        <Logo size="sm" />
        <nav className={styles.mobileNav}>
          {NAV.map((item) => (
            <NavLink
              key={item.to}
              to={item.to}
              className={({ isActive }) => cn(isActive && styles.mobileNavActive)}
            >
              {item.label}
            </NavLink>
          ))}
        </nav>
        <button className={styles.logout} onClick={logout}>
          Выйти
        </button>
      </div>

      <div className={styles.content}>
        <Outlet />
      </div>
    </div>
  )
}
