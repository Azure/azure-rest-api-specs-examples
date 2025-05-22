using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateGatewayHostnameConfiguration.json
// this example is just showing the usage of "GatewayHostnameConfiguration_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ApiManagementGatewayHostnameConfigurationResource
ApiManagementGatewayHostnameConfigurationCollection collection = apiManagementGateway.GetApiManagementGatewayHostnameConfigurations();

// invoke the operation
string hcId = "default";
ApiManagementGatewayHostnameConfigurationData data = new ApiManagementGatewayHostnameConfigurationData
{
    Hostname = "*",
    CertificateId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1",
    IsClientCertificateRequired = false,
    IsTls1_0Enabled = false,
    IsTls1_1Enabled = false,
    IsHttp2_0Enabled = true,
};
ArmOperation<ApiManagementGatewayHostnameConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, hcId, data);
ApiManagementGatewayHostnameConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementGatewayHostnameConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
