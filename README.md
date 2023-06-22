
# fhirscope
> SMART on FHIR app use OAuth scopes to request access. Launch context is a negotiation where a client asks for specific launch context parameters then a server can decide which launch context parameters to provide, using the client’s request as an input into the decision process.

### Build

```bash
make build
```

## Usage
    
Permission to read and search any resource for the current patient (notice the wildcard which will match any resource type)

```bash
bin/fhirscope patient/*.rs	
```

Perform bulk data export across all available data within a FHIR server	

```bash
bin/fhirscope patient/*.rs	
```

Alert engine to monitor all lab observations in a health system	
    
```bash
bin/fhirscope system/Observation.rs
```

## CLI
```bash
  fhirscope <scope>
  Example: 
          fhirscope patient/Observation.rs
          fhirscope system/*.cruds
```

[Specification](https://build.fhir.org/ig/HL7/smart-app-launch/scopes-and-launch-context.html)

