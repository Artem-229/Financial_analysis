import { Card } from '../Card/Card'
import { cn } from '../../lib/cn'
import styles from './StatCard.module.css'

interface StatCardProps {
  label: string
  value: string
  hint?: string
  trend?: 'up' | 'down'
}

export function StatCard({ label, value, hint, trend }: StatCardProps) {
  return (
    <Card className={styles.card}>
      <span className={styles.label}>{label}</span>
      <span className={styles.value}>{value}</span>
      {hint && (
        <span className={cn(styles.hint, trend === 'up' && styles.up, trend === 'down' && styles.down)}>
          {hint}
        </span>
      )}
    </Card>
  )
}
