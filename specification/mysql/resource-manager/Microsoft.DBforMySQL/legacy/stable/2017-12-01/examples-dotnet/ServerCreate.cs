using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;
using Azure.ResourceManager.MySql.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerCreate.json
// this example is just showing the usage of "Servers_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MySqlServerResource
MySqlServerCollection collection = resourceGroupResource.GetMySqlServers();

// invoke the operation
string serverName = "mysqltestsvc4";
MySqlServerCreateOrUpdateContent content = new MySqlServerCreateOrUpdateContent(new MySqlServerPropertiesForDefaultCreate("cloudsa", "<administratorLoginPassword>")
{
    SslEnforcement = MySqlSslEnforcementEnum.Enabled,
    StorageProfile = new MySqlStorageProfile()
    {
        BackupRetentionDays = 7,
        GeoRedundantBackup = MySqlGeoRedundantBackup.Enabled,
        StorageInMB = 128000,
    },
}, new AzureLocation("westus"))
{
    Sku = new MySqlSku("GP_Gen5_2")
    {
        Tier = MySqlSkuTier.GeneralPurpose,
        Capacity = 2,
        Family = "Gen5",
    },
    Tags =
    {
    ["ElasticServer"] = "1",
    },
};
ArmOperation<MySqlServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, content);
MySqlServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
