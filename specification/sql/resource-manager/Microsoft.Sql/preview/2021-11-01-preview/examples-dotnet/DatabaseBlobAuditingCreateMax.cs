using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DatabaseBlobAuditingCreateMax.json
// this example is just showing the usage of "DatabaseBlobAuditingPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseBlobAuditingPolicyResource created on azure
// for more information of creating SqlDatabaseBlobAuditingPolicyResource, please refer to the document of SqlDatabaseBlobAuditingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "blobauditingtest-4799";
string serverName = "blobauditingtest-6440";
string databaseName = "testdb";
BlobAuditingPolicyName blobAuditingPolicyName = BlobAuditingPolicyName.Default;
ResourceIdentifier sqlDatabaseBlobAuditingPolicyResourceId = SqlDatabaseBlobAuditingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, blobAuditingPolicyName);
SqlDatabaseBlobAuditingPolicyResource sqlDatabaseBlobAuditingPolicy = client.GetSqlDatabaseBlobAuditingPolicyResource(sqlDatabaseBlobAuditingPolicyResourceId);

// invoke the operation
SqlDatabaseBlobAuditingPolicyData data = new SqlDatabaseBlobAuditingPolicyData()
{
    RetentionDays = 6,
    AuditActionsAndGroups =
    {
    "DATABASE_LOGOUT_GROUP","DATABASE_ROLE_MEMBER_CHANGE_GROUP","UPDATE on database::TestDatabaseName by public"
    },
    IsStorageSecondaryKeyInUse = false,
    IsAzureMonitorTargetEnabled = true,
    QueueDelayMs = 4000,
    State = BlobAuditingPolicyState.Enabled,
    StorageEndpoint = "https://mystorage.blob.core.windows.net",
    StorageAccountAccessKey = "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    StorageAccountSubscriptionId = Guid.Parse("00000000-1234-0000-5678-000000000000"),
};
ArmOperation<SqlDatabaseBlobAuditingPolicyResource> lro = await sqlDatabaseBlobAuditingPolicy.UpdateAsync(WaitUntil.Completed, data);
SqlDatabaseBlobAuditingPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlDatabaseBlobAuditingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
