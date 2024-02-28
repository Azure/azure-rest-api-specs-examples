using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/BackupsUnderBackupVault_List.json
// this example is just showing the usage of "Backups_ListByVault" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppBackupVaultResource created on azure
// for more information of creating NetAppBackupVaultResource, please refer to the document of NetAppBackupVaultResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string backupVaultName = "backupVault1";
ResourceIdentifier netAppBackupVaultResourceId = NetAppBackupVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupVaultName);
NetAppBackupVaultResource netAppBackupVault = client.GetNetAppBackupVaultResource(netAppBackupVaultResourceId);

// get the collection of this NetAppBackupVaultBackupResource
NetAppBackupVaultBackupCollection collection = netAppBackupVault.GetNetAppBackupVaultBackups();

// invoke the operation and iterate over the result
await foreach (NetAppBackupVaultBackupResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppBackupData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
