using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DisconnectedOperations.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DisconnectedOperations;

// Generated from example definition: 2025-06-01-preview/DisconnectedOperations_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "DisconnectedOperation_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "51DB5DE7-A66C-4789-BFFF-9F75C95A0201";
string resourceGroupName = "rgdisconnectedOperations";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DisconnectedOperationResource
DisconnectedOperationCollection collection = resourceGroupResource.GetDisconnectedOperations();

// invoke the operation
string name = "demo-resource";
DisconnectedOperationData data = new DisconnectedOperationData(new AzureLocation("eastus"))
{
    Properties = new DisconnectedOperationProperties(null, DisconnectedOperationsBillingModel.Capacity, DisconnectedOperationsConnectionIntent.Disconnected),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<DisconnectedOperationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DisconnectedOperationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DisconnectedOperationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
