[![Deploy](https://get.pulumi.com/new/button.svg)](https://app.pulumi.com/new)

# cooking-app

The runtime is golang for pulumi config and dotnet for the functions

Monorepo for Cloud Functions with cooking app in flutter.

To get started, update the pulumi stack configuration gcp project and region should be updated.

To deploy 

```bash
gcloud auth application-default login
pulumi up
```

## Todo

* Add test cases with github actions with codecov (no unit tests for flutter)
* Add multiple functions
* Add environment variables to functions (do some testing)
* Get bare bones flutter add
###  References


* https://github.com/pulumi/examples/blob/master/gcp-go-functions/main.go
* https://cloud.google.com/functions/docs/first-dotnet#creating_a_function
* https://github.com/GoogleCloudPlatform/dotnet-docs-samples/tree/master/functions/helloworld
* https://www.pulumi.com/docs/reference/pkg/gcp/cloudfunctions/function/#environmentvariables_go