import LinearSearch from "./array/LinearSearch";
import BinarySearch, { RecursiveBinarySearch } from "./array/BinarySearch";
import BubbleSort from "./sort/BubbleSort";
import Queue from "./sort/Queue";
import Stack from "./sort/Stack";

let s = new Stack
s.push(5)
s.push(7)
s.push(9)
s.push(1)
s.push(8)
console.log(s)
console.log(s.top?.prev?.prev?.prev)