package main

type Point struct {
	x, y int
}

type SortStep struct {
	array  []Point
	p1, p2 int
	x, y   []int // point of start
}

type SortSteps struct {
	steps []SortStep
}

func (s *SortSteps) append(step []int, p1, p2 int) {
	stepPoints := make([]Point, len(step))
	i := 0
	for x, y := range step {
		stepPoints[i] = Point{x, y}
		i++
	}
	newStep := SortStep{array: stepPoints, p1: p1, p2: p2, x: make([]int, len(step)), y: make([]int, len(step))}
	s.steps = append(s.steps, newStep)
}

func (s *SortSteps) swapAndSave(steps []int, p1, p2 int) {
	if p1 == p2 {
		return
	}
	steps[p1], steps[p2] = steps[p2], steps[p1]
	s.append(steps, p1, p2)
}

func (s *SortSteps) setPadding(x, y int) {
	for i, _ := range s.steps {
		for j, _ := range s.steps[i].array {
			s.steps[i].x[j] = x
			s.steps[i].y[j] = y
		}
	}
}

func (s *SortSteps) merge(o *SortSteps) *SortSteps {
	for len(s.steps) < len(o.steps) {
		s.steps = append(s.steps, s.steps[len(s.steps)-1])
	}

	for len(o.steps) < len(s.steps) {
		o.steps = append(o.steps, o.steps[len(o.steps)-1])
	}

	for i := 0; i < len(s.steps); i++ {
		s.steps[i].array = append(s.steps[i].array, o.steps[i].array...)
		s.steps[i].x = append(s.steps[i].x, o.steps[i].x...)
		s.steps[i].y = append(s.steps[i].y, o.steps[i].y...)
	}

	return s
}
