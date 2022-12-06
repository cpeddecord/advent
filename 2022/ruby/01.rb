class First
  def initialize(input)
    @input = []

    elf_calories = 0
    # add a trailing 0 to deal with EOF
    input.split("\n").push("0").each do |s|
      n = s.to_i
      if n == 0
        @input.push(elf_calories)
        elf_calories = 0
      end
      elf_calories += n
    end
  end

  def get_most_calories
    @input.max
  end

  def get_top_three_calories
    @input.max(3).sum
  end
end
