using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StreamAnalytics.Models;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/Cluster_ListStreamingJobs.json
// this example is just showing the usage of "Clusters_ListStreamingJobs" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamAnalyticsClusterResource created on azure
// for more information of creating StreamAnalyticsClusterResource, please refer to the document of StreamAnalyticsClusterResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "sjrg";
string clusterName = "testcluster";
ResourceIdentifier streamAnalyticsClusterResourceId = StreamAnalyticsClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
StreamAnalyticsClusterResource streamAnalyticsCluster = client.GetStreamAnalyticsClusterResource(streamAnalyticsClusterResourceId);

// invoke the operation and iterate over the result
await foreach (StreamAnalyticsClusterJob item in streamAnalyticsCluster.GetStreamingJobsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
