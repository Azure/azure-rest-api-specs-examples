using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/stable/2025-09-01/examples/NetworkSiblingSet_Update.json
// this example is just showing the usage of "NetAppResource_UpdateNetworkSiblingSet" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("eastus");
UpdateNetworkSiblingSetContent content = new UpdateNetworkSiblingSetContent("9760acf5-4638-11e7-9bdb-020073ca3333", new ResourceIdentifier("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet"), "12345_44420.8001578125", NetAppNetworkFeature.Standard);
ArmOperation<NetworkSiblingSet> lro = await subscriptionResource.UpdateNetworkSiblingSetNetAppResourceAsync(WaitUntil.Completed, location, content);
NetworkSiblingSet result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
