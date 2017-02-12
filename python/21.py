from aenum import Enum
import random

Suit=['CLUB', 'DIAMOND', 'HEART', 'SPADE']
  
class Rank(Enum):
  __order__ = 'ACE TWO THREE FOUR FIVE SIX SEVEN EIGHT NINE TEN JACK QUEEN KING'
  ACE = 1
  TWO = 2
  THREE = 3
  FOUR = 4
  FIVE = 5
  SIX = 6
  SEVEN = 7
  EIGHT = 8
  NINE = 9
  TEN = 10
  JACK = 10
  QUEEN = 10
  KING = 10
 
class Card(object):  
  def __init__ (self, rank, suit):
    self.rank = rank
    self.suit = suit

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
  cards=[]
  dealer=False
  name=''
  
  def __init__ (self,name,dealer=False):
    self.name = name
    self.dealer = dealer
   
  def add_to_hand (self, card):
    self.cards.append(card)
  
  def get_current_hand(self):
    return self.cards
    
    
deck = Deck()
deck.shuffle()

p1=Player('ME')
dealer=Player('Dealer', True)
p1.add_to_hand(deck.draw_card())

print(p1.get_current_hand()[0].rank)

