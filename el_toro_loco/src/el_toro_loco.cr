require "toro"
require "uuid"
require "uuid/json"
require "http/client"

lib CONSTANTS
  STATUS_RESPONSE = {"status": "ok"}
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
  HTTP::Client.post(CONSTANTS::REFEREE_WON_URL, body: { name: "El Toro Loco", date: Time.utc }.to_json)
end

def fibonacci(fib_number : Int32)
  if fib_number <= 1
    return fib_number
  end
  fibonacci(fib_number-2) + fibonacci(fib_number-1)
end

def attack_response(fib_number : Int32)
  { "uuid": UUID.random, "fib": fibonacci(fib_number) }
end

def connection_is_ok : Bool
  HTTP::Client.get(CONSTANTS::REFEREE_URL).status_code == 200 && HTTP::Client.get(CONSTANTS::OPPONENT_STATUS_URL).status_code == 200
end

class App < Toro::Router
  def routes
    on "status" do
      get do
        json CONSTANTS::STATUS_RESPONSE
      end
    end
    on "test" do
      get do
        result = CONSTANTS::STATUS_RESPONSE

        result = { "status": "not connected" } unless connection_is_ok

        json result
      end
    end
    on "combat" do
      get do
        attack_jab
        attack_hook

        json CONSTANTS::STATUS_RESPONSE
      end
    end
    on "jab" do
      get do
        attack_jab
        attack_jab

        json attack_response(2)
      end
    end
    on "cross" do
      get do
        attack_jab
        attack_jab
        attack_cross

        json attack_response(4)
      end
    end
    on "hook" do
      get do
        attack_hook
        attack_hook
        attack_uppercut

        json attack_response(8)
      end
    end
    on "uppercut" do
      get do
        attack_cross
        attack_hook
        attack_uppercut

        json attack_response(16)
      end
    end
    on "warmup" do
      get do
        json attack_response(16)
      end
    end
  end
end

System.cpu_count.times do |i|
  Process.fork do
    App.run 3000 do |server|
      server.bind_tcp("0.0.0.0", 8001, true)
      server.listen
    end
  end
end

sleep
