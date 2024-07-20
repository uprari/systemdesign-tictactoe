package game

type Marker interface {
	Mark(s Spotter)
}

type NotifyWinner interface {
	DeclareWinner(winnerName string)
}

type BaseEvaluator struct {
	NextMarker    Marker
	WinnerTracker map[string]int
}

func (b *BaseEvaluator) init(m Marker) {
	b.NextMarker = m
	b.WinnerTracker = make(map[string]int)
}

type RowEvaluators struct {
	BaseEvaluator
	notify NotifyWinner
	row    int
	max    int
}

func NewRowEvaluator(row, max int, m Marker, n NotifyWinner) Marker {
	re := &RowEvaluators{row: row, max: max, notify: n}
	re.init(m)
	return re
}
func (r *RowEvaluators) Mark(s Spotter) {
	row, _ := s.GetPosition()
	if row == r.row {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == r.max {
			r.notify.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

type ColEvaluators struct {
	BaseEvaluator
	notify NotifyWinner
	col    int
	max    int
}

func NewColEvaluator(col, max int, m Marker, n NotifyWinner) Marker {
	ce := &ColEvaluators{col: col, max: max, notify: n}
	ce.init(m)
	return ce
}

func (r *ColEvaluators) Mark(s Spotter) {
	_, col := s.GetPosition()
	if col == r.col {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == r.max {
			r.notify.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

type LRDiagEvaluators struct {
	BaseEvaluator
	notify NotifyWinner
	max    int
}

func NewDiagEvaluator(max int, m Marker, n NotifyWinner) Marker {
	lrde := &LRDiagEvaluators{max: max, notify: n}
	lrde.init(m)
	return lrde
}

func (r *LRDiagEvaluators) Mark(s Spotter) {
	row, col := s.GetPosition()
	if col == row {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == r.max {
			r.notify.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

type RLDiagEvaluators struct {
	BaseEvaluator
	notify NotifyWinner
	max    int
}

func NewRLDiagEvaluator(max int, m Marker, n NotifyWinner) Marker {
	rlde := &RLDiagEvaluators{max: max, notify: n}
	rlde.init(m)
	return rlde
}

func (r *RLDiagEvaluators) Mark(s Spotter) {
	row, col := s.GetPosition()
	if col+row == r.max-1 {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == r.max {
			r.notify.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

func BuildEvaluator(m Marker, size int, n NotifyWinner) Marker {
	for row := 0; row < size; row++ {
		rowEvaluator := NewRowEvaluator(row, size, m, n)
		m = rowEvaluator
	}
	for col := 0; col < size; col++ {
		colEvaluator := NewColEvaluator(col, size, m, n)
		m = colEvaluator
	}
	lrDiagEval := NewDiagEvaluator(size, m, n)
	rlDiagEval := NewRLDiagEvaluator(size, lrDiagEval, n)
	return rlDiagEval
}
