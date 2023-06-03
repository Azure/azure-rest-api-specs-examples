using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.OperationalInsights;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersGet.json
// this example is just showing the usage of "Clusters_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsClusterResource created on azure
// for more information of creating OperationalInsightsClusterResource, please refer to the document of OperationalInsightsClusterResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "oiautorest6685";
string clusterName = "oiautorest6685";
ResourceIdentifier operationalInsightsClusterResourceId = OperationalInsightsClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
OperationalInsightsClusterResource operationalInsightsCluster = client.GetOperationalInsightsClusterResource(operationalInsightsClusterResourceId);

// invoke the operation
OperationalInsightsClusterResource result = await operationalInsightsCluster.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OperationalInsightsClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
