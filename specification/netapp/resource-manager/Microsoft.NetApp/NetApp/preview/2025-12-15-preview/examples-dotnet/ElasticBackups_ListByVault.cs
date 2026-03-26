using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-12-15-preview/examples/ElasticBackups_ListByVault.json
// this example is just showing the usage of "ElasticBackups_ListByVault" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppElasticBackupVaultResource created on azure
// for more information of creating NetAppElasticBackupVaultResource, please refer to the document of NetAppElasticBackupVaultResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string backupVaultName = "backupVault1";
ResourceIdentifier netAppElasticBackupVaultResourceId = NetAppElasticBackupVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupVaultName);
NetAppElasticBackupVaultResource netAppElasticBackupVault = client.GetNetAppElasticBackupVaultResource(netAppElasticBackupVaultResourceId);

// get the collection of this NetAppElasticBackupResource
NetAppElasticBackupCollection collection = netAppElasticBackupVault.GetNetAppElasticBackups();

// invoke the operation and iterate over the result
await foreach (NetAppElasticBackupResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppElasticBackupData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
