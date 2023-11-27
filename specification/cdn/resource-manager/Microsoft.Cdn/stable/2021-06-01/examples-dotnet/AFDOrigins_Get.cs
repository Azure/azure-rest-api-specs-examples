using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_Get.json
// this example is just showing the usage of "FrontDoorOrigins_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorOriginGroupResource created on azure
// for more information of creating FrontDoorOriginGroupResource, please refer to the document of FrontDoorOriginGroupResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string originGroupName = "origingroup1";
ResourceIdentifier frontDoorOriginGroupResourceId = FrontDoorOriginGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, originGroupName);
FrontDoorOriginGroupResource frontDoorOriginGroup = client.GetFrontDoorOriginGroupResource(frontDoorOriginGroupResourceId);

// get the collection of this FrontDoorOriginResource
FrontDoorOriginCollection collection = frontDoorOriginGroup.GetFrontDoorOrigins();

// invoke the operation
string originName = "origin1";
NullableResponse<FrontDoorOriginResource> response = await collection.GetIfExistsAsync(originName);
FrontDoorOriginResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    FrontDoorOriginData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
