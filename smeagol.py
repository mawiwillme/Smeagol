

def usage():
	print("Did not understand input")
	menu()
def menu():
	print("1. sessions - list active sessions")
	print("2. interact <agent> - interact with agent using given agent id")
	print("3. listener <start,stop,configure,list> - configue and start listeners")
	print("4. exit")

def main():
	menu()
	command = input(">")
	print(command)

if __name__ == "__main__":
	main()
