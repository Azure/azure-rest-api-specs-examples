package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/replicas/codepackages/get_logs.json
func ExampleCodePackageClient_GetContainerLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmesh.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCodePackageClient().GetContainerLogs(ctx, "sbz_demo", "sbzDocApp", "sbzDocService", "0", "sbzDocCode", &armservicefabricmesh.CodePackageClientGetContainerLogsOptions{Tail: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerLogs = armservicefabricmesh.ContainerLogs{
	// 	Content: to.Ptr(" * Running on http://0.0.0.0:8080/ (Press CTRL+C to quit)\n * Downloading style https://assets-cdn.github.com/assets/frameworks-8f281eb0a8d2308ceb36e714ba3c3aec.css\n * Downloading style https://assets-cdn.github.com/assets/github-a698da0d53574b056d3c79ac732d4a70.css\n * Downloading style https://assets-cdn.github.com/assets/site-83dc1f7ebc9c7461fe1eab799b56c4c4.css\n * Cached all downloads in /root/.grip/cache-4.5.2\n167.220.0.83 - - [06/Apr/2018 07:16:02] \"GET / HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:02] \"GET /__/grip/asset/frameworks-8f281eb0a8d2308ceb36e714ba3c3aec.css HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:02] \"GET /__/grip/asset/site-83dc1f7ebc9c7461fe1eab799b56c4c4.css HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:02] \"GET /__/grip/asset/github-a698da0d53574b056d3c79ac732d4a70.css HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:02] \"GET /__/grip/static/octicons/octicons.css HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:03] \"GET /__/grip/static/octicons/octicons.woff2?ef21c39f0ca9b1b5116e5eb7ac5eabe6 HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:03] \"GET /__/grip/static/favicon.ico HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:05] \"GET /seabreeze-index.md HTTP/1.1\" 200 -\n167.220.0.83 - - [06/Apr/2018 07:16:09] \"GET /seabreeze-api-application_get.md HTTP/1.1\" 200 -\n"),
	// }
}
