filename = "4_input.txt"
# filename = "example.txt"


class Scratchcards:
    def __init__(self) -> None:
        self.cards: dict[str, tuple[list[int]]] = {}
        self.card_points: list[int] = []

        self.get_cards()
        self.calc_win()
        points = sum(self.card_points)
        print(points)

    def get_cards(self):
        with open(filename) as file:
            for line in file:
                front, numbers = line.split("|")
                card, win_num = front.split(":")

                numbers = numbers.strip().split()
                win_num = win_num.strip().split()

                numbers = list(map(int, numbers))
                win_num = list(map(int, win_num))

                self.cards[card] = (win_num, numbers)

    def calc_win(self):
        for card, numbers in self.cards.items():
            worth = 0
            win_num, num = numbers
            for win in win_num:
                if win in num:
                    if worth:
                        worth *= 2
                    else:
                        worth = 1
            self.card_points.append(worth)
            """"""


if __name__ == "__main__":
    cards = Scratchcards()
