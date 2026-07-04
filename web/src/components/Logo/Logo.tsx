import { cn } from '../../lib/cn'
import styles from './Logo.module.css'

export function Logo({ size = 'md' }: { size?: 'sm' | 'md' }) {
  return (
    <div className={cn(styles.logo, size === 'sm' && styles.sm)}>
      <span className={styles.mark}>К</span>
      <span className={styles.word}>Кошелёк</span>
    </div>
  )
}
