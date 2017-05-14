from StringIO import StringIO   
import pycurl
import json

class nasa_request(object):

  def __init__(self):
    self.base_url = 'https://www.nasa.gov'
    
  def query(
      self,
      url, 
      http_request_type = 'GET',
      request_params = {}
    ):
    curl = pycurl.Curl()
    curl.setopt(curl.URL, url)
    if http_request_type == 'GET':
      for key, value in request_params.iteritems():
        url += key + '&' + value 
    elif http_request_type == 'POST':
      curl.setopt(pycurl.POST, 1)
      curl.setopt(curl.POSTFIELDS, json.dump(request_params))
    response = StringIO()
    curl.setopt(curl.WRITEFUNCTION, response.write)
    curl.perform()
    curl.close()
    return json.loads(response.getvalue())
    
  #https://www.nasa.gov/api/1/query/ubernodes.json?unType%5B%5D=image&routes%5B%5D=1446&page=0&pageSize=2
  def get_ubernodes(self):
    target_url = self.base_url +  '/api/1/query/ubernodes.json?'
    params = {
        'unType[]': 'image',
        'routes[]': '1446',
        'page':     '0',
        'pageSize': '24'
      }
    ubernodes = self.query(
        target_url,
        'GET',
        params
      )
    return ubernodes['ubernodes']
  
  #EG: https://www.nasa.gov/api/1/record/node/401700.json
  def get_record_node(self, node = 401700):
    target_url = self.base_url + '/api/1/record/node/' + str(node) +'.json'
    return self.query(target_url, 'GET')
    
  #Get the url for the latest image of the day from nasa.
  def get_latest_iotd(self):
    ubernodes = self.get_ubernodes()
    latest = ubernodes[0]['nid']
    node = self.get_record_node(latest)
    return self.base_url + node['images'][0]['fullWidthFeature']
    
nasa = nasa_request()
print(nasa.get_latest_iotd())