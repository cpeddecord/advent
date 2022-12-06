require "minitest/autorun"
require_relative "../helpers"
require_relative "../04"

class TestCampCleanup < Minitest::Test
  SAMPLE_INPUT = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8"

  include Helpers

  def setup
    @camp_cleanup = CampCleanup.new(SAMPLE_INPUT)
  end

  def test_find_full_intersections
    assert_equal(2, @camp_cleanup.find_full_intersections)
  end

  def test_find_any_intersections
    assert_equal(4, @camp_cleanup.find_any_intersections)
  end

  # def test_with_input
  #   input = get_input_by_day(4)

  #   camp = CampCleanup.new(input)
  #   solution_one = camp.find_full_intersections
  #   solution_two = camp.find_any_intersections

  #   puts solution_one
  #   puts solution_two
  # end
end
