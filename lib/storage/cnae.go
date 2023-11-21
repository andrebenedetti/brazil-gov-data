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
