using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationServer.json
// this example is just showing the usage of "AuthorizationServer_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementServiceResource created on azure
// for more information of creating ApiManagementServiceResource, please refer to the document of ApiManagementServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementServiceResourceId = ApiManagementServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementServiceResource apiManagementService = client.GetApiManagementServiceResource(apiManagementServiceResourceId);

// get the collection of this ApiManagementAuthorizationServerResource
ApiManagementAuthorizationServerCollection collection = apiManagementService.GetApiManagementAuthorizationServers();

// invoke the operation
string authsid = "newauthServer";
ApiManagementAuthorizationServerData data = new ApiManagementAuthorizationServerData()
{
    Description = "test server",
    AuthorizationMethods =
    {
    AuthorizationMethod.Get
    },
    TokenEndpoint = "https://www.contoso.com/oauth2/token",
    DoesSupportState = true,
    DefaultScope = "read write",
    BearerTokenSendingMethods =
    {
    BearerTokenSendingMethod.AuthorizationHeader
    },
    ResourceOwnerUsername = "un",
    ResourceOwnerPassword = "pwd",
    DisplayName = "test2",
    UseInTestConsole = false,
    UseInApiDocumentation = true,
    ClientRegistrationEndpoint = "https://www.contoso.com/apps",
    AuthorizationEndpoint = "https://www.contoso.com/oauth2/auth",
    GrantTypes =
    {
    GrantType.AuthorizationCode,GrantType.Implicit
    },
    ClientId = "1",
    ClientSecret = "2",
};
ArmOperation<ApiManagementAuthorizationServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authsid, data);
ApiManagementAuthorizationServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementAuthorizationServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
