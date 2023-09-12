"""
Generate golang structs from csv files. Columns of a csv file:
- Parameter (name of the parameter in the API)
- Go Parameter (how should be the parameter named in the go struct)
- Type (type of the parameter according to the API documentation)
- Description (optional description of the parameter)
""" 

import csv
import sys
import textwrap


TYPES = {
    "Bool": "enum.Bool",
    "integer": "int",
    "UNIX Timestamp": "TODO",
}


def parse_file(path: str) -> list[dict[str, str]]:
    rsl = []
    with open(path) as f:
        raw = [line.strip() for line in f]
    i = 0
    while i < len(raw):
        item = {}
        for j in range(3):
            if j == 0:
                item["name"] = raw[i + 0]
            elif j == 1:
                item["type"] = raw[i + 1]
            elif j == 2:
                item["description"] = raw[i + 2]
            if i + j + 1 == len(raw):
                break
        rsl.append(item)
        i += 4
    return rsl

def fmt_go_name(data: dict[str, str]) -> str:
    return data["name"].removesuffix("*").title()

def fmt_api_name(data: dict[str, str]) -> str:
    return data["name"].removesuffix("*")

def fmt_type(data: dict[str, str]) -> str:
    raw = data["type"]
    if raw in TYPES:
        return TYPES[raw]
    return raw

def fmt_comment(data: dict[str, str]) -> str:
    required = data["name"].endswith("*")
    if "description" not in data and not required:
        return ""
    raw = data["description"]
    if raw is None or raw == "" and not required:
        return ""
    if required:
        raw = f"Required. {raw}"
    comment_lines = textwrap.wrap(
        raw, width=73, fix_sentence_endings=True
    )
    comment_lines[0] = comment_lines[0].capitalize()
    for i in range(len(comment_lines)):
        comment_lines[i] = "// " + comment_lines[i]
    comment = "\n".join(comment_lines)

    return f"{comment}\n"

def generate_client_line(item):
    return f"{fmt_comment(item)}{fmt_go_name(item)} {fmt_type(item)} `qs:\"{fmt_api_name(item)},omitempty\"`"


def help():
    print("Usage: generate_structs.py PATH")
    print("Help: generate_structs.py --help\n")
    print("INPUT FILE: Select the table in the API documentation and paste it to a text file.")
    print("Struct of the text file: <PARAM_NAME>\\n<TYPE>\\n<DESCRIPTION>\\n\\n <Next item>")


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Invalid arguments.") 
        help()
        sys.exit(1)
    if sys.argv[1] == "--help":
        help()
        sys.exit(0)
    print("type NAME struct {")
    data = parse_file(sys.argv[1]) 
    for item in data:
        print(generate_client_line(item))
    print("}")