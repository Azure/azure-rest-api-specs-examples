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

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ServerCreateWithDataEncryptionEnabled.json
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
string serverName = "pgtestsvc4";
PostgreSqlFlexibleServerData data = new PostgreSqlFlexibleServerData(new AzureLocation("westus"))
{
    Sku = new PostgreSqlFlexibleServerSku("Standard_D4s_v3", PostgreSqlFlexibleServerSkuTier.GeneralPurpose),
    Identity = new PostgreSqlFlexibleServerUserAssignedIdentity(PostgreSqlFlexibleServerIdentityType.UserAssigned)
    {
        UserAssignedIdentities =
        {
        ["/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"] = new UserAssignedIdentity(),
        },
    },
    AdministratorLogin = "cloudsa",
    AdministratorLoginPassword = "password",
    Version = PostgreSqlFlexibleServerVersion.Ver12,
    Storage = new PostgreSqlFlexibleServerStorage()
    {
        StorageSizeInGB = 512,
        AutoGrow = StorageAutoGrow.Disabled,
    },
    DataEncryption = new PostgreSqlFlexibleServerDataEncryption()
    {
        PrimaryKeyUri = new Uri("https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787"),
        PrimaryUserAssignedIdentityId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
        GeoBackupKeyUri = new Uri(""),
        GeoBackupUserAssignedIdentityId = "",
        KeyType = PostgreSqlFlexibleServerKeyType.AzureKeyVault,
    },
    Backup = new PostgreSqlFlexibleServerBackupProperties()
    {
        BackupRetentionDays = 7,
        GeoRedundantBackup = PostgreSqlFlexibleServerGeoRedundantBackupEnum.Disabled,
    },
    Network = new PostgreSqlFlexibleServerNetwork()
    {
        DelegatedSubnetResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
        PrivateDnsZoneArmResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone.postgres.database.azure.com"),
    },
    HighAvailability = new PostgreSqlFlexibleServerHighAvailability()
    {
        Mode = PostgreSqlFlexibleServerHighAvailabilityMode.ZoneRedundant,
    },
    AvailabilityZone = "1",
    CreateMode = PostgreSqlFlexibleServerCreateMode.Create,
    Tags =
    {
    ["ElasticServer"] = "1",
    },
};
ArmOperation<PostgreSqlFlexibleServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, data);
PostgreSqlFlexibleServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
