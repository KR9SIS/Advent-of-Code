filename = "5_input.txt"
# filename = "example.txt"


class SeedMapper:
    def __init__(self) -> None:
        self.file: list[str] = []
        self.seeds: list[int] = []
        self.maps: dict[str, list[set[int]]] = {}
        self.seed_local: list[int] = []

        self.read_file()
        self.organise_maps()
        self.path_seeds()
        self.low_local = min(self.seed_local)
        print(self.low_local)

    def read_file(self):
        with open(filename) as file:
            doc = file.read()
            self.file = doc.split("\n\n")

    def get_seeds(self):
        seeds = self.file.pop(0)
        _, seeds = seeds.split(":")
        seeds = seeds.split()
        self.seeds = list(map(int, seeds))

    def organise_maps(self):
        self.get_seeds()
        for item in self.file:
            header, val = item.split(":")
            val = val.strip()
            val = val.split("\n")
            old_val = val
            val = []
            for i in old_val:
                nums = i.split()
                nums = list(map(int, nums))
                set_range = nums.pop()
                dest, src = nums
                new_val = []
                for i in nums:
                    start = i
                    end = i + set_range
                    rng = (start, end)
                    new_val.append(rng)

                new_val.append((dest, src))
                val.append(new_val)

            self.maps[header] = val

    def check_map(self, seed: int, map: str) -> int:
        def check_line(seed: int, found: tuple[int], relative: tuple[int]) -> int:
            diff = relative[0] - found[0]

            seed = seed + diff

            # found = sorted(found)
            # relative = sorted(relative)

            # index = found.index(seed)
            # seed = relative[index]

            return seed

        lines = self.maps[map]

        for line in lines:
            # dest, src = line[-1]
            # dest + (seed - src)
            src = line[1]
            start, end = src

            if start <= seed <= end:
                return check_line(seed, line[1], line[0])

        return seed

    def path_seeds(self):
        for seed in self.seeds:
            for map in self.maps:
                seed = self.check_map(seed, map)

            self.seed_local.append(seed)


if __name__ == "__main__":
    mapping = SeedMapper()
