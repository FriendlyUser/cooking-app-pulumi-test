package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudfunctions"
	// "github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"fmt"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// LEGACY
		// Create a bucket.
		// bucket, err := storage.NewBucket(ctx, "bucket", nil)
		// if err != nil {
		// 	return err
		// }

		// // Create an object in our bucket with our function.
		// bucketObjectArgs := &storage.BucketObjectArgs{
		// 	Bucket: bucket.Name,
		// 	Source: pulumi.NewFileArchive("functions"),
		// }
		// bucketObject, err := storage.NewBucketObject(ctx, "go-zip", bucketObjectArgs)
		// if err != nil {
		// 	return err
		// }

		// load files from cloud source repository
		// Set arguments for creating the function resource.
		fmt.Println("Failing here")
		args := &cloudfunctions.FunctionArgs{
			Runtime:             pulumi.String("dotnet3"),
			SourceRepository: &cloudfunctions.FunctionSourceRepositoryArgs{
				Url: pulumi.String("https://source.cloud.google.com/cooking-app-dli/github_friendlyuser_cooking-app"),
			},
			// SourceRepository: &cloudfunctions.FunctionSourceRepositoryArgs{
			// 	Url: "https://source.cloud.google.com/cooking-app-dli/github_friendlyuser_cooking-app?authuser=0",
			// },
			EntryPoint:          pulumi.String("HandleAsync"),
			TriggerHttp:         pulumi.Bool(true),
			AvailableMemoryMb:   pulumi.Int(128),
		}


		// Create the function using the args.
		function, err := cloudfunctions.NewFunction(ctx, "basicFunction", args)
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