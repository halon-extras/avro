# Avro

This plugin allows encoding and decoding of data in the avro format. It's based on the [Goavro](https://github.com/linkedin/goavro/) maintained by Linkedin.

## Installation

Follow the [instructions](https://docs.halon.io/manual/comp_install.html#installation) in our manual to add our package repository and then run the below command.

### Ubuntu

```
apt-get install halon-extras-avro
```

### RHEL

```
yum install halon-extras-avro
```

## Exported functions

These functions needs to be [imported](https://docs.halon.io/hsl/structures.html#import) from the `extras://avro` module path.

### avro_encode(schema, data)

Encode data in binary arvo format according to the schema (works with any data serializable to JSON).

**Params**

- schema `string` - The avro schema
- data `any` - The data

**Returns**

On success it will return a string that contains the binary avro data according to the schema. On error an exception will be thrown.

**Example**

```
import { avro_encode } from "extras://avro";
$schema = ''{"type": "record", "name": "LoginEvent", "fields": [{"name": "Username", "type": "string"}]}'';
$data = ["Username" => "batman"];
$binary = avro_encode($schema, $data);
```

### avro_decode(schema, avro)

Decode binary arvo data to an object type according to the schema (works with any data serializable to JSON).

**Params**

- schema `string` - The avro schema
- avro `string` - The data

**Returns**

On success it will return the object type from the decoded binary avro data according to the schema. On error an exception will be thrown.

**Example**

```
import { avro_encode, avro_decode } from "extras://avro";
$schema = ''{"type": "record", "name": "LoginEvent", "fields": [{"name": "Username", "type": "string"}]}'';
$data = ["Username" => "batman"];
$binary = avro_encode($schema, $data);
echo avro_decode($schema, $binary);
```