import type { ReactNode } from 'react'
import { Logo } from '../Logo/Logo'
import styles from './AuthLayout.module.css'

interface AuthLayoutProps {
  title: string
  subtitle: string
  children: ReactNode
}

export function AuthLayout({ title, subtitle, children }: AuthLayoutProps) {
  return (
    <div className={styles.wrap}>
      <aside className={styles.aside}>
        <Logo />
        <div className={styles.pitch}>
          <h1>Тратьте осознанно</h1>
          <p>
            Заносите траты, а «Кошелёк» сам разложит их по категориям, посчитает суммы и раз в неделю
            пришлёт понятный разбор.
          </p>
        </div>
        <ul className={styles.points}>
          <li>Без ручных таблиц и подсчётов в уме</li>
          <li>Категории — по кнопке или свободным текстом</li>
          <li>Отчёты в личном кабинете и в Telegram</li>
        </ul>
      </aside>
      <main className={styles.main}>
        <div className={styles.formCol}>
          <div className={styles.mobileLogo}>
            <Logo size="sm" />
          </div>
          <h2 className={styles.title}>{title}</h2>
          <p className={styles.subtitle}>{subtitle}</p>
          {children}
        </div>
      </main>
    </div>
  )
}
