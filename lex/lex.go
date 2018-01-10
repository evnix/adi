package lex

type lex struct {  
    buffer   string
    ptr   int
}


func isSpace(c rune){
    if c == ' ' || c == '\t'{
      return True 
	} else {
      return False
	}
}
  
  func isOp(c rune){
    if ord(c) >=33 and ord(c) <= 47:
      return True
    elif ord(c) >=58 and ord(c) <= 64:
      return True
    elif ord(c) == 96:
      return True
    elif ord(c) >=123 and ord(c) <= 126:
      return True
    else:
      return False
  }
    
  func isChar(self,c):
    if ord(c) >=65 and ord(c) <= 90:
      return True
    elif ord(c) >=97 and ord(c) <= 122:
      return True
    elif ord(c) >=91 and ord(c) <= 94:
      return True
    elif ord(c) == 95:
      return True
    else:
      return False
      
  func process_str(self,str):
    buffer = ""
    while self.ptr < len(str):
      if self.isSpace(str[self.ptr]) or self.isOp(str[self.ptr]):
        break
      else:
        buffer = buffer + str[self.ptr]
        self.ptr = self.ptr + 1
    print "[" + buffer +"]",
      
    
  func process(self,str):     
    while self.ptr<len(str):
        c = str[self.ptr] 
        if self.isChar(c):
          self.process_str(str)
        self.ptr = self.ptr + 1