I don't like this solution one bit, but couldn't figure out the edge cases with the proper maths approach within a reasonable amount of time.

This more iterative approach could be improved by batching a bit more (if you're going left 60 from 50, you can subtract 50 to get to 0 in one go, track answers as needed, then subtract the remaining 10).

The best option would be the proper maths approach to avoid iterations, but for some reason it just wouldn't gel with my brain on the day.

I typically split answers into a part1 and part2 function, but I combined them here as they overlapped so much.