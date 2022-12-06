class Rucksacks
  def initialize(input)
    @input = input.split("\n").map { |s| s.split("") }
  end

  def byte_val_to_priority(s)
    n = s.bytes.at(0)
    n >= 97 ? n - 96 : n - 38
  end

  def sum_priorities
    @input
      .map { |s| s.each_slice(s.length / 2).to_a }
      .sum do |m|
        x = m.at(0)
        y = m.at(1)

        byte_val_to_priority((x & y).first)
      end
  end

  def sum_group_priorities
    sum = 0
    arr = @input.dup

    until arr.empty?
      a, b, c = arr.pop(3)
      sum += byte_val_to_priority((a & b & c).first)
    end

    sum
  end
end
