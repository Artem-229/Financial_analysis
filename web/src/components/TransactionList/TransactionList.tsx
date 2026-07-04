import type { Transaction } from '../../types/transaction'
import { CATEGORIES } from '../../data/categories'
import { formatCurrency, formatDate } from '../../lib/format'
import styles from './TransactionList.module.css'

interface TransactionListProps {
  items: Transaction[]
  onDelete?: (id: string) => void
}

export function TransactionList({ items, onDelete }: TransactionListProps) {
  if (items.length === 0) {
    return <p className={styles.empty}>Ничего не найдено.</p>
  }

  return (
    <div className={styles.list}>
      <div className={`${styles.row} ${onDelete ? styles.withActions : ''} ${styles.head}`}>
        <span>Дата</span>
        <span>Заметка</span>
        <span>Категория</span>
        <span className={styles.amountCol}>Сумма</span>
        {onDelete && <span />}
      </div>
      {items.map((tx) => {
        const category = CATEGORIES.find((c) => c.id === tx.categoryId)
        return (
          <div className={`${styles.row} ${onDelete ? styles.withActions : ''}`} key={tx.id}>
            <span className={styles.date}>{formatDate(tx.date)}</span>
            <span className={styles.note}>{tx.note}</span>
            <span className={styles.category}>
              <span className={styles.dot} style={{ background: category?.color }} />
              {category?.label}
            </span>
            <span className={styles.amount}>{formatCurrency(tx.amount)}</span>
            {onDelete && (
              <button
                type="button"
                className={styles.deleteButton}
                onClick={() => onDelete(tx.id)}
                aria-label="Удалить транзакцию"
                title="Удалить"
              >
                ✕
              </button>
            )}
          </div>
        )
      })}
    </div>
  )
}
