using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationAADClientCred.json
// this example is just showing the usage of "Authorization_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationContractResource created on azure
// for more information of creating AuthorizationContractResource, please refer to the document of AuthorizationContractResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authorizationProviderId = "aadwithclientcred";
string authorizationId = "authz1";
ResourceIdentifier authorizationContractResourceId = AuthorizationContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authorizationProviderId, authorizationId);
AuthorizationContractResource authorizationContract = client.GetAuthorizationContractResource(authorizationContractResourceId);

// invoke the operation
AuthorizationContractData data = new AuthorizationContractData
{
    AuthorizationType = ApiManagementAuthorizationType.OAuth2,
    OAuth2GrantType = OAuth2GrantType.AuthorizationCode,
    Parameters =
    {
    ["clientId"] = "clientsecretid",
    ["clientSecret"] = "clientsecretvalue"
    },
};
ArmOperation<AuthorizationContractResource> lro = await authorizationContract.UpdateAsync(WaitUntil.Completed, data);
AuthorizationContractResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
