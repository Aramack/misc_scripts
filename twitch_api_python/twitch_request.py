class Twitch_Request(object):

  def __init__(self):
    self.base_url = 'https://api.twitch.tv/kraken/'
    self.http_request_header = 'Accept: application/vnd.twitchtv.v3+json'
	
  def get_games_top(self, limit=10, offset=0):
    games_on_twitch = {
        'League Of Legends': 20000,
        'Dark Souls III': 15000,
        'DOTA 2': 10000
      }
    for game in games_on_twitch.keys():
      # If statement:
      if games_on_twitch[game] > 18000:
        # String concatenation with '+':
        print('Most watched ' + game + ':' + str(games_on_twitch[game]))

