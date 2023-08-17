## B+ tree  (deletion not implemented)
this is the first part that we are going to make, for b+tree tutorials you can check online on the insertio part and other stuff.
As we are going to store this in disk and not in memory these are not straight forward at all, what we would try to achieve is that we will make the btree traversal and insertion as inmemory but the 
for actual traversal we will add a wrapper and make sure that the b+tree and b+diskWrapper is decoupled.

The node structure will be stored as bytes array and the below node structure will tell where to access which type of bytes

##Node Structure
1. type_of_node (leaf/internal) = 2B space required
2. number of elements in node = 2B (keeping 2 bytes, can be 4 as well) 
3. pointers to next(child) Nodes = number of elements in node * 8Byte 
4. list of offsets pointing to each elements - number of elements in node * 8Byte
5. K,Value elements keys as - Key = 2B value = 2B 

These are not in memory pointers these are 64 bit integers representing the offet and page in disk.
check code in package bTree


we make sure to try to fit the key value in a single page in disk

##Insertion through mutation 
1. go from root to each node until you find the required node
2. check the size of node if greater than or equal to maxsize/2 +1 then split the node and add one element from the node to the upper node 
3. do this recursively 

##Insertion without mutation
1. go from root to each node until you find the required node
2. copy the node path
2. copy the older node and create 2 new nodes from there
3. go through node path and create the nodes again 
4. now attach this node path to the root


