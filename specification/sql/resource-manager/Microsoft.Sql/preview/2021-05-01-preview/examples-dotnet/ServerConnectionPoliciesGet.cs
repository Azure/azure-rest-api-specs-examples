using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ServerConnectionPoliciesGet.json
// this example is just showing the usage of "ServerConnectionPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerConnectionPolicyResource created on azure
// for more information of creating SqlServerConnectionPolicyResource, please refer to the document of SqlServerConnectionPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "rgtest-12";
string serverName = "servertest-6285";
ConnectionPolicyName connectionPolicyName = ConnectionPolicyName.Default;
ResourceIdentifier sqlServerConnectionPolicyResourceId = SqlServerConnectionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, connectionPolicyName);
SqlServerConnectionPolicyResource sqlServerConnectionPolicy = client.GetSqlServerConnectionPolicyResource(sqlServerConnectionPolicyResourceId);

// invoke the operation
SqlServerConnectionPolicyResource result = await sqlServerConnectionPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerConnectionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
