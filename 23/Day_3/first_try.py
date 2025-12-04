from string import punctuation

# filename = "3_input.txt"
filename = "example.txt"


class find_parts:
    def __init__(self) -> None:
        self.valid_num = []
        self.numbers_in_line: dict[int, list[int]] = {}
        self.lines: list[str] = []
        self.line_number = 0
        self.curr_line_1: str = ""
        self.curr_line_2: str = ""
        self.line1_index = 0
        self.line2_index = 0
        self.string_index = 0
        self.last_catch_index = 0

        self.get_file_lines(filename)
        for line_index in range(len(self.lines)):
            self.line1_index = line_index
            self.line2_index = line_index + 1
            self.get_two_lines()
            self.comp_lines()
        ""

    def get_file_lines(self, filename):
        with open(filename) as file:
            for line in file:
                self.lines.append(line.strip())

    def get_two_lines(self):
        """
        Finds the index of a specific line and the line below it and returns those.
        """
        try:
            self.curr_line_2 = self.lines[self.line2_index]
            self.curr_line_1 = self.lines[self.line1_index]
        except IndexError:
            """"""

    def catch(
        self, checked_line: str, index: int, catch_widht: int, start=False, end=False
    ):
        val1 = catch_widht / 2
        val2 = val1
        if isinstance(val1, float):
            val1 = int(val1 + 0.5)
            val2 = int(val2 - 0.5)
        if start:
            checked_line[: index + val1]
        elif end:
            checked_line[index - val1 :]
        else:
            catch = checked_line[index - val1 : index + val2]

        return catch

    def checks(self, check, checked_line, cl_index):
        try:
            num = int(check)  # See if check is unbroken int i.e. not 1..2
            self.numbers_in_line[cl_index].append(num)
        except ValueError:
            check: str = self.catch(checked_line, self.string_index, 5)
            val = check
            num_not_found = True
            offset = 0
            while num_not_found:
                if val[0] in punctuation and val[-1] in punctuation:
                    val = int(val.strip(punctuation))
                    self.numbers_in_line[cl_index].append(val)
                    num_not_found = False
                elif val[0] in punctuation:
                    offset += 1
                    index = self.string_index + offset
                    val = self.catch(checked_line, index, 5)
                elif val[-1] in punctuation:
                    offset -= 1
                    index = self.string_index + offset
                    val = self.catch(checked_line, index, 5)
                else:
                    val = []
                    for i in check:
                        val.append(i)
                    seat = val.index(".")
                    if seat < 2:
                        offset += 1
                        index = self.string_index + offset
                        val = self.catch(checked_line, index, 5)
                    elif seat > 2:
                        offset -= 1
                        index = self.string_index + offset
                        val = self.catch(checked_line, index, 5)
                    elif seat == 2:
                        offset = -2
                        val = self.catch(checked_line, self.string_index - offset, 5)
                        # TODO imporve this code .split() doesn't work
                        val = self.catch(checked_line, self.string_index, 7)
                        val = val.strip(punctuation).split(".")
                        for i in val:
                            i = int(i)
                            self.numbers_in_line[cl_index].append(val)
                        num_not_found = False

                    # val = val.split(".")
                    # val = list(map(int, val))
                    # val = max(val)

            # TODO Make work with unbroken numbers such as ..806.540..

            ""

            # chech_parts = check.split(".")
            # for num in chech_parts:
            #     num = num.strip(punctuation)
            #     if num and num not in punctuation:
            #         num = int(num)
            #         if self.numbers_in_line[cl_index]:
            #             last_num = self.numbers_in_line[cl_index][-1]
            #         if last_num != num:
            #             self.numbers_in_line[cl_index].append(num)
            #     else:
            #         pass

    def line_check(self, i, checked_line, cl_index):
        """
        Checks if i1 is not a "." or an integer and if it is then it checks the corresponding index of line2
        """
        if i != ".":
            try:
                int(i)  # Filter awway numbers
            except ValueError:
                # If neither line start nor end
                if self.string_index != 1 and self.string_index != len(checked_line):
                    check: str = self.catch(checked_line, self.string_index, 3)

                    self.checks(check, checked_line, cl_index)
                # Account for index error at line start
                elif self.string_index == 1:
                    check = self.catch(checked_line, self.string_index, 3, start=True)
                    self.checks(check, checked_line, cl_index)

                # Account for index error at line end
                elif self.string_index == len(checked_line):
                    check: str = self.catch(
                        checked_line, self.string_index, 3, end=True
                    )
                    self.checks(check, checked_line, cl_index)

                    # TODO First step is to see if there are any values across or diagonally to the symbol, and then if there is I need to get that entire number.

    def comp_lines(self):
        """
        Iterate over line 1 and 2 and if line 1 shows something other than a "."
        then it checks what is in line 2 and if that is a number, then logs it as a valid num
        """
        curr_line1 = self.curr_line_1
        curr_line1_index = self.line1_index
        if curr_line1_index not in self.numbers_in_line:
            self.numbers_in_line[curr_line1_index] = []

        curr_line2 = self.curr_line_2
        curr_line2_index = self.line2_index
        if curr_line2_index not in self.numbers_in_line:
            self.numbers_in_line[curr_line2_index] = []

        for i1, i2 in zip(self.curr_line_1, self.curr_line_2):
            self.string_index += 1
            self.line_check(i1, curr_line2, curr_line2_index)
            self.line_check(i2, curr_line1, curr_line1_index)
        self.string_index = 0

        ""


if __name__ == "__main__":
    parts = find_parts()
