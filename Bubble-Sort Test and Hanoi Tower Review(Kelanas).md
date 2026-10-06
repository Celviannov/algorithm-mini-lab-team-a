# Bubble-Sort, BFS (Bredth-Fisrt-Search), and Tower Hanoi Review
# Alghoritms MiniLab Team A

# Bubble Sort & Tower of Hanoi Review

Kelana: I tested Kevin's Bubble-Sort code, and reviewed Mahendra's Tower of Hanoi.

## Bubble Sort

I predicted that the array will sort to [1, 2, 3, 5, 7, 9]. Since the index length is now 6, it will take 5 passes to fully sort. The largest number (9) will bubble to the end in Pass 1, then 7, then 5.

I changed the values and the index length (from 5 to 6). Because the array is longer, Bubble Sort needed 5 passes instead of 4. The largest numbers (9, 7, 5) bubbled to the end one by one in the first few passes. The final sorted result is [1, 2, 3, 5, 7, 9].

## Tower Hanoi

I ran the code and it works, sorting [5, 2, 8, 1, 4] into [1, 2, 4, 5, 8] in 4 passes. The algorithm uses nested loops where the inner loop compares adjacent numbers and swaps them if the left is bigger, so the biggest number bubbles to the end each pass.

However, there is a bug. The code does not stop early if the array is already sorted, so if we run [1, 2, 3, 4, 5] it will still do 4 passes for nothing. We should add a swapped boolean flag to break out early, making the best case O(n) instead of O(n²).

I also suggest testing edge cases like empty array, single element, and duplicates. And please add the README with the 5 steps: predict, run, observe, change, explain.

# Tower Hanoi and BFS Review

celvian: I tested Mahen's Tower of Hanoi, and reviewed Kelana's BFS code.

## Tower Hanoi

I tried disks = -1 just for fun, and the program crashed. I asked Kelana and figured out why the code only stops when the number of disks becomes 0. But if you start at −1, it goes −2, −3, −4… and never reaches 0, so it keeps calling itself forever until the program dies. I also tried disks = 0 — that one didn't crash, but it just said "Done! Total moves: 0" which is kind of pointless.

 ## BFS Bredth-First-Search

all the tests passed. The most interesting one was the disconnected graph, a graph that's split into two separate pieces. BFS starting from A only found the nodes in A's piece. The nodes in the other piece never got a distance. At first I thought this was a bug, but actually it's correct, BFS can't jump to a part of the graph that isn't connected.

# Bubble Sort & BFS Review

Mahendra: I Tested Celvian's Bubble sort, And Reviewed Kelana's BFS code

## Bubble Sort

All the tests passed successfully. The numbers were sorted correctly from [5, 2, 8, 1, 4] into ascending order. The interesting part was the swapped variable, which allows the program to stop early when no swaps are needed. This makes the sorting process more efficient.

## BSF (Bredth-First-Search)

All the tests passed successfully. BFS correctly visited the connected nodes and ignored the unreachable node. The test with cycles also worked correctly because BFS avoids visiting the same node twice.


