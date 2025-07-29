using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/BackupsUnderBackupVault_Get.json
// this example is just showing the usage of "Backups_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppBackupVaultResource created on azure
// for more information of creating NetAppBackupVaultResource, please refer to the document of NetAppBackupVaultResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string backupVaultName = "backupVault1";
ResourceIdentifier netAppBackupVaultResourceId = NetAppBackupVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupVaultName);
NetAppBackupVaultResource netAppBackupVault = client.GetNetAppBackupVaultResource(netAppBackupVaultResourceId);

// get the collection of this NetAppBackupVaultBackupResource
NetAppBackupVaultBackupCollection collection = netAppBackupVault.GetNetAppBackupVaultBackups();

// invoke the operation
string backupName = "backup1";
NullableResponse<NetAppBackupVaultBackupResource> response = await collection.GetIfExistsAsync(backupName);
NetAppBackupVaultBackupResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppBackupData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
