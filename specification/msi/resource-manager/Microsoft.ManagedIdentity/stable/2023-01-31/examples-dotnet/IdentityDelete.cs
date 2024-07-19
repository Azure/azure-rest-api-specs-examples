using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedServiceIdentities.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedServiceIdentities;

// Generated from example definition: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/IdentityDelete.json
// this example is just showing the usage of "UserAssignedIdentities_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this UserAssignedIdentityResource created on azure
// for more information of creating UserAssignedIdentityResource, please refer to the document of UserAssignedIdentityResource
string subscriptionId = "subid";
string resourceGroupName = "rgName";
string resourceName = "resourceName";
ResourceIdentifier userAssignedIdentityResourceId = UserAssignedIdentityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
UserAssignedIdentityResource userAssignedIdentity = client.GetUserAssignedIdentityResource(userAssignedIdentityResourceId);

// invoke the operation
await userAssignedIdentity.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
