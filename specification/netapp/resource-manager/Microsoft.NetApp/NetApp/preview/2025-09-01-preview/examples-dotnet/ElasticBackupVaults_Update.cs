using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/ElasticBackupVaults_Update.json
// this example is just showing the usage of "ElasticBackupVaults_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticBackupVaultResource created on azure
// for more information of creating ElasticBackupVaultResource, please refer to the document of ElasticBackupVaultResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string backupVaultName = "backupVault1";
ResourceIdentifier elasticBackupVaultResourceId = ElasticBackupVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupVaultName);
ElasticBackupVaultResource elasticBackupVault = client.GetElasticBackupVaultResource(elasticBackupVaultResourceId);

// invoke the operation
ElasticBackupVaultPatch patch = new ElasticBackupVaultPatch
{
    Tags =
    {
    ["Tag1"] = "Value1"
    },
};
ArmOperation<ElasticBackupVaultResource> lro = await elasticBackupVault.UpdateAsync(WaitUntil.Completed, patch);
ElasticBackupVaultResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticBackupVaultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
