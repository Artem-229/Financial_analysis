const currencyFormatter = new Intl.NumberFormat('ru-RU', {
  style: 'currency',
  currency: 'RUB',
  maximumFractionDigits: 0,
})

const dateFormatter = new Intl.DateTimeFormat('ru-RU', {
  day: '2-digit',
  month: 'short',
})

export function formatCurrency(amount: number) {
  return currencyFormatter.format(amount)
}

export function formatDate(iso: string) {
  return dateFormatter.format(new Date(iso))
}
