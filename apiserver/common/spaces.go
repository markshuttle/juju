// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package common

import (
	"github.com/juju/errors"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/context"
)

// ReloadSpacesState defines an in situ point of use type for ReloadSpaces
type ReloadSpacesState interface {
	// SaveSpacesFromProvider loads providerSpaces into state.
	SaveSpacesFromProvider([]network.SpaceInfo) error

	// SaveSubnetsFromProvider loads subnets into state.
	SaveSubnetsFromProvider([]network.SubnetInfo, string) error
}

// ReloadSpaces loads spaces and subnets from provider specified by environ into state.
// Currently it's an append-only operation, no spaces/subnets are deleted.
func ReloadSpaces(ctx context.ProviderCallContext, state ReloadSpacesState, environ environs.BootstrapEnviron) error {
	netEnviron, ok := environs.SupportsNetworking(environ)
	if !ok {
		return errors.NotSupportedf("spaces discovery in a non-networking environ")
	}

	canDiscoverSpaces, err := netEnviron.SupportsSpaceDiscovery(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	if canDiscoverSpaces {
		spaces, err := netEnviron.Spaces(ctx)
		if err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(state.SaveSpacesFromProvider(spaces))
	}

	logger.Debugf("environ does not support space discovery, falling back to subnet discovery")
	subnets, err := netEnviron.Subnets(ctx, instance.UnknownId, nil)
	if err != nil {
		return errors.Trace(err)
	}
	return errors.Trace(state.SaveSubnetsFromProvider(subnets, ""))
}
