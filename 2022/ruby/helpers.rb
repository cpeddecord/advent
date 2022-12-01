require 'dotenv/load'
require 'uri'
require 'faraday'

module Helpers
  def get_input_by_day(day, year = 2022)
    uri = URI("https://adventofcode.com/")

    conn = Faraday.new(
      url: uri,
      headers:  {"cookie": "session=#{ENV['SESSION_TOKEN']}"}
    )

    res = conn.get("#{year}/day/#{day}/input")

    res.body
  end
end
