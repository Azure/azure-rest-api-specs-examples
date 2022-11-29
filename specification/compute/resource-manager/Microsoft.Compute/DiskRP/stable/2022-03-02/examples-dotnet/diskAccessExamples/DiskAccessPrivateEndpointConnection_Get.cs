using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Get.json
// this example is just showing the usage of "DiskAccesses_GetAPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

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
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
