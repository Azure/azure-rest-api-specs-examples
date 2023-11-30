using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kubernetes;
using Azure.ResourceManager.Kubernetes.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2022-05-01-preview/examples/GetClusterExample.json
// this example is just showing the usage of "ConnectedCluster_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
string resourceGroupName = "k8sc-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ConnectedClusterResource
ConnectedClusterCollection collection = resourceGroupResource.GetConnectedClusters();

// invoke the operation
string clusterName = "testCluster";
NullableResponse<ConnectedClusterResource> response = await collection.GetIfExistsAsync(clusterName);
ConnectedClusterResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ConnectedClusterData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
