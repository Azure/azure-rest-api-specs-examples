using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2025-02-01/examples/ClustersCreate.json
// this example is just showing the usage of "Clusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "53bc36c5-91e1-4d09-92c9-63b89e571926";
string resourceGroupName = "oiautorest6685";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this OperationalInsightsClusterResource
OperationalInsightsClusterCollection collection = resourceGroupResource.GetOperationalInsightsClusters();

// invoke the operation
string clusterName = "oiautorest6685";
OperationalInsightsClusterData data = new OperationalInsightsClusterData(new AzureLocation("eastus"))
{
    Sku = new OperationalInsightsClusterSku
    {
        Capacity = OperationalInsightsClusterCapacity.TenHundred,
        Name = OperationalInsightsClusterSkuName.CapacityReservation,
    },
    Tags =
    {
    ["tag1"] = "val1"
    },
};
ArmOperation<OperationalInsightsClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
OperationalInsightsClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OperationalInsightsClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
