using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/ClusterMetricsConfigurations_Delete.json
// this example is just showing the usage of "MetricsConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ClusterMetricsConfigurationResource created on azure
// for more information of creating ClusterMetricsConfigurationResource, please refer to the document of ClusterMetricsConfigurationResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
string metricsConfigurationName = "default";
ResourceIdentifier clusterMetricsConfigurationResourceId = ClusterMetricsConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, metricsConfigurationName);
ClusterMetricsConfigurationResource clusterMetricsConfiguration = client.GetClusterMetricsConfigurationResource(clusterMetricsConfigurationResourceId);

// invoke the operation
await clusterMetricsConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
