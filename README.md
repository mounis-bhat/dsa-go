# DSA-Go

A comprehensive implementation of data structures and algorithms in Go, featuring clean code, extensive testing, and practical examples.

## 📁 Project Structure

```
├── cmd/
│   └── main.go              # Main entry poin| Group Anagrams        | Hash Map/Sorting     | O(n×k×log k)    | O(n×k)           |
| Top K Frequent        | Hash Map/Heap        | O(n log k)      | O(n)             |
├── internal/
│   ├── avltree/             # AVL Tree implementation
│   ├── binarytree/          # Binary Search Tree
│   ├── bubblesort/          # Bubble Sort algorithm
│   ├── linkedlist/          # Singly Linked List
│   ├── neetcode150/         # Neetcode 150 problem solutions
│   ├── queue/               # Queue data structure
│   ├── selectionsort/       # Selection Sort algorithm
│   └── stack/               # Stack data structure
├── pkg/
│   └── utils/               # Utility functions with Go generics
├── go.mod                   # Go module file
└── Makefile                 # Build automation
```

## 🚀 Features

### Data Structures

#### 1. **Stack** ([internal/stack](internal/stack/stack.go))

- LIFO (Last In, First Out) implementation
- Operations: `Push`, `Pop`, `IsEmpty`, `Display`
- Error handling for empty stack operations
- Time Complexity: Push O(1), Pop O(1)

#### 2. **Queue** ([internal/queue](internal/queue/queue.go))

- FIFO (First In, First Out) implementation
- Operations: `Enqueue`, `Dequeue`, `Peek`, `IsEmpty`, `Display`
- Error handling for empty queue operations
- Time Complexity: Enqueue O(1), Dequeue O(n)

#### 3. **Linked List** ([internal/linkedlist](internal/linkedlist/linked_list.go))

- Singly linked list with comprehensive operations
- **Core Operations**: `Append`, `Prepend`, `Insert`, `Delete`, `Reverse`
- **Access Methods**: `Get`, `Search`, `Display`
- **Stack/Queue Operations**: `Pop`, `Shift` (with return values and error handling)
- **Utility Methods**: `Size`, `Clear`, `IsEmpty`
- **Error Handling**: Returns boolean success indicators for operations that can fail
- Time Complexity: Prepend/Append O(1), Insert/Delete/Search O(n), Get O(n)

#### 4. **Binary Search Tree** ([internal/binarytree](internal/binarytree/binary_tree.go))

- Classic BST implementation with recursive operations
- Tree traversals: `InOrder`, `PreOrder`, `PostOrder`, `LevelOrder`
- Search functionality with O(log n) average case
- Insert operation maintains BST property

#### 5. **AVL Tree** ([internal/avltree](internal/avltree/avl_tree.go))

- Self-balancing binary search tree
- Automatic rotations (Left, Right, Left-Right, Right-Left) for balance maintenance
- All BST operations with O(log n) guarantee
- Tree visualization feature with `Visualize()` method
- Uses utility functions from [`pkg/utils`](pkg/utils/utils.go) for generic operations

### Sorting Algorithms

#### 1. **Bubble Sort** ([internal/bubblesort](internal/bubblesort/bubble_sort.go))

- Time Complexity: O(n²) worst case, O(n) best case
- Space Complexity: O(1)
- Optimized with early termination when no swaps occur
- In-place sorting algorithm

#### 2. **Selection Sort** ([internal/selectionsort](internal/selectionsort/selection_sort.go))

- Time Complexity: O(n²) in all cases
- Space Complexity: O(1)
- In-place sorting algorithm
- Finds minimum element and swaps with current position

### Neetcode 150 Solutions

