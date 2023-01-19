using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ResourceMover;
using Azure.ResourceManager.ResourceMover.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveCollections_Create.json
// this example is just showing the usage of "MoveCollections_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MoverResourceSetResource
MoverResourceSetCollection collection = resourceGroupResource.GetMoverResourceSets();

// invoke the operation
string moverResourceSetName = "movecollection1";
MoverResourceSetData data = new MoverResourceSetData(new AzureLocation("eastus2"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Properties = new MoverResourceSetProperties(new AzureLocation("eastus"), new AzureLocation("westus")),
};
ArmOperation<MoverResourceSetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, moverResourceSetName, data);
MoverResourceSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MoverResourceSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
