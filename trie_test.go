package trie

import "testing"

func TestIsEmpty(t *testing.T) {

	emptyTrie := NewTrie()

	empty := emptyTrie.isEmpty()
	if empty != true {
		t.Errorf("Expected 'isEmpty' to be 'true'")
	}
}

func TestIsNotEmpty(t *testing.T) {

	test := NewTrie()

	inserted := test.Insert("romane")
	if inserted != true {
		t.Errorf("Expected inserted to be 'true'")
	}

	empty := test.isEmpty()
	if empty != false {
		t.Errorf("Expected 'isEmpty' to be 'false'")
	}
}

func TestInsert(t *testing.T) {

	insertTests := []struct {
		name          string
		value         string
		trie          Trie
		expectedCount int
		inserted      bool
	}{
		{
			name:          "insert into empty trie (empty string)",
			value:         "",
			trie:          getTrie(0, 'r'),
			expectedCount: 0,
			inserted:      false,
		},
		{
			name:          "insert into empty trie (string less than 2 characters)",
			value:         "a",
			trie:          getTrie(0, 'r'),
			expectedCount: 0,
			inserted:      false,
		},
		{
			name:          "insert into empty trie (trimmed string less than 2 characters)",
			value:         " \r\n",
			trie:          getTrie(0, 'r'),
			expectedCount: 0,
			inserted:      false,
		},
		{
			name:          "insert into empty trie",
			value:         "romane",
			trie:          getTrie(0, 'r'),
			expectedCount: 1,
			inserted:      true,
		},
		{
			name:          "insert into trie existing element",
			value:         "romane",
			trie:          getTrie(1, 'r'),
			expectedCount: 1,
			inserted:      false,
		},
		{
			name:          "insert into trie with one element",
			value:         "romanus",
			trie:          getTrie(1, 'r'),
			expectedCount: 2,
			inserted:      true,
		},
		{
			name:          "insert into trie with two elements",
			value:         "romulus",
			trie:          getTrie(2, 'r'),
			expectedCount: 3,
			inserted:      true,
		},
		{
			name:          "insert into trie with three elements",
			value:         "rubens",
			trie:          getTrie(3, 'r'),
			expectedCount: 4,
			inserted:      true,
		},
		{
			name:          "insert into trie with four elements",
			value:         "ruber",
			trie:          getTrie(4, 'r'),
			expectedCount: 5,
			inserted:      true,
		},
		{
			name:          "insert into trie with five elements",
			value:         "rubicon",
			trie:          getTrie(5, 'r'),
			expectedCount: 6,
			inserted:      true,
		},
		{
			name:          "insert into trie with six elements",
			value:         "rubicundus",
			trie:          getTrie(6, 'r'),
			expectedCount: 7,
			inserted:      true,
		},
	}

	for _, test := range insertTests {
		inserted := test.trie.Insert(test.value)
		if inserted != test.inserted {
			t.Errorf("test '%s': expected inserted to be %t", test.name, test.inserted)
		}
		if test.trie.Count() != test.expectedCount {
			t.Errorf("test '%s': expected count to be %d, but was %d", test.name, test.expectedCount, test.trie.Count())
		}
	}
}

func getTrie(nodes int, prefix byte) Trie {

	emptyTrie := NewTrie()
	if nodes == 0 {
		return emptyTrie
	}
	if prefix == 'r' {
		if nodes == 1 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{
					{value: "omane", childCount: 0}},
				childCount: 1}}, count: 1}
		}
		if nodes == 2 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "oman",
					children: []*Node{{value: "e", childCount: 0},
						{value: "us", childCount: 0}},
					childCount: 2}}, childCount: 1}}, count: 2}
		}
		if nodes == 3 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2}}, childCount: 1}}, count: 3}
		}
		if nodes == 4 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2},
					{value: "ubens", childCount: 0}}, childCount: 1}}, count: 4}
		}
		if nodes == 5 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2},
					{value: "ube",
						children: []*Node{{value: "ns", childCount: 0},
							{value: "r", childCount: 0}},
						childCount: 2}}, childCount: 1}}, count: 5}
		}
		if nodes == 6 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2},
					{value: "ub",
						children: []*Node{{value: "e",
							children: []*Node{{value: "ns", childCount: 0},
								{value: "r", childCount: 0}}, childCount: 2},
							{value: "icon", childCount: 0}}, childCount: 2}},
				childCount: 1}}, count: 6}
		}
		if nodes == 7 {
			return Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2},
					{value: "ub",
						children: []*Node{{value: "e",
							children: []*Node{{value: "ns", childCount: 0},
								{value: "r", childCount: 0}}, childCount: 2},
							{value: "ic",
								children: []*Node{{value: "on", childCount: 0},
									{value: "undus", childCount: 0}},
								childCount: 2}},
						childCount: 2}}, childCount: 1}}, count: 7}
		}
	}
	return emptyTrie
}

