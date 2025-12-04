filename = "2_input.txt"

min_cubes: dict[int, dict[str, int]] = {}


def main():
    content = get_content(filename)
    for game, sets in zip(content["games"], content["sets"]):
        check_set(game, sets)
    powers = get_pow()
    answer = sum(powers)
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
    for i in sets:
        for cubes in i:
            amount, color = cubes.split()
            amount = int(amount)
            if game not in min_cubes:
                min_cubes[game] = {}
            if color not in min_cubes[game]:
                min_cubes[game][color] = 0
            if min_cubes[game][color] < amount:
                min_cubes[game][color] = amount


def get_pow():
    powers = []
    for game in min_cubes:
        g_powers = None
        for color in min_cubes[game].values():
            if g_powers:
                g_powers *= color
            else:
                g_powers = color
        powers.append(g_powers)

    return powers


if __name__ == "__main__":
    main()
