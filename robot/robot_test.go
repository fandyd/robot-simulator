package robot_test

import (
	. "robot-simulator/robot"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Robot", func() {
	Describe("Invalid Command", func() {
		robot := initialiseRobot()
		_, err := robot.Perform("JUMP")
		It("should return err", func() {
			Expect(err != nil).To(Equal(true))
		})
	})

	Describe("Place Command", func() {
		robot := initialiseRobot()
		Context("with Invalid Direction", func() {
			_, err := robot.Perform("PLACE 0,0,CENTER")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})
		Context("with negative X", func() {
			_, err := robot.Perform("PLACE -1,0,NORTH")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})
		Context("with negative Y", func() {
			_, err := robot.Perform("PLACE 0,-1,NORTH")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("with X over the table", func() {
			_, err := robot.Perform("PLACE 6,0,NORTH")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("with Y over the table", func() {
			_, err := robot.Perform("PLACE 0,6,NORTH")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("with valid param", func() {
			_, err := robot.Perform("PLACE 1,2,NORTH")
			It("should return no error", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.X == 1).To(Equal(true))
				Expect(robot.Y == 2).To(Equal(true))
				Expect(robot.F == "NORTH").To(Equal(true))
			})
		})
	})

	Describe("Move Command", func() {
		Context("without placing the robot on the table", func() {
			robot := initialiseRobot()
			_, err := robot.Perform("MOVE")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Describe("with invalid move", func() {
			Context("when facing north", func() {
				robot := placeRobot(0, 4, "NORTH")
				_, err := robot.Perform("MOVE")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
			Context("when facing south", func() {
				robot := placeRobot(0, 0, "SOUTH")
				_, err := robot.Perform("MOVE")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
			Context("when facing east", func() {
				robot := placeRobot(4, 0, "EAST")
				_, err := robot.Perform("MOVE")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
			Context("when facing west", func() {
				robot := placeRobot(0, 0, "WEST")
				_, err := robot.Perform("MOVE")
				It("should return err", func() {
					Expect(err != nil).To(Equal(true))
				})
			})
		})

		Describe("with valid move", func() {
			Context("when facing north", func() {
				robot := placeRobot(0, 3, "NORTH")
				_, err := robot.Perform("MOVE")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(robot.X == 0).To(Equal(true))
					Expect(robot.Y == 4).To(Equal(true))
					Expect(robot.F == "NORTH").To(Equal(true))
				})
			})

			Context("when facing south", func() {
				robot := placeRobot(0, 1, "SOUTH")
				_, err := robot.Perform("MOVE")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(robot.X == 0).To(Equal(true))
					Expect(robot.Y == 0).To(Equal(true))
					Expect(robot.F == "SOUTH").To(Equal(true))
				})
			})

			Context("when facing east", func() {
				robot := placeRobot(3, 0, "EAST")
				_, err := robot.Perform("MOVE")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(robot.X == 4).To(Equal(true))
					Expect(robot.Y == 0).To(Equal(true))
					Expect(robot.F == "EAST").To(Equal(true))
				})
			})

			Context("when facing west", func() {
				robot := placeRobot(1, 0, "WEST")
				_, err := robot.Perform("MOVE")
				It("should be successful", func() {
					Expect(err == nil).To(Equal(true))
					Expect(robot.X == 0).To(Equal(true))
					Expect(robot.Y == 0).To(Equal(true))
					Expect(robot.F == "WEST").To(Equal(true))
				})
			})
		})
	})

	Describe("Left", func() {
		Context("when robot is not place on the table", func() {
			robot := initialiseRobot()
			_, err := robot.Perform("LEFT")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("when robot facing North", func() {
			robot := placeRobot(0, 0, "NORTH")
			_, err := robot.Perform("LEFT")
			It("should face West", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "WEST").To(Equal(true))
			})
		})

		Context("when robot facing South", func() {
			robot := placeRobot(0, 0, "SOUTH")
			_, err := robot.Perform("LEFT")
			It("should face East", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "EAST").To(Equal(true))
			})
		})

		Context("when robot facing West", func() {
			robot := placeRobot(0, 0, "WEST")
			_, err := robot.Perform("LEFT")
			It("should face South", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "SOUTH").To(Equal(true))
			})
		})

		Context("when robot facing East", func() {
			robot := placeRobot(0, 0, "EAST")
			_, err := robot.Perform("LEFT")
			It("should face NORTH", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "NORTH").To(Equal(true))
			})
		})
	})

	Describe("Right Command", func() {
		Context("when robot is not place on the table", func() {
			robot := initialiseRobot()
			_, err := robot.Perform("RIGHT")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("when robot facing North", func() {
			robot := placeRobot(0, 0, "NORTH")
			_, err := robot.Perform("RIGHT")
			It("should face East", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "EAST").To(Equal(true))
			})
		})

		Context("when robot facing South", func() {
			robot := placeRobot(0, 0, "SOUTH")
			_, err := robot.Perform("RIGHT")
			It("should face West", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "WEST").To(Equal(true))
			})
		})

		Context("when robot facing West", func() {
			robot := placeRobot(0, 0, "WEST")
			_, err := robot.Perform("RIGHT")
			It("should face North", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "NORTH").To(Equal(true))
			})
		})

		Context("when robot facing East", func() {
			robot := placeRobot(0, 0, "EAST")
			_, err := robot.Perform("RIGHT")
			It("should face SOUTH", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.F == "SOUTH").To(Equal(true))
			})
		})
	})

	Describe("Report Command", func() {
		Context("when robot is not place on the table", func() {
			robot := initialiseRobot()
			_, err := robot.Perform("REPORT")
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})

		Context("when robot is not place on the table", func() {
			robot := placeRobot(1, 1, "NORTH")
			report, err := robot.Perform("REPORT")
			It("should return err", func() {
				Expect(err == nil).To(Equal(true))
				Expect(report == "1,1,NORTH").To(Equal(true))
			})
		})
	})

	Describe("Initiate Robot", func() {
		Context("with invalid data", func() {
			_, err := NewRobot(-1, 0)
			It("should return err", func() {
				Expect(err != nil).To(Equal(true))
			})
		})
		Context("with valid data", func() {
			robot, err := NewRobot(1, 3)
			It("should return err", func() {
				Expect(err == nil).To(Equal(true))
				Expect(robot.Table.Width == 1).To(Equal(true))
				Expect(robot.Table.Length == 3).To(Equal(true))
			})
		})
	})
})

func placeRobot(x int, y int, f string) Robot {
	robot := initialiseRobot()
	robot.Place(x, y, f)
	return robot
}

func initialiseRobot() Robot {
	robot, _ := NewRobot(5, 5)
	return robot
}
