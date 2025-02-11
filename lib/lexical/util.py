import typing

class token:
    def __init__(self, obj: str, type: str):
        self.id = -1 # the id is assigned when push into the token list
        self.obj = obj
        self.type = type

    def __eq__(self, t) -> bool:
        return t.obj == self.obj

    def __str__(self) -> str:
        return f'<\'{self.obj}\', {self.type}, {self.id}>'

    def same_type(self, t) -> bool:
        return self.type == t.type
    
class token_table:
    def __init__(self):
        self.table = {
            'identifier': [],
            'integer': [],
            'string': [],
            'keyword': [],
            'operator': [],
            'punc': [],
        }
        self.id_boundary = {
            'identifier': [0, 1],
            'integer': [1, 2],
            'string': [2, 3],
            'keyword': [3, 10],
            'operator': [10, 19],
            'punc': [20, 27],
        }
        self.current_id = {
            'identifier': 0,
            'integer': 1,
            'string': 2,
            'keyword': 3,
            'operator': 10,
            'punc': 20,
        }
        self.allow_increment = {
            'identifier': False,
            'integer': False,
            'string': False,
            'keyword': True,
            'operator': True,
            'punc': True,
        }

    def __str__(self):
        ret = ''
        for x in self.table:
            ret += f'{x}: ['+ ''.join([tok.__str__() + ', ' for tok in self.table[x]]) + ']\n'
        return ret

    def is_exist(self, tok: token) -> token | bool:
        # check if @tok already inside this token table
        # if such token exist, the token inside the table is returned, otherwise False is returned

        for t in self.table[tok.type]:
            if t == tok:
                return t, True
        
        return None, False

    def insert(self, t: token) -> bool:
        # return True if @t successfully insert into the list

        # assigne id to @t
        if not self.assign_token_id(t):
            return False
        
        # push @t to the list
        _, exist = self.is_exist(t)
        if not exist:
            self.table[t.type].append(t)
        return True
     
    def increment_token_id(self, type: str) -> int:
        # return a new token type id if the id is increment successfully
        # it only allows @type's id is inside the boundary and allow to increment the token id

        if type not in self.current_id.keys():
            print(f'Error: no such token type, type: {type}')
            return -1
        
        if not self.allow_increment[type]:
            return self.current_id[type]

        if self.current_id[type] + 1 >= self.id_boundary[type][1]:
            print('Error: all id of type {type} are assigned')
            return -1
        
        self.current_id[type] += 1
        return self.current_id[type]

    def assign_token_id(self, tok: token) -> bool:
        # return True if the token is assigned an id successfully
        # if @tok already inside this token table, the token's id is copied from the one inside the token table

        t, exist = self.is_exist(tok)
        if exist:
            tok.id = t.id
            return True
 
        new_id = self.increment_token_id(tok.type)
        if new_id == -1:
            return False
        
        tok.id = new_id
        return True

class token_list:
    def __init__(self):
        self.list = []
        self.pt = 0

    def __str__(self) -> str:
        ret = ''.join([tok.__str__() + '\n' for tok in self.list])
        return ret
    
    def get_list(self):
        return self.list

    def push(self, t: token) -> bool:

        # return True if @t is successfully insert into the list

        if not isinstance(t, token):
            return False
        self.list.append(t)
        return True

    def current(self) -> token | None:

        # this return the current token, and move the pointer forward
        # if no more token to return, None is returned

        if self.pt >= len(self.list):
            return None
        self.pt += 1
        return self.list[self.pt - 1]
        