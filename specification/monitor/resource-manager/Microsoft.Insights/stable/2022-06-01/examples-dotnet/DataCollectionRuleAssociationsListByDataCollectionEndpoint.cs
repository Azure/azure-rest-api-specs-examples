using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionRuleAssociationsListByDataCollectionEndpoint.json
// this example is just showing the usage of "DataCollectionRuleAssociations_ListByDataCollectionEndpoint" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataCollectionEndpointResource created on azure
// for more information of creating DataCollectionEndpointResource, please refer to the document of DataCollectionEndpointResource
string subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
string resourceGroupName = "myResourceGroup";
string dataCollectionEndpointName = "myDataCollectionEndpointName";
ResourceIdentifier dataCollectionEndpointResourceId = DataCollectionEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataCollectionEndpointName);
DataCollectionEndpointResource dataCollectionEndpoint = client.GetDataCollectionEndpointResource(dataCollectionEndpointResourceId);

// invoke the operation and iterate over the result
await foreach (DataCollectionRuleAssociationResource item in dataCollectionEndpoint.GetDataCollectionRuleAssociationsAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataCollectionRuleAssociationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
