using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/NetworkSiblingSet_Query.json
// this example is just showing the usage of "NetAppResource_QueryNetworkSiblingSet" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("eastus");
QueryNetworkSiblingSetContent content = new QueryNetworkSiblingSetContent("9760acf5-4638-11e7-9bdb-020073ca3333", new ResourceIdentifier("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet"));
NetworkSiblingSet result = await subscriptionResource.QueryNetworkSiblingSetNetAppResourceAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
