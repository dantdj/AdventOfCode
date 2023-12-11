Part one is a bit rocky - it was originally meant to be a breadth-first search implementation,
but I couldn't figure out how to avoid double-counting the steps (because you're doing two paths simultaneously),
so I just hacked it and halved the total number of steps taken, when really I could have just gone round the loop until the start
came round again and halved that (maybe minus 1 first, not sure). All the nodes get marked as discovered correctly either way.
