using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGatewayHostnameConfiguration.json
// this example is just showing the usage of "GatewayHostnameConfiguration_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementGatewayHostnameConfigurationResource created on azure
// for more information of creating ApiManagementGatewayHostnameConfigurationResource, please refer to the document of ApiManagementGatewayHostnameConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string gatewayId = "gw1";
string hcId = "default";
ResourceIdentifier apiManagementGatewayHostnameConfigurationResourceId = ApiManagementGatewayHostnameConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, gatewayId, hcId);
ApiManagementGatewayHostnameConfigurationResource apiManagementGatewayHostnameConfiguration = client.GetApiManagementGatewayHostnameConfigurationResource(apiManagementGatewayHostnameConfigurationResourceId);

// invoke the operation
bool result = await apiManagementGatewayHostnameConfiguration.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
