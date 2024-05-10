export default function BinarySearch(haystack: number[], needle: number): boolean {
    let low = 0
    let middle = 0
    let high = haystack.length
    do {
        middle = Math.floor(low + (high - low) / 2)
        if (haystack[middle] === needle) return true
        else if (haystack[middle] < needle) {
            low = middle + 1
        }
        else {
            high = middle
        }
    } while (low < high)
    return false
}

export function RecursiveBinarySearch(haystack: number[], high: number, low: number, needle: number): boolean {
    if (low > high) return false

    let middle = Math.floor((high + low) / 2)
    let value = haystack[middle]

    if (value === needle) return true

    else if (value < needle) return RecursiveBinarySearch(haystack, high, middle + 1, needle)

    else if (value > needle) return RecursiveBinarySearch(haystack, middle - 1, low, needle)

    return false
}