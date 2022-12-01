require "minitest/autorun"
require_relative '../00'

class TestHello < Minitest::Test
  def setup
    @hello = Hello.new
  end

  def test_say_hi
    assert_equal "hi", @hello.say_hi()
  end

end
