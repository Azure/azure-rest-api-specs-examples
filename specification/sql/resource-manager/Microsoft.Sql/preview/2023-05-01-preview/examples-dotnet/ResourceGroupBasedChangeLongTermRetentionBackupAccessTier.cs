using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ResourceGroupBasedChangeLongTermRetentionBackupAccessTier.json
// this example is just showing the usage of "LongTermRetentionBackups_ChangeAccessTierByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupLongTermRetentionBackupResource created on azure
// for more information of creating ResourceGroupLongTermRetentionBackupResource, please refer to the document of ResourceGroupLongTermRetentionBackupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroupName";
AzureLocation locationName = new AzureLocation("japaneast");
string longTermRetentionServerName = "serverName";
string longTermRetentionDatabaseName = "databaseName";
string backupName = "55555555-6666-7777-8888-999999999999;131637960820000000;Archive";
ResourceIdentifier resourceGroupLongTermRetentionBackupResourceId = ResourceGroupLongTermRetentionBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, locationName, longTermRetentionServerName, longTermRetentionDatabaseName, backupName);
ResourceGroupLongTermRetentionBackupResource resourceGroupLongTermRetentionBackup = client.GetResourceGroupLongTermRetentionBackupResource(resourceGroupLongTermRetentionBackupResourceId);

// invoke the operation
ChangeLongTermRetentionBackupAccessTierParameters changeLongTermRetentionBackupAccessTierParameters = new ChangeLongTermRetentionBackupAccessTierParameters("Hot", "Copy");
ArmOperation<ResourceGroupLongTermRetentionBackupResource> lro = await resourceGroupLongTermRetentionBackup.ChangeAccessTierByResourceGroupAsync(WaitUntil.Completed, changeLongTermRetentionBackupAccessTierParameters);
ResourceGroupLongTermRetentionBackupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LongTermRetentionBackupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
