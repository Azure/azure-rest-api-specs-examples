using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateGateway.json
// this example is just showing the usage of "Gateway_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementGatewayResource created on azure
// for more information of creating ApiManagementGatewayResource, please refer to the document of ApiManagementGatewayResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string gatewayId = "gw1";
ResourceIdentifier apiManagementGatewayResourceId = ApiManagementGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, gatewayId);
ApiManagementGatewayResource apiManagementGateway = client.GetApiManagementGatewayResource(apiManagementGatewayResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementGatewayData data = new ApiManagementGatewayData
{
    LocationData = new ResourceLocationDataContract("my location"),
    Description = "my gateway 1",
};
ApiManagementGatewayResource result = await apiManagementGateway.UpdateAsync(ifMatch, data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
