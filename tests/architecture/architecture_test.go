package architecture

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

const entitiesPackage = "btcapp/src/entities" // most clear package

func getHelpers() []string {
	return []string{
		"btcapp/src/file_storage",
		"btcapp/src/gmail_notifier",
		"btcapp/src/logger",
		"btcapp/src/providers",
		"btcapp/src/settings",
	}
}

func getServices() []string {
	return []string{
		"btcapp/src/usecases/currency_rate",
		"btcapp/src/usecases/notifier",
		"btcapp/src/usecases/subscription",
	}
}

func TestArchitecture(t *testing.T) {
	t.Run("The entities package must have no dependencies", func(t *testing.T) {
		mockT := new(testingT)

		for _, p := range getHelpers() {
			archtest.Package(t, entitiesPackage).ShouldNotDependOn(p)
			assertNoError(t, mockT)
		}

		for _, p := range getServices() {
			archtest.Package(t, entitiesPackage).ShouldNotDependOn(p)
			assertNoError(t, mockT)
		}
	})

	t.Run("Services must be independent among ourselves", func(t *testing.T) {
		for _, p1 := range getServices() {
			for _, p2 := range getServices() {
				if p1 == p2 {
					continue
				}

				archtest.Package(t, p1).ShouldNotDependOn(p2)
			}
		}
	})

	t.Run("Helpers must not know about services", func(t *testing.T) {
		for _, s := range getServices() {
			for _, h := range getHelpers() {
				archtest.Package(t, h).ShouldNotDependOn(s)
			}
		}
	})
}
