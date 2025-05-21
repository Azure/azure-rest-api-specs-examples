using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2025-04-15/examples/AFDEndpoints_PurgeContent.json
// this example is just showing the usage of "FrontDoorEndpoints_PurgeContent" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
FrontDoorPurgeContent content = new FrontDoorPurgeContent(new string[] { "/folder1" })
{
    Domains = { "endpoint1-abcdefghijklmnop.z01.azurefd.net" },
};
await frontDoorEndpoint.PurgeContentAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
