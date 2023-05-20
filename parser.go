package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Context int
type Operation int

const (
	PatientContext Context = iota
	UserContext
	SystemContext

	CreateOperation Operation = iota
	ReadOperation
	UpdateOperation
	DeleteOperation
	SearchOperation

	Wildcard = "*"
)

type Scope struct {
	Context    Context           `json:"context"`
	Operations []Operation       `json:"operations"`
	Resource   string            `json:"resource"`
	Params     map[string]string `json:"params"`
}

func main() {
	v := "0.0.1"
	usage := `Usage:
  fhirscope <scope>
  Example: 
  	  fhirscope patient/Observation.rs
	  fhirscope system/*.cruds
	`

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg == "-h" || arg == "--help" {
		fmt.Println(usage)
		os.Exit(0)
	}

	if arg == "-V" || arg == "--version" {
		fmt.Println(v)
		os.Exit(0)
	}

	scope, err := Parse(arg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := json.MarshalIndent(scope, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}

func Parse(scope string) (Scope, error) {
	parts := strings.Split(scope, "/")

	if len(parts) < 2 {
		return Scope{}, fmt.Errorf("invalid scope: %s", scope)
	}

	out := Scope{
		Params: make(map[string]string),
	}

	ctx := parts[0]

	switch ctx {
	case "patient":
		out.Context = PatientContext
	case "user":
		out.Context = UserContext
	case "system":
		out.Context = SystemContext
	default:
		return Scope{}, fmt.Errorf("invalid context expected patient, user or system got: %s", ctx)
	}

	rest := strings.Split(parts[1], ".")

	if len(rest) < 2 {
		return Scope{}, fmt.Errorf("invalid resource or operation: %s", parts[1])
	}

	// resource
	if rest[0] == Wildcard {
		out.Resource = Wildcard
	} else {
		var found bool
		for _, resource := range SupportedResources {
			if resource == rest[0] {
				out.Resource = resource
				found = true
				break
			}
		}
		if !found {
			return Scope{}, fmt.Errorf("invalid resource: %s", rest[0])
		}
	}

	opsWithParams := strings.Split(rest[1], "?")

	if opsWithParams[0] == Wildcard {
		// for backwards compatibility with v1
		out.Operations = append(out.Operations, CreateOperation, ReadOperation, UpdateOperation, DeleteOperation, SearchOperation)
	} else if opsWithParams[0] == "read" {
		out.Operations = append(out.Operations, ReadOperation, SearchOperation)
	} else if opsWithParams[0] == "write" {
		out.Operations = append(out.Operations, CreateOperation, UpdateOperation, DeleteOperation)
	} else {
		ops := strings.Split(opsWithParams[0], "")

		for _, op := range ops {
			switch op {
			case "c":
				out.Operations = append(out.Operations, CreateOperation)
			case "r":
				out.Operations = append(out.Operations, ReadOperation)
			case "u":
				out.Operations = append(out.Operations, UpdateOperation)
			case "d":
				out.Operations = append(out.Operations, DeleteOperation)
			case "s":
				out.Operations = append(out.Operations, SearchOperation)
			default:
				return Scope{}, fmt.Errorf("invalid operation: %s", op)
			}
		}
	}

	if len(opsWithParams) > 1 {
		params := strings.Split(opsWithParams[1], "&")

		for _, param := range params {
			parts := strings.Split(param, "=")

			if len(parts) != 2 {
				return Scope{}, fmt.Errorf("invalid param: %s", param)
			}

			out.Params[parts[0]] = parts[1]
		}
	}
	return out, nil
}

func (c Context) MarshalJSON() ([]byte, error) {
	value := c.String()
	return json.Marshal(value)
}

func (o Operation) MarshalJSON() ([]byte, error) {
	value := o.String()
	return json.Marshal(value)
}

func (c Context) String() string {
	switch c {
	case PatientContext:
		return "patient"
	case UserContext:
		return "user"
	case SystemContext:
		return "system"
	default:
		return ""
	}
}

func (o Operation) String() string {
	switch o {
	case CreateOperation:
		return "c"
	case ReadOperation:
		return "r"
	case UpdateOperation:
		return "u"
	case DeleteOperation:
		return "d"
	case SearchOperation:
		return "s"
	default:
		return ""
	}
}

var SupportedResources = []string{
	"Binary",
	"Bundle",
	"CanonicalResource",
	"CapabilityStatement",
	"CodeSystem",
	"Condition",
	"DomainResource",
	"Immunization",
	"Location",
	"MetadataResource",
	"Observation",
	"OperationDefinition",
	"OperationOutcome",
	"Organization",
	"Parameters",
	"Patient",
	"Practitioner",
	"Questionnaire",
	"QuestionnaireResponse",
	"RelatedPerson",
	"Resource",
	"SearchParameter",
	"StructureDefinition",
	"ValueSet",
	"ActivityDefinition",
	"AuditEvent",
	"Composition",
	"Coverage",
	"CoverageEligibilityRequest",
	"CoverageEligibilityResponse",
	"DocumentReference",
	"Encounter",
	"HealthcareService",
	"ImagingStudy",
	"ImplementationGuide",
	"Library",
	"List",
	"Measure",
	"MeasureReport",
	"Medication",
	"MedicationRequest",
	"MedicationStatement",
	"MessageHeader",
	"NamingSystem",
	"PaymentNotice",
	"PaymentReconciliation",
	"Person",
	"PlanDefinition",
	"PractitionerRole",
	"Procedure",
	"Provenance",
	"RequestOrchestration",
	"ServiceRequest",
	"StructureMap",
	"TestScript",
	"AllergyIntolerance",
	"Appointment",
	"AppointmentResponse",
	"Basic",
	"CompartmentDefinition",
	"ConceptMap",
	"DiagnosticReport",
	"Group",
	"MedicinalProductDefinition",
	"Schedule",
	"Slot",
	"Subscription",
	"Task",
	"VisionPrescription",
	"Account",
	"AdministrableProductDefinition",
	"AdverseEvent",
	"BiologicallyDerivedProduct",
	"CarePlan",
	"CareTeam",
	"Claim",
	"ClaimResponse",
	"ClinicalUseDefinition",
	"Communication",
	"CommunicationRequest",
	"Consent",
	"DetectedIssue",
	"Device",
	"Endpoint",
	"EpisodeOfCare",
	"ExplanationOfBenefit",
	"FamilyMemberHistory",
	"Goal",
	"GraphDefinition",
	"GuidanceResponse",
	"Ingredient",
	"ManufacturedItemDefinition",
	"MedicationAdministration",
	"MedicationDispense",
	"NutritionOrder",
	"PackagedProductDefinition",
	"RegulatedAuthorization",
	"RiskAssessment",
	"Specimen",
	"SubscriptionStatus",
	"SubscriptionTopic",
	"Substance",
	"ActorDefinition",
	"ArtifactAssessment",
	"BodyStructure",
	"ChargeItem",
	"ChargeItemDefinition",
	"Citation",
	"ClinicalImpression",
	"Contract",
	"DeviceDefinition",
	"DeviceMetric",
	"DeviceRequest",
	"DeviceUsage",
	"Evidence",
	"EvidenceVariable",
	"ExampleScenario",
	"Flag",
	"ImagingSelection",
	"ImmunizationEvaluation",
	"ImmunizationRecommendation",
	"MedicationKnowledge",
	"MessageDefinition",
	"MolecularSequence",
	"NutritionIntake",
	"NutritionProduct",
	"ObservationDefinition",
	"OrganizationAffiliation",
	"Requirements",
	"SpecimenDefinition",
	"SubstanceDefinition",
	"SupplyDelivery",
	"SupplyRequest",
	"TerminologyCapabilities",
	"TestReport",
	"Transport",
	"VerificationResult",
	"BiologicallyDerivedProductDispense",
	"ConditionDefinition",
	"DeviceAssociation",
	"DeviceDispense",
	"EncounterHistory",
	"EnrollmentRequest",
	"EnrollmentResponse",
	"EventDefinition",
	"EvidenceReport",
	"FormularyItem",
	"GenomicStudy",
	"InsurancePlan",
	"InventoryItem",
	"InventoryReport",
	"Invoice",
	"Linkage",
	"Permission",
	"ResearchStudy",
	"ResearchSubject",
	"SubstanceNucleicAcid",
	"SubstancePolymer",
	"SubstanceProtein",
	"SubstanceReferenceInformation",
	"SubstanceSourceMaterial",
	"TestPlan",
}
