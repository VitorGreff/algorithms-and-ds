import LinearSearch from "./array/LinearSearch";
import BinarySearch, { RecursiveBinarySearch } from "./array/BinarySearch";
import BubbleSort from "./sort/BubbleSort";
import Queue from "./sort/Queue";

// console.log(LinearSearch([1, 4, 5, 6, 7], 4))
// console.log(BinarySearch([1, 4, 5, 6, 7], 11))
// console.log(RecursiveBinarySearch([1, 4, 5, 6, 7], 5, 0, 4))
// BubbleSort([3, 4, 45, 6, 2, 1, 234, 5, 2, 1, 2])

let q = new Queue
q.enqueue(2)
q.enqueue(4)
console.log(q.peek())