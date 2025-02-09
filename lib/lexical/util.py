import typing

class token:
    def __init__(self, obj: str, type: str):
        self.id = -1 # the id is assigned when push into the token list
        self.obj = obj
        self.type = type

    def __eq__(self, t) -> bool:
        return t.obj == self.obj
    
    def same_type(self, t) -> bool:
        return self.type == t.type
    
class token_list:
    def __init__(self):
        self.l = {
            'identifier': [],
            'integer': [],
            'string': [],
            'keyword': [],
            'operators': [],
            'punc': [],
        }
        self.id_boundary = {
            'identifier': [0, 1],
            'integer': [1, 2],
            'string': [2, 3],
            'keyword': [3, 10],
            'operators': [10, 19],
            'punc': [20, 27],
        }
        self.current_id = {
            'identifier': 0,
            'integer': 1,
            'string': 2,
            'keyword': 3,
            'operators': 10,
            'punc': 20,
        }
        self.allow_inc = {
            'identifier': False,
            'integer': False,
            'string': False,
            'keyword': True,
            'operators': True,
            'punc': True,
        }

    def push(self, t: token) -> bool:
        # return True if @t successfully push into the list

        # assigne id to @t
        if not self.assign_token_id(t):
            return False
        
        # check if @t already in the list
        for x in self.l[t.type]:
            if x == t:
                return False

        # push @t to the list
        self.l[t.type].append(t)
        return True
     
    def increment_token_id(self, type: str) -> int:
        # return a new token type id if the id is increment successfully

        if type not in self.current_id.keys():
            return -1
        
        if not self.allow_increment[type]:
            return self.current_id[type]

        if self.current_id[type] + 1 >= self.id_boundary[type]:
            return -1
        
        self.current_id[type] += 1
        return self.current_id[type]

    def assign_token_id(self, t: token) -> bool:
        # return True if the token is assigned an id successfully

        new_id = self.increment_token_id(t.type)
        if new_id == -1:
            return False
        
        t.id = new_id