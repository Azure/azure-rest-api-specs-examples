using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/BareMetalMachineKeySets_Get.json
// this example is just showing the usage of "BareMetalMachineKeySets_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudClusterResource created on azure
// for more information of creating NetworkCloudClusterResource, please refer to the document of NetworkCloudClusterResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
ResourceIdentifier networkCloudClusterResourceId = NetworkCloudClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
NetworkCloudClusterResource networkCloudCluster = client.GetNetworkCloudClusterResource(networkCloudClusterResourceId);

// get the collection of this NetworkCloudBareMetalMachineKeySetResource
NetworkCloudBareMetalMachineKeySetCollection collection = networkCloudCluster.GetNetworkCloudBareMetalMachineKeySets();

// invoke the operation
string bareMetalMachineKeySetName = "bareMetalMachineKeySetName";
NullableResponse<NetworkCloudBareMetalMachineKeySetResource> response = await collection.GetIfExistsAsync(bareMetalMachineKeySetName);
NetworkCloudBareMetalMachineKeySetResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkCloudBareMetalMachineKeySetData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
