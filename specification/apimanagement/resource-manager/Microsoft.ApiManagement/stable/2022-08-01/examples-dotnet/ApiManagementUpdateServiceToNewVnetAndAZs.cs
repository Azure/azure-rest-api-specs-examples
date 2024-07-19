using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateServiceToNewVnetAndAZs.json
// this example is just showing the usage of "ApiManagementService_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
ApiManagementServicePatch patch = new ApiManagementServicePatch()
{
    Sku = new ApiManagementServiceSkuProperties(ApiManagementServiceSkuType.Premium, 3),
    Zones =
    {
    "1","2","3"
    },
    PublicIPAddressId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/publicip-apim-japan-east"),
    VirtualNetworkConfiguration = new VirtualNetworkConfiguration()
    {
        SubnetResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-apim-japaneast/subnets/apim2"),
    },
    AdditionalLocations =
    {
    new AdditionalLocation(new AzureLocation("Australia East"),new ApiManagementServiceSkuProperties(ApiManagementServiceSkuType.Premium,3))
    {
    Zones =
    {
    "1","2","3"
    },
    PublicIPAddressId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/apim-australia-east-publicip"),
    VirtualNetworkConfiguration = new VirtualNetworkConfiguration()
    {
    SubnetResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apimaeavnet/subnets/default"),
    },
    }
    },
    VirtualNetworkType = VirtualNetworkType.External,
};
ArmOperation<ApiManagementServiceResource> lro = await apiManagementService.UpdateAsync(WaitUntil.Completed, patch);
ApiManagementServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
