type Node<T> = {
    value: T
    prev?: Node<T>
}
export default class Stack<T>{
    public size: number
    public top?: Node<T>

    constructor() {
        this.size = 0
        this.top = undefined
    }

    push(item?: T): void {
        const node = { value: item } as Node<T>
        this.size++
        if (!this.top) {
            this.top = node
            return
        }

        node.prev = this.top
        this.top = node

    }

    pop(): T | undefined {
        this.size = Math.max(this.size - 1, 0)
        if (this.size === 0) {
            const top = this.top?.value
            this.top = undefined
            return top
        }

        const top = this.top?.value
        this.top = this.top?.prev
        return top
    }

    peek(): T | undefined {
        return this.top?.value
    }
} 