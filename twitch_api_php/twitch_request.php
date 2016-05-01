<?php
class twitch_request {
  /** 
   * The url of the twitch api
   *
   * @var string
   */   
  private $base_url;
  
  /** 
   * The http_request_header
   *
   * @var string
   */   
  private $http_request_header; 

  /** 
   * The constructor for a twitch_request object
   *
   * @return void
   */   
  public function __construct() {
    $this->base_url = 'https://api.twitch.tv/kraken/';
    $this->http_request_header = 'Accept: application/vnd.twitchtv.v3+json';
  }

  /**
   * Private query function to curl the twitch api
   *
   * @param $url               The restful endpoint to curl
   * @param $http_request_type The http request type (GET, POST, PUT, etc...)
   * @param $request_params    The request parameters
   *
   * @return array 
   */
  private function query(
      $url, 
      $http_request_type = 'GET',
      $request_params = ''
    ) {
    $curl_base = curl_init();
    switch($http_request_type) {
      case 'GET':
        $url .= '?' . $request_params;
        break;
      case 'POST':
        curl_setopt($curl_base, CURLOPT_POST, true);
        curl_setopt($curl_base, CURLOPT_POSTFIELDS, true);
        break;
      case 'PUT':
        curl_setopt($curl_base, CURLOPT_PUT, true);
        break;
      default:
        //No http request type specified.
        //Return an empty array to fullfil the contract.
        return [];
    }
    curl_setopt($curl_base, CURLOPT_URL, $url);	
    curl_setopt($curl_base, CURLOPT_HEADER, $this->http_request_header);
    curl_setopt($curl_base, CURLOPT_RETURNTRANSFER, TRUE);

    $result = json_decode(curl_exec($curl_base));
    curl_close($curl_base);
    return $result;
  }
  
  /**
   * Get the top games being streamed on twitch
   * https://github.com/justintv/Twitch-API/blob/master/v3_resources/games.md#get-gamestop
   *
   * @param $limit  the number of results
   * @param $offset the page offset ($limit*offset TO ($limit+1)*offset) 
   *
   * @return array
   */
  public function get_games_top($limit = 10, $offset = 0) {
    return $this->query(
        $this->base_url . '/games/top',
        'GET',
        'limit=' . $limit . '&offset=' . $offset
      );
  }

  /**@todo implement the rest of the API request functions.*/
}


?>