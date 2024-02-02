package keeper

import (
	"context"

	"crud/x/crud/types"
    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) ListSpecificPosts(ctx context.Context, req *types.QueryListSpecificPostsRequest) (*types.QueryListSpecificPostsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

    var posts []types.Post
    pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
        var post types.Post
        if err := k.cdc.Unmarshal(value, &post); err != nil {
            return false, err
        }

        if post.EventType == req.EventType {
            if accumulate {
                posts = append(posts, post)
            }
            return true, nil
        }
        return false, nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListSpecificPostsResponse{Post: posts, Pagination: pageRes}, nil
}
