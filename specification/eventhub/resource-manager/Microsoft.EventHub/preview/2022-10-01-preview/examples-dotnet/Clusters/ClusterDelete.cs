using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/Clusters/ClusterDelete.json
// this example is just showing the usage of "Clusters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsClusterResource created on azure
// for more information of creating EventHubsClusterResource, please refer to the document of EventHubsClusterResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "myResourceGroup";
string clusterName = "testCluster";
ResourceIdentifier eventHubsClusterResourceId = EventHubsClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
EventHubsClusterResource eventHubsCluster = client.GetEventHubsClusterResource(eventHubsClusterResourceId);

// invoke the operation
await eventHubsCluster.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
