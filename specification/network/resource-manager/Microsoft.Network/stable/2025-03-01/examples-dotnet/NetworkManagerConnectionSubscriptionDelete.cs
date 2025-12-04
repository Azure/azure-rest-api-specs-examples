using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerConnectionSubscriptionDelete.json
// this example is just showing the usage of "SubscriptionNetworkManagerConnections_Delete" operation, for the dependent resources, they will have to be created separately.

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
await subscriptionNetworkManagerConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
