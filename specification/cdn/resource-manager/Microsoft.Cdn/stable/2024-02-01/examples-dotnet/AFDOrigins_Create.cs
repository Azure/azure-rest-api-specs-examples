using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_Create.json
// this example is just showing the usage of "FrontDoorOrigins_Create" operation, for the dependent resources, they will have to be created separately.

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
FrontDoorOriginData data = new FrontDoorOriginData()
{
    HostName = "host1.blob.core.windows.net",
    HttpPort = 80,
    HttpsPort = 443,
    OriginHostHeader = "host1.foo.com",
    EnabledState = EnabledState.Enabled,
};
ArmOperation<FrontDoorOriginResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, originName, data);
FrontDoorOriginResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorOriginData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
