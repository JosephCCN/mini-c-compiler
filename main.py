import argparse

from lib.lexical.lexical_analysis import lexical_analysis

def main():
    parser = argparse.ArgumentParser(description='This is a mini c compiler')
    parser.add_argument('-s', '--source', help='Source file')
    args = parser.parse_args()
    src = open(args.source)
    src = src.read()

    la = lexical_analysis()
    tok_lst = la.run(src)
    if tok_lst is None:
        return
    print(tok_lst)
    print(la.tok_table)

if __name__ == '__main__':
    main()