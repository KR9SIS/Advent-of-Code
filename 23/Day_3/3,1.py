from string import punctuation, digits

filename = "3_input.txt"
# filename = "example.txt"
# filename = "test.txt"  # p1: 925, p2: 6756

"""
Check if 
"""


class find_parts:
    def __init__(self) -> None:
        # Removes . from punctuation and makes it a set
        self.symbols = set(list(punctuation))
        self.symbols.remove(".")
        self.lines: list[str] = []
        self.symbol_ind_per_line: dict[int, list[list[int, str]]] = {}
        self.number_ind_in_line: dict[int, list[int, list[int]]] = {}
        self.parts = []

        self.get_file_lines(filename)
        self.get_symbols_and_num_in_lines()
        self.compare_lines()
        parts = sum(self.parts)
        print(parts)

    def get_file_lines(self, filename):
        with open(filename) as file:
            for line in file:
                self.lines.append(line.strip())

    def get_symbols_and_num_in_lines(self):
        """
        Start by iterating over every line
        then we check if the symbol is something other than "."
        If it is a symbol or number, then we note that down along with it's corresponding index
        """
        for index, line in enumerate(self.lines):
            self.symbol_ind_per_line[index] = []
            self.number_ind_in_line[index] = []

            numbers_in_line = []
            corresponding_index = []
            num_part = []

            for l_index, symbol in enumerate(line):
                if symbol in digits:
                    # Finds every number and index of said number in the line
                    num_part.append(symbol)
                    corresponding_index.append(l_index)

                elif symbol not in digits:
                    if symbol in self.symbols:
                        # Finds the index of every symbol in the line
                        self.symbol_ind_per_line[index].append((l_index, symbol))

                    num_part = "".join(num_part)
                    if num_part:
                        numbers_in_line.append([int(num_part), corresponding_index])
                    corresponding_index = []
                    num_part = []

            num_part = "".join(num_part)
            if num_part:
                numbers_in_line.append([int(num_part), corresponding_index])
            corresponding_index = []
            num_part = []

            self.number_ind_in_line[index] = numbers_in_line

    def compare_lines(self):
        """
        Iterate over the symbols in each line getting its index
        See if any number in their line the one above or below has an index equal to +-1
        If it does, then we mark it as an engine part.
        """

        def cmp_index(symbol_index, line_index):
            checked_line = self.number_ind_in_line[line_index]
            left = symbol_index - 1
            right = symbol_index + 1

            for number, indexes in checked_line:
                if symbol_index in indexes:
                    self.parts.append(number)
                elif left in indexes:
                    self.parts.append(number)
                elif right in indexes:
                    self.parts.append(number)

        for key, symbols in self.symbol_ind_per_line.items():
            if symbols:
                try:
                    self.number_ind_in_line[key - 1]
                    index_line_above = key - 1
                except KeyError:
                    index_line_above = None
                try:
                    self.number_ind_in_line[key + 1]
                    index_line_below = key + 1
                except KeyError:
                    index_line_below = None

                for symbol_index, _ in symbols:
                    if isinstance(index_line_above, int):
                        cmp_index(symbol_index, index_line_above)
                    if isinstance(index_line_below, int):
                        cmp_index(symbol_index, index_line_below)

                    cmp_index(symbol_index, key)


if __name__ == "__main__":
    parts = find_parts()
