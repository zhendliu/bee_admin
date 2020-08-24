package	model	
import (	
"time"	
)	
// Baseinfo [...]		
type	Baseinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Name	string	`gorm:"column:name;type:text"`			
IP	string	`gorm:"column:ip;type:text"`			
Port	string	`gorm:"column:port;type:text"`			
User	string	`gorm:"column:user;type:text"`			
Pwd	string	`gorm:"column:pwd;type:text"`			
Dname	string	`gorm:"column:dname;type:text"`			
Createtime	time.Time	`gorm:"column:createtime;type:datetime"`			
}		
// Cphinfo [...]		
type	Cphinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Hgp	string	`gorm:"column:hgp;type:text"`			
Htp	string	`gorm:"column:htp;type:text"`			
Ctrl	string	`gorm:"column:ctrl;type:text"`			
Cycle	int	`gorm:"column:cycle;type:int"`			
Sw	int	`gorm:"column:sw;type:int"`			
Protocolid	int	`gorm:"column:protocolid;type:int"`			
Createtime	time.Time	`gorm:"column:createtime;type:datetime"`			
Info	string	`gorm:"column:info;type:text"`			
}		
// Ctrlinfo [...]		
type	Ctrlinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Sensorlist	int	`gorm:"column:sensorlist;type:int"`			
Sensorname	string	`gorm:"column:sensorname;type:text"`			
Info	string	`gorm:"column:info;type:text"`			
Createtme	time.Time	`gorm:"column:createtme;type:timestamp"`			
}		
// Ctrllist [...]		
type	Ctrllist	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Createtime	time.Time	`gorm:"column:createtime;type:timestamp"`			
Size	int	`gorm:"column:size;type:int"`			
Name	string	`gorm:"column:name;type:text"`			
}		
// Fileinfo [...]		
type	Fileinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Name	string	`gorm:"column:name;type:text"`			
Path	string	`gorm:"column:path;type:text"`			
Fname	string	`gorm:"column:fname;type:text"`			
Createtime	time.Time	`gorm:"column:createtime;type:datetime"`			
}		
// Redisinfo [...]		
type	Redisinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Name	string	`gorm:"column:name;type:text"`			
IP	string	`gorm:"column:ip;type:text"`			
Port	string	`gorm:"column:port;type:text"`			
Auth	string	`gorm:"column:auth;type:text"`			
Queuename	string	`gorm:"column:queuename;type:text"`			
Createtime	time.Time	`gorm:"column:createtime;type:datetime"`
}		
// Sensorinfo [...]		
type	Sensorinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Sensorlist	int	`gorm:"column:sensorlist;type:int"`			
Sensorname	string	`gorm:"column:sensorname;type:text"`			
Savename	string	`gorm:"column:savename;type:text"`			
Transform	string	`gorm:"column:transform;type:text"`			
Unit	string	`gorm:"column:unit;type:text"`			
Maxv	string	`gorm:"column:maxv;type:text"`			
Info	string	`gorm:"column:info;type:text"`			
Createtme	time.Time	`gorm:"column:createtme;type:datetime"`			
}		
// Sensorlist [...]		
type	Sensorlist	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Createtime	time.Time	`gorm:"column:createtime;type:datetime"`			
Size	int	`gorm:"column:size;type:int"`			
Name	string	`gorm:"column:name;type:text"`			
Tname	string	`gorm:"column:tname;type:text"`			
}		
// Server [...]		
type	Server	struct {		
Port	int	`gorm:"column:port;type:int"`			
Info	string	`gorm:"column:info;type:text"`			
}		
// Sphinfo [...]		
type	Sphinfo	struct {		
ID	uint	`gorm:"primary_key;column:id;type:int unsigned;not null"`			
Hgp	string	`gorm:"column:hgp;type:text"`			
Htp	string	`gorm:"column:htp;type:text"`			
Ctrl	string	`gorm:"column:ctrl;type:text"`			
Cycle	int	`gorm:"column:cycle;type:int"`			
Sw	int	`gorm:"column:sw;type:int"`			
Protocolid	int	`gorm:"column:protocolid;type:int"`			
Databaseid	string	`gorm:"column:databaseid;type:text"`			
Fileid	int	`gorm:"column:fileid;type:int"`			
Redisid	int	`gorm:"column:redisid;type:int"`			
Createtime	time.Time	`gorm:"column:createtime;type:datetime"`			
Info	string	`gorm:"column:info;type:text"`			
}		

