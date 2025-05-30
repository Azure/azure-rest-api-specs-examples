using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupDelete.json
// this example is just showing the usage of "LongTermRetentionManagedInstanceBackups_DeleteByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupLongTermRetentionManagedInstanceBackupResource created on azure
// for more information of creating ResourceGroupLongTermRetentionManagedInstanceBackupResource, please refer to the document of ResourceGroupLongTermRetentionManagedInstanceBackupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testResourceGroup";
AzureLocation locationName = new AzureLocation("japaneast");
string managedInstanceName = "testInstance";
string databaseName = "testDatabase";
string backupName = "55555555-6666-7777-8888-999999999999;131637960820000000";
ResourceIdentifier resourceGroupLongTermRetentionManagedInstanceBackupResourceId = ResourceGroupLongTermRetentionManagedInstanceBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, locationName, managedInstanceName, databaseName, backupName);
ResourceGroupLongTermRetentionManagedInstanceBackupResource resourceGroupLongTermRetentionManagedInstanceBackup = client.GetResourceGroupLongTermRetentionManagedInstanceBackupResource(resourceGroupLongTermRetentionManagedInstanceBackupResourceId);

// invoke the operation
await resourceGroupLongTermRetentionManagedInstanceBackup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
