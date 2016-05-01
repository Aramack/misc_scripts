<?php
/**
 * A helper class for sending UDP packets to an open statsd connection.
 */
class statsd_helper {
  /**
   * The url of the statsd host
   *
   * @var string
   */
   private $url = 'localhost';
   
   /**
    * The port the statsd host is listening on
    *
    * @var int 
    */
   private $port = 8125;
   
  /**
   * Construct a statsd_helper object
   *
   * @param $url  The url of the statsd device
   * @param $port The port of the statsd interface is listening on
   *
   * @return void
   */
  private function __construct($url, $port) {
    $this->url = $url;
    $this->port = $port;
  }
  
  /**
   * Get an instance of a statsd_helper
   *
   * @return \statsd_helper;
   */
  public static function get_instance() {
    return new statsd_helper('localhost', 8125);
  }
  /**
   * send a message to the statsd port
   *
   * @param $message string period delimited statsd metric
   * @param $value   int    the value of this metric
   *
   * @return void
   */ 
  public function send_metric($message, $value) {
    $message_body = $message . ':' . $value . '|c';
    $udp_socket = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);
	
    socket_sendto(
        $udp_socket,
        $message_body,
        strlen($message_body),
        0,
	    $this->url,
		$this->port
	  )
	
	socket_close($udp_socket);
  }
?>