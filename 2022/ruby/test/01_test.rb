require "minitest/autorun"
require_relative '../helpers'
require_relative '../01'

class TestFirst < Minitest::Test
  SAMPLE_INPUT = "1000
  2000
  3000

  4000

  5000
  6000

  7000
  8000
  9000

  10000"

  include Helpers

  def setup
    @first = First.new(SAMPLE_INPUT)
  end

  def test_get_most_calories
    assert_equal 24000, @first.get_most_calories
  end

  def test_get_most_calories
    assert_equal 45000, @first.get_top_three_calories
  end

  # def test_with_input
  #   input = get_input_by_day(1)

  #   first = First.new(input)
  #   solution_one = first.get_most_calories
  #   solution_two = first.get_top_three_calories


  #   puts solution_one
  #   puts solution_two
  # end
end
