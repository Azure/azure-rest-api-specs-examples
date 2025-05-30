using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ManagedInstanceLongTermRetentionPolicyGet.json
// this example is just showing the usage of "ManagedInstanceLongTermRetentionPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceLongTermRetentionPolicyResource created on azure
// for more information of creating ManagedInstanceLongTermRetentionPolicyResource, please refer to the document of ManagedInstanceLongTermRetentionPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testResourceGroup";
string managedInstanceName = "testInstance";
string databaseName = "testDatabase";
ManagedInstanceLongTermRetentionPolicyName policyName = ManagedInstanceLongTermRetentionPolicyName.Default;
ResourceIdentifier managedInstanceLongTermRetentionPolicyResourceId = ManagedInstanceLongTermRetentionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, policyName);
ManagedInstanceLongTermRetentionPolicyResource managedInstanceLongTermRetentionPolicy = client.GetManagedInstanceLongTermRetentionPolicyResource(managedInstanceLongTermRetentionPolicyResourceId);

// invoke the operation
ManagedInstanceLongTermRetentionPolicyResource result = await managedInstanceLongTermRetentionPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceLongTermRetentionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
