import csv
import textwrap

INPUT_FILE = "params.csv"


def generate_line(row):
    comment_lines = textwrap.wrap(
        row["Description"], width=73, fix_sentence_endings=True)
    comment_lines[0] = comment_lines[0].capitalize()
    for i in range(len(comment_lines)):
        comment_lines[i] = "// " + comment_lines[i]
    comment = "\n".join(comment_lines)

    return f"{comment}\n{row['Go Parameter']} {row['Values']} `qs:\"{row['Parameter']},omitempty\"`"


if __name__ == "__main__":
    data = csv.DictReader(open(INPUT_FILE), delimiter=";")
    rsl = ""
    for row in data:
        rsl = f"{rsl}\n{generate_line(row)}"
    print(rsl)
