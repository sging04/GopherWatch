import time

def countApples():
    i = 0;
    while True:
        print(f'Apple number {i}')
        i += 1
        time.sleep(0.4) # sleeping 400 miliseconds here
def countOranges():
    i = 0;
    while True:
        print(f'Orange number {i}')
        i += 1
        time.sleep(0.4)

def main():
    countApples()
    countOranges() # will never get here because not concurrent

if __name__ == "__main__":
    main()