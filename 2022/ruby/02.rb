class GameRound
  @player_x = 0
  @player_y = 0

  def initialize(x, y)
    score_by_play = { r: 1, p: 2, s: 3 }
    score_by_game = { w: 6, d: 3, l: 0 }
    tree = {
      r: { p: :l, s: :w },
      p: { r: :w, s: :l },
      s: { r: :l, p: :w },
    }

    @player_x ||= 0
    @player_y ||= 0

    @player_x += score_by_play[x]
    @player_y += score_by_play[y]

    if x == y
      @player_x += score_by_game[:d]
      @player_y += score_by_game[:d]
      return
    end

    @player_x += score_by_game[tree[x][y]]
    @player_y += score_by_game[tree[y][x]]
  end

  def scores
    [@player_x, @player_y]
  end
end

class RockPaperScissors
  def initialize(input)
    symbols = { "A": :r, "B": :p, "C": :s, "X": :r, "Y": :p, "Z": :s }
    guide = { "X": :l, "Y": :d, "Z": :w }
    guide_to_symbol = {
      r: { w: :s, l: :p, d: :r },
      p: { w: :s, l: :r, d: :p },
      s: { w: :r, l: :p, d: :s },
    }

    @input = input.split("\n").map do |s|
      splits = s.split("\s")
      x = splits.at(0).to_sym
      y = splits.at(1).to_sym

      player_x = symbols[x]
      player_y = symbols[y]
      player_y_from_guide = guide_to_symbol[player_x][guide[y]]

      [player_x, player_y, player_y_from_guide]
    end
  end

  def find_total_player_score
    @input.sum do |input|
      GameRound.new(input.at(0), input.at(1))
        .scores.at(1)
    end
  end

  def find_total_player_score_from_guide
    @input.sum do |input|
      GameRound.new(input.at(0), input.at(2))
        .scores.at(1)
    end
  end
end
