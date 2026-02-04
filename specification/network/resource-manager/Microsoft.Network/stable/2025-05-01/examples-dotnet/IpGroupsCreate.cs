using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/IpGroupsCreate.json
// this example is just showing the usage of "IpGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IPGroupResource
IPGroupCollection collection = resourceGroupResource.GetIPGroups();

// invoke the operation
string ipGroupsName = "ipGroups1";
IPGroupData data = new IPGroupData
{
    IPAddresses = { "13.64.39.16/32", "40.74.146.80/31", "40.74.147.32/28" },
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<IPGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ipGroupsName, data);
IPGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IPGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
