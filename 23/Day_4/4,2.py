filename = "4_input.txt"
# filename = "example.txt"


class Scratchcards:
    """
    Game rules:
        if card 1 has four matches, then you get A copy of cards 2 til 5
        Then you calculate the copies created by your copies. Keep doing this until no more copies are generated.

    """

    def __init__(self) -> None:
        self.cards: dict[int, tuple[list[int], int]] = {}
        self.count = 0

        self.get_cards()
        self.card_counter()
        print(self.count)

    def calc_wins(self, card, win_num, num):
        wins = 0
        win_list = []
        for win in win_num:
            if win in num:
                wins += 1
        if wins:
            copies = card
            for i in range(wins):
                copies += 1
                win_list.append(copies)

        return win_list

    def get_cards(self):
        with open(filename) as file:
            for line in file:
                front, numbers = line.split("|")
                card, win_num = front.split(":")
                _, card = card.split()
                card = int(card)

                numbers = numbers.strip().split()
                win_num = win_num.strip().split()

                numbers = list(map(int, numbers))
                win_num = list(map(int, win_num))
                win_list = self.calc_wins(card, win_num, numbers)
                self.cards[card] = [win_list, 0]

    def card_counter(self):
        for card in self.cards:
            self.check_tree(card)
        self.get_sum()

    def check_tree(self, root):
        """
        Start by iterating over a cards wins
            See if the win has wins and if it does check those
                When no more wins are found, return, and update each cards worth
        """
        self.cards[root][1] += 1
        win_list = self.cards[root][0]

        for copy in win_list:
            self.check_tree(copy)

    def get_sum(self):
        for card in self.cards.values():
            self.count += card[1]


if __name__ == "__main__":
    cards = Scratchcards()
