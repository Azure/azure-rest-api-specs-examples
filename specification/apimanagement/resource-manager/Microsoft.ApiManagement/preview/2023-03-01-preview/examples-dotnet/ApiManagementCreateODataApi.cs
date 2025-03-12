using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementCreateODataApi.json
// this example is just showing the usage of "Api_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ApiResource
ApiCollection collection = apiManagementService.GetApis();

// invoke the operation
string apiId = "tempgroup";
ApiCreateOrUpdateContent content = new ApiCreateOrUpdateContent
{
    Description = "apidescription5200",
    ApiType = ApiType.Odata,
    DisplayName = "apiname1463",
    ServiceLink = "https://services.odata.org/TripPinWebApiService",
    Path = "odata-api",
    Protocols = { ApiOperationInvokableProtocol.Http, ApiOperationInvokableProtocol.Https },
    Value = "https://services.odata.org/TripPinWebApiService/$metadata",
    Format = ContentFormat.OdataLink,
};
ArmOperation<ApiResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, apiId, content);
ApiResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
