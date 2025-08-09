using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.Models;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMin.json
// this example is just showing the usage of "ServerSecurityAlertPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlServerResource created on azure
// for more information of creating PostgreSqlServerResource, please refer to the document of PostgreSqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securityalert-4799";
string serverName = "securityalert-6440";
ResourceIdentifier postgreSqlServerResourceId = PostgreSqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlServerResource postgreSqlServer = client.GetPostgreSqlServerResource(postgreSqlServerResourceId);

// get the collection of this PostgreSqlServerSecurityAlertPolicyResource
PostgreSqlServerSecurityAlertPolicyCollection collection = postgreSqlServer.GetPostgreSqlServerSecurityAlertPolicies();

// invoke the operation
PostgreSqlSecurityAlertPolicyName securityAlertPolicyName = PostgreSqlSecurityAlertPolicyName.Default;
PostgreSqlServerSecurityAlertPolicyData data = new PostgreSqlServerSecurityAlertPolicyData
{
    State = PostgreSqlServerSecurityAlertPolicyState.Disabled,
    SendToEmailAccountAdmins = true,
};
ArmOperation<PostgreSqlServerSecurityAlertPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, securityAlertPolicyName, data);
PostgreSqlServerSecurityAlertPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlServerSecurityAlertPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
