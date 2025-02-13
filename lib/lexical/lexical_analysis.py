import typing
import re

from .util import token_table, token_list, token

class lexical_analysis:
    def __init__(self):
        self.regex = {
            'identifier': '[_a-zA-Z][_a-zA-Z0-9]*',
            'integer': '[+-]?[0-9]+',
            'string': '".*"',
            'character': '\'.\'',
            'library': '<.*>',
            'keyword': 'int|char|string|main|for|while|else if|if|else|return|include|define',
            'operator': '=|-|\+|\*|\\|>=|<=|>|<|==',
            'punc': '{|}|;|\(|\)|,|\[|\]|#',
        }
        self.order = ['keyword', 'library', 'identifier', 'string', 'integer', 'character', 'operator', 'punc']
        self.tok_table = token_table()

    def run(self, src: str) -> token_list | None:
        # this will run the lexical analysis and return the token list
        tok_lst = token_list()
        while len(src) > 0:
            src = src.strip()
            matched = False
            for type in self.order:
                res = re.match(self.regex[type], src)
                if res is not None:
                    obj = res.group() # fetch the element back
                    tok = token(obj = obj, type = type) # cast the obj into token
                    if not self.tok_table.insert(tok): # try to insert it to token table
                        print(f'Error: lexical analysis, {tok} cannot insert to token table')
                        return None
                    # insert tok into token list
                    if not tok_lst.push(tok):
                        print(f'Error: lexical analysis, {tok} cannot insert to token list')
                        return None

                    # remove the found token from src
                    src = src[len(obj):]
                    matched = True
                    break
            
            if not matched:
                print(f'Error: lexical analysis, {src}')
                return None
        
        return tok_lst