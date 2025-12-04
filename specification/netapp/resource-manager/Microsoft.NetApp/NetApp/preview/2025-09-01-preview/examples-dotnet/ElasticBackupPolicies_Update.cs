using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/ElasticBackupPolicies_Update.json
// this example is just showing the usage of "ElasticBackupPolicies_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticBackupPolicyResource created on azure
// for more information of creating ElasticBackupPolicyResource, please refer to the document of ElasticBackupPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string backupPolicyName = "backupPolicyName";
ResourceIdentifier elasticBackupPolicyResourceId = ElasticBackupPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupPolicyName);
ElasticBackupPolicyResource elasticBackupPolicy = client.GetElasticBackupPolicyResource(elasticBackupPolicyResourceId);

// invoke the operation
ElasticBackupPolicyPatch patch = new ElasticBackupPolicyPatch
{
    Properties = new ElasticBackupPolicyUpdateProperties
    {
        DailyBackupsToKeep = 5,
        WeeklyBackupsToKeep = 10,
        MonthlyBackupsToKeep = 10,
        PolicyState = ElasticBackupPolicyState.Enabled,
    },
};
ArmOperation<ElasticBackupPolicyResource> lro = await elasticBackupPolicy.UpdateAsync(WaitUntil.Completed, patch);
ElasticBackupPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticBackupPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
