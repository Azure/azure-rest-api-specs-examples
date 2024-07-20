using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/privateEndpoint/PrivateEndpointConnection_Delete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputePrivateEndpointConnectionResource created on azure
// for more information of creating HybridComputePrivateEndpointConnectionResource, please refer to the document of HybridComputePrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myResourceGroup";
string scopeName = "myPrivateLinkScope";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier hybridComputePrivateEndpointConnectionResourceId = HybridComputePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName, privateEndpointConnectionName);
HybridComputePrivateEndpointConnectionResource hybridComputePrivateEndpointConnection = client.GetHybridComputePrivateEndpointConnectionResource(hybridComputePrivateEndpointConnectionResourceId);

// invoke the operation
await hybridComputePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
