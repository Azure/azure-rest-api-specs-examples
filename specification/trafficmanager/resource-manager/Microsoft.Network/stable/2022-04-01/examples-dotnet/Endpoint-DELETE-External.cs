using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.TrafficManager;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-DELETE-External.json
// this example is just showing the usage of "Endpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficManagerEndpointResource created on azure
// for more information of creating TrafficManagerEndpointResource, please refer to the document of TrafficManagerEndpointResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
string profileName = "azsmnet6386";
string endpointType = "ExternalEndpoints";
string endpointName = "azsmnet7187";
ResourceIdentifier trafficManagerEndpointResourceId = TrafficManagerEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointType, endpointName);
TrafficManagerEndpointResource trafficManagerEndpoint = client.GetTrafficManagerEndpointResource(trafficManagerEndpointResourceId);

// invoke the operation
await trafficManagerEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
