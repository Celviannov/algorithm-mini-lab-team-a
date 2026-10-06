

# Tower of Hanoi experiment:

Predict: With 3 disks, I think it will take 7 moves. The first move should be disk 1 from A to C. The biggest disk should move only once near the end.

Run: We run the program with 3 disks. It prints each move and says total moves: 7, expected: 7.

Observe: It took exactly 7 moves. First move was A to C. Disk 3 moved at move 4 from A to C. The formula 2^3 - 1 gave 7. So the prediction was correct.

Change: We change the number of disks from 3 to 4.

Explain: With 4 disks, it takes 15 moves. The pattern is 2^n - 1. Every extra disk doubles the moves and adds one. The algorithm still works the same way: move n-1 disks to the helper, move the biggest disk, then move n-1 disks to the target.

# Bubble Sort experiment:

Predict: Input is [64, 34, 25, 12, 22]. I think the sorted result will be [12, 22, 25, 34, 64]. Swaps happen when the left number is bigger than the right number. The first comparison is 64 and 34, so they swap. The 64 should move to the end after the first pass.

Run: We run the program. It prints each pass and the final sorted list.

Observe: Pass 1 moved 64 to the end. Pass 3 already had the sorted list. Pass 4 made no swaps. Final sorted array was [12, 22, 25, 34, 64]. This matched the prediction.

Change: We change the input to [9, 3, 7, 2, 1].

Explain: The new input has 9 at the start, which is the biggest number. So 9 bubbles to the end in the first pass. More swaps happen because the array is less sorted. The algorithm still compares adjacent numbers and swaps if the left is bigger. Final sorted array is [1, 2, 3, 7, 9].

# BFS experiment:

Predict: The graph has nodes Andeevka, Toretsk, Kosntantinovka, Pokrovsk, Mirnograd. Starting from Andeevka, I think BFS will visit Andeevka, then Toretsk and Kosntantinovka, then Pokrovsk and Mirnograd. So the order should be Andeevka -> Toretsk -> Kosntantinovka -> Pokrovsk -> Mirnograd. Distances: Andeevka 0, Toretsk 1, Kosntantinovka 1, Pokrovsk 2, Mirnograd 2.

Run: We run the program. It prints the BFS order and distances. It also runs tests and all 8 tests pass.

Observe: The order matched the prediction. BFS visited all neighbours of Andeevka first, then moved to the next level. Distances matched. The tests for single node, disconnected graph, cycle, and empty graph also passed.

Change: We change the start node from Andeevka to Pokrovsk.

Explain: Starting from Pokrovsk changes the order. Now level 1 is Toretsk and Mirnograd. Level 2 is Andeevka and Kosntantinovka. So the new order is Pokrovsk -> Toretsk -> Mirnograd -> Andeevka -> Kosntantinovka. BFS still works the same way: it uses a queue, visits all neighbours of the current level, then goes deeper. The visited set stops cycles from looping forever.

That is all three experiments in order.

