filename = "1_input.txt"
# filename = "example.txt"

DIGITS = {
    "zero": ["0", "zero"],
    "one": ["1", "one"],
    "two": ["2", "two"],
    "three": ["3", "three"],
    "four": ["4", "four"],
    "five": ["5", "five"],
    "six": ["6", "six"],
    "seven": ["7", "seven"],
    "eight": ["8", "eight"],
    "nine": ["9", "nine"],
}
PARTS = {
    "zero": set(("z", "ze", "zer")),
    "one": set(("o", "on")),
    "two": set(("t", "tw")),
    "three": set(("t", "th", "thr", "thre")),
    "four": set(("f", "fo", "fou")),
    "five": set(("f", "fi", "fiv")),
    "six": set(("s", "si")),
    "seven": set(("s", "se", "sev", "seve")),
    "eight": set(("e", "ei", "eig", "eigh")),
    "nine": set(("n", "ni", "nin")),
}


def main(filename):
    content = get_file(filename)
    numbers = []
    for line in content:
        l_num = check_num(line)
        n_order = check_order(line, l_num)
        f_l = first_last(n_order)
        numbers.append(f_l)
    answer = sum(numbers)
    print(answer)


def get_file(filename):
    with open(filename) as file:
        content = []
        for line in file:
            line = line.lower()
            content.append(line)
    return content


def check_num(line):
    line_num = []
    for number in DIGITS.values():
        for version in number:
            if version in line:
                line_num.append(version)
    ""
    return line_num


def check_order(line, numbers: list[str]):
    """Checks the order of the numbers found by iterating over the line and matching the number in the line with the one"""
    num_order = []
    n_parts = set()
    for key in numbers:
        try:
            n_parts = n_parts.union(PARTS[key])
        except:
            """"""
    letters: str = ""
    for char in line:
        letters += char
        if letters in numbers:
            num_order.append(letters)
            if letters[-1] not in n_parts:
                letters = ""
            else:
                letters = letters[-1]
        elif letters in n_parts:
            pass
        elif letters[-1] in numbers:
            num_order.append(letters[-1])
            letters = ""
        elif letters[-1] in n_parts:
            letters = letters[-1]

        else:
            letters = ""
    return num_order


def first_last(numbers):
    integers = []
    for i in numbers:
        try:
            int(i)  # Check whether the element is i.e. a "1" or a "one"
            integers.append(i)
        except:
            num = DIGITS[i][0]
            integers.append(num)

    first = integers[0]
    last = integers[-1]
    firstNlast = first + last
    return int(firstNlast)


if __name__ == "__main__":
    main(filename)
