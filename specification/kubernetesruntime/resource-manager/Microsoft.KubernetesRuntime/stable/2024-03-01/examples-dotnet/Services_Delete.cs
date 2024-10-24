using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerOrchestratorRuntime.Models;
using Azure.ResourceManager.ContainerOrchestratorRuntime;

// Generated from example definition: 2024-03-01/Services_Delete.json
// this example is just showing the usage of "ServiceResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectedClusterServiceResource created on azure
// for more information of creating ConnectedClusterServiceResource, please refer to the document of ConnectedClusterServiceResource
string resourceUri = "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1";
string serviceName = "storageclass";
ResourceIdentifier connectedClusterServiceResourceId = ConnectedClusterServiceResource.CreateResourceIdentifier(resourceUri, serviceName);
ConnectedClusterServiceResource connectedClusterService = client.GetConnectedClusterServiceResource(connectedClusterServiceResourceId);

// invoke the operation
await connectedClusterService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
