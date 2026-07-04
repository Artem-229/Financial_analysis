export interface Category {
  id: string
  label: string
  color: string
}

export const CATEGORIES: Category[] = [
  { id: 'groceries', label: 'Продукты', color: '#8fae6b' },
  { id: 'transport', label: 'Транспорт', color: '#6b9ac4' },
  { id: 'housing', label: 'Жильё', color: '#c9a24a' },
  { id: 'entertainment', label: 'Развлечения', color: '#b07cc6' },
  { id: 'health', label: 'Здоровье', color: '#c1666b' },
  { id: 'clothing', label: 'Одежда', color: '#4fb0a5' },
  { id: 'utilities', label: 'Коммуналка', color: '#9a9ea5' },
  { id: 'education', label: 'Образование', color: '#6bc48a' },
  { id: 'travel', label: 'Путешествия', color: '#d18a4a' },
  { id: 'other', label: 'Другое', color: '#7d828a' },
]
