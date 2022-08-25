using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Delete.json
// this example is just showing the usage of "DiskAccesses_DeleteAPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ComputePrivateEndpointConnectionResource created on azure
// for more information of creating ComputePrivateEndpointConnectionResource, please refer to the document of ComputePrivateEndpointConnectionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskAccessName = "myDiskAccess";
string privateEndpointConnectionName = "myPrivateEndpointConnection";
ResourceIdentifier computePrivateEndpointConnectionResourceId = ComputePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskAccessName, privateEndpointConnectionName);
ComputePrivateEndpointConnectionResource computePrivateEndpointConnection = client.GetComputePrivateEndpointConnectionResource(computePrivateEndpointConnectionResourceId);

// invoke the operation
await computePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
