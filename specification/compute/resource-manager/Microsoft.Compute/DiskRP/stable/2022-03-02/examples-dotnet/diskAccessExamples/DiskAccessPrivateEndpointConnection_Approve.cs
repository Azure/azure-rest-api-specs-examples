using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Approve.json
// this example is just showing the usage of "DiskAccesses_UpdateAPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.
            
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
ComputePrivateEndpointConnectionData data = new ComputePrivateEndpointConnectionData()
{
    ConnectionState = new ComputePrivateLinkServiceConnectionState()
    {
        Status = ComputePrivateEndpointServiceConnectionStatus.Approved,
        Description = "Approving myPrivateEndpointConnection",
    },
};
ArmOperation<ComputePrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
ComputePrivateEndpointConnectionResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ComputePrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
