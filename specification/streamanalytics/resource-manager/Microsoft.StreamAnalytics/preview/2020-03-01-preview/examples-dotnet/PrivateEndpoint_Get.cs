using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/PrivateEndpoint_Get.json
// this example is just showing the usage of "PrivateEndpoints_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this StreamAnalyticsPrivateEndpointResource
StreamAnalyticsPrivateEndpointCollection collection = streamAnalyticsCluster.GetStreamAnalyticsPrivateEndpoints();

// invoke the operation
string privateEndpointName = "testpe";
NullableResponse<StreamAnalyticsPrivateEndpointResource> response = await collection.GetIfExistsAsync(privateEndpointName);
StreamAnalyticsPrivateEndpointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StreamAnalyticsPrivateEndpointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
