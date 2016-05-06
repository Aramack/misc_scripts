# A list is not an array
games_on_twitch = ['League Of Legends','Dark Souls III', 'DOTA 2']

for game in games_on_twitch:
  # IndentationError if not properly indented
  print(game)
 
 
# slice inclusive_index:exclusive_index
for game in games_on_twitch[0:1]:
  # String concatenation with '+'
  print('Most watched:' + game)