using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RedisEnterprise;
using Azure.ResourceManager.RedisEnterprise.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseCreate.json
// this example is just showing the usage of "RedisEnterprise_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this RedisEnterpriseClusterResource
RedisEnterpriseClusterCollection collection = resourceGroupResource.GetRedisEnterpriseClusters();

// invoke the operation
string clusterName = "cache1";
RedisEnterpriseClusterData data = new RedisEnterpriseClusterData(new AzureLocation("West US"), new RedisEnterpriseSku(RedisEnterpriseSkuName.EnterpriseFlashF300)
{
    Capacity = 3,
})
{
    Zones =
    {
    "1","2","3"
    },
    MinimumTlsVersion = RedisEnterpriseTlsVersion.Tls1_2,
    Tags =
    {
    ["tag1"] = "value1",
    },
};
ArmOperation<RedisEnterpriseClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
RedisEnterpriseClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisEnterpriseClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
