using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationAADClientCred.json
// this example is just showing the usage of "Authorization_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationProviderContractResource created on azure
// for more information of creating AuthorizationProviderContractResource, please refer to the document of AuthorizationProviderContractResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authorizationProviderId = "aadwithclientcred";
ResourceIdentifier authorizationProviderContractResourceId = AuthorizationProviderContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authorizationProviderId);
AuthorizationProviderContractResource authorizationProviderContract = client.GetAuthorizationProviderContractResource(authorizationProviderContractResourceId);

// get the collection of this AuthorizationContractResource
AuthorizationContractCollection collection = authorizationProviderContract.GetAuthorizationContracts();

// invoke the operation
string authorizationId = "authz1";
AuthorizationContractData data = new AuthorizationContractData()
{
    AuthorizationType = ApiManagementAuthorizationType.OAuth2,
    OAuth2GrantType = OAuth2GrantType.AuthorizationCode,
    Parameters =
    {
    ["clientId"] = "53790925-fdd3-4b80-bc7a-4c3aaf25801d",
    ["clientSecret"] = "xxxxxxxxxxxxxxx~xxxxxxxxx",
    },
};
ArmOperation<AuthorizationContractResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authorizationId, data);
AuthorizationContractResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
