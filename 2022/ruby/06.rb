class TuningTrouble
  def initialize(input)
    @input = input
  end

  def find_packet_start(eom = 4)
    arr = []
    h = {}
    v = 0
    index = 0

    while arr.length < eom
      if index >= @input.length
        return
      end

      c = @input[index]
      sym = c.to_sym

      if !arr.include?(c) && arr.length == (eom - 1) && h.length == (eom - 1)
        return index + 1
      end

      if arr.length == (eom - 1)
        shifted = arr.shift
        s_sym = shifted.to_sym
        h[s_sym] -= 1
        if h[s_sym] <= 0
          h.delete(s_sym)
        end
      end

      arr.push(c)
      h[sym] ||= 0
      h[sym] += 1

      index += 1
    end

    index
  end
end
