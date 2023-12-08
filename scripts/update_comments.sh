#!/usr/bin/env python3

#
# Somehow `comment: ""` section is not reflected.
# Here is the middle script to help register add column comments
# in each table
#

import os
import re
import mysql.connector

user='moomin'
password='moomin'
database='test'

dirpath = 'models'
files   = [f for f in os.listdir(dirpath) if os.path.isfile(os.path.join(dirpath, f))]

def to_snake_case(input_string):
    i = 0
    words = ''
    for char in input_string.strip():
        if(char.isupper()):
            char = char.lower()

            if(i == 1 and input_string[0].isupper()):
                char = '_' + char
            elif(i > 0 and input_string[i-1].islower()):
                char = '_' + char

        words += char
        i = i +1
    print(words)
    return words

def to_plural(table):
    if(table == 'reply'):
        return 'replies'
    return table + 's'

def format_table(line):
    return line.replace('type', '')\
               .replace('struct', '')\
               .replace('{', '')\
               .strip()\
               .replace('/', '')\
               .strip()

queries = []
for file in files:
    with open(f"models/{file}", 'r') as file:
        lines = file.readlines()
        table = ''
        for line in lines:
            line = line.strip()
            if(re.search(r'type.*struct.*{', line)):
                table = to_snake_case(format_table(line))
            if('comment:' in line):
                data_type_match = re.search(r'gorm:"type:(.*?)"', line)

                data_type = ''
                if data_type_match:
                    data_type = data_type_match.group(1).split(';')[0]

                column = to_snake_case(line.split(' ')[0].replace('/', '').strip())

                comment_match = re.search(r'comment:\s+"(.*?)"', line)
                if comment_match:
                    comment = comment_match.group(1)
                else:
                    comment = None

                queries.append(f"ALTER TABLE {to_plural(table)} MODIFY {column} {data_type} COMMENT '{comment}';\n")

with open('query.sql', 'w') as file:
    file.writelines(queries)

os.system(f"mysql -u {user} -p{password} {database} < query.sql")
os.system(f"mysql -u {user} -p{password} -e \"SELECT table_name, column_name, column_comment FROM information_schema.columns WHERE table_schema = '{database}';\"")
