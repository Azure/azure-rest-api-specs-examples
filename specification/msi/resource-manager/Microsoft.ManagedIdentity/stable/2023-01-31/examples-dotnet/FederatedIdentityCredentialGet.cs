using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedServiceIdentities;

// Generated from example definition: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialGet.json
// this example is just showing the usage of "FederatedIdentityCredentials_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this UserAssignedIdentityResource created on azure
// for more information of creating UserAssignedIdentityResource, please refer to the document of UserAssignedIdentityResource
string subscriptionId = "c267c0e7-0a73-4789-9e17-d26aeb0904e5";
string resourceGroupName = "rgName";
string resourceName = "resourceName";
ResourceIdentifier userAssignedIdentityResourceId = UserAssignedIdentityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
UserAssignedIdentityResource userAssignedIdentity = client.GetUserAssignedIdentityResource(userAssignedIdentityResourceId);

// get the collection of this FederatedIdentityCredentialResource
FederatedIdentityCredentialCollection collection = userAssignedIdentity.GetFederatedIdentityCredentials();

// invoke the operation
string federatedIdentityCredentialResourceName = "ficResourceName";
NullableResponse<FederatedIdentityCredentialResource> response = await collection.GetIfExistsAsync(federatedIdentityCredentialResourceName);
FederatedIdentityCredentialResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    FederatedIdentityCredentialData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
