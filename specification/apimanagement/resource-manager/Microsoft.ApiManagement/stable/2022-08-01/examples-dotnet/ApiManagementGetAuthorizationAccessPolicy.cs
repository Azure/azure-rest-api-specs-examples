using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorizationAccessPolicy.json
// this example is just showing the usage of "AuthorizationAccessPolicy_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationAccessPolicyContractResource created on azure
// for more information of creating AuthorizationAccessPolicyContractResource, please refer to the document of AuthorizationAccessPolicyContractResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authorizationProviderId = "aadwithauthcode";
string authorizationId = "authz1";
string authorizationAccessPolicyId = "fe0bed83-631f-4149-bd0b-0464b1bc7cab";
ResourceIdentifier authorizationAccessPolicyContractResourceId = AuthorizationAccessPolicyContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authorizationProviderId, authorizationId, authorizationAccessPolicyId);
AuthorizationAccessPolicyContractResource authorizationAccessPolicyContract = client.GetAuthorizationAccessPolicyContractResource(authorizationAccessPolicyContractResourceId);

// invoke the operation
AuthorizationAccessPolicyContractResource result = await authorizationAccessPolicyContract.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationAccessPolicyContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
