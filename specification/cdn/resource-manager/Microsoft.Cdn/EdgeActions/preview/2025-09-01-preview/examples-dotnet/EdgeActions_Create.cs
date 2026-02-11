using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EdgeActions.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.EdgeActions;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/EdgeActions/preview/2025-09-01-preview/examples/EdgeActions_Create.json
// this example is just showing the usage of "EdgeActions_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EdgeActionResource
EdgeActionCollection collection = resourceGroupResource.GetEdgeActions();

// invoke the operation
string edgeActionName = "edgeAction1";
EdgeActionData data = new EdgeActionData(new AzureLocation("global"), new EdgeActionSkuType("Standard", "Standard"));
ArmOperation<EdgeActionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, edgeActionName, data);
EdgeActionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeActionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
