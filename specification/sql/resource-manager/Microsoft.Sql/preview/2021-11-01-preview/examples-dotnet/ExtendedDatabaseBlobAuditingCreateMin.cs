using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseBlobAuditingCreateMin.json
// this example is just showing the usage of "ExtendedDatabaseBlobAuditingPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "blobauditingtest-4799";
string serverName = "blobauditingtest-6440";
string databaseName = "testdb";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// get the collection of this ExtendedDatabaseBlobAuditingPolicyResource
ExtendedDatabaseBlobAuditingPolicyCollection collection = sqlDatabase.GetExtendedDatabaseBlobAuditingPolicies();

// invoke the operation
BlobAuditingPolicyName blobAuditingPolicyName = BlobAuditingPolicyName.Default;
ExtendedDatabaseBlobAuditingPolicyData data = new ExtendedDatabaseBlobAuditingPolicyData
{
    State = BlobAuditingPolicyState.Enabled,
    StorageEndpoint = "https://mystorage.blob.core.windows.net",
    StorageAccountAccessKey = "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
};
ArmOperation<ExtendedDatabaseBlobAuditingPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, blobAuditingPolicyName, data);
ExtendedDatabaseBlobAuditingPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExtendedDatabaseBlobAuditingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
