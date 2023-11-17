#!/usr/bin/env python3

import yaml

# Load the YAML data
with open("test_data.yml", "r") as yaml_file:
    data = yaml.safe_load(yaml_file)

# Define SQL insert statement templates
insert_template = "INSERT INTO {} ({}) VALUES ({})"

# Function to convert data to SQL insert statements
def convert_to_insert(table_name, record):
    columns = ", ".join(record.keys())
    values = ", ".join([f"'{value}'" if isinstance(value, str) else str(value) for value in record.values()])
    return insert_template.format(table_name, columns, values)

# Function to generate SQL insert statements for a table
def generate_inserts(table_name, records):
    inserts = []
    for record in records:
        insert = convert_to_insert(table_name, record)
        inserts.append(insert)
    return inserts

queries = ''

# Process each table
for table_name, records in data.items():
    inserts = generate_inserts(table_name, records)
    print("\n".join(inserts))
    print("\n")
    queries += "; ".join(inserts) + '; '

print(queries)
