package test

import (
	"testing"

	"github.com/mstrYoda/go-arctest/pkg/arctest"
)

// Test Case 1: Same as documentation
// const (
// 	domainPattern         = "^domain/.*$"
// 	applicationPattern    = "^app/.*$"
// 	infrastructurePattern = "^infra/.*$"
// )

// Test Case 2: Without ^
const (
	domainPattern         = "domain.*$"
	applicationPattern    = "app.*$"
	infrastructurePattern = "infra.*$"
)

func TestArchitecture(t *testing.T) {
	arch, err := arctest.New("../")
	if err != nil {
		t.Fatalf("Failed to create architecture: %v", err)
	}

	err = arch.ParsePackages()
	if err != nil {
		t.Fatalf("Failed to parse packages: %v", err)
	}

	// Architecture Rule: DOMAIN should not depend on INFRA
	domainDoesNotDependOnInfra, err := arch.DoesNotDependOn(domainPattern, infrastructurePattern)
	if err != nil {
		t.Fatalf("Failed to create dependency rule: %v", err)
	}

	// Architecture Rule: APPLICATION should not depend on INFRA
	appDoesNotDependOnInfra, err := arch.DoesNotDependOn(applicationPattern, infrastructurePattern)
	if err != nil {
		t.Fatalf("Failed to create dependency rule: %v", err)
	}

	// Expectation: FAIL
	appDoesNotDependOnDomain, err := arch.DoesNotDependOn(applicationPattern, domainPattern)
	if err != nil {
		t.Fatalf("Failed to create dependency rule: %v", err)
	}

	// Validate dependencies
	valid, violations := arch.ValidateDependenciesWithRules([]*arctest.DependencyRule{
		domainDoesNotDependOnInfra,
		appDoesNotDependOnInfra,
		appDoesNotDependOnDomain,
	})
	if !valid {
		for _, violation := range violations {
			t.Errorf("Dependency violation: %s", violation)
		}
	}
}
