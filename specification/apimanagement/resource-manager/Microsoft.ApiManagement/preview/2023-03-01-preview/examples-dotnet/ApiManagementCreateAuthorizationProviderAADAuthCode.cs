using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementCreateAuthorizationProviderAADAuthCode.json
// this example is just showing the usage of "AuthorizationProvider_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AuthorizationProviderContractResource created on azure
// for more information of creating AuthorizationProviderContractResource, please refer to the document of AuthorizationProviderContractResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authorizationProviderId = "aadwithauthcode";
ResourceIdentifier authorizationProviderContractResourceId = AuthorizationProviderContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authorizationProviderId);
AuthorizationProviderContractResource authorizationProviderContract = client.GetAuthorizationProviderContractResource(authorizationProviderContractResourceId);

// invoke the operation
AuthorizationProviderContractData data = new AuthorizationProviderContractData()
{
    DisplayName = "aadwithauthcode",
    IdentityProvider = "aad",
    Oauth2 = new AuthorizationProviderOAuth2Settings()
    {
        RedirectUri = new Uri("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
        GrantTypes = new AuthorizationProviderOAuth2GrantTypes()
        {
            AuthorizationCode =
            {
            ["clientId"] = "59790825-fdd3-4b10-bc7a-4c3aaf25801d",
            ["clientSecret"] = "xxxxxxxxxxxxxxxxxxxxxxxx",
            ["resourceUri"] = "https://graph.microsoft.com",
            ["scopes"] = "User.Read.All Group.Read.All",
            },
        },
    },
};
ArmOperation<AuthorizationProviderContractResource> lro = await authorizationProviderContract.UpdateAsync(WaitUntil.Completed, data);
AuthorizationProviderContractResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AuthorizationProviderContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
