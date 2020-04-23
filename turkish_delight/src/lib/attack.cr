require "http/client"

module TurkishAttack
  referee_url = "https://enqfc8y2t9fo.x.pipedream.net"
  referee_won_url = referee_url + "/won"

  opponent_url = "https://enqfc8y2t9fo.x.pipedream.net"
  opponent_status_url = opponent_url + "/status"
  opponent_jab_url = opponent_url + "/jab"
  opponent_cross_url = opponent_url + "/cross"
  opponent_hook_url = opponent_url + "/hook"
  opponent_uppercut_url = opponent_url + "/uppercut"

  def self.attack(url : String)
    # TODO
    true
  end

  def self.attack_jab
    attack opponent_jab_url
  end

  def self.attack_cross
    attack opponent_cross_url
  end

  def self.attack_hook
    attack opponent_hook_url
  end

  def self.attack_uppercut
    attack opponent_uppercut_url
  end

  def self.notify_referee_won
    HTTP::Client.post(referee_won_url, body: { name: "Turkish Delight" }.to_json)
  end
end
