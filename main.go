package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudfunctions"
	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/FriendlyUser/cooking-app/types"
	"path/filepath"
	"fmt"
)


func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := storage.NewBucket(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		// loop through function config
		p := filepath.FromSlash("functions/helloworld")
		// Create an object in our bucket with our function.
		bucketObjectArgs := &storage.BucketObjectArgs{
			Bucket: bucket.Name,
			Source: pulumi.NewFileArchive(p),
		}
		bucketObject, err := storage.NewBucketObject(ctx, "helloworld-zip", bucketObjectArgs)
		if err != nil {
			return err
		}
		cfg := &types.FuncConfig{

		}
		if cfg == nil {
			fmt.Println(cfg)
		}
		// load files from cloud source repository
		// Set arguments for creating the function resource.
		args := &cloudfunctions.FunctionArgs{
			Runtime:             pulumi.String("dotnet3"),
			SourceArchiveBucket: bucket.Name,
			SourceArchiveObject: bucketObject.Name,
			// https://codelabs.developers.google.com/codelabs/cloud-functions-csharp#3
			EntryPoint:          pulumi.String("HelloHttp.Function"),
			TriggerHttp:         pulumi.Bool(true),
			AvailableMemoryMb:   pulumi.Int(128),
		}
// cooking-app-dli/github_friendlyuser_cooking-app/main/functions/helloworld/Function.cs

		// Create the function using the args.
		function, err := cloudfunctions.NewFunction(ctx, "HelloHttp", args)
		if err != nil {
			return err
		}

		_, err = cloudfunctions.NewFunctionIamMember(ctx, "invoker", &cloudfunctions.FunctionIamMemberArgs{
			Project:       function.Project,
			Region:        function.Region,
			CloudFunction: function.Name,
			Role:          pulumi.String("roles/cloudfunctions.invoker"),
			Member:        pulumi.String("allUsers"),
		})
		if err != nil {
			return err
		}

		// Export the trigger URL.
		ctx.Export("function", function.HttpsTriggerUrl)
		return nil
	})
}