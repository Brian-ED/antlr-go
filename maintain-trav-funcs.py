import re

def p(x):
  print(x)
  return x

r=r"type [A-Z][a-zA-Z0-9]*Context struct \{(?:\n[^}]*)*\}"

with open("parsing/calculator_parser.go","r") as f:
  s = f.read()

with open("parser/Calculator.g4", "r") as f:
  gs = f.read()

names = re.findall("# (.*)", gs)
names = {i[0].upper()+i[1:]+"Context" for i in names} # names formatted to go's type names

bodies = re.findall(r, s, re.MULTILINE)

varNames = [re.findall("type (.*) struct", i)[0] for i in bodies]

VarNamesParsingCases = [i for i in varNames if i in names]

contentOfStructs = [i.removeprefix("type "+n+" struct {\n\t").removesuffix("\n}").split("\n\t") for i,n in zip(bodies, varNames)]

begin = "// PYTHON: AUTOPLACE BEGIN"
end = "// PYTHON: AUTOPLACE END"
with open("main.go", "r+") as f:
  t = f.read()
  a, b = t.split('\n'+begin)
  b, c = b.split('\n'+end+'\n')
  aL = a.split('\n')
  bL = b.split('\n')
  cL = c.split('\n')

  prefix = "func (l *MyCalculatorListener) {}{} (ctx *parsing.{})"
  newB = [
    prefix.format(direc, name.removesuffix("Context"), name) + " {}"
      for name in VarNamesParsingCases
      for direc in ["Enter", "Exit"]
      if name not in c
  ]
  useless = [i for i in re.findall(r"func \(l \*MyCalculatorListener\) ([A-Z][A-Za-z0-9]*) \([a-z][a-zA-Z0-9]* \*parsing\.",c) if i not in [*map((lambda x:"Enter"+x.removesuffix("Context")),VarNamesParsingCases), *map((lambda x:"Exit"+x.removesuffix("Context")),VarNamesParsingCases)]]
  if useless:
    print("Useless:",useless)

with open("main.go", "w") as f:
  f.write('\n'.join(aL+[begin]+newB+[end]+cL))
