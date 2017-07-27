package robot_test

import (
	. "robot-simulator/robot"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table", func() {
	Describe("Initiate Table", func() {
		Context("with invalid data", func() {
			_, err := NewTable(-1, 0)
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})
		Context("with valid data", func() {
			table, err := NewTable(1, 3)
			It("should return err", func() {
				Expect(err == nil).To(Equal(true))
				Expect(table.Width == 1).To(Equal(true))
				Expect(table.Length == 3).To(Equal(true))
			})
		})
	})
})
