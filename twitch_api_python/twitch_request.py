# A list is not an array
games_on_twitch = [
    # A tuple:
    ('League Of Legends', 20000),
    ('Dark Souls III', 15000),
    ('DOTA 2', 10000)
  ]

for game in games_on_twitch:
  # IndentationError if not properly indented
  print(game[0] + ':' + str(game[1]))
 
 
# slice inclusive_index:exclusive_index
for game in games_on_twitch[0:2]:
  # If statement:
  if game[1] > 18000:
    # String concatenation with '+':
    print('Most watched:' + str(game[0]))