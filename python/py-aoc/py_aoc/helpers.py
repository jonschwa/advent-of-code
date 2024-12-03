import pathlib

PATH = pathlib.Path(__file__).parent.absolute()


def read_lines(datapath):
    with open(f"{PATH}/data/{datapath}") as f:
        line = f.readline()
        while line:
            yield line
            line = f.readline()


def read_file_to_list(datapath, strip=True):
    file = open(f"{PATH}/data/{datapath}")
    file_list = file.readlines()
    if strip:
        return [line.strip() for line in file_list]
    return file_list

def fetch_data(path) -> list:
    return read_file_to_list(path)

def get_data(year: int, day: int) -> list[list, list]:
    test_data = fetch_data(f'{year}/day{day}/test.txt')
    data = fetch_data(f'{year}/day{day}/data.txt')

    return [test_data, data]