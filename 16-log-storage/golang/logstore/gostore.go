package logstore

// OrderID type
type OrderID int

// Logger interface
type Logger interface {
	Record(OrderID)
	GetLast(int) OrderID
}

// Log type
type Log struct {
	db    []OrderID
	cap   int
	next  int
	count int
}

// NewLog function
func NewLog(cap int) Logger {
	return &Log{make([]OrderID, cap), cap, 0, 0}
}

// Record function
func (l *Log) Record(id OrderID) {
	l.db[l.next] = id
	l.next = (l.next + 1) % l.cap

	if l.count < l.cap {
		l.count++
	}
}

// GetLast function
func (l *Log) GetLast(i int) OrderID {
	if i <= 0 || i > l.cap || i > l.count {
		panic("logger error: index out of range")
	}

	if l.count < l.cap { // before first round
		return l.db[i-1]
	}

	return l.db[(l.next+i-1)%l.cap]
}
