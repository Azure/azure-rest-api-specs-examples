using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StorageCache;
using Azure.ResourceManager.StorageCache.Models;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/Caches_CreateOrUpdate_ldap_only.json
// this example is just showing the usage of "Caches_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StorageCacheResource
StorageCacheCollection collection = resourceGroupResource.GetStorageCaches();

// invoke the operation
string cacheName = "sc1";
StorageCacheData data = new StorageCacheData(new AzureLocation("westus"))
{
    SkuName = "Standard_2G",
    CacheSizeGB = 3072,
    Subnet = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1"),
    UpgradeSettings = new StorageCacheUpgradeSettings()
    {
        EnableUpgradeSchedule = true,
        ScheduledOn = DateTimeOffset.Parse("2022-04-26T18:25:43.511Z"),
    },
    EncryptionSettings = new StorageCacheEncryptionSettings()
    {
        KeyEncryptionKey = new StorageCacheEncryptionKeyVaultKeyReference(new Uri("https://keyvault-cmk.vault.azure.net/keys/key2048/test"), new WritableSubResource()
        {
            Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
        }),
    },
    SecurityAccessPolicies =
    {
    new NfsAccessPolicy("default",new NfsAccessRule[]
    {
    new NfsAccessRule(NfsAccessRuleScope.Default,NfsAccessRuleAccess.ReadWrite)
    {
    AllowSuid = false,
    AllowSubmountAccess = true,
    EnableRootSquash = false,
    }
    })
    },
    DirectoryServicesSettings = new StorageCacheDirectorySettings()
    {
        UsernameDownload = new StorageCacheUsernameDownloadSettings()
        {
            EnableExtendedGroups = true,
            UsernameSource = StorageCacheUsernameSourceType.Ldap,
            LdapServer = "192.0.2.12",
            LdapBaseDN = "dc=contosoad,dc=contoso,dc=local",
            Credentials = new StorageCacheUsernameDownloadCredential()
            {
                BindDistinguishedName = "cn=ldapadmin,dc=contosoad,dc=contoso,dc=local",
                BindPassword = "<bindPassword>",
            },
        },
    },
    Tags =
    {
    ["Dept"] = "Contoso",
    },
};
ArmOperation<StorageCacheResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, cacheName, data);
StorageCacheResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageCacheData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
