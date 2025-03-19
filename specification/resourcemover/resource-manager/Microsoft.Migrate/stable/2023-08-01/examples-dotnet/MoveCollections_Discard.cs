using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ResourceMover.Models;
using Azure.ResourceManager.ResourceMover;

// Generated from example definition: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Discard.json
// this example is just showing the usage of "MoveCollections_Discard" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
MoverDiscardContent content = new MoverDiscardContent(new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1") })
{
    ValidateOnly = false,
};
ArmOperation<MoverOperationStatus> lro = await moverResourceSet.DiscardAsync(WaitUntil.Completed, content: content);
MoverOperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
