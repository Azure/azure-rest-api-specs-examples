using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.TrafficManager;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Endpoint-GET-External-WithSubnetMapping.json
// this example is just showing the usage of "Endpoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficManagerEndpointResource created on azure
// for more information of creating TrafficManagerEndpointResource, please refer to the document of TrafficManagerEndpointResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager2191";
string profileName = "azuresdkfornetautoresttrafficmanager8224";
string endpointType = "ExternalEndpoints";
string endpointName = "My%20external%20endpoint";
ResourceIdentifier trafficManagerEndpointResourceId = TrafficManagerEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointType, endpointName);
TrafficManagerEndpointResource trafficManagerEndpoint = client.GetTrafficManagerEndpointResource(trafficManagerEndpointResourceId);

// invoke the operation
TrafficManagerEndpointResource result = await trafficManagerEndpoint.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrafficManagerEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
