using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;
using Azure.ResourceManager.ApiManagement.Models;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadOpenIdConnectProvider.json
// this example is just showing the usage of "OpenIdConnectProvider_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementOpenIdConnectProviderResource created on azure
// for more information of creating ApiManagementOpenIdConnectProviderResource, please refer to the document of ApiManagementOpenIdConnectProviderResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string openId = "templateOpenIdConnect2";
ResourceIdentifier apiManagementOpenIdConnectProviderResourceId = ApiManagementOpenIdConnectProviderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, openId);
ApiManagementOpenIdConnectProviderResource apiManagementOpenIdConnectProvider = client.GetApiManagementOpenIdConnectProviderResource(apiManagementOpenIdConnectProviderResourceId);

// invoke the operation
bool result = await apiManagementOpenIdConnectProvider.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
