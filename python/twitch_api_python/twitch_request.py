from StringIO import StringIO   
import pycurl
import json

class Twitch_Request(object):

  def __init__(self):
    self.base_url = 'https://api.twitch.tv/kraken/'
    self.http_request_header = 'Accept: application/vnd.twitchtv.v3+json'

  def query(
      self,
      url, 
      http_request_type = 'GET',
      request_params = ''
    ):
    my_curl = pycurl.Curl()
    my_curl.setopt(my_curl.URL, url)
    if http_request_type == 'GET':
      url += '?' + request_params
    elif http_request_type == 'POST':
      my_curl.setopt(pycurl.POST, 1)
      my_curl.setopt(my_curl.POSTFIELDS, request_params)
    response = StringIO()
    my_curl.setopt(my_curl.WRITEFUNCTION, response.write)
    my_curl.perform()
    my_curl.close()
    return json.loads(response.getvalue())
	
	
  def get_games_top(self, limit=10, offset=0):
    pagination_params = 'limit=' + str(limit) + '&offset=' + str(offset)
    api_endpoint = self.base_url + 'games/top'
    return self.query(
        api_endpoint,
        'GET',
        pagination_params
      )
