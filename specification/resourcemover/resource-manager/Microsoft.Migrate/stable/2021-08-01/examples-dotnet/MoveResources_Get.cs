using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceMover;
using Azure.ResourceManager.ResourceMover.Models;

// Generated from example definition: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveResources_Get.json
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
bool result = await collection.ExistsAsync(moverResourceName);

Console.WriteLine($"Succeeded: {result}");
