import { useState } from 'react'
import { Card } from '../../components/Card/Card'
import { CategoryChip } from '../../components/CategoryChip/CategoryChip'
import { Input } from '../../components/Input/Input'
import { Button } from '../../components/Button/Button'
import { TransactionList } from '../../components/TransactionList/TransactionList'
import { CATEGORIES } from '../../data/categories'
import { useTransactions } from '../../hooks/useTransactions'
import styles from './History.module.css'

export function HistoryPage() {
  const [activeCategory, setActiveCategory] = useState<string | null>(null)
  const [search, setSearch] = useState('')
  const { transactions, removeTransaction } = useTransactions()

  const items = transactions.filter((tx) => {
    const matchesCategory = !activeCategory || tx.categoryId === activeCategory
    const matchesSearch = tx.note.toLowerCase().includes(search.toLowerCase())
    return matchesCategory && matchesSearch
  })

  return (
    <div className={styles.page}>
      <header className={styles.header}>
        <h1>История трат</h1>
        <p>Все операции, которые вы занесли вручную или которые распознала система.</p>
      </header>

      <Card>
        <div className={styles.filters}>
          <Input
            placeholder="Поиск по заметке"
            wrapperClassName={styles.search}
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
          <div className={styles.chips}>
            <CategoryChip
              label="Все"
              color="var(--text-faint)"
              active={activeCategory === null}
              onClick={() => setActiveCategory(null)}
            />
            {CATEGORIES.map((cat) => (
              <CategoryChip
                key={cat.id}
                label={cat.label}
                color={cat.color}
                active={activeCategory === cat.id}
                onClick={() => setActiveCategory(cat.id)}
              />
            ))}
          </div>
        </div>

        <div className={styles.tableScroll}>
          <TransactionList items={items} onDelete={removeTransaction} />
        </div>

        <div className={styles.pagination}>
          <Button variant="secondary" size="sm" disabled>
            ← Назад
          </Button>
          <span className={styles.pageInfo}>Показаны все {items.length} операций</span>
          <Button variant="secondary" size="sm" disabled>
            Вперёд →
          </Button>
        </div>
      </Card>
    </div>
  )
}
