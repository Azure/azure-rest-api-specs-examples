using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.FrontDoor;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorGet.json
// this example is just showing the usage of "FrontDoors_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
FrontDoorResource result = await frontDoor.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
