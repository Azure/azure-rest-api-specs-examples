using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Monitor;
using Azure.ResourceManager.Monitor.Models;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateEndpointConnectionUpdate.json
// this example is just showing the usage of "PrivateEndpointConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MonitorPrivateEndpointConnectionResource created on azure
// for more information of creating MonitorPrivateEndpointConnectionResource, please refer to the document of MonitorPrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "MyResourceGroup";
string scopeName = "MyPrivateLinkScope";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier monitorPrivateEndpointConnectionResourceId = MonitorPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName, privateEndpointConnectionName);
MonitorPrivateEndpointConnectionResource monitorPrivateEndpointConnection = client.GetMonitorPrivateEndpointConnectionResource(monitorPrivateEndpointConnectionResourceId);

// invoke the operation
MonitorPrivateEndpointConnectionData data = new MonitorPrivateEndpointConnectionData()
{
    ConnectionState = new MonitorPrivateLinkServiceConnectionState()
    {
        Status = MonitorPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Approved by johndoe@contoso.com",
    },
};
ArmOperation<MonitorPrivateEndpointConnectionResource> lro = await monitorPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
MonitorPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MonitorPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
