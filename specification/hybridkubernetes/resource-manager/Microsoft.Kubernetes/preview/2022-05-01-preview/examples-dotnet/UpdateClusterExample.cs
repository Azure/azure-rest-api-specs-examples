using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kubernetes;
using Azure.ResourceManager.Kubernetes.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2022-05-01-preview/examples/UpdateClusterExample.json
// this example is just showing the usage of "ConnectedCluster_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectedClusterResource created on azure
// for more information of creating ConnectedClusterResource, please refer to the document of ConnectedClusterResource
string subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
string resourceGroupName = "k8sc-rg";
string clusterName = "testCluster";
ResourceIdentifier connectedClusterResourceId = ConnectedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ConnectedClusterResource connectedCluster = client.GetConnectedClusterResource(connectedClusterResourceId);

// invoke the operation
ConnectedClusterPatch patch = new ConnectedClusterPatch()
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
ConnectedClusterResource result = await connectedCluster.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectedClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
