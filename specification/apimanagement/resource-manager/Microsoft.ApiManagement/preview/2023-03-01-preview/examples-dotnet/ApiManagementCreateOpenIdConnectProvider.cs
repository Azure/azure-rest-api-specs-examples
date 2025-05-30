using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementCreateOpenIdConnectProvider.json
// this example is just showing the usage of "OpenIdConnectProvider_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ApiManagementOpenIdConnectProviderResource
ApiManagementOpenIdConnectProviderCollection collection = apiManagementService.GetApiManagementOpenIdConnectProviders();

// invoke the operation
string openId = "templateOpenIdConnect3";
ApiManagementOpenIdConnectProviderData data = new ApiManagementOpenIdConnectProviderData
{
    DisplayName = "templateoidprovider3",
    MetadataEndpoint = "https://oidprovider-template3.net",
    ClientId = "oidprovidertemplate3",
    ClientSecret = "x",
    UseInTestConsole = false,
    UseInApiDocumentation = true,
};
ArmOperation<ApiManagementOpenIdConnectProviderResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, openId, data);
ApiManagementOpenIdConnectProviderResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementOpenIdConnectProviderData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
