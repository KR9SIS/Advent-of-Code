from time import time

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
        seeds = list(map(int, seeds))
        for combo in range(0, len(seeds), 2):
            start = seeds[combo]
            rng = seeds[combo + 1]
            end = start + rng
            self.seeds.append((start, end))

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
            src = line[1]
            start, end = src

            if start <= seed <= end:
                return check_line(seed, line[1], line[0])

        return seed

    def path_seeds(self):
        for num_big, combo in enumerate(self.seeds):
            seed_start, seed_end = combo
            big_locale = self.calc_loading(num_big, len(self.seeds), "Seeds")
            for num_small, seed in enumerate(range(seed_start, seed_end)):
                self.calc_loading(
                    num_small, seed_end - seed_start, "Seed range", big_locale
                )
                for path in self.maps:
                    seed = self.check_map(seed, path)

                self.seed_local.append(seed)

    def calc_loading(
        self, num: int, end: int, small_locale: str, big_locale: float = ""
    ):
        left = num / end
        print(f"{left:.4f}%\t{small_locale} {big_locale}")
        return left


if __name__ == "__main__":
    start_time = time()
    mapping = SeedMapper()
    end_time = time()
    run_time = end_time - start_time
    print(f"Run time was: {run_time} in seconds")
