using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/PutRoleDefinition.json
// this example is just showing the usage of "RoleDefinitions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationRoleDefinitionResource created on azure
// for more information of creating AuthorizationRoleDefinitionResource, please refer to the document of AuthorizationRoleDefinitionResource
string scope = "scope";
ResourceIdentifier roleDefinitionId = new ResourceIdentifier("roleDefinitionId");
ResourceIdentifier authorizationRoleDefinitionResourceId = AuthorizationRoleDefinitionResource.CreateResourceIdentifier(scope, roleDefinitionId);
AuthorizationRoleDefinitionResource authorizationRoleDefinition = client.GetAuthorizationRoleDefinitionResource(authorizationRoleDefinitionResourceId);

// invoke the operation
AuthorizationRoleDefinitionData data = new AuthorizationRoleDefinitionData();
ArmOperation<AuthorizationRoleDefinitionResource> lro = await authorizationRoleDefinition.UpdateAsync(WaitUntil.Completed, data);
AuthorizationRoleDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationRoleDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
