using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ExtendedDatabaseBlobAuditingGet.json
// this example is just showing the usage of "ExtendedDatabaseBlobAuditingPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExtendedDatabaseBlobAuditingPolicyResource created on azure
// for more information of creating ExtendedDatabaseBlobAuditingPolicyResource, please refer to the document of ExtendedDatabaseBlobAuditingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "blobauditingtest-6852";
string serverName = "blobauditingtest-2080";
string databaseName = "testdb";
BlobAuditingPolicyName blobAuditingPolicyName = BlobAuditingPolicyName.Default;
ResourceIdentifier extendedDatabaseBlobAuditingPolicyResourceId = ExtendedDatabaseBlobAuditingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, blobAuditingPolicyName);
ExtendedDatabaseBlobAuditingPolicyResource extendedDatabaseBlobAuditingPolicy = client.GetExtendedDatabaseBlobAuditingPolicyResource(extendedDatabaseBlobAuditingPolicyResourceId);

// invoke the operation
ExtendedDatabaseBlobAuditingPolicyResource result = await extendedDatabaseBlobAuditingPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExtendedDatabaseBlobAuditingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
