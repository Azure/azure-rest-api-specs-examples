using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/MigrationRecoveryPoints_Get.json
// this example is just showing the usage of "MigrationRecoveryPoints_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this MigrationRecoveryPointResource
MigrationRecoveryPointCollection collection = siteRecoveryMigrationItem.GetMigrationRecoveryPoints();

// invoke the operation
string migrationRecoveryPointName = "b22134ea-620c-474b-9fa5-3c1cb47708e3";
NullableResponse<MigrationRecoveryPointResource> response = await collection.GetIfExistsAsync(migrationRecoveryPointName);
MigrationRecoveryPointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MigrationRecoveryPointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
