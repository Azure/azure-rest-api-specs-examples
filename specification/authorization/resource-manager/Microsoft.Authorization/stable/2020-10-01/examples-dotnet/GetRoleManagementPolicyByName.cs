using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleManagementPolicyByName.json
// this example is just showing the usage of "RoleManagementPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this RoleManagementPolicyResource
string scope = "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
RoleManagementPolicyCollection collection = client.GetRoleManagementPolicies(scopeId);

// invoke the operation
string roleManagementPolicyName = "570c3619-7688-4b34-b290-2b8bb3ccab2a";
bool result = await collection.ExistsAsync(roleManagementPolicyName);

Console.WriteLine($"Succeeded: {result}");
