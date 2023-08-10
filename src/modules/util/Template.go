package util

import (
	log "bits/modules/common/log"
)


type Template struct {
	logger *log.CommonLogger
}

type TemplateST struct {
    key string
    content string
}

func (t Template) Execute() bool {

	t.logger = log.GetLogger()

	templateST := []TemplateST{
{
"ALIAS Command",
`
## MSTR COMMAND

MSTR_HOME=/home/mstr/MicroStrategy
MSTR_LOG=/home/mstr/MicroStrategy/log
export PATH=$PATH:$MSTR_HOME/bin

alias cdmbin='cd $MSTR_HOME/bin'
alias cdmlog='cd $MSTR_LOG'
alias mlog='tail -100f $MSTR_LOG/DSSErrors.log'

alias mgs='$MSTR_HOME/bin/mstrctl -s IntelligenceServer gs'
alias mstop='$MSTR_HOME/bin/mstrctl -s IntelligenceServer stop'
alias mterm='$MSTR_HOME/bin/mstrctl -s IntelligenceServer term'
alias mstart='$MSTR_HOME/bin/mstrctl -s IntelligenceServer start'

alias pdfgs='$MSTR_HOME/bin/mstrctl -s PDFExportService gs'
alias pdfstop='$MSTR_HOME/install/Export/pdfexporter.sh stop'
alias pdfstart='$MSTR_HOME/install/Export/pdfexporter.sh start'

alias emgs='$MSTR_HOME/bin/mstrctl -s EMService gs'
alias emstop='$MSTR_HOME/bin/mstrctl -s EMService stop'
alias emstart='$MSTR_HOME/bin/mstrctl -s EMService start'
`,
},

{
"EM_Response.ini",
`
[EMProjectHeader]
EMProject=1
EMProjectEncryptPwd=0
EMProjectDSSUser=Administrator
EMProjectDSSPwd=
EMProjectPkgFile=/sw/mstr/MicroStrategy/install/OOTB-EM.mmp
EMProjectDSNName=MSTR_STAT
EMProjectDSNUserName=mstr
EMProjectDSNUserPwd=mstr
EMProjectDSNPrefix=
`,
},

{
"tomcat-users.xml",
`
<role rolename="admin"/>
<role rolename="manager"/>
<role rolename="mstrWebAdmin"/>
<user username="admin" password="admin" roles="admin,manager,mstrWebAdmin"/>
`,
},

{
"ODBC(MYSQL BASIC)",
`
[MYSQLDB]
DriverUnicodeType=1
Description=MySQL ODBC 5.x Driver
PORT=3306
DATABASE=MSTRDB
CHARSET=utf8
SERVER=XXX.XXX.XXX.XXX
OPTION=2
SOCKET=
Driver=/home/mstr/mysql_odbc/mysql_odbc_8.0.21/lib/libmyodbc8w.so
`,
},

}

	   for _, template := range templateST {

			t.logger.BasicPrint("----------------------")
			t.logger.BasicPrint(template.key)
			t.logger.BasicPrint("----------------------")
			t.logger.BasicPrint(template.content)

	   }

	return true
}
