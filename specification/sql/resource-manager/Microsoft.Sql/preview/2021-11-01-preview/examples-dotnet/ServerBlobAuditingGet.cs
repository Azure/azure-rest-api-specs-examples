using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerBlobAuditingGet.json
// this example is just showing the usage of "ServerBlobAuditingPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerBlobAuditingPolicyResource created on azure
// for more information of creating SqlServerBlobAuditingPolicyResource, please refer to the document of SqlServerBlobAuditingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "blobauditingtest-4799";
string serverName = "blobauditingtest-6440";
BlobAuditingPolicyName blobAuditingPolicyName = BlobAuditingPolicyName.Default;
ResourceIdentifier sqlServerBlobAuditingPolicyResourceId = SqlServerBlobAuditingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, blobAuditingPolicyName);
SqlServerBlobAuditingPolicyResource sqlServerBlobAuditingPolicy = client.GetSqlServerBlobAuditingPolicyResource(sqlServerBlobAuditingPolicyResourceId);

// invoke the operation
SqlServerBlobAuditingPolicyResource result = await sqlServerBlobAuditingPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerBlobAuditingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
