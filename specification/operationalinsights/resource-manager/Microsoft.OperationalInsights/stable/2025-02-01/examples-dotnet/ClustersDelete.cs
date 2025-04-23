using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2025-02-01/examples/ClustersDelete.json
// this example is just showing the usage of "Clusters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsClusterResource created on azure
// for more information of creating OperationalInsightsClusterResource, please refer to the document of OperationalInsightsClusterResource
string subscriptionId = "53bc36c5-91e1-4d09-92c9-63b89e571926";
string resourceGroupName = "oiautorest6685";
string clusterName = "oiautorest6685";
ResourceIdentifier operationalInsightsClusterResourceId = OperationalInsightsClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
OperationalInsightsClusterResource operationalInsightsCluster = client.GetOperationalInsightsClusterResource(operationalInsightsClusterResourceId);

// invoke the operation
await operationalInsightsCluster.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
