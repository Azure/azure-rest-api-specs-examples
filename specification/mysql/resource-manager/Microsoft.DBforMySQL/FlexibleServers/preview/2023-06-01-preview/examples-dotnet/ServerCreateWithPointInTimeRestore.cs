using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MySql.FlexibleServers;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2023-06-01-preview/examples/ServerCreateWithPointInTimeRestore.json
// this example is just showing the usage of "Servers_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TargetResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MySqlFlexibleServerResource
MySqlFlexibleServerCollection collection = resourceGroupResource.GetMySqlFlexibleServers();

// invoke the operation
string serverName = "targetserver";
MySqlFlexibleServerData data = new MySqlFlexibleServerData(new AzureLocation("SoutheastAsia"))
{
    Sku = new MySqlFlexibleServerSku("Standard_D14_v2", MySqlFlexibleServerSkuTier.GeneralPurpose),
    CreateMode = MySqlFlexibleServerCreateMode.PointInTimeRestore,
    SourceServerResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver"),
    RestorePointInTime = DateTimeOffset.Parse("2021-06-24T00:00:37.467Z"),
    Tags =
    {
    ["num"] = "1",
    },
};
ArmOperation<MySqlFlexibleServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, data);
MySqlFlexibleServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
