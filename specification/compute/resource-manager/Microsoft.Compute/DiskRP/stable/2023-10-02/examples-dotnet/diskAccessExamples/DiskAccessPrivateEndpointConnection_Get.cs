using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Get.json
// this example is just showing the usage of "DiskAccesses_GetAPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskAccessResource created on azure
// for more information of creating DiskAccessResource, please refer to the document of DiskAccessResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskAccessName = "myDiskAccess";
ResourceIdentifier diskAccessResourceId = DiskAccessResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskAccessName);
DiskAccessResource diskAccess = client.GetDiskAccessResource(diskAccessResourceId);

// get the collection of this ComputePrivateEndpointConnectionResource
ComputePrivateEndpointConnectionCollection collection = diskAccess.GetComputePrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "myPrivateEndpointConnection";
NullableResponse<ComputePrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
ComputePrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ComputePrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
