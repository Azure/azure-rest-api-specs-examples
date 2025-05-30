using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/PrivateEndpointConnection_UpdateByHostPool.json
// this example is just showing the usage of "PrivateEndpointConnections_UpdateByHostPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HostPoolResource created on azure
// for more information of creating HostPoolResource, please refer to the document of HostPoolResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostPool1";
ResourceIdentifier hostPoolResourceId = HostPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName);
HostPoolResource hostPool = client.GetHostPoolResource(hostPoolResourceId);

// get the collection of this HostPoolPrivateEndpointConnectionResource
HostPoolPrivateEndpointConnectionCollection collection = hostPool.GetHostPoolPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b";
DesktopVirtualizationPrivateEndpointConnection connection = new DesktopVirtualizationPrivateEndpointConnection
{
    ConnectionState = new DesktopVirtualizationPrivateLinkServiceConnectionState
    {
        Status = DesktopVirtualizationPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Approved by admin@consoto.com",
        ActionsRequired = "None",
    },
};
ArmOperation<HostPoolPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, connection);
HostPoolPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DesktopVirtualizationPrivateEndpointConnectionDataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
