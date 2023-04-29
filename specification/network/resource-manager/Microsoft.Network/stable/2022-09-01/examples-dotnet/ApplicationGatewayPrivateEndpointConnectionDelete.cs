using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ApplicationGatewayPrivateEndpointConnectionDelete.json
// this example is just showing the usage of "ApplicationGatewayPrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationGatewayPrivateEndpointConnectionResource created on azure
// for more information of creating ApplicationGatewayPrivateEndpointConnectionResource, please refer to the document of ApplicationGatewayPrivateEndpointConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string applicationGatewayName = "appgw";
string connectionName = "connection1";
ResourceIdentifier applicationGatewayPrivateEndpointConnectionResourceId = ApplicationGatewayPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, applicationGatewayName, connectionName);
ApplicationGatewayPrivateEndpointConnectionResource applicationGatewayPrivateEndpointConnection = client.GetApplicationGatewayPrivateEndpointConnectionResource(applicationGatewayPrivateEndpointConnectionResourceId);

// invoke the operation
await applicationGatewayPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
