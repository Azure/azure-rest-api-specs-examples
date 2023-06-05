using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorFrontendEndpointGet.json
// this example is just showing the usage of "FrontendEndpoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorResource created on azure
// for more information of creating FrontDoorResource, please refer to the document of FrontDoorResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string frontDoorName = "frontDoor1";
ResourceIdentifier frontDoorResourceId = FrontDoorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, frontDoorName);
FrontDoorResource frontDoor = client.GetFrontDoorResource(frontDoorResourceId);

// get the collection of this FrontendEndpointResource
FrontendEndpointCollection collection = frontDoor.GetFrontendEndpoints();

// invoke the operation
string frontendEndpointName = "frontendEndpoint1";
bool result = await collection.ExistsAsync(frontendEndpointName);

Console.WriteLine($"Succeeded: {result}");
