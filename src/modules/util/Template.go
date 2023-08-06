package util

type Template struct {
}

type TemplateST struct {
    key string
    content string
}

func (template Template) Execute() bool {

	logger := log.CommonLogger{}

	templateST := []TemplateST{
        {"EM_Response.ini",
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
		{"tomcat-users.xml",
			`
			<role rolename="admin"/>
			<role rolename="manager"/>
			<role rolename="mstrWebAdmin"/>
			<user username="admin" password="admin" roles="admin,manager,mstrWebAdmin"/>
		`,
   		},
   	}

	   for _, command := range TemplateST {

	return true
}
