using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceMover;
using Azure.ResourceManager.ResourceMover.Models;

// Generated from example definition: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Delete.json
// this example is just showing the usage of "MoveResources_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MoverResource created on azure
// for more information of creating MoverResource, please refer to the document of MoverResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string moverResourceSetName = "movecollection1";
string moverResourceName = "moveresourcename1";
ResourceIdentifier moverResourceId = MoverResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, moverResourceSetName, moverResourceName);
MoverResource moverResource = client.GetMoverResource(moverResourceId);

// invoke the operation
ArmOperation<MoverOperationStatus> lro = await moverResource.DeleteAsync(WaitUntil.Completed);
MoverOperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
