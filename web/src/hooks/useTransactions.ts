import { useCallback, useEffect, useState } from 'react'
import { useAuth } from '../auth/AuthContext'
import { createTransaction, deleteTransaction, listTransactions, type TransactionDTO } from '../lib/api'
import { getTransactionDate, removeTransactionDate, setTransactionDate } from '../lib/txDates'
import type { Transaction } from '../types/transaction'

function toTransaction(dto: TransactionDTO): Transaction {
  return {
    id: dto.id,
    note: dto.item,
    categoryId: dto.class,
    amount: dto.price,
    date: getTransactionDate(dto.id) ?? new Date().toISOString().slice(0, 10),
  }
}

export interface NewTransactionInput {
  item: string
  price: number
  class: string
  date: string
}

export function useTransactions() {
  const { token } = useAuth()
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    if (!token) return
    setLoading(true)
    setError(null)
    try {
      const items = await listTransactions(token)
      setTransactions(items.map(toTransaction).sort((a, b) => b.date.localeCompare(a.date)))
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Не удалось загрузить транзакции')
    } finally {
      setLoading(false)
    }
  }, [token])

  useEffect(() => {
    refresh()
  }, [refresh])

  const addTransaction = useCallback(
    async (input: NewTransactionInput) => {
      if (!token) return
      const created = await createTransaction(token, { item: input.item, price: input.price, class: input.class })
      setTransactionDate(created.id, input.date)
      setTransactions((prev) => [toTransaction(created), ...prev])
    },
    [token],
  )

  const removeTransaction = useCallback(
    async (id: string) => {
      if (!token) return
      await deleteTransaction(token, id)
      removeTransactionDate(id)
      setTransactions((prev) => prev.filter((tx) => tx.id !== id))
    },
    [token],
  )

  return { transactions, loading, error, addTransaction, removeTransaction }
}
