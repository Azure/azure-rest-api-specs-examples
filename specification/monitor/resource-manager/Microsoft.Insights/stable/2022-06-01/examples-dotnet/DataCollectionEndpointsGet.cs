using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionEndpointsGet.json
// this example is just showing the usage of "DataCollectionEndpoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataCollectionEndpointResource
DataCollectionEndpointCollection collection = resourceGroupResource.GetDataCollectionEndpoints();

// invoke the operation
string dataCollectionEndpointName = "myCollectionEndpoint";
NullableResponse<DataCollectionEndpointResource> response = await collection.GetIfExistsAsync(dataCollectionEndpointName);
DataCollectionEndpointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataCollectionEndpointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
