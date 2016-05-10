import pycurl

class Twitch_Request(object):

  def __init__(self):
    self.base_url = 'https://api.twitch.tv/kraken/'
    self.http_request_header = 'Accept: application/vnd.twitchtv.v3+json'

  def query(
	  url, 
      http_request_type = 'GET',
      request_params = ''
    ):
	my_curl = pycurl.Curl()
	my_curl.setopt(my_curl.URL, url)
	if http_request_type == 'GET':
	  url += '?' + request_params
	elif http_request_type == 'POST':
	  my_curl.setopt(my_curl.POSTFIELDS, request_params)
	my_curl.perform()
	my_curl.close()
	
	
  def get_games_top(self, limit=10, offset=0):
    self.query(
	    self.base_url + '/games/top',
        'GET',
        'limit=' + limit + '&offset=' + offset
	  )
