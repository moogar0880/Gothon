// Grammar package
package Parser

// A label of an arc
type Label struct {
	// An int denoting the type of this Label
	lb_type int
	// The string content of this label
	lb_str string
}

// A list of Labels
type LabelList struct {
	ll_label []Label
}
func (list *LabelList) AddLabel(t int, str string) int {
	var i int
	new_label := new(Label)
	new_label.lb_type = t
	new_label.lb_str = str

	for i = 0; i < len(list.ll_label); i++ {
		match := list.ll_label[i].lb_str == str
		if list.ll_label[i].lb_type == t && match {
			return i
		}
	}

	if list.ll_label == nil {
		list.ll_label := make([]Label, 0)
	}
	append(list.ll_label, new_label)
	return i
}
func (list *LabelList) FindLabel(t int, str string) int {
	new_label := new(Label)
	new_label.lb_type = t
	new_label.lb_str = str

	for i := 0; i < len(list.ll_label); i++ {
		if list.ll_label[i].lb_type == t {
			return i
		}
	}
	return 0
}

// An arc from one state to another
type Arc struct {
	// Label of this arc and State where this arc goes to
	a_lbl, a_arrow int
}

// A state in a DFA
type State struct {
	/* s_lower: Lowest label index
	   s_upper: Highest label index
	   s_accel: Accelerator
	   s_accept: Nonzero for accepting state */
	s_narcs, s_lower, s_upper, s_accel, s_accept int
	// Array of arcs
	s_arc []Arc
}

// A DFA
type DFA struct {
	// Non-terminal this represents, initial state, n_states
	d_type, d_initial, d_nstates int
	// For printing
	d_name string
	// Array of states
	d_state []State
	// Bitset
	bitset []int
}
// Add a new State to this DFA's array of states
func (dfa *DFA) AddState() int {
	s := new(State)
	append(dfa.d_state, s)
	return dfa.d_state
}
// Add an Arc from from to to with the specified label
func (dfa *DFA) AddArc(from, to, lbl int) {
	s := dfa.d_state[from]
	new_arc = new(Arc)
	new_arc.a_lbl = lbl
	new_arc.a_arrow = to

	if s.s_arc == nil {
		s.s_arc := make([]Arc, 0)
	}
	append(s.s_arc, new_arc)
}

// A Grammar
type Grammar struct {
	// g_start: Start symbol of the grammar
	// g_accel: Set if accelerators present
	g_ndfas, g_start, g_accel int
	// Array of DFAs
	g_dfa []DFA
	g_ll LabelList
}
// Add a new DFA to this Grammar's list of DFAs
func (grammar *Grammar) AddDFA(t int, name string) DFA {
	d := new(DFA)
	d.d_type = t
	d.d_name = name
	d.d_initial = -1

	if grammar.g_dfa == nil {
		grammar.g_dfa := make([]DFA, 0)
	}
	append(grammar.g_dfa, d)

	return d
}
func (grammar *Grammar) AddDFA(label Label) {

}
