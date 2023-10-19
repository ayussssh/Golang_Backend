import re

# Define a function to decode the packet
def decode_packet(data, encoding='utf-8'):
    if len(data) != 44:
        raise ValueError("Invalid packet size. Expected 44 bytes, got {}.".format(len(data)))

    decoded = {
        'Short1': int.from_bytes(data[0:2], byteorder='big', signed=True),
        'Characters1': data[2:14].decode(encoding),
        'SingleByte': data[14],
        'Characters2': data[15:23].decode(encoding),
        'Short2': int.from_bytes(data[23:25], byteorder='big', signed=True),
        'Characters3': data[25:40].decode(encoding),
        'Long': int.from_bytes(data[40:44], byteorder='big', signed=True)
    }

    return decoded

# Input in the format [\x** \x** ...]
input_str = input("Enter the packet in the format [ \xD2 \x5B ....]: ")

# Extract hexadecimal values using regular expressions
hex_values = re.findall(r'\\x([0-9a-fA-F]{2})', input_str)

# Convert hex values to bytes
byte_values = bytes(int(value, 16) for value in hex_values)

# Check if the byte_values length is 44
if len(byte_values) != 44:
    print("Invalid input. Please provide a valid packet in the specified format.")
else:
    decoded_msg = decode_packet(byte_values)

    # Format the output as specified
    output = "Decoded struct: {{{}, \"{}\", {}, \"{}\", {}, \"{}\", , {}}}".format(
        decoded_msg['Short1'], decoded_msg['Characters1'],
        decoded_msg['SingleByte'], decoded_msg['Characters2'],
        decoded_msg['Short2'], decoded_msg['Characters3'],
        decoded_msg['Long']
    )

    print(output)
