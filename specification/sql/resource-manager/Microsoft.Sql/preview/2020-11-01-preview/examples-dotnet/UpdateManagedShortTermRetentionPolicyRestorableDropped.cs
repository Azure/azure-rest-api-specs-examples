using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/UpdateManagedShortTermRetentionPolicyRestorableDropped.json
// this example is just showing the usage of "ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorableDroppedManagedDatabaseResource created on azure
// for more information of creating RestorableDroppedManagedDatabaseResource, please refer to the document of RestorableDroppedManagedDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup";
string managedInstanceName = "testsvr";
string restorableDroppedDatabaseId = "testdb,131403269876900000";
ResourceIdentifier restorableDroppedManagedDatabaseResourceId = RestorableDroppedManagedDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, restorableDroppedDatabaseId);
RestorableDroppedManagedDatabaseResource restorableDroppedManagedDatabase = client.GetRestorableDroppedManagedDatabaseResource(restorableDroppedManagedDatabaseResourceId);

// get the collection of this ManagedRestorableDroppedDbBackupShortTermRetentionPolicyResource
ManagedRestorableDroppedDbBackupShortTermRetentionPolicyCollection collection = restorableDroppedManagedDatabase.GetManagedRestorableDroppedDbBackupShortTermRetentionPolicies();

// invoke the operation
ManagedShortTermRetentionPolicyName policyName = ManagedShortTermRetentionPolicyName.Default;
ManagedBackupShortTermRetentionPolicyData data = new ManagedBackupShortTermRetentionPolicyData()
{
    RetentionDays = 14,
};
ArmOperation<ManagedRestorableDroppedDbBackupShortTermRetentionPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyName, data);
ManagedRestorableDroppedDbBackupShortTermRetentionPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedBackupShortTermRetentionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
