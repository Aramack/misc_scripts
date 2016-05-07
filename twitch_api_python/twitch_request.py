# A list is not an array
list_tuple = [
    # A tuple, is an immutable list?
    ('League Of Legends', 20000),
    ('Dark Souls III', 15000),
    ('DOTA 2', 10000)
  ]


games_on_twitch = {
    'League Of Legends': 20000,
    'Dark Souls III': 15000,
    'DOTA 2': 10000
  }
  
for game in list_tuple:
  # IndentationError if not properly indented
  print(game[0] + ':' + str(game[1]))
 
 
# slice inclusive_index:exclusive_index
for game in games_on_twitch.keys():
  # If statement:
  if games_on_twitch[game] > 18000:
    # String concatenation with '+':
    print('Most watched ' + game + ':' + str(games_on_twitch[game]))