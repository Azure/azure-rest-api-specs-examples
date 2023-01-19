using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetRoleDefinitionByName.json
// this example is just showing the usage of "RoleDefinitions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this AuthorizationRoleDefinitionResource
string scope = "scope";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
AuthorizationRoleDefinitionCollection collection = client.GetAuthorizationRoleDefinitions(scopeId);

// invoke the operation
ResourceIdentifier roleDefinitionId = new ResourceIdentifier("roleDefinitionId");
bool result = await collection.ExistsAsync(roleDefinitionId);

Console.WriteLine($"Succeeded: {result}");
