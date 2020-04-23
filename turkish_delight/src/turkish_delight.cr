require "kemal"
require "uuid"
require "uuid/json"

# require "./lib/attack"

logging false

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
  # TurkishAttack.attack_jab
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
