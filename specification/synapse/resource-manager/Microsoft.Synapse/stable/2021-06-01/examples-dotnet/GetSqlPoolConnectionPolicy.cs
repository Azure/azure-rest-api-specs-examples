using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolConnectionPolicy.json
// this example is just showing the usage of "SqlPoolConnectionPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSqlPoolConnectionPolicyResource created on azure
// for more information of creating SynapseSqlPoolConnectionPolicyResource, please refer to the document of SynapseSqlPoolConnectionPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "blobauditingtest-6852";
string workspaceName = "blobauditingtest-2080";
string sqlPoolName = "testdb";
SqlPoolConnectionPolicyName connectionPolicyName = SqlPoolConnectionPolicyName.Default;
ResourceIdentifier synapseSqlPoolConnectionPolicyResourceId = SynapseSqlPoolConnectionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName, connectionPolicyName);
SynapseSqlPoolConnectionPolicyResource synapseSqlPoolConnectionPolicy = client.GetSynapseSqlPoolConnectionPolicyResource(synapseSqlPoolConnectionPolicyResourceId);

// invoke the operation
SynapseSqlPoolConnectionPolicyResource result = await synapseSqlPoolConnectionPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseSqlPoolConnectionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
