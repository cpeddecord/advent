require "minitest/autorun"
require_relative "../helpers"
require_relative "../05"

class TestSupplyStacks < Minitest::Test
  SAMPLE_INPUT = "    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2"

  include Helpers

  def setup
    @stacks = SupplyStacks.new(SAMPLE_INPUT)
  end

  def test_find_top_of_stacks
    assert_equal("CMZ", @stacks.find_top_of_stacks)
  end

  def test_find_top_of_stacks_without_reverse
    assert_equal("MCD", @stacks.find_top_of_stacks_without_reverse)
  end

  # def test_with_input
  #   input = get_input_by_day(5)

  #   stacks = SupplyStacks.new(input)
  #   solution_one = stacks.find_top_of_stacks
  #   solution_two = stacks.find_top_of_stacks_without_reverse

  #   puts solution_one
  #   puts solution_two
  # end
end
