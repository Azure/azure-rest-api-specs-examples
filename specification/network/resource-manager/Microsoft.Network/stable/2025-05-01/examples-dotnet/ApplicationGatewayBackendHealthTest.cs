using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ApplicationGatewayBackendHealthTest.json
// this example is just showing the usage of "ApplicationGateways_BackendHealthOnDemand" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationGatewayResource created on azure
// for more information of creating ApplicationGatewayResource, please refer to the document of ApplicationGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string applicationGatewayName = "appgw";
ResourceIdentifier applicationGatewayResourceId = ApplicationGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, applicationGatewayName);
ApplicationGatewayResource applicationGateway = client.GetApplicationGatewayResource(applicationGatewayResourceId);

// invoke the operation
ApplicationGatewayOnDemandProbe probeRequest = new ApplicationGatewayOnDemandProbe
{
    Protocol = ApplicationGatewayProtocol.Http,
    Path = "/",
    Timeout = 30,
    PickHostNameFromBackendHttpSettings = true,
    BackendAddressPoolId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendaddressPools/MFAnalyticsPool"),
    BackendHttpSettingsId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/MFPoolSettings"),
};
ArmOperation<ApplicationGatewayBackendHealthOnDemand> lro = await applicationGateway.BackendHealthOnDemandAsync(WaitUntil.Completed, probeRequest);
ApplicationGatewayBackendHealthOnDemand result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
