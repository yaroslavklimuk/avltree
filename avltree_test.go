package avltree

import (
	"fmt"
	"testing"
)

func TestAVLTree_Insert(t1 *testing.T) {
	type fields struct {
		value              int
		rebalanceThreshold int
		left               *AVLTree
		right              *AVLTree
	}
	tests := []struct {
		name      string
		origTree  AVLTree
		newValue  int
		expResult AVLTree
	}{
		{
			name: "case 1",
			origTree: AVLTree{
				value:              4,
				rebalanceThreshold: 1,
				left: &AVLTree{
					value:              3,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              2,
						rebalanceThreshold: 1,
					},
				},
				right: &AVLTree{
					value:              7,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              5,
						rebalanceThreshold: 1,
					},
				},
			},
			newValue: 6,
			expResult: AVLTree{
				value:              4,
				rebalanceThreshold: 1,
				left: &AVLTree{
					value:              3,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              2,
						rebalanceThreshold: 1,
					},
				},
				right: &AVLTree{
					value:              6,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              5,
						rebalanceThreshold: 1,
					},
					right: &AVLTree{
						value:              7,
						rebalanceThreshold: 1,
					},
				},
			},
		},

		{
			name: "case 2",
			origTree: AVLTree{
				value:              4,
				rebalanceThreshold: 1,
				left: &AVLTree{
					value:              3,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              2,
						rebalanceThreshold: 1,
					},
				},
				right: &AVLTree{
					value:              6,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              5,
						rebalanceThreshold: 1,
					},
					right: &AVLTree{
						value:              7,
						rebalanceThreshold: 1,
					},
				},
			},
			newValue: 9,
			expResult: AVLTree{
				value:              6,
				rebalanceThreshold: 1,
				left: &AVLTree{
					value:              4,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              3,
						rebalanceThreshold: 1,
						left: &AVLTree{
							value:              2,
							rebalanceThreshold: 1,
						},
					},
					right: &AVLTree{
						value:              5,
						rebalanceThreshold: 1,
					},
				},
				right: &AVLTree{
					value:              7,
					rebalanceThreshold: 1,
					right: &AVLTree{
						value:              9,
						rebalanceThreshold: 1,
					},
				},
			},
		},

		{
			name: "case 3",
			origTree: AVLTree{
				value:              4,
				rebalanceThreshold: 1,
				left: &AVLTree{
					value:              3,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              2,
						rebalanceThreshold: 1,
					},
				},
				right: &AVLTree{
					value:              6,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              5,
						rebalanceThreshold: 1,
					},
					right: &AVLTree{
						value:              7,
						rebalanceThreshold: 1,
						right: &AVLTree{
							value:              9,
							rebalanceThreshold: 1,
						},
					},
				},
			},
			newValue: 10,
			expResult: AVLTree{
				value:              6,
				rebalanceThreshold: 1,
				left: &AVLTree{
					value:              4,
					rebalanceThreshold: 1,
					left: &AVLTree{
						value:              3,
						rebalanceThreshold: 1,
						left: &AVLTree{
							value:              2,
							rebalanceThreshold: 1,
						},
					},
					right: &AVLTree{
						value:              5,
						rebalanceThreshold: 1,
					},
				},
				right: &AVLTree{
					value:              7,
					rebalanceThreshold: 1,
					right: &AVLTree{
						value:              9,
						rebalanceThreshold: 1,
						right: &AVLTree{
							value:              10,
							rebalanceThreshold: 1,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tt.origTree
			t.Insert(tt.newValue)

			if !TreesAreEqual(t, &tt.expResult) {
				t1.Errorf("The trees are not equal. Want:\n%s\n, got:\n%s\n",
					showTree(&tt.expResult, 0), showTree(t, 0))
			}
		})
	}
}

func showTree(tr *AVLTree, level int) string {
	result := fmt.Sprintf("v:%d\n", tr.value)
	if tr.left != nil {
		result += fmt.Sprintf("-")
		for i := 0; i <= level; i++ {
			result += fmt.Sprintf("-")
		}
		result += fmt.Sprintf("left:")
		result += showTree(tr.left, level+2)
	}
	if tr.right != nil {
		result += fmt.Sprintf("-")
		for i := 0; i <= level; i++ {
			result += fmt.Sprintf("-")
		}
		result += fmt.Sprintf("right:")
		result += showTree(tr.right, level+2)
	}
	return result
}

func TestTreesAreEqual(t *testing.T) {
	type args struct {
		one *AVLTree
		two *AVLTree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				one: &AVLTree{
					value: 7,
					left: &AVLTree{
						value: 5,
					},
				},
				two: &AVLTree{
					value: 7,
					left: &AVLTree{
						value: 5,
					},
				},
			},
			want: true,
		},

		{
			name: "case 2",
			args: args{
				one: &AVLTree{
					value: 4,
					left: &AVLTree{
						value: 3,
						left: &AVLTree{
							value: 2,
						},
					},
					right: &AVLTree{
						value: 6,
						left: &AVLTree{
							value: 5,
						},
						right: &AVLTree{
							value: 7,
							right: &AVLTree{
								value: 9,
							},
						},
					},
				},
				two: &AVLTree{
					value: 4,
					left: &AVLTree{
						value: 3,
						left: &AVLTree{
							value: 2,
						},
					},
					right: &AVLTree{
						value: 6,
						left: &AVLTree{
							value: 5,
						},
						right: &AVLTree{
							value: 7,
							right: &AVLTree{
								value: 9,
							},
						},
					},
				},
			},
			want: true,
		},

		{
			name: "case 3",
			args: args{
				one: &AVLTree{
					value: 4,
					left: &AVLTree{
						value: 3,
						left: &AVLTree{
							value: 2,
						},
					},
					right: &AVLTree{
						value: 6,
						left: &AVLTree{
							value: 5,
						},
						right: &AVLTree{
							value: 7,
							right: &AVLTree{
								value: 9,
							},
						},
					},
				},
				two: &AVLTree{
					value: 4,
					left: &AVLTree{
						value: 3,
						left: &AVLTree{
							value: 2,
							left: &AVLTree{
								value: 1,
							},
						},
					},
					right: &AVLTree{
						value: 6,
						left: &AVLTree{
							value: 5,
						},
						right: &AVLTree{
							value: 7,
							right: &AVLTree{
								value: 9,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreesAreEqual(tt.args.one, tt.args.two); got != tt.want {
				t.Errorf("TreesAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
