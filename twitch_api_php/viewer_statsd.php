<?php
 
require_once './twitch_request.php';
require_once './statsd_helper.php';

$request = new twitch_request();
$statsd = \statsd_helper::get_instance();
$games_on_twitch = $request->get_games_top();

foreach($games_on_twitch->top as $game) {
  echo($game->game->name . ':' . $game->viewers . PHP_EOL);
  $statsd->send_message('top10.' . $game->game->name, $game->viewer);
}


?>