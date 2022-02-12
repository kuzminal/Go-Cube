package node

type Node struct {
	Name            string
	Ip              string
	Memory          int
	MemoryAllocated int
	Cores           int
	Disk            int
	DiskAllocated   int
	TaskCount       int
	Role            string
}