var findTests = []struct {
	name   string
	value  string
	trie   Trie
	found  bool
	isLeaf bool
}{
	{
		name:   "find nonexistent element in trie",
		value:  "romanus",
		trie:   getTrie(1, 'r'),
		found:  false,
		isLeaf: false,
	},
	{
		name:   "find trimmed nonexistent element in trie",
		value:  " \r\n",
		trie:   getTrie(1, 'r'),
		found:  false,
		isLeaf: false,
	},
	{
		name:   "find existing element in trie",
		value:  "romane",
		trie:   getTrie(1, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find trimmed existing element in trie",
		value:  " romane\n",
		trie:   getTrie(1, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing element in trie with two elements",
		value:  "romane",
		trie:   getTrie(2, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing second element in trie with two elements",
		value:  "romanus",
		trie:   getTrie(2, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing element in trie with three elements",
		value:  "romane",
		trie:   getTrie(3, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing second element in trie with three elements",
		value:  "romanus",
		trie:   getTrie(3, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing third element in trie with three elements",
		value:  "romulus",
		trie:   getTrie(3, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing element in trie with four elements",
		value:  "romane",
		trie:   getTrie(4, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing second element in trie with four elements",
		value:  "romanus",
		trie:   getTrie(4, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing third element in trie with four elements",
		value:  "romulus",
		trie:   getTrie(4, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fourth element in trie with four elements",
		value:  "rubens",
		trie:   getTrie(4, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing element in trie with five elements",
		value:  "romane",
		trie:   getTrie(5, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing second element in trie with five elements",
		value:  "romanus",
		trie:   getTrie(5, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing third element in trie with five elements",
		value:  "romulus",
		trie:   getTrie(5, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fourth element in trie with five elements",
		value:  "rubens",
		trie:   getTrie(5, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fifth element in trie with five elements",
		value:  "ruber",
		trie:   getTrie(5, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing element in trie with six elements",
		value:  "romane",
		trie:   getTrie(6, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing second element in trie with six elements",
		value:  "romanus",
		trie:   getTrie(6, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing third element in trie with six elements",
		value:  "romulus",
		trie:   getTrie(6, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fourth element in trie with six elements",
		value:  "rubens",
		trie:   getTrie(6, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fifth element in trie with six elements",
		value:  "ruber",
		trie:   getTrie(6, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing sixth element in trie with six elements",
		value:  "rubicon",
		trie:   getTrie(6, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing element in trie with seven elements",
		value:  "romane",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing second element in trie with seven elements",
		value:  "romanus",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing third element in trie with seven elements",
		value:  "romulus",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fourth element in trie with seven elements",
		value:  "rubens",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing fifth element in trie with seven elements",
		value:  "ruber",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing sixth element in trie with seven elements",
		value:  "rubicon",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
	{
		name:   "find existing seventh element in trie with seven elements",
		value:  "rubicundus",
		trie:   getTrie(7, 'r'),
		found:  true,
		isLeaf: true,
	},
}

func TestFind(t *testing.T) {

	for _, test := range findTests {
		found, n := test.trie.Find(test.value)
		if found != test.found {
			t.Errorf("test '%s': expected found to be %t", test.name, test.found)
		}
		if found {
			if n.IsLeaf() != test.isLeaf {
				t.Errorf("test '%s': expected isLeaf to be %t", test.name, test.isLeaf)
			}
		}
	}
}

func BenchmarkInsert(b *testing.B) {

	trie := NewTrie()

	insertBenchmarks := []struct {
		name          string
		value         string
		expectedCount int
		inserted      bool
	}{
		{
			name:          "insert into empty trie (empty string)",
			value:         "",
			expectedCount: 0,
			inserted:      false,
		},
		{
			name:          "insert into empty trie (string less than 2 characters)",
			value:         "a",
			expectedCount: 0,
			inserted:      false,
		},
		{
			name:          "insert into empty trie (trimmed string less than 2 characters)",
			value:         " \r\n",
			expectedCount: 0,
			inserted:      false,
		},
		{
			name:          "insert into empty trie",
			value:         "romane",
			expectedCount: 1,
			inserted:      true,
		},
		{
			name:          "insert into trie existing element",
			value:         "romane",
			expectedCount: 1,
			inserted:      false,
		},
		{
			name:          "insert into trie with one element",
			value:         "romanus",
			expectedCount: 2,
			inserted:      true,
		},
		{
			name:          "insert into trie with two elements",
			value:         "romulus",
			expectedCount: 3,
			inserted:      true,
		},
		{
			name:          "insert into trie with three elements",
			value:         "rubens",
			expectedCount: 4,
			inserted:      true,
		},
		{
			name:          "insert into trie with four elements",
			value:         "ruber",
			expectedCount: 5,
			inserted:      true,
		},
		{
			name:          "insert into trie with five elements",
			value:         "rubicon",
			expectedCount: 6,
			inserted:      true,
		},
		{
			name:          "insert into trie with six elements",
			value:         "rubicundus",
			expectedCount: 7,
			inserted:      true,
		},
	}

	for _, benchmark := range insertBenchmarks {
		inserted := trie.Insert(benchmark.value)
		if inserted != benchmark.inserted {
			b.Errorf("benchmark '%s': expected inserted to be %t", benchmark.name, benchmark.inserted)
		}
		if trie.Count() != benchmark.expectedCount {
			b.Errorf("benchmark '%s': expected count to be %d, but was %d", benchmark.name, benchmark.expectedCount, trie.Count())
		}
	}
}

var benchmarkFound bool
var benchmarkN *Node

func BenchmarkFind(b *testing.B) {

	for _, benchmark := range findTests {
		found, n := benchmark.trie.Find(benchmark.value)
		if found != benchmark.found {
			b.Errorf("benchmark '%s': expected found to be %t", benchmark.name, benchmark.found)
		}
		if found {
			if n.IsLeaf() != benchmark.isLeaf {
				b.Errorf("benchmark '%s': expected isLeaf to be %t", benchmark.name, benchmark.isLeaf)
			}
		}
		benchmarkFound = found
		benchmarkN = n
	}
}
