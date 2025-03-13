using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationMigrationItems_TestMigrateCleanup.json
// this example is just showing the usage of "ReplicationMigrationItems_TestMigrateCleanup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryMigrationItemResource created on azure
// for more information of creating SiteRecoveryMigrationItemResource, please refer to the document of SiteRecoveryMigrationItemResource
string subscriptionId = "cb53d0c3-bd59-4721-89bc-06916a9147ef";
string resourceGroupName = "resourcegroup1";
string resourceName = "migrationvault";
string fabricName = "vmwarefabric1";
string protectionContainerName = "vmwareContainer1";
string migrationItemName = "virtualmachine1";
ResourceIdentifier siteRecoveryMigrationItemResourceId = SiteRecoveryMigrationItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, migrationItemName);
SiteRecoveryMigrationItemResource siteRecoveryMigrationItem = client.GetSiteRecoveryMigrationItemResource(siteRecoveryMigrationItemResourceId);

// invoke the operation
TestMigrateCleanupContent content = new TestMigrateCleanupContent(new TestMigrateCleanupProperties
{
    Comments = "Test Failover Cleanup",
});
ArmOperation<SiteRecoveryMigrationItemResource> lro = await siteRecoveryMigrationItem.TestMigrateCleanupAsync(WaitUntil.Completed, content);
SiteRecoveryMigrationItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryMigrationItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
