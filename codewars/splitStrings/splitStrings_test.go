// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
package splitstrings

var _ = Describe("Split Strings", func() {
	It("Basic tests", func() {
		Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})
