using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorDisableHttps.json
// this example is just showing the usage of "FrontendEndpoints_DisableHttps" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontendEndpointResource created on azure
// for more information of creating FrontendEndpointResource, please refer to the document of FrontendEndpointResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string frontDoorName = "frontDoor1";
string frontendEndpointName = "frontendEndpoint1";
ResourceIdentifier frontendEndpointResourceId = FrontendEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, frontDoorName, frontendEndpointName);
FrontendEndpointResource frontendEndpoint = client.GetFrontendEndpointResource(frontendEndpointResourceId);

// invoke the operation
await frontendEndpoint.DisableHttpsAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
