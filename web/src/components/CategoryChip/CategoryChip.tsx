import type { CSSProperties } from 'react'
import { cn } from '../../lib/cn'
import styles from './CategoryChip.module.css'

interface CategoryChipProps {
  label: string
  color: string
  active?: boolean
  onClick?: () => void
}

export function CategoryChip({ label, color, active, onClick }: CategoryChipProps) {
  return (
    <button
      type="button"
      className={cn(styles.chip, active && styles.active)}
      style={{ '--chip-color': color } as CSSProperties}
      onClick={onClick}
    >
      <span className={styles.dot} />
      {label}
    </button>
  )
}
