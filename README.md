# LRU Cache

Fully associative and in-memory least recently used (LRU) cache implementation.

The cache works by using two data structures:

- Map
- Linked List

Cache entries are fetched and stored in the map using the user-defined key.

Bookkeeping

The linked list is used for constant-time tracking of the least and most
recently used entries. The head of the list serves as the least recently used
object and the tail as most recently used.

For each Get operation that results in a cache hit, the element is moved to the
end of the linked list to reflect it's most recently used status.

For Put operations, whether its a new entry being added or an update to an
existing element, the entry is moved to the tail of the list.

## Demonstration

```text
LRUCache(3)

+-----+-----+-----+
|     |     |     |
+-----+-----+-----+

Put(1,1)

+-----+-----+-----+
|     |     | 1,1 |
+-----+-----+-----+

Put(2,2)

+-----+-----+-----+
|     | 1,1 | 2,2 |
+-----+-----+-----+

Put(3,3)

+-----+-----+-----+
| 1,1 | 2,2 | 3,3 |
+-----+-----+-----+

Get(1)

+-----+-----+-----+
| 2,2 | 3,3 | 1,1 |
+-----+-----+-----+

Put(3,4)

+-----+-----+-----+
| 2,2 | 1,1 | 3,4 |
+-----+-----+-----+

Put(5,5)

+-----+-----+-----+
| 1,1 | 3,4 | 5,5 |
+-----+-----+-----+
```
