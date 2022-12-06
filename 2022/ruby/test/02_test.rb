require "minitest/autorun"
require_relative "../helpers"
require_relative "../02"

class TestRockPaperScissors < Minitest::Test
  SAMPLE_INPUT = "A Y
B X
C Z"

  include Helpers

  def setup
    @rps = RockPaperScissors.new(SAMPLE_INPUT)
  end

  def test_find_total_score
    assert_equal(15, @rps.find_total_player_score)
  end

  def test_find_score_from_guide
    assert_equal(12, @rps.find_total_player_score_from_guide)
  end

  # def test_with_input
  #   input = get_input_by_day(2)

  #   rps = RockPaperScissors.new(input)
  #   solution_one = rps.find_total_player_score
  #   solution_two = rps.find_total_player_score_from_guide

  #   puts solution_one
  #   puts solution_two
  # end
end
