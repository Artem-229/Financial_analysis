import { useState, type FormEvent } from 'react'
import { Card } from '../../components/Card/Card'
import { CategoryChip } from '../../components/CategoryChip/CategoryChip'
import { Input } from '../../components/Input/Input'
import { Button } from '../../components/Button/Button'
import { CATEGORIES } from '../../data/categories'
import { ApiError } from '../../lib/api'
import type { NewTransactionInput } from '../../hooks/useTransactions'
import styles from './AddExpenseForm.module.css'

const today = new Date().toISOString().slice(0, 10)

interface AddExpenseFormProps {
  onSubmit: (input: NewTransactionInput) => Promise<void>
}

export function AddExpenseForm({ onSubmit }: AddExpenseFormProps) {
  const [selected, setSelected] = useState<string | null>(null)
  const [custom, setCustom] = useState(false)
  const [customCategory, setCustomCategory] = useState('')
  const [amount, setAmount] = useState('')
  const [date, setDate] = useState(today)
  const [note, setNote] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [submitting, setSubmitting] = useState(false)

  const category = custom ? customCategory.trim() : selected

  async function handleSubmit(e: FormEvent) {
    e.preventDefault()

    const price = Number(amount)
    if (!note.trim() || !category || !price || price <= 0 || !date) {
      setError('Заполните сумму, заметку, дату и категорию')
      return
    }

    setError(null)
    setSubmitting(true)
    try {
      await onSubmit({ item: note.trim(), price, class: category, date })
      setAmount('')
      setNote('')
      setSelected(null)
      setCustom(false)
      setCustomCategory('')
      setDate(today)
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'Не удалось сохранить трату')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <Card>
      <h3 className={styles.title}>Новая трата</h3>
      <p className={styles.subtitle}>Сумма, заметка и категория — остальное посчитаем сами.</p>

      <form onSubmit={handleSubmit}>
        <div className={styles.row}>
          <Input
            label="Сумма"
            placeholder="0 ₽"
            inputMode="decimal"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
          />
          <Input label="Дата" type="date" value={date} onChange={(e) => setDate(e.target.value)} />
        </div>
        <div className={styles.field}>
          <Input
            label="Заметка"
            placeholder="Например: обед в кафе"
            value={note}
            onChange={(e) => setNote(e.target.value)}
          />
        </div>

        <div className={styles.categoryBlock}>
          <span className={styles.categoryLabel}>Категория</span>
          <div className={styles.chips}>
            {CATEGORIES.map((cat) => (
              <CategoryChip
                key={cat.id}
                label={cat.label}
                color={cat.color}
                active={selected === cat.id}
                onClick={() => {
                  setSelected(cat.id)
                  setCustom(false)
                }}
              />
            ))}
            <CategoryChip
              label="Свой вариант"
              color="var(--text-faint)"
              active={custom}
              onClick={() => {
                setCustom(true)
                setSelected(null)
              }}
            />
          </div>
          {custom && (
            <div className={styles.customInput}>
              <Input
                placeholder="Введите свою категорию"
                value={customCategory}
                onChange={(e) => setCustomCategory(e.target.value)}
              />
            </div>
          )}
        </div>

        {error && <p className={styles.error}>{error}</p>}

        <Button type="submit" fullWidth disabled={submitting} className={styles.submit}>
          {submitting ? 'Сохраняем…' : 'Добавить трату'}
        </Button>
      </form>
    </Card>
  )
}
