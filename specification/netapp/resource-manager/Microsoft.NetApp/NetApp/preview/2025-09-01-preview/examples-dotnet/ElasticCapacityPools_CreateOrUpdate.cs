using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/ElasticCapacityPools_CreateOrUpdate.json
// this example is just showing the usage of "ElasticCapacityPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticAccountResource created on azure
// for more information of creating ElasticAccountResource, please refer to the document of ElasticAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
ResourceIdentifier elasticAccountResourceId = ElasticAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
ElasticAccountResource elasticAccount = client.GetElasticAccountResource(elasticAccountResourceId);

// get the collection of this ElasticCapacityPoolResource
ElasticCapacityPoolCollection collection = elasticAccount.GetElasticCapacityPools();

// invoke the operation
string poolName = "pool1";
ElasticCapacityPoolData data = new ElasticCapacityPoolData(new AzureLocation("eastus"))
{
    Properties = new ElasticCapacityPoolProperties(4398046511104L, ElasticServiceLevel.ZoneRedundant, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
    {
        ActiveDirectoryConfigResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/activeDirectoryConfigs/activeDirectoryConfig1"),
    },
    Zones = { "1", "2", "3" },
};
ArmOperation<ElasticCapacityPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, poolName, data);
ElasticCapacityPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticCapacityPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
