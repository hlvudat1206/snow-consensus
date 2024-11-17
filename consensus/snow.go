package consensus

import "log"

//configurable
const (
	SampleSize    = 10 // Number of peers to query per round
	Alpha         = 6  //Sufficient majority
	MaxIterations = 50 // Confidence Threshold
)

type Snow struct {
	preference string // "A" or "B"
	confidence int
	accepted   bool
}

func NewSnow(initialPreference string) *Snow {
	return &Snow{
		preference: initialPreference,
		confidence: 0,
		accepted:   false,
	}
}

func (s *Snow) Sample(preferences []string) {
	if s.accepted {
		return
	}
	log.Printf("Current preference: %s", s.preference)

	count := map[string]int{}
	for _, pref := range preferences {
		count[pref]++
	}

	for pref, cnt := range count {
		if cnt >= Alpha {
			if s.preference == pref { //the new preference is the same as the old preference
				s.confidence++
			} else {
				s.preference = pref //the new preference is different then the old preference
				s.confidence = 1
			}
			break
		} else {
			s.confidence = 0 //If no response gets a quorum (an α majority of the same response)
		}
	}

	if s.confidence >= MaxIterations { // if the confidence exceeds MaxIterations that is accepted
		s.accepted = true
	}
}

func (s *Snow) IsAccepted() bool {
	return s.accepted
}

func (s *Snow) GetPreference() string {
	return s.preference
}
