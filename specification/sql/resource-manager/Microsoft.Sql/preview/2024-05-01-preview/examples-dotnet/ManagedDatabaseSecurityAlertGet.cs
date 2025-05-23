using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/ManagedDatabaseSecurityAlertGet.json
// this example is just showing the usage of "ManagedDatabaseSecurityAlertPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseSecurityAlertPolicyResource created on azure
// for more information of creating ManagedDatabaseSecurityAlertPolicyResource, please refer to the document of ManagedDatabaseSecurityAlertPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securityalert-6852";
string managedInstanceName = "securityalert-2080";
string databaseName = "testdb";
SqlSecurityAlertPolicyName securityAlertPolicyName = new SqlSecurityAlertPolicyName("default");
ResourceIdentifier managedDatabaseSecurityAlertPolicyResourceId = ManagedDatabaseSecurityAlertPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, securityAlertPolicyName);
ManagedDatabaseSecurityAlertPolicyResource managedDatabaseSecurityAlertPolicy = client.GetManagedDatabaseSecurityAlertPolicyResource(managedDatabaseSecurityAlertPolicyResourceId);

// invoke the operation
ManagedDatabaseSecurityAlertPolicyResource result = await managedDatabaseSecurityAlertPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedDatabaseSecurityAlertPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
