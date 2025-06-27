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

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/ServerCreateGeoRestoreWithDataEncryptionEnabled.json
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

// get the collection of this PostgreSqlFlexibleServerResource
PostgreSqlFlexibleServerCollection collection = resourceGroupResource.GetPostgreSqlFlexibleServers();

// invoke the operation
string serverName = "pgtestsvc5geo";
PostgreSqlFlexibleServerData data = new PostgreSqlFlexibleServerData(new AzureLocation("eastus"))
{
    Identity = new PostgreSqlFlexibleServerUserAssignedIdentity(PostgreSqlFlexibleServerIdentityType.UserAssigned)
    {
        UserAssignedIdentities =
        {
        ["/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-usermanagedidentity"] = new UserAssignedIdentity(),
        ["/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"] = new UserAssignedIdentity()
        },
    },
    DataEncryption = new PostgreSqlFlexibleServerDataEncryption
    {
        PrimaryKeyUri = new Uri("https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787"),
        PrimaryUserAssignedIdentityId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
        GeoBackupKeyUri = new Uri("https://test-geo-kv.vault.azure.net/keys/test-key1/66f57315bab34b0189daa113fbc78787"),
        GeoBackupUserAssignedIdentityId = "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-usermanagedidentity",
        KeyType = PostgreSqlFlexibleServerKeyType.AzureKeyVault,
    },
    SourceServerResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourcepgservername"),
    PointInTimeUtc = DateTimeOffset.Parse("2021-06-27T00:04:59.4078005+00:00"),
    CreateMode = PostgreSqlFlexibleServerCreateMode.GeoRestore,
};
ArmOperation<PostgreSqlFlexibleServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, data);
PostgreSqlFlexibleServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
