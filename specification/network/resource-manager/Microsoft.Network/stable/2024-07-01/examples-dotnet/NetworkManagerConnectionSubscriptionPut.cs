using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerConnectionSubscriptionPut.json
// this example is just showing the usage of "SubscriptionNetworkManagerConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionNetworkManagerConnectionResource created on azure
// for more information of creating SubscriptionNetworkManagerConnectionResource, please refer to the document of SubscriptionNetworkManagerConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string networkManagerConnectionName = "TestNMConnection";
ResourceIdentifier subscriptionNetworkManagerConnectionResourceId = SubscriptionNetworkManagerConnectionResource.CreateResourceIdentifier(subscriptionId, networkManagerConnectionName);
SubscriptionNetworkManagerConnectionResource subscriptionNetworkManagerConnection = client.GetSubscriptionNetworkManagerConnectionResource(subscriptionNetworkManagerConnectionResourceId);

// invoke the operation
NetworkManagerConnectionData data = new NetworkManagerConnectionData
{
    NetworkManagerId = new ResourceIdentifier("/subscriptions/subscriptionC/resourceGroup/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager"),
};
ArmOperation<SubscriptionNetworkManagerConnectionResource> lro = await subscriptionNetworkManagerConnection.UpdateAsync(WaitUntil.Completed, data);
SubscriptionNetworkManagerConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkManagerConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
