using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationMigrationItems_Create.json
// this example is just showing the usage of "ReplicationMigrationItems_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryProtectionContainerResource created on azure
// for more information of creating SiteRecoveryProtectionContainerResource, please refer to the document of SiteRecoveryProtectionContainerResource
string subscriptionId = "cb53d0c3-bd59-4721-89bc-06916a9147ef";
string resourceGroupName = "resourcegroup1";
string resourceName = "migrationvault";
string fabricName = "vmwarefabric1";
string protectionContainerName = "vmwareContainer1";
ResourceIdentifier siteRecoveryProtectionContainerResourceId = SiteRecoveryProtectionContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName);
SiteRecoveryProtectionContainerResource siteRecoveryProtectionContainer = client.GetSiteRecoveryProtectionContainerResource(siteRecoveryProtectionContainerResourceId);

// get the collection of this SiteRecoveryMigrationItemResource
SiteRecoveryMigrationItemCollection collection = siteRecoveryProtectionContainer.GetSiteRecoveryMigrationItems();

// invoke the operation
string migrationItemName = "virtualmachine1";
SiteRecoveryMigrationItemCreateOrUpdateContent content = new SiteRecoveryMigrationItemCreateOrUpdateContent(new EnableMigrationProperties(new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationPolicies/vmwarepolicy1"), new VMwareCbtEnableMigrationContent(
    new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.OffAzure/VMwareSites/vmwaresite1/machines/virtualmachine1"),
    new VMwareCbtDiskContent[]
{
new VMwareCbtDiskContent("disk1", "true", new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.Storage/storageAccounts/logStorageAccount1"), "logStorageSas")
},
    new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.OffAzure/VMwareSites/vmwaresite1/runasaccounts/dataMoverRunAsAccount1"),
    new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.OffAzure/VMwareSites/vmwaresite1/runasaccounts/snapshotRunAsAccount1"),
    new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1"),
    new ResourceIdentifier("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1"))));
ArmOperation<SiteRecoveryMigrationItemResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, migrationItemName, content);
SiteRecoveryMigrationItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryMigrationItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
