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
    # print(words)
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
for file_path in files:
    with open(f"models/{file_path}", 'r') as file:

        lines = file.readlines()
        table = ''
        for line in lines:
            line = line.strip()
            if(re.search(r'type.*struct.*{', line)):
                table = to_snake_case(format_table(line))
            if('gorm:' in line):
                if('comment:' in line):
                    data_type_match = re.search(r'gorm:"type:(.*?)"', line)

                    column = to_snake_case(line.split(' ')[0].replace('/', '').strip())

                    data_type = ''
                    if data_type_match:
                        data_type = data_type_match.group(1).split(';')[0]
                    else:
                        if(column == 'required' or column == 'boolean_value'):
                            data_type = 'bool'


                    comment_match = re.search(r'comment:\s+"(.*?)"', line)
                    comment = ''
                    if comment_match:
                        comment = comment_match.group(1)
                    else:
                        if('comment:' in line):
                           comment = line.split('comment:')[1].replace('"', '').replace("'", "").replace("`", "")

                    print(f"comment found: {file_path} [{column}][{comment}]")
                    queries.append(f"ALTER TABLE {to_plural(table)} MODIFY {column} {data_type} COMMENT \"{comment}\";\n")
                else:
                    print(f"comment not found: {file_path} [{line}]")
                    pass


queries = '; '.join(queries)
print('=========== query =============')
print(queries)
print('===========       =============')
with open('query.sql', 'w') as file:
    file.writelines(queries)

os.system(f"mysql -u {user} -p{password} {database} < query.sql")
os.system(f"mysql -u {user} -p{password} -e \"SELECT table_name, column_name, column_comment FROM information_schema.columns WHERE table_schema = '{database}';\"")
