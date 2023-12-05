using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ServerUpdateWithDataEncryptionEnabled.json
// this example is just showing the usage of "Servers_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerResource created on azure
// for more information of creating PostgreSqlFlexibleServerResource, please refer to the document of PostgreSqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "pgtestsvc4";
ResourceIdentifier postgreSqlFlexibleServerResourceId = PostgreSqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlFlexibleServerResource postgreSqlFlexibleServer = client.GetPostgreSqlFlexibleServerResource(postgreSqlFlexibleServerResourceId);

// invoke the operation
PostgreSqlFlexibleServerPatch patch = new PostgreSqlFlexibleServerPatch()
{
    Sku = new PostgreSqlFlexibleServerSku("Standard_D8s_v3", PostgreSqlFlexibleServerSkuTier.GeneralPurpose),
    Identity = new PostgreSqlFlexibleServerUserAssignedIdentity(PostgreSqlFlexibleServerIdentityType.UserAssigned)
    {
        UserAssignedIdentities =
        {
        ["/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-usermanagedidentity"] = new UserAssignedIdentity(),
        ["/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"] = new UserAssignedIdentity(),
        },
    },
    AdministratorLoginPassword = "newpassword",
    Backup = new PostgreSqlFlexibleServerBackupProperties()
    {
        BackupRetentionDays = 20,
    },
    DataEncryption = new PostgreSqlFlexibleServerDataEncryption()
    {
        PrimaryKeyUri = new Uri("https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787"),
        PrimaryUserAssignedIdentityId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
        GeoBackupKeyUri = new Uri("https://test-geo-kv.vault.azure.net/keys/test-key1/66f57315bab34b0189daa113fbc78787"),
        GeoBackupUserAssignedIdentityId = "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-usermanagedidentity",
        KeyType = PostgreSqlFlexibleServerKeyType.AzureKeyVault,
    },
    CreateMode = PostgreSqlFlexibleServerCreateModeForUpdate.Update,
};
ArmOperation<PostgreSqlFlexibleServerResource> lro = await postgreSqlFlexibleServer.UpdateAsync(WaitUntil.Completed, patch);
PostgreSqlFlexibleServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
