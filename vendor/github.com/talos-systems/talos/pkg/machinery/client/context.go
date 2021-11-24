// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package client

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// WithNodes wraps the context with metadata to send request to set of nodes.
func WithNodes(ctx context.Context, nodes ...string) context.Context {
	md, _ := metadata.FromOutgoingContext(ctx)

	// overwrite any previous nodes in the context metadata with new value
	md = md.Copy()
	md.Set("nodes", nodes...)

	return metadata.NewOutgoingContext(ctx, md)
}
