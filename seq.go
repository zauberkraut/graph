/* Functions over graphic sequences. */
package graph

import "sort"

type DegSeq sort.IntSlice

/* Decides whether a degree sequence is graphic in linear time using the
   Erdős–Gallai characterization. The behavior is undefined if the sequence
   contains negative degrees or isn't in nonincreasing order.

   Returns true iff d is even and satisfies the following for all 1 <= k <= n:
   \sum_{i=1}^k d_i \leq k(k-1) + \sum_{j=k+1}^n \text{min}\{d_j, k\} */
func (d DegSeq) Graphic() bool {
	maybe := true
	n := len(d)

	var lhs, rhs int
	for k, j, dProd := 1, n-1, 0; k <= n; k, dProd = k+1, dProd+2 {
		i := k - 1 // the index of d_i
		deg := d[i]
		lhs += deg   // fold the next degree into the sum on the left-hand side
		rhs += dProd // handle the right-hand product term: (k+1)k - k(k-1) = 2k

		switch { // grow and trim the sum on the right-hand side
		case j > i: // still growing
			for ; j > i && k > d[j]; j-- { // find the next slice of area
			}
			rhs += j - i // add the slice to the sum
			fallthrough
		case j == i: // in case i and j meet at a degree greater than k...
			rhs -= i // don't subtract more than we ever added to the sum
		default: // downhill
			rhs -= deg
		}

		if lhs > rhs {
			maybe = false
			break
		}
	}

	// lhs now equals the degree sum; reject it if it's odd
	return maybe && lhs&1 == 0
}
