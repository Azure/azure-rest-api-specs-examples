using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceUpdate;
using Azure.ResourceManager.DeviceUpdate.Models;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateEndpointConnections/PrivateEndpointConnection_ListByAccount.json
// this example is just showing the usage of "PrivateEndpointConnections_ListByAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceUpdateAccountResource created on azure
// for more information of creating DeviceUpdateAccountResource, please refer to the document of DeviceUpdateAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
ResourceIdentifier deviceUpdateAccountResourceId = DeviceUpdateAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
DeviceUpdateAccountResource deviceUpdateAccount = client.GetDeviceUpdateAccountResource(deviceUpdateAccountResourceId);

// get the collection of this DeviceUpdatePrivateEndpointConnectionResource
DeviceUpdatePrivateEndpointConnectionCollection collection = deviceUpdateAccount.GetDeviceUpdatePrivateEndpointConnections();

// invoke the operation and iterate over the result
await foreach (DeviceUpdatePrivateEndpointConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DeviceUpdatePrivateEndpointConnectionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
