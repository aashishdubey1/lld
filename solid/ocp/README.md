# OCP => Open Close Principle 

it states that software entities (classes,modules,function,etc.) should be **open** for extension, 
but **Closed** for modification. 

_what tf is that mean ?_ 
- okay so somehow we have to make shitty look code, looks good 
- in our extractor.go file, we originally dumped everything into one massive function.
- that means in future if we need to support more formats to extract, 
- we have to modify the main function 
**so Why bad is modifying main function ?** 
  1. because we noob coder don't want to touch working code 
  - if that code is working and running in production, why tf u want to touch that 
  2. also if you keep writing like this it will turn into a monster function that no one likes to touch 
  - they kick u out of the company 
- so for all the reason we follow the OCP rule 
- for that we create Abstract interface separate struct 
- so we can add new features without risking anything 

Why these two mf SRP and OCP sounds same 
**SRP** => by breaking your main function into separate structs
(CSVExporter, JSONExporter, PDFExporter),
we gave each struct one job and one reason to change.

**OCP** => Because those responsibilities are cleanly separated into isolated structs,
our overarching application logic is now closed for modification.
_achieving SRP is often the exact mechanism that allows you to follow OCP._
