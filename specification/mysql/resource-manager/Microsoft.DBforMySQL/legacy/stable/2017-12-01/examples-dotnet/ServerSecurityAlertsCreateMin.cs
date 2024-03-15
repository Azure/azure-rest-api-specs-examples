using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;
using Azure.ResourceManager.MySql.Models;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerSecurityAlertsCreateMin.json
// this example is just showing the usage of "ServerSecurityAlertPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerResource created on azure
// for more information of creating MySqlServerResource, please refer to the document of MySqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securityalert-4799";
string serverName = "securityalert-6440";
ResourceIdentifier mySqlServerResourceId = MySqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlServerResource mySqlServer = client.GetMySqlServerResource(mySqlServerResourceId);

// get the collection of this MySqlServerSecurityAlertPolicyResource
MySqlServerSecurityAlertPolicyCollection collection = mySqlServer.GetMySqlServerSecurityAlertPolicies();

// invoke the operation
MySqlSecurityAlertPolicyName securityAlertPolicyName = MySqlSecurityAlertPolicyName.Default;
MySqlServerSecurityAlertPolicyData data = new MySqlServerSecurityAlertPolicyData()
{
    State = MySqlServerSecurityAlertPolicyState.Disabled,
    SendToEmailAccountAdmins = true,
};
ArmOperation<MySqlServerSecurityAlertPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, securityAlertPolicyName, data);
MySqlServerSecurityAlertPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlServerSecurityAlertPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
