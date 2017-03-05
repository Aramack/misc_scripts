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
  hand=[]
  dealer=False
  name=''
  
  def __init__ (self,name,dealer=False):
    self.name = name
    self.dealer = dealer
   
  def add_to_hand (self, card):
    self.hand.append(card)
  
  def get_current_hand(self):
    return self.hand
    
    
deck = Deck()
deck.shuffle()

p1=Player('ME')
dealer=Player('Dealer', True)
p1.add_to_hand(deck.draw_card())
dealer.add_to_hand(deck.draw_card())
p1.add_to_hand(deck.draw_card())
dealer.add_to_hand(deck.draw_card())

print(p1.get_current_hand()[0])

