using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
// this example is just showing the usage of "AuthorizationProvider_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementServiceResource created on azure
// for more information of creating ApiManagementServiceResource, please refer to the document of ApiManagementServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementServiceResourceId = ApiManagementServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementServiceResource apiManagementService = client.GetApiManagementServiceResource(apiManagementServiceResourceId);

// get the collection of this AuthorizationProviderContractResource
AuthorizationProviderContractCollection collection = apiManagementService.GetAuthorizationProviderContracts();

// invoke the operation
string authorizationProviderId = "eventbrite";
AuthorizationProviderContractData data = new AuthorizationProviderContractData
{
    DisplayName = "eventbrite",
    IdentityProvider = "oauth2",
    Oauth2 = new AuthorizationProviderOAuth2Settings
    {
        RedirectUri = new Uri("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
        GrantTypes = new AuthorizationProviderOAuth2GrantTypes
        {
            AuthorizationCode =
            {
            ["authorizationUrl"] = "https://www.eventbrite.com/oauth/authorize",
            ["clientId"] = "clientid",
            ["clientSecret"] = "clientsecretvalue",
            ["refreshUrl"] = "https://www.eventbrite.com/oauth/token",
            ["scopes"] = null,
            ["tokenUrl"] = "https://www.eventbrite.com/oauth/token"
            },
        },
    },
};
ArmOperation<AuthorizationProviderContractResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authorizationProviderId, data);
AuthorizationProviderContractResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationProviderContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
