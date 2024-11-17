package consensus

const (
	SampleSize    = 10
	Alpha         = 6
	MaxIterations = 50
)

type Snow struct {
	preference string
	counter    int
	accepted   bool
}

func NewSnow(initialPreference string) *Snow {
	return &Snow{
		preference: initialPreference,
		counter:    0,
		accepted:   false,
	}
}

func (s *Snow) Sample(preferences []string) {
	if s.accepted {
		return
	}

	count := map[string]int{}
	for _, pref := range preferences {
		count[pref]++
	}

	for pref, cnt := range count {
		if cnt >= Alpha {
			if s.preference == pref {
				s.counter++
			} else {
				s.preference = pref
				s.counter = 1
			}
			break
		}
	}

	if s.counter >= MaxIterations {
		s.accepted = true
	}
}

func (s *Snow) IsAccepted() bool {
	return s.accepted
}

func (s *Snow) GetPreference() string {
	return s.preference
}
