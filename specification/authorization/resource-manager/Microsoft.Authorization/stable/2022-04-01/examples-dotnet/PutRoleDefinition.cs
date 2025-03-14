using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/PutRoleDefinition.json
// this example is just showing the usage of "RoleDefinitions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this AuthorizationRoleDefinitionResource
string scope = "scope";
AuthorizationRoleDefinitionCollection collection = client.GetAuthorizationRoleDefinitions(new ResourceIdentifier(scope));

// invoke the operation
ResourceIdentifier roleDefinitionId = new ResourceIdentifier("roleDefinitionId");
AuthorizationRoleDefinitionData data = new AuthorizationRoleDefinitionData();
ArmOperation<AuthorizationRoleDefinitionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, roleDefinitionId, data);
AuthorizationRoleDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationRoleDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
