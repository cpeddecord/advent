require "minitest/autorun"
require_relative "../helpers"
require_relative "../06"

class TestTuningTrouble < Minitest::Test
  SAMPLE_INPUT = [
    ["mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19],
    ["bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23],
    ["nppdvjthqldpwncqszvftbrmjlhg", 6, 23],
    ["nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29],
    ["zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26],
  ]

  include Helpers

  def test_find_packet_start
    SAMPLE_INPUT.each do |i|
      tuner = TuningTrouble.new(i.at(0))
      assert_equal(i.at(1), tuner.find_packet_start)
    end
  end

  def test_find_packet_message
    SAMPLE_INPUT.each do |i|
      tuner = TuningTrouble.new(i.at(0))
      assert_equal(i.at(2), tuner.find_packet_start(14))
    end
  end

  def test_with_input
    input = get_input_by_day(6)

    tuner = TuningTrouble.new(input)
    solution_one = tuner.find_packet_start
    solution_two = tuner.find_packet_start(14)

    puts solution_one
    puts solution_two
  end
end
