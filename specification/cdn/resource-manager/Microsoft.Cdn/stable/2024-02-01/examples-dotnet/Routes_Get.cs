using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Routes_Get.json
// this example is just showing the usage of "FrontDoorRoutes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorRouteResource created on azure
// for more information of creating FrontDoorRouteResource, please refer to the document of FrontDoorRouteResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
string routeName = "route1";
ResourceIdentifier frontDoorRouteResourceId = FrontDoorRouteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName, routeName);
FrontDoorRouteResource frontDoorRoute = client.GetFrontDoorRouteResource(frontDoorRouteResourceId);

// invoke the operation
FrontDoorRouteResource result = await frontDoorRoute.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorRouteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
