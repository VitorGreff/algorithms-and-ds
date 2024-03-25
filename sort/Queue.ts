type Node<T> = {
    value: T
    next?: Node<T>
}
export default class Queue<T>{
    public lenght: number
    public head?: Node<T>
    public tail?: Node<T>

    constructor() {
        this.head = this.tail = undefined
        this.lenght = 0
    }

    enqueue(item?: T): void {
        const node = { value: item } as Node<T>
        this.lenght++
        if (!this.tail) {
            this.tail = this.head = node
            return
        }

        this.tail.next = node
        this.tail = node
    }

    deque(): T | undefined {
        if (!this.head) {
            return undefined
        }

        this.lenght--

        const head = this.head.value
        this.head = this.head.next
        if (this.head == undefined) {
            this.tail = undefined
        }

        return head
    }

    peek(): T | undefined { // returns the first element
        return this.head?.value
    }
}