using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreateGeoRestoreMode.json
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

// get the collection of this PostgreSqlServerResource
PostgreSqlServerCollection collection = resourceGroupResource.GetPostgreSqlServers();

// invoke the operation
string serverName = "targetserver";
PostgreSqlServerCreateOrUpdateContent content = new PostgreSqlServerCreateOrUpdateContent(new PostgreSqlServerPropertiesForGeoRestore(new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforPostgreSQL/servers/sourceserver")), new AzureLocation("westus"))
{
    Sku = new PostgreSqlSku("GP_Gen5_2")
    {
        Tier = PostgreSqlSkuTier.GeneralPurpose,
        Capacity = 2,
        Family = "Gen5",
    },
    Tags =
    {
    ["ElasticServer"] = "1",
    },
};
ArmOperation<PostgreSqlServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, content);
PostgreSqlServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
