using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Routes_Get.json
// this example is just showing the usage of "FrontDoorRoutes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorEndpointResource created on azure
// for more information of creating FrontDoorEndpointResource, please refer to the document of FrontDoorEndpointResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
ResourceIdentifier frontDoorEndpointResourceId = FrontDoorEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName);
FrontDoorEndpointResource frontDoorEndpoint = client.GetFrontDoorEndpointResource(frontDoorEndpointResourceId);

// get the collection of this FrontDoorRouteResource
FrontDoorRouteCollection collection = frontDoorEndpoint.GetFrontDoorRoutes();

// invoke the operation
string routeName = "route1";
bool result = await collection.ExistsAsync(routeName);

Console.WriteLine($"Succeeded: {result}");
