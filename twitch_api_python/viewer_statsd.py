from twitch_request import Twitch_Request

my_request = Twitch_Request()
top_games = my_request.get_games_top(my_request)

print (top_games)
