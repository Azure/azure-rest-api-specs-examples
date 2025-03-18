using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceMover.Models;
using Azure.ResourceManager.ResourceMover;

// Generated from example definition: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Get.json
// this example is just showing the usage of "MoveResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MoverResourceSetResource created on azure
// for more information of creating MoverResourceSetResource, please refer to the document of MoverResourceSetResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string moverResourceSetName = "movecollection1";
ResourceIdentifier moverResourceSetResourceId = MoverResourceSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, moverResourceSetName);
MoverResourceSetResource moverResourceSet = client.GetMoverResourceSetResource(moverResourceSetResourceId);

// get the collection of this MoverResource
MoverResourceCollection collection = moverResourceSet.GetMoverResources();

// invoke the operation
string moverResourceName = "moveresourcename1";
NullableResponse<MoverResource> response = await collection.GetIfExistsAsync(moverResourceName);
MoverResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MoverResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
