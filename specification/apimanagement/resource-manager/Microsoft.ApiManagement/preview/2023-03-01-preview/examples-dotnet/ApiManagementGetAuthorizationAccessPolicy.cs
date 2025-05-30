using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementGetAuthorizationAccessPolicy.json
// this example is just showing the usage of "AuthorizationAccessPolicy_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationContractResource created on azure
// for more information of creating AuthorizationContractResource, please refer to the document of AuthorizationContractResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authorizationProviderId = "aadwithauthcode";
string authorizationId = "authz1";
ResourceIdentifier authorizationContractResourceId = AuthorizationContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authorizationProviderId, authorizationId);
AuthorizationContractResource authorizationContract = client.GetAuthorizationContractResource(authorizationContractResourceId);

// get the collection of this AuthorizationAccessPolicyContractResource
AuthorizationAccessPolicyContractCollection collection = authorizationContract.GetAuthorizationAccessPolicyContracts();

// invoke the operation
string authorizationAccessPolicyId = "fe0bed83-631f-4149-bd0b-0464b1bc7cab";
NullableResponse<AuthorizationAccessPolicyContractResource> response = await collection.GetIfExistsAsync(authorizationAccessPolicyId);
AuthorizationAccessPolicyContractResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AuthorizationAccessPolicyContractData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