A comprehensive collection of coding interview problems from the popular [Neetcode 150](https://neetcode.io/) list.

### Neetcode 150 Solutions

A comprehensive collection of coding interview problems from the popular [Neetcode 150](https://neetcode.io/) list.

#### Progress Overview ([internal/neetcode150](internal/neetcode150/))

| #   | Problem               | Difficulty | Pattern                 | Time         | Space  | Status |
| --- | --------------------- | ---------- | ----------------------- | ------------ | ------ | ------ |
| 1   | Two Sum               | Easy       | Hash Map                | O(n)         | O(n)   | ✅     |
| 2   | Valid Anagram         | Easy       | Frequency Count         | O(n)         | O(1)   | ✅     |
| 3   | Two Sum II            | Medium     | Two Pointers            | O(n)         | O(1)   | ✅     |
| 4   | Majority Element      | Easy       | Boyer-Moore/Hash Map    | O(n)         | O(1)   | 🚧     |
| 5   | Contains Duplicate II | Easy       | Sliding Window/Hash Map | O(n)         | O(k)   | ✅     |
| 6   | Group Anagrams        | Medium     | Hash Map/String Sorting | O(n×k×log k) | O(n×k) | ✅     |
| 7   | Top K Frequent        | Medium     | Hash Map/Heap           | O(n log k)   | O(n)   | 🚧     |

#### Categories Covered:

- **Arrays & Hashing**: Two Sum, Valid Anagram, Majority Element, Contains Duplicate II, Group Anagrams, Top K Frequent
- **Two Pointers**: Two Sum II
- **Sliding Window**: Contains Duplicate II
- **String Processing**: Group Anagrams
- **Heap/Priority Queue**: Top K Frequent

For detailed problem descriptions, examples, and approaches, see the individual source files in [`internal/neetcode150/`](internal/neetcode150/).

### Utility Functions ([pkg/utils](pkg/utils/utils.go))

Generic utility functions leveraging Go 1.18+ generics:

- `Compare[T]` - Generic comparison function for int, float64, string
- `Swap[T]` - Generic swap function for any type
- `Max[T]` / `Min[T]` - Find maximum/minimum values for numeric types
- `Contains[T]` - Check if slice contains element (comparable types)
- `RemoveDuplicates[T]` - Remove duplicate elements from slice

## 🛠️ Getting Started

### Prerequisites

- Go 1.22.5 or later

### Installation

```bash
git clone https://github.com/mounis-bhat/dsa-go.git
cd dsa-go
go mod tidy
```

### Building the Project

```bash
# Build the main binary
make build

# Or manually
go build -o dsa ./cmd
```

### Running Tests

```bash
# Run all tests
make test

# Run specific module tests
make test-stack
make test-queue
make test-linkedlist
make test-binarytree
make test-avltree
make test-bubble
make test-selection
make test-neetcode150

# Run specific Neetcode150 problem tests
make test-two-sum
make test-two-sum-ii
make test-valid-anagram
make test-majority-element
make test-contains-duplicate-ii
make test-group-anagrams
make test-top-k-frequent
```

### Running the Application

```bash
make run
```

## 🧪 Testing

Each module includes comprehensive test suites with 100% coverage:

- **Unit Tests**: Every data structure and algorithm has corresponding test files
- **Edge Cases**: Tests cover empty collections, single elements, and boundary conditions
- **Error Handling**: Validation of error conditions and proper error returns
- **Performance**: Tests verify expected behavior under various data sizes

Example test execution:

```bash
# Run all tests with verbose output
go test -v ./...

# Run tests for a specific package
go test -v ./internal/stack

# Run tests with coverage
go test -cover ./...
```

## 📊 Time & Space Complexities

### Data Structures

| Data Structure | Insert   | Delete   | Search   | Space |
| -------------- | -------- | -------- | -------- | ----- |
| Stack          | O(1)     | O(1)     | O(n)     | O(n)  |
| Queue          | O(1)     | O(n)\*   | O(n)     | O(n)  |
| Linked List    | O(1)\*\* | O(n)     | O(n)     | O(n)  |
| Binary Tree    | O(n)     | O(n)     | O(n)     | O(n)  |
| AVL Tree       | O(log n) | O(log n) | O(log n) | O(n)  |

\*Queue dequeue is O(n) due to slice shifting  
\*\*O(1) for prepend/append, O(n) for arbitrary position

### Sorting Algorithms

| Algorithm      | Best Case | Average Case | Worst Case | Space |
| -------------- | --------- | ------------ | ---------- | ----- |
| Bubble Sort    | O(n)      | O(n²)        | O(n²)      | O(1)  |
| Selection Sort | O(n²)     | O(n²)        | O(n²)      | O(1)  |

### Neetcode 150 Solutions

| Problem               | Pattern          | Time Complexity | Space Complexity |
| --------------------- | ---------------- | --------------- | ---------------- |
| Two Sum               | Hash Map         | O(n)            | O(n)             |
| Two Sum II            | Two Pointers     | O(n)            | O(1)             |
| Valid Anagram         | Frequency Count  | O(n)            | O(1)             |
| Majority Element      | Boyer-Moore      | O(n)            | O(1)             |
| Contains Duplicate II | Sliding Window   | O(n)            | O(k)             |
| Group Anagrams        | Hash Map/Sorting | O(n×k×log k)    | O(n×k)           |

## 🎯 Usage Examples

### Stack Example

```go
s := stack.NewStack()
s.Push(1)
s.Push(2)
element, err := s.Pop() // Returns 2
fmt.Println(s.Display()) // [1]
```

### Linked List Example

```go
ll := linkedlist.NewLinkedList()
ll.Append(1)
ll.Append(2)
ll.Prepend(0)
fmt.Println(ll.Display()) // [0, 1, 2]

ll.Insert(3, 3) // Insert 3 at index 3
ll.Reverse()
fmt.Println(ll.Display()) // [3, 2, 1, 0]

element, exists := ll.Pop() // Remove last element
if exists {
    fmt.Printf("Popped: %d\n", element) // Popped: 0
}

index := ll.Search(2) // Find index of value 2
fmt.Printf("Element 2 is at index: %d\n", index) // Element 2 is at index: 1
```

### AVL Tree Example

```go
avl := avltree.NewAVLTree()
avl.Insert(10)
avl.Insert(5)
avl.Insert(15)
avl.Visualize() // Prints tree structure
fmt.Println(avl.InOrder()) // [5, 10, 15]
```

### Neetcode 150 Example

```go
// Two Sum problem
nums := []int{2, 7, 11, 15}
target := 9
indices := neetcode150.TwoSum(nums, target) // Returns [0, 1]

// Two Sum II problem (sorted array, 1-indexed)
numbers := []int{2, 7, 11, 15}
indices = neetcode150.TwoSumII(numbers, target) // Returns [1, 2]

// Valid Anagram problem
isAnagram := neetcode150.ValidAnagram("anagram", "nagaram") // Returns true

// Majority Element problem
majority := neetcode150.MajorityElement([]int{3, 2, 3}) // Returns 3

// Contains Duplicate II problem
hasDuplicateWithinK := neetcode150.ContainsDuplicateII([]int{1, 2, 3, 1}, 3) // Returns true

// Group Anagrams problem
anagramGroups := neetcode150.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
// Returns [["bat"], ["nat","tan"], ["ate","eat","tea"]] (order may vary)

// Top K Frequent Elements problem
topK := neetcode150.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
// Returns [1, 2] (order may vary)
```

### Utility Functions Example

```go
// Generic comparison
result := utils.Compare(5, 3) // Returns 1

// Generic swap
a, b := 10, 20
utils.Swap(&a, &b) // a=20, b=10

// Remove duplicates
nums := []int{1, 2, 2, 3, 3, 4}
unique := utils.RemoveDuplicates(nums) // [1, 2, 3, 4]
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go conventions and best practices
- Add comprehensive tests for new features
- Update documentation for API changes
- Ensure all tests pass before submitting PR

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🎯 Future Enhancements

- [ ] Add more sorting algorithms (Quick Sort, Merge Sort, Heap Sort)
- [ ] Implement Hash Table/HashMap data structure
- [ ] Add Graph data structures (Adjacency List/Matrix)
- [ ] Implement graph algorithms (DFS, BFS, Dijkstra)
- [ ] Add more LeetCode solutions by category
- [ ] Create interactive CLI demos for each data structure
- [ ] Add benchmarking tests and performance analysis
- [ ] Implement priority queue and heap data structures
- [ ] Add trie (prefix tree) implementation

## 📚 Learning Resources

This project serves as a practical implementation guide for:

- **Go Programming**: Modern Go features including generics, error handling, and testing
- **Data Structures**: Fundamental computer science concepts with real implementations
- **Algorithm Design**: Analysis of time/space complexity and optimization techniques
- **Test-Driven Development**: Comprehensive testing strategies in Go
- **Software Engineering**: Clean code principles and project organization

Each implementation includes:

- Clear, readable code with comprehensive comments
- Proper error handling and edge case management
- Extensive test coverage with multiple test scenarios
- Performance considerations and complexity analysis

Perfect for computer science students, Go developers, and anyone looking to strengthen their understanding of fundamental algorithms and data structures.

## 🏗️ Architecture

The project follows Go best practices with a clean architecture:

- **`cmd/`**: Application entry points
- **`internal/`**: Private application code (data structures & algorithms)
- **`pkg/`**: Public library code (utilities that could be imported by other projects)
- **Comprehensive testing**: Each package has corresponding test files
- **Makefile**: Automation for common development tasks

This structure makes the codebase maintainable, testable, and easy to extend with new features.
