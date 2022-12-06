class SupplyStacks
  def initialize(input)
    @stacks = []
    input.split("\n\n").at(0).split("\n").each do |l|
      foo = l.chars.each_slice(4).map(&:join).map do |s|
        /[A-Z]/.match(s)
      end

      foo.each_with_index do |a, i|
        @stacks[i] ||= []
        unless a.nil?
          @stacks[i].push(a[0])
        end
      end
    end

    @stacks.each_with_index { |a, i| @stacks[i] = a.reverse }

    @instructions = input.split("\n\n").at(1).split("\n").map do |l|
      l.scan(/\d+/).each_with_index.map { |s, i| i == 0 ? s.to_i : s.to_i - 1 }
    end
  end

  def find_top_of_stacks
    st = @stacks.map(&:dup)

    @instructions.each do |t|
      quant, from, dest = t

      popped = st[from].pop(quant).reverse
      st[dest].push(*popped)
    end

    st.map(&:last).join("")
  end

  def find_top_of_stacks_without_reverse
    st = @stacks.map(&:dup)

    @instructions.each do |t|
      quant, from, dest = t

      popped = st[from].pop(quant)
      st[dest].push(*popped)
    end

    st.map(&:last).join("")
  end
end
