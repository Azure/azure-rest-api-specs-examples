using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Monitor;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionEndpointsDelete.json
// this example is just showing the usage of "DataCollectionEndpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataCollectionEndpointResource created on azure
// for more information of creating DataCollectionEndpointResource, please refer to the document of DataCollectionEndpointResource
string subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
string resourceGroupName = "myResourceGroup";
string dataCollectionEndpointName = "myCollectionEndpoint";
ResourceIdentifier dataCollectionEndpointResourceId = DataCollectionEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataCollectionEndpointName);
DataCollectionEndpointResource dataCollectionEndpoint = client.GetDataCollectionEndpointResource(dataCollectionEndpointResourceId);

// invoke the operation
await dataCollectionEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
