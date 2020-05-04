// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package rule

import (
	"context"
	"sync"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/hook"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/privacy"
)

// DenyUpdateRule is a mutation rule that denies update many operations.
func DenyUpdateRule() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		if m.Op() == ent.OpUpdate {
			return privacy.Denyf("ent/privacy: update operation not allowed")
		}
		return privacy.Skip
	})
}

// DenyPlanetSelfLinkRule is a mutation rule rule that prevents rule self link via neighbor edge.
func DenyPlanetSelfLinkRule() privacy.MutationRule {
	return privacy.PlanetMutationRuleFunc(func(ctx context.Context, m *ent.PlanetMutation) error {
		if !m.Op().Is(ent.OpUpdateOne) {
			return privacy.Skip
		}
		id, exists := m.ID()
		if !exists {
			return privacy.Denyf("ent/privacy: rule id not provided")
		}
		for _, neighbor := range m.NeighborsIDs() {
			if id == neighbor {
				return privacy.Denyf("ent/privacy: planet %d cannot have itself as a neighbor", id)
			}
		}
		return privacy.Skip
	})
}

// FilterZeroPlanetAgeRule is a query rule that filters out planet with age equal to zero.
func FilterZeroAgePlanetRule() privacy.QueryRule {
	return privacy.PlanetQueryRuleFunc(func(ctx context.Context, q *ent.PlanetQuery) error {
		q.Where(planet.AgeNEQ(0))
		return privacy.Skip
	})
}

// FilterIrregularGalaxyRule is a query rule that filters out irregular galaxies.
func FilterIrregularGalaxyRule() privacy.QueryRule {
	return privacy.GalaxyQueryRuleFunc(func(ctx context.Context, q *ent.GalaxyQuery) error {
		q.Where(galaxy.TypeNEQ(galaxy.TypeIrregular))
		return privacy.Skip
	})
}

var logger = struct {
	logf func(string, ...interface{})
	sync.RWMutex
}{
	logf: func(string, ...interface{}) {},
}

// SetMutationLogFunc overrides the logging function used by LogPlanetMutationHook.
func SetMutationLogFunc(f func(string, ...interface{})) func(string, ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	logf := logger.logf
	logger.logf = f
	return logf
}

// LogPlanetMutationHook returns a hook logging planet mutations.
func LogPlanetMutationHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.PlanetFunc(func(ctx context.Context, m *ent.PlanetMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			logger.RLock()
			defer logger.RUnlock()
			logger.logf("planet mutation: type %s, value %v, err %v", m.Op(), value, err)
			return value, err
		})
	}
}
