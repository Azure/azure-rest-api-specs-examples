using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedServiceIdentities;

// Generated from example definition: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/SystemAssignedIdentityGet.json
// this example is just showing the usage of "SystemAssignedIdentities_GetByScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SystemAssignedIdentityResource created on azure
// for more information of creating SystemAssignedIdentityResource, please refer to the document of SystemAssignedIdentityResource
string scope = "subscriptions/subId/resourceGroups/resourceGroupName/providers/Resource.Provider/resourceType/resourceName/identities/default";
ResourceIdentifier systemAssignedIdentityResourceId = SystemAssignedIdentityResource.CreateResourceIdentifier(scope);
SystemAssignedIdentityResource systemAssignedIdentity = client.GetSystemAssignedIdentityResource(systemAssignedIdentityResourceId);

// invoke the operation
SystemAssignedIdentityResource result = await systemAssignedIdentity.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SystemAssignedIdentityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
