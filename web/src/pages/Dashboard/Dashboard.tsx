import { Link } from 'react-router-dom'
import { useAuth } from '../../auth/AuthContext'
import { StatCard } from '../../components/StatCard/StatCard'
import { Card } from '../../components/Card/Card'
import { TransactionList } from '../../components/TransactionList/TransactionList'
import { useTransactions } from '../../hooks/useTransactions'
import { AddExpenseForm } from './AddExpenseForm'
import styles from './Dashboard.module.css'

export function DashboardPage() {
  const { username } = useAuth()
  const { transactions, addTransaction, removeTransaction } = useTransactions()
  const recent = transactions.slice(0, 5)

  return (
    <div className={styles.page}>
      <header className={styles.header}>
        <h1>Здравствуйте, {username}</h1>
        <p>Вот как выглядят ваши траты в этом месяце.</p>
      </header>

      <div className={styles.stats}>
        <StatCard label="Траты в этом месяце" value="48 260 ₽" hint="↑ 12% к прошлому месяцу" trend="up" />
        <StatCard label="Топ категория" value="Продукты" hint="36% от всех трат" />
        <StatCard label="Транзакций" value="23" hint="за последние 30 дней" />
        <StatCard label="Аномалии" value="1" hint="трата в 3 раза выше обычной" trend="up" />
      </div>

      <div className={styles.columns}>
        <AddExpenseForm onSubmit={addTransaction} />
        <Card className={styles.recentCard}>
          <div className={styles.recentHeader}>
            <h3>Последние траты</h3>
            <Link to="/history" className={styles.allLink}>
              Смотреть всё →
            </Link>
          </div>
          <div className={styles.tableScroll}>
            <TransactionList items={recent} onDelete={removeTransaction} />
          </div>
        </Card>
      </div>
    </div>
  )
}
