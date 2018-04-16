package trie

import "testing"

func TestInsert(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		trie          Trie
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
			name:          "insert into empty trie",
			value:         "romane",
			expectedCount: 1,
			inserted:      true,
		},
		{
			name:  "insert into trie existing element",
			value: "romane",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{
					{value: "omane", childCount: 0}},
				childCount: 1}}, count: 1},
			expectedCount: 1,
			inserted:      false,
		},
		{
			name:  "insert into trie with one element",
			value: "romanus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{
					{value: "omane", childCount: 0}},
				childCount: 1}}, count: 1},
			expectedCount: 2,
			inserted:      true,
		},
		{
			name:  "insert into trie with two elements",
			value: "romulus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "oman",
					children: []*Node{{value: "e", childCount: 0},
						{value: "us", childCount: 0}},
					childCount: 2}}, childCount: 1}}, count: 2},
			expectedCount: 3,
			inserted:      true,
		},
		{
			name:  "insert into trie with three elements",
			value: "rubens",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}}}, childCount: 1}}, count: 3},
			expectedCount: 4,
			inserted:      true,
		},
		{
			name:  "insert into trie with four elements",
			value: "ruber",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}},
					{value: "ubens", childCount: 0}}, childCount: 1}}, count: 4},
			expectedCount: 5,
			inserted:      true,
		},
		{
			name:  "insert into trie with five elements",
			value: "rubicon",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}},
					{value: "ube",
						children: []*Node{{value: "ns", childCount: 0},
							{value: "r", childCount: 0}},
						childCount: 2}}, childCount: 1}}, count: 5},
			expectedCount: 6,
			inserted:      true,
		},
		{
			name:  "insert into trie with six elements",
			value: "rubicundus",
			trie: Trie{child: []*Node{{
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
				childCount: 1}}, count: 6},
			expectedCount: 7,
			inserted:      true,
		},
	}

	for _, test := range tests {
		inserted := test.trie.Insert(test.value)
		if inserted != test.inserted {
			t.Errorf("test '%s': expected inserted to be %t", test.name, test.inserted)
		}
		if test.trie.Count() != test.expectedCount {
			t.Errorf("test '%s': expected count to be %d, but was %d", test.name, test.expectedCount, test.trie.Count())
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		trie   Trie
		found  bool
		isLeaf bool
	}{
		{
			name:  "find nonexistent element in trie",
			value: "romanus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{
					{value: "omane", childCount: 0}},
				childCount: 1}}, count: 1},
			found:  false,
			isLeaf: false,
		},
		{
			name:  "find existing element in trie",
			value: "romane",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{
					{value: "omane", childCount: 0}},
				childCount: 1}}, count: 1},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find trimmed existing element in trie",
			value: " romane\n",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{
					{value: "omane", childCount: 0}},
				childCount: 1}}, count: 1},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing element in trie with two elements",
			value: "romane",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "oman",
					children: []*Node{{value: "e", childCount: 0},
						{value: "us", childCount: 0}},
					childCount: 2}}, childCount: 1}}, count: 2},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing second element in trie with two elements",
			value: "romanus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "oman",
					children: []*Node{{value: "e", childCount: 0},
						{value: "us", childCount: 0}},
					childCount: 2}}, childCount: 1}}, count: 2},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing element in trie with three elements",
			value: "romane",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2}}, childCount: 1}}, count: 3},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing second element in trie with three elements",
			value: "romanus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2}}, childCount: 1}}, count: 3},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing third element in trie with three elements",
			value: "romulus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}, childCount: 2}}, childCount: 1}}, count: 3},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing fourth element in trie with four elements",
			value: "rubens",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}},
					{value: "ubens", childCount: 0}}, childCount: 1}}, count: 4},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing fifth element in trie with five elements",
			value: "ruber",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}},
					{value: "ube",
						children: []*Node{{value: "ns", childCount: 0},
							{value: "r", childCount: 0}},
						childCount: 2}}, childCount: 1}}, count: 5},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing sixth element in trie with six elements",
			value: "rubicon",
			trie: Trie{child: []*Node{{
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
				childCount: 1}}, count: 6},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing sixth element in trie with seven elements",
			value: "rubicon",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}},
					{value: "ub",
						children: []*Node{{value: "e",
							children: []*Node{{value: "ns", childCount: 0},
								{value: "r", childCount: 0}}, childCount: 2},
							{value: "ic",
								children: []*Node{{value: "on", childCount: 0},
									{value: "undus", childCount: 0}},
								childCount: 2}},
						childCount: 2}}, childCount: 1}}, count: 7},
			found:  true,
			isLeaf: true,
		},
		{
			name:  "find existing seventh element in trie with seven elements",
			value: "rubicundus",
			trie: Trie{child: []*Node{{
				value: "r",
				children: []*Node{{value: "om",
					children: []*Node{{value: "an",
						children: []*Node{{value: "e", childCount: 0},
							{value: "us", childCount: 0}},
						childCount: 2},
						{value: "ulus", childCount: 0}}},
					{value: "ub",
						children: []*Node{{value: "e",
							children: []*Node{{value: "ns", childCount: 0},
								{value: "r", childCount: 0}}, childCount: 2},
							{value: "ic",
								children: []*Node{{value: "on", childCount: 0},
									{value: "undus", childCount: 0}},
								childCount: 2}},
						childCount: 2}}, childCount: 1}}, count: 7},
			found:  true,
			isLeaf: true,
		},
	}

	for _, test := range tests {
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
