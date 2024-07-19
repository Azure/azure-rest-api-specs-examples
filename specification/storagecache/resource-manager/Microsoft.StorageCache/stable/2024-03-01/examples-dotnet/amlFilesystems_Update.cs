using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StorageCache.Models;
using Azure.ResourceManager.StorageCache;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/amlFilesystems_Update.json
// this example is just showing the usage of "amlFilesystems_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AmlFileSystemResource created on azure
// for more information of creating AmlFileSystemResource, please refer to the document of AmlFileSystemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string amlFileSystemName = "fs1";
ResourceIdentifier amlFileSystemResourceId = AmlFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, amlFileSystemName);
AmlFileSystemResource amlFileSystem = client.GetAmlFileSystemResource(amlFileSystemResourceId);

// invoke the operation
AmlFileSystemPatch patch = new AmlFileSystemPatch()
{
    Tags =
    {
    ["Dept"] = "ContosoAds",
    },
    KeyEncryptionKey = new StorageCacheEncryptionKeyVaultKeyReference(new Uri("https://examplekv.vault.azure.net/keys/kvk/3540a47df75541378d3518c6a4bdf5af"), new WritableSubResource()
    {
        Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
    }),
    MaintenanceWindow = new AmlFileSystemUpdatePropertiesMaintenanceWindow()
    {
        DayOfWeek = MaintenanceDayOfWeekType.Friday,
        TimeOfDayUTC = "22:00",
    },
    RootSquashSettings = new AmlFileSystemRootSquashSettings()
    {
        Mode = AmlFileSystemSquashMode.All,
        NoSquashNidLists = "10.0.0.[5-6]@tcp;10.0.1.2@tcp",
        SquashUID = 99,
        SquashGID = 99,
    },
};
ArmOperation<AmlFileSystemResource> lro = await amlFileSystem.UpdateAsync(WaitUntil.Completed, patch);
AmlFileSystemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AmlFileSystemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
