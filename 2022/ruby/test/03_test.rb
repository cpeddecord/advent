require "minitest/autorun"
require_relative "../helpers"
require_relative "../03"

class TestRucksacks < Minitest::Test
  SAMPLE_INPUT = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw"

  include Helpers

  def setup
    @sacks = Rucksacks.new(SAMPLE_INPUT)
  end

  def test_byte_val_to_priority
    results = ["a", "A", "z", "Z"].map do |s|
      @sacks.byte_val_to_priority(s)
    end

    assert_equal([1, 27, 26, 52], results)
  end

  def test_sum_priorities
    assert_equal(157, @sacks.sum_priorities)
  end

  def test_sum_group_priorities
    assert_equal(70, @sacks.sum_group_priorities)
  end

  # def test_with_input
  #   input = get_input_by_day(3)

  #   sacks = Rucksacks.new(input)
  #   solution_one = sacks.sum_priorities
  #   solution_two = sacks.sum_group_priorities

  #   puts solution_one
  #   puts solution_two
  # end
end
