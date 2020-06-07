# Standard Unix file to CSV conversion program 

Provide metadata and convert standard Unix files into csv format - it is dependent on ordered structure for Input file  
Metadata example attached  
Ouput csv will be in current working directory    

### How to use  
Update unixmeta.json with required fields  
go run unix2csv.go  
output and csv file will be hosted in current pwd  
 
### Pending Features   
Socket or rest connection to send csv to remote location  
Potential for auth function to allow this in a secure manner  

