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

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ServerUpdateWithAadAuthEnabled.json
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
    AdministratorLoginPassword = "newpassword",
    Storage = new PostgreSqlFlexibleServerStorage()
    {
        StorageSizeInGB = 1024,
        AutoGrow = StorageAutoGrow.Disabled,
        Tier = PostgreSqlManagedDiskPerformanceTier.P30,
    },
    Backup = new PostgreSqlFlexibleServerBackupProperties()
    {
        BackupRetentionDays = 20,
    },
    AuthConfig = new PostgreSqlFlexibleServerAuthConfig()
    {
        ActiveDirectoryAuth = PostgreSqlFlexibleServerActiveDirectoryAuthEnum.Enabled,
        PasswordAuth = PostgreSqlFlexibleServerPasswordAuthEnum.Enabled,
        TenantId = Guid.Parse("tttttt-tttt-tttt-tttt-tttttttttttt"),
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
