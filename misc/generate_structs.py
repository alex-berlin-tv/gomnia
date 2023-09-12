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

INPUT_FILE = "params.csv"


def generate_client_line(row):
    comment_lines = textwrap.wrap(
        row["Description"], width=73, fix_sentence_endings=True)
    comment_lines[0] = comment_lines[0].capitalize()
    for i in range(len(comment_lines)):
        comment_lines[i] = "// " + comment_lines[i]
    comment = "\n".join(comment_lines)

    return f"{comment}\n{row['Go Parameter']} {row['Type']} `qs:\"{row['Parameter']},omitempty\"`"


def client_cmd(path):
    data = csv.DictReader(open(path), delimiter=";")
    rsl = ""
    for row in data:
        rsl = f"{rsl}\n{generate_client_line(row)}"
    print(rsl)

def notification_cmd(path):
    print("Not implemented, just do it manually, fnord.")

def template_cmd(typ):
    if typ == "client":
        print("Parameter;Go Parameter;Type;Description")
    elif typ == "notification":
        print("Not implemented.")
    else:
        print("No valid type, supported: client|notification")
        sys.exit(1)


if __name__ == "__main__":
    if len(sys.argv) in range(2, 3):
        print("Invalid arguments. Allowed:")
        print("client PATH\t\t\t generate code for the client API")
        print("notification PATH\t\t generate code for the notification gateway")
        print("template [client|notification]\t outputs an empty csv template")
        sys.exit(1)
    
    if sys.argv[1] == "client":
        client_cmd(sys.argv[1])
    elif sys.argv[1] == "notification":
        notification_cmd(sys.argv[1])
    elif sys.argv[1] == "template":
        template_cmd(sys.argv[2])
    else:
        print("Invalid command")
        print("Valid client|notification|template")
        sys.exit(1)
    
