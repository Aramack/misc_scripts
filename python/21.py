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
    self.hidden_card = None
    
   
  def add_to_hand (self, card):
    if (self.dealer and (self.hidden_card == None)):
      self.hidden_card = card
    else:
      self.hand.append(card)
  
  def get_current_hand(self):
    return self.hand
  
  def hand_value(self):
    total = 0
    for card in self.hand:
      total += Rank[card.rank]
    if self.dealer:
      total += Rank[self.hidden_card.rank]
    return total
    
class Game(object):
  def __init__(self):
    # Instantiate the deck:
    self.deck = Deck()
    self.deck.shuffle()
    self.players = []
    self.dealer = Player('Dealer', True)
    player_count = input('Number of players: ')
    for int in range(0, player_count):
      self.players.append(Player(raw_input('Player Name: ')))
    self.init_hands()
      
  def init_hands(self):
    for player in self.players:
      for itt in range(0,2):
        player.add_to_hand(self.deck.draw_card())
    for itt in range(0,2):
      self.dealer.add_to_hand(self.deck.draw_card())  
      
  def player_turn(self):
    for player in self.players:
      stay = False    
      while not stay:
        print(player.name + '\'s hand:')
        for card in player.get_current_hand():
          print(card)
        action = raw_input('(H)it, (S)tay: ')
        if(action == 'H'):
          card  = self.deck.draw_card()
          print('Player ' + player.name + ' was dealt a ' + str(card))
          player.add_to_hand(card)
          total = player.hand_value()
          print('Total: ' + str(total))
          if (total > 21):
            print('Bust')
            stay = True        
        elif(action == 'S'):
          stay = True
        else:
          print('Unknown action')

  def dealer_turn(self):
    while (self.dealer.hand_value() < 16):
      card = self.deck.draw_card()
      self.dealer.add_to_hand(card)
      print('Player ' + self.dealer.name + ' was dealt a ' + str(card))
      if (self.dealer.hand_value()> 21):
        print('Dealer busts!')
        return
  
  def winner(self):
    # Who won?
    print('Dealer\'s hand:')
    print(str(self.dealer.hidden_card))
    for card in self.dealer.get_current_hand():
      print(str(card))
    dealers_total = self.dealer.hand_value()
    dealer_bust = dealers_total > 21;
    print('Total: ' + str(dealers_total))
    for player in self.players:
      player_total = player.hand_value()
      if player_total > 21:
        print('Player ' + player.name + ': Bust')
      elif player_total > dealers_total or dealer_bust:
        print('Player ' + player.name + ': Winner')
      elif player_total == dealers_total:
        print('Player ' + player.name + ': Push')
      elif player_total < dealers_total:
        print('Player ' + player.name + ': Loser')
      
game = Game()
game.player_turn()
game.dealer_turn()
game.winner()
exit(0)



