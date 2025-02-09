import argparse

def main():
    parser = argparse.ArgumentParser(description='This is a mini c compiler')
    parser.add_argument('-s', '--source', help='Source file')
    args = parser.parse_args()
    src = open(args.source)
    src = src.read()
    

if __name__ == '__main__':
    main()