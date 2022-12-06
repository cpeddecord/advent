class CampCleanup
  def initialize(input)
    @input = input
      .split("\n")
      .map do |line|
        line
          .split(",")
          .map do |sect|
            a, b = sect.split("-")
            (a..b).to_a
          end
      end
  end

  def find_full_intersections
    @input.sum do |s|
      a, b = s

      intersection = a & b

      intersection == a || intersection == b ? 1 : 0
    end
  end

  def find_any_intersections
    @input.sum do |s|
      a, b = s

      intersection = a & b

      intersection.empty? ? 0 : 1
    end
  end
end
