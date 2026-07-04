import { forwardRef, type InputHTMLAttributes } from 'react'
import { cn } from '../../lib/cn'
import styles from './Input.module.css'

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string
  error?: string
  wrapperClassName?: string
}

export const Input = forwardRef<HTMLInputElement, InputProps>(function Input(
  { label, error, id, className, wrapperClassName, name, ...rest },
  ref,
) {
  const inputId = id ?? name
  return (
    <label className={cn(styles.field, wrapperClassName)} htmlFor={inputId}>
      {label && <span className={styles.label}>{label}</span>}
      <input
        id={inputId}
        name={name}
        ref={ref}
        className={cn(styles.input, error && styles.inputError, className)}
        {...rest}
      />
      {error && <span className={styles.error}>{error}</span>}
    </label>
  )
})
