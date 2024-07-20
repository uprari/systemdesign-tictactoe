package game

type Marker interface {
	Mark(s Spotter)
}

type NotifyWinner interface {
	DeclareWinner(winnerName string)
}

type BaseEvaluator struct {
	NextMarker Marker
}

type RowEvaluators struct {
	BaseEvaluator
	notify        NotifyWinner
	row           int
	max           int
	WinnerTracker map[string]int
}

func NewRowEvaluator(row, max int, m Marker, n NotifyWinner) Marker {
	return &RowEvaluators{row: row, max: max, NextMarker: m, notify: n}
}
func (r *RowEvaluators) Mark(s Spotter) {
	row, _ := s.GetPosition()
	if row == r.row {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == max {
			r.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

type ColEvaluators struct {
	BaseEvaluator
	notify        NotifyWinner
	col           int
	max           int
	WinnerTracker map[string]int
}

func NewColEvaluator(col, max int, m Marker, n NotifyWinner) Marker {
	return &RowEvaluators{col: col, max: max, NextMarker: m, notify: n}
}

func (r *ColEvaluators) Mark(s Spotter) {
	_, col := s.GetPosition()
	if col == r.col {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == max {
			r.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

type LRDiagEvaluators struct {
	BaseEvaluator
	notify        NotifyWinner
	max           int
	WinnerTracker map[string]int
}

func NewDiagEvaluator(max int, m Marker, n NotifyWinner) Marker {
	return &LRDiagEvaluators{max: max, NextMarker: m, notify: n}
}

func (r *ColEvaluators) Mark(s Spotter) {
	row, col := s.GetPosition()
	if col == row {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == max {
			r.DeclareWinner(s.GetPlayerName())
		}
		r.WinnerTracker[s.GetPlayerName()] = val
	}
	r.NextMarker.Mark(s)
}

type RLDiagEvaluators struct {
	BaseEvaluator
	notify        NotifyWinner
	max           int
	WinnerTracker map[string]int
}

func NewRLDiagEvaluator(max int, m Marker, n NotifyWinner) Marker {
	return &LRDiagEvaluators{max: max, NextMarker: m, notify: n}
}

func (r *ColEvaluators) Mark(s Spotter) {
	row, _ := s.GetPosition()
	if col+row == r.max-1 {
		val, ok := r.WinnerTracker[s.GetPlayerName()]
		if !ok {
			r.WinnerTracker[s.GetPlayerName()] = 1
		}
		val++
		if val == max {
			r.DeclareWinner(s.GetPlayerName())
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
		m = colvaluator
	}
	lrDiagEval := NewDiagEvaluator(size, m, n)
	rlDiagEval := NewRLDiagEvaluator(size, lrDiagEval, n)
}
