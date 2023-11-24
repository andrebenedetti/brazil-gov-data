package storage

// CNAE stands for Classificação Nacional de Atividades Econômicas, or
// National Classification of Economic Activities.
// Each entry is a numeric code and a string that identifies an economic activity.
// For example, this CNAE identifies nursery activities:
// "8650001";"Atividades de enfermagem"

// We store codes as strings to handle CNAES in the form "0112199"
type Cnae struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

// This is what the first few companies look like in the government data source:
// "41273596";"RODRIGO JOSE FERREIRA LOPES 05010247941";"2135";"50";"10000,00";"01";""
// "41273597";"PACHARRUS QUEIROZ DA COSTA E SILVA 03618384335";"2135";"50";"5000,00";"01";""
// "41273598";"GLORIA VIANA DIAS DA SILVA 13118961716";"2135";"50";"1100,00";"01";""
// "41273599";"ANA PAULA DA SILVA DE PAULA 04659802957";"2135";"50";"2000,00";"01";""
// "41273600";"41.273.600 AVANILSON BRUNO MATIAS DA SILVA";"2135";"50";"50000,00";"01";""
// "41273601";"GABRIELA HELENA FACINI DA SILVA 47022415838";"2135";"50";"2000,00";"01";""

// The first item, named "basic cnpj" by the government, is composed of the 8 first cnpj digits.
// On day-to-day business, people usually use the full CNPJ in the form 33.000.167/0001-01.

// We won't use this struct as is, no more. Rather, we'll join this dataset with the "Estabelecimentos" dataset
type CompanyOld struct {
	BasicCnpj                        string `json:"basicCnpj"`
	Name                             string `json:"name"`
	LegalNature                      string `json:"legalNature"`
	LegalRepresentativeQualification string `json:"legalRepresentativeQualification"`
	NominalCapital                   string `json:"nominalCapital"`
	Size                             string `json:"size"`
	FederalEntityInCharge            string `json:"federalEntityInCharge"`
}

// "20456681";"0001";"04";"1";"MARCELO CONSTRUTOR";"08";"20180201";"67";"";"";"20140616";"4399103";"";"RUA";"26, quadra 25 lote 07";"s/n";"";"Cidade Verde";"78310000";"MT";"9883";"65";"99470550";"";"";"";"";"marceloguilhermino88@hotmail.com";"";""
// "13966672";"0001";"06";"1";"CLL ASSESSORIA FARMACEUTICA";"02";"20191213";"00";"";"";"20110628";"7490199";"4712100,4751201,4753900,5819100,5911101,5920100,6201502,6202300,6204000,6209100,6311900,6319400,6810202,7020400,7311400,7312200,7319002,7319004,8599604,9002701";"AVENIDA";"KENNEDY";"914";"SALA  11";"JARDIM DO MAR";"09726253";"SP";"7075";"11";"71667800";"";"";"11";"26427380";"ATENDIMENTO@CLL.COM.BR";"";""
// "13966779";"0001";"46";"1";"TRANSPORTES IGLEZIAS";"02";"20110707";"00";"";"";"20110707";"4930201";"4930202";"RUA";"MARIA ZILDA SALUSTRIANO DE FREITAS";"511";"";"JD SUMAREZINHO";"13185012";"SP";"2951";"19";"32720466";"";"";"";"";"escritorioiris@escritorioiris.com.br";"";""
// "13966914";"0001";"53";"1";"";"02";"20110708";"00";"";"";"20110708";"8220200";"8291100,7319002,4752100";"RUA";"AZEVEDO PORTUGAL";"484";"";"CENTRO";"85010200";"PR";"7583";"42";"36231960";"";"";"";"";"EXATACONTABIL1@UOL.COM.BR";"";""
// "13967023";"0001";"11";"1";"";"08";"20180201";"67";"";"";"20110719";"4530703";"4520001,4530704";"RUA";"ORLANDO JOSE PEREIRA";"155";"";"CENTRO";"16240000";"SP";"7069";"18";"36441516";"";"";"";"";"";"";""
// "20456684";"0001";"48";"1";"";"08";"20191014";"01";"";"";"20140616";"1412601";"";"RUA";"EURICO CORREA";"334";"";"RECANTO AZUL";"14990000";"SP";"6531";"17";"96073717";"";"";"";"";"ecirapua@hotmail.com";"";""
// "13967224";"0001";"19";"1";"J K STUDIO";"02";"20110719";"00";"";"";"20110719";"7319002";"8230001";"RUA";"JOAO BEUX SOBRINHO";"283";"";"PERPETUO SOCORRO";"89990000";"SC";"8333";"49";"33442141";"";"";"";"";"";"";""
// "13967339";"0001";"03";"1";"";"02";"20110719";"00";"";"";"20110719";"9602501";"";"RUA";"FRANCISCO SA";"86";"";"CENTRO";"39800127";"MG";"5371";"33";"35223049";"";"";"";"";"glaubergmp@hotmail.com";"";""
// "13967464";"0001";"13";"1";"METAL AGUIA";"02";"20110701";"00";"";"";"20110701";"2599302";"4672900,4744001,2542000,2543800,2593400";"RUA";"SANTO INACIO";"496";"      QD 17 LT 03";"IPIRANGA";"74453280";"GO";"9373";"62";"35973294";"";"";"";"";"";"";""
// "13967563";"0001";"03";"1";"QSIMULADOS";"02";"20110711";"00";"";"";"20110711";"8599605";"4761001,4762800,8599699";"RUA";"PARANA";"569";"ANDAR 1                   SALA  11";"SAO CRISTOVAO";"85813010";"PR";"7493";"47";"89013869";"";"";"";"";"CONTATO@BEABADOCONCURSO.COM.BR";"";""

// The first 3 fields compose the entire CNPJ, which we'll use as primary key
// We'll merge "Estabelecimentos" and "Empresas" datasets to build this table
type Company struct {
	Cnpj         string `json:"cnpj"`
	Name         string `json:"name"`
	PhantasyName string `json:"phantasyName"`
}
