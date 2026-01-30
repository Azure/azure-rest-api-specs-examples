using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersCreateReplica.json
// this example is just showing the usage of "Servers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PostgreSqlFlexibleServerResource
PostgreSqlFlexibleServerCollection collection = resourceGroupResource.GetPostgreSqlFlexibleServers();

// invoke the operation
string serverName = "exampleserver";
PostgreSqlFlexibleServerData data = new PostgreSqlFlexibleServerData(new AzureLocation("eastus"))
{
    Identity = new PostgreSqlFlexibleServerUserAssignedIdentity(PostgreSqlFlexibleServerIdentityType.UserAssigned)
    {
        UserAssignedIdentities =
        {
        ["/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity"] = new UserAssignedIdentity()
        },
    },
    DataEncryption = new PostgreSqlFlexibleServerDataEncryption
    {
        PrimaryKeyUri = new Uri("https://exampleprimarykeyvault.vault.azure.net/keys/examplekey/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
        PrimaryUserAssignedIdentityId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity"),
        GeoBackupKeyUri = new Uri(""),
        GeoBackupUserAssignedIdentityId = "",
        KeyType = PostgreSqlFlexibleServerKeyType.AzureKeyVault,
    },
    SourceServerResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/examplesourceserver"),
    PointInTimeUtc = DateTimeOffset.Parse("2025-06-01T18:35:22.123456Z"),
    CreateMode = PostgreSqlFlexibleServerCreateMode.Replica,
};
ArmOperation<PostgreSqlFlexibleServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, data);
PostgreSqlFlexibleServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
