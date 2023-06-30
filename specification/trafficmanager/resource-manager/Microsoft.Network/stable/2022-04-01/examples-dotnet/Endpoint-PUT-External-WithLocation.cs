using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.TrafficManager;
using Azure.ResourceManager.TrafficManager.Models;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-WithLocation.json
// this example is just showing the usage of "Endpoints_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficManagerProfileResource created on azure
// for more information of creating TrafficManagerProfileResource, please refer to the document of TrafficManagerProfileResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager1421";
string profileName = "azsmnet6386";
ResourceIdentifier trafficManagerProfileResourceId = TrafficManagerProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
TrafficManagerProfileResource trafficManagerProfile = client.GetTrafficManagerProfileResource(trafficManagerProfileResourceId);

// get the collection of this TrafficManagerEndpointResource
TrafficManagerEndpointCollection collection = trafficManagerProfile.GetTrafficManagerEndpoints();

// invoke the operation
string endpointType = "ExternalEndpoints";
string endpointName = "azsmnet7187";
TrafficManagerEndpointData data = new TrafficManagerEndpointData()
{
    Target = "foobar.contoso.com",
    EndpointStatus = TrafficManagerEndpointStatus.Enabled,
    EndpointLocation = "North Europe",
    Name = "azsmnet7187",
    ResourceType = new ResourceType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
};
ArmOperation<TrafficManagerEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, endpointType, endpointName, data);
TrafficManagerEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrafficManagerEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
