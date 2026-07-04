const STORAGE_KEY = 'tx-dates'

function readMap(): Record<string, string> {
  try {
    return JSON.parse(localStorage.getItem(STORAGE_KEY) ?? '{}')
  } catch {
    return {}
  }
}

function writeMap(map: Record<string, string>) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(map))
}

export function setTransactionDate(id: string, date: string) {
  const map = readMap()
  map[id] = date
  writeMap(map)
}

export function getTransactionDate(id: string): string | undefined {
  return readMap()[id]
}

export function removeTransactionDate(id: string) {
  const map = readMap()
  delete map[id]
  writeMap(map)
}
