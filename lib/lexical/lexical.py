import typing
import re

from util import token_list, token

class lexical_analysis:
    def __init__(self, src: str):
        self.src = src
        self.regex = {
            'integer': '[+-]?[0-9]+'
        }

    def run(self) -> token_list:
        # this will run the lexical analysis and return the token list
        pass