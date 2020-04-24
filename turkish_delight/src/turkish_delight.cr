require "kemal"
require "uuid"
require "uuid/json"
require "http/client"

logging false
Kemal.config do |cfg|
  cfg.env = "production"
  cfg.serve_static = false
  cfg.logging = false
end

lib CONSTANTS
  STATUS_RESPONSE = {"status": "ok"}.to_json
  REFEREE_URL = ENV["REFEREE_URL"]
  REFEREE_WON_URL = REFEREE_URL + "/won"

  OPPONENT_URL = URI.parse(ENV["OPPONENT_URL"])
  OPPONENT_STATUS_URL = URI.parse(ENV["OPPONENT_URL"] + "/status")
  OPPONENT_JAB_PATH = "/jab"
  OPPONENT_CROSS_PATH = "/cross"
  OPPONENT_HOOK_PATH = "/hook"
  OPPONENT_UPPERCUT_PATH = "/uppercut"

  OPPONENT_HTTP_CLIENT = HTTP::Client.new OPPONENT_URL
end

CONSTANTS::OPPONENT_HTTP_CLIENT.connect_timeout = 4.seconds

def attack(path : String)
  spawn do
    begin
      CONSTANTS::OPPONENT_HTTP_CLIENT.get(path)
    rescue ex : IO::TimeoutError
      notify_referee_won
      puts "TIMEOUT! WON"
      puts ex
    end
  end
end

def attack_jab
  attack CONSTANTS::OPPONENT_JAB_PATH
end

def attack_cross
  attack CONSTANTS::OPPONENT_CROSS_PATH
end

def attack_hook
  attack CONSTANTS::OPPONENT_HOOK_PATH
end

def attack_uppercut
  attack CONSTANTS::OPPONENT_UPPERCUT_PATH
end

def notify_referee_won
  HTTP::Client.post(CONSTANTS::REFEREE_WON_URL, body: { name: "Turkish Delight", date: Time.utc }.to_json)
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

def connection_is_ok : Bool
  HTTP::Client.get(CONSTANTS::REFEREE_URL).status_code == 200 && HTTP::Client.get(CONSTANTS::OPPONENT_STATUS_URL).status_code == 200
end

get "/status" do |env|
  env.response.content_type = "application/json"
  CONSTANTS::STATUS_RESPONSE
end

get "/test" do |env|
  env.response.content_type = "application/json"
  result = CONSTANTS::STATUS_RESPONSE

  result = { "status": "not connected" }.to_json unless connection_is_ok

  result
end

get "/combat" do |env|
  env.response.content_type = "application/json"
  CONSTANTS::STATUS_RESPONSE
end

get "/jab" do |env|
  env.response.content_type = "application/json"
  attack_jab
  attack_jab

  attack_response(2)
end

get "/cross" do |env|
  env.response.content_type = "application/json"
  attack_jab
  attack_jab
  attack_cross

  attack_response(4)
end

get "/hook" do |env|
  env.response.content_type = "application/json"
  attack_hook
  attack_hook
  attack_uppercut

  attack_response(8)
end

get "/uppercut" do |env|
  env.response.content_type = "application/json"
  attack_cross
  attack_hook
  attack_uppercut

  attack_response(16)
end

get "/warmup" do |env|
  env.response.content_type = "application/json"

  attack_response(16)
end

System.cpu_count.times do |i|
  Process.fork do
    Kemal.run do |config|
      server = config.server.not_nil!
      server.bind_tcp "0.0.0.0", 3000, reuse_port: true
    end
  end
end

sleep
