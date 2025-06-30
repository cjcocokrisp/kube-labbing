import random
import string

characters = string.ascii_letters + string.digits
random_str = ''.join(random.choice(characters) for _ in range(32))
print(random_str)