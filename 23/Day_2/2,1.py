filename = "2_input.txt"

BAG = {"red": 12, "green": 13, "blue": 14}

valid_g = []


def main():
    content = get_content(filename)
    for game, sets in zip(content["games"], content["sets"]):
        check_set(game, sets)
    answer = sum(valid_g)
    print(answer)
    ""


def get_content(filename):
    with open(filename) as file:
        content = {"games": [], "sets": []}
        for line in file:
            game, g_sets = line.split(":")
            g_sets = g_sets.strip().split(";")
            sets = [s.split(",") for s in g_sets]

            _, game = game.split()

            content["games"].append(int(game))
            content["sets"].append(sets)

    return content


def check_set(game: list[str], sets: list[list[list[str]]]):
    game_v = True
    for i in sets:
        for cubes in i:
            amount, color = cubes.split()
            amount = int(amount)
            try:
                bag_am = BAG[color]
                if amount > bag_am:
                    game_v = False
            except KeyError:
                game_v = False
    if game_v:
        valid_g.append(game)


if __name__ == "__main__":
    main()
