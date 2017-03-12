import random

Suit=['Clubs', 'Diamonds', 'Hearts', 'Spades']
  
Rank = {
  'Ace': 1,
  'Two': 2,
  'Three': 3,
  'Four':  4,
  'Five': 5,
  'Six': 6,
  'Sevem': 7,
  'Eight': 8,
  'Nine': 9,
  'Ten': 10,
  'Jack': 10,
  'Queen': 10,
  'King': 10,
}

class Card(object):  
  def __init__ (self, rank, suit):
    self.rank = rank
    self.suit = suit
    
  def __str__(self):
    return self.rank + ' of ' + self.suit 

class Deck(object):
  cards = []
  
  def __init__ (self):
    for rank in Rank:
      for suit in Suit:
        self.cards.append(Card(rank, suit))
        
  def shuffle (self):
    random.shuffle(self.cards)
    
  def draw_card(self):
    return self.cards.pop(0)       

class Player(object):  
  def __init__ (self,name,dealer=False):
    self.hand=[]
    self.name = name
    self.dealer = dealer
   
  def add_to_hand (self, card):
    self.hand.append(card)
  
  def get_current_hand(self):
    return self.hand

# Instantiate the deck:
deck = Deck()
deck.shuffle()

# Create the players:
players = []
players.append(Player('ME'))
players.append(Player('Dealer', True))

# Deal initial hands:
for player in players:
  for itt in range(0,2):
    player.add_to_hand(deck.draw_card())
 
for player in players:
  stay = False
  while not stay:
    print(player.name + ' hand:')
    for card in player.get_current_hand():
      print(card)
    action = raw_input('Hit, Stay: ')
    if(action == 'H'):
      player.add_to_hand(deck.draw_card())
    elif(action == 'S'):
      stay = True
    else:
      print('Unknown action')
  



