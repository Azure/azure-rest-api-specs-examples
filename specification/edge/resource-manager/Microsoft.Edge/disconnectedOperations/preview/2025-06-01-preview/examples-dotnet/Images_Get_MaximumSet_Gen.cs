using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DisconnectedOperations;

// Generated from example definition: 2025-06-01-preview/Images_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Image_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DisconnectedOperationResource created on azure
// for more information of creating DisconnectedOperationResource, please refer to the document of DisconnectedOperationResource
string subscriptionId = "301DCB09-82EC-4777-A56C-6FFF26BCC814";
string resourceGroupName = "rgdisconnectedoperations";
string name = "bT62l-KS7g1-uh";
ResourceIdentifier disconnectedOperationResourceId = DisconnectedOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DisconnectedOperationResource disconnectedOperation = client.GetDisconnectedOperationResource(disconnectedOperationResourceId);

// get the collection of this DisconnectedOperationsImageResource
DisconnectedOperationsImageCollection collection = disconnectedOperation.GetDisconnectedOperationsImages();

// invoke the operation
string imageName = "2P6";
NullableResponse<DisconnectedOperationsImageResource> response = await collection.GetIfExistsAsync(imageName);
DisconnectedOperationsImageResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DisconnectedOperationsImageData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
