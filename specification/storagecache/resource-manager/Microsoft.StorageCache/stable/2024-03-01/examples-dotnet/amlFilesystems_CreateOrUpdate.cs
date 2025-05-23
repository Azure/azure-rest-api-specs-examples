using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StorageCache.Models;
using Azure.ResourceManager.StorageCache;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/amlFilesystems_CreateOrUpdate.json
// this example is just showing the usage of "amlFilesystems_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AmlFileSystemResource
AmlFileSystemCollection collection = resourceGroupResource.GetAmlFileSystems();

// invoke the operation
string amlFileSystemName = "fs1";
AmlFileSystemData data = new AmlFileSystemData(new AzureLocation("eastus"))
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1")] = new UserAssignedIdentity()
        },
    },
    SkuName = "AMLFS-Durable-Premium-250",
    Zones = { "1" },
    StorageCapacityTiB = 16,
    FilesystemSubnet = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub",
    KeyEncryptionKey = new StorageCacheEncryptionKeyVaultKeyReference(new Uri("https://examplekv.vault.azure.net/keys/kvk/3540a47df75541378d3518c6a4bdf5af"), new WritableSubResource
    {
        Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
    }),
    MaintenanceWindow = new AmlFileSystemPropertiesMaintenanceWindow
    {
        DayOfWeek = MaintenanceDayOfWeekType.Friday,
        TimeOfDayUTC = "22:00",
    },
    Hsm = new AmlFileSystemPropertiesHsm
    {
        Settings = new AmlFileSystemHsmSettings("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/containername", "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/loggingcontainername")
        {
            ImportPrefixesInitial = { "/" },
        },
    },
    RootSquashSettings = new AmlFileSystemRootSquashSettings
    {
        Mode = AmlFileSystemSquashMode.All,
        NoSquashNidLists = "10.0.0.[5-6]@tcp;10.0.1.2@tcp",
        SquashUID = 99L,
        SquashGID = 99L,
    },
    Tags =
    {
    ["Dept"] = "ContosoAds"
    },
};
ArmOperation<AmlFileSystemResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, amlFileSystemName, data);
AmlFileSystemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AmlFileSystemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
