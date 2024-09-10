using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementCreateServiceWithoutLegacyConfigurationApi.json
// this example is just showing the usage of "ApiManagementService_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ApiManagementServiceResource
ApiManagementServiceCollection collection = resourceGroupResource.GetApiManagementServices();

// invoke the operation
string serviceName = "apimService1";
ApiManagementServiceData data = new ApiManagementServiceData(new AzureLocation("Central US"), new ApiManagementServiceSkuProperties(ApiManagementServiceSkuType.Basic, 1), "apim@autorestsdk.com", "autorestsdk")
{
    LegacyApi = LegacyApiState.Disabled,
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    ["tag3"] = "value3",
    },
};
ArmOperation<ApiManagementServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceName, data);
ApiManagementServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
