package _1656_design_an_ordered_stream

type OrderedStream struct {
	values     []string
	chunkStart int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{values: make([]string, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.values[idKey-1] = value
	chunkEnd := this.chunkStart
	for this.values[chunkEnd] != "" {
		chunkEnd += 1
		if chunkEnd == len(this.values) {
			return this.values[this.chunkStart:chunkEnd]
		}
	}
	chunk := this.values[this.chunkStart:chunkEnd]
	this.chunkStart = chunkEnd
	return chunk
}
