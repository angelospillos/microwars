require "kemal"
require "uuid"
require "uuid/json"
require "http/client"

module TurkishDelight

  logging false

  referee_url = "https://enqfc8y2t9fo.x.pipedream.net"
  referee_won_url = referee_url + "/won"

  opponent_url = "https://enqfc8y2t9fo.x.pipedream.net"
  opponent_status_url = opponent_url + "/status"
  opponent_jab_url = opponent_url + "/jab"
  opponent_cross_url = opponent_url + "/cross"
  opponent_hook_url = opponent_url + "/hook"
  opponent_uppercut_url = opponent_url + "/uppercut"

  lib CONSTANTS
    STATUS_RESPONSE = {"status": "ok"}.to_json
  end

  def fibonacci(fib_number : Int32)
    if fib_number <= 1
      return fib_number
    end
    fibonacci(fib_number-2) + fibonacci(fib_number-1)
  end

  def attack_response(fib_number : Int32)
    { "uuid": UUID.random, "fib": fibonacci(fib_number) }.to_json
  end

  def attack(url : String)
    # TODO
    true
  end

  def attack_jab
    attack opponent_jab_url
  end

  def attack_cross
    attack opponent_cross_url
  end

  def attack_hook
    attack opponent_hook_url
  end

  def attack_uppercut
    attack opponent_uppercut_url
  end

  def notify_referee_won
    HTTP::Client.post(referee_won_url, body: { name: "Turkish Delight" }.to_json)
  end

  def connection_is_ok : Bool
    true
  end

  get "/status" do |env|
    env.response.content_type = "application/json"
    CONSTANTS::STATUS_RESPONSE
  end

  get "/test" do |env|
    env.response.content_type = "application/json"

    # return CONSTANTS::STATUS_RESPONSE if true

    { "status": "not connected" }.to_json
  end

  get "/combat" do |env|
    env.response.content_type = "application/json"
    CONSTANTS::STATUS_RESPONSE
  end

  get "/jab" do |env|
    env.response.content_type = "application/json"
    attack_jab
    # attack_jab

    attack_response(2)
  end

  get "/cross" do |env|
    env.response.content_type = "application/json"
    # attack_jab
    # attack_jab
    # attack_cross

    attack_response(4)
  end

  get "/hook" do |env|
    env.response.content_type = "application/json"
    # attack_hook
    # attack_hook
    # attack_uppercut

    attack_response(8)
  end

  get "/uppercut" do |env|
    env.response.content_type = "application/json"
    # attack_cross
    # attack_hook
    # attack_uppercut

    attack_response(16)
  end

  Kemal.run
end
