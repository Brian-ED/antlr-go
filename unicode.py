import tangled_up_in_unicode as u
assert "APL FUNCTIONAL SYMBOL CIRCLE BACKSLASH" == u.name("⍉")
print(u.name("⊆"))
print(u.name("⊂"))
print(u.name("⊃"))
print(u.name("⊇"))

def upper(x:str):
  return "".join(chr(int(h, base=16)) if (h:=u.uppercase(i)) else i for i in x)

def lower(x:str):
  return "".join(chr(int(h, base=16)) if (h:=u.lowercase(i)) else i for i in x)

def info(xin):
  return list(zip(*[[
#   u.name(x),
    u.combining(x),
    u.mirrored(x),
    u.decomposition(x),
    u.category(x),
    u.bidirectional(x),
    u.east_asian_width(x),
    u.script(x),
    u.block(x),
    u.age(x),
    u.combining_long(x),
    u.category_long(x),
    u.bidirectional_long(x),
    u.east_asian_width_long(x),
    u.script_abbr(x),
    u.block_abbr(x),
    u.age_long(x),
    u.prop_list(x),
#   u.titlecase(x),
#   u.lowercase(x),
#   u.uppercase(x),
  ] for x in xin]))

from sys import maxunicode
#print(*info("[]⟨⟩⍉⌽←→↑↓字"), sep="\n")
#print(maxunicode)
