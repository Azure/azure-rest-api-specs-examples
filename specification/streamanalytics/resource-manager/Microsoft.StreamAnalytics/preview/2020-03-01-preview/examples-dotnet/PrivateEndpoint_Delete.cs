using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/PrivateEndpoint_Delete.json
// this example is just showing the usage of "PrivateEndpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamAnalyticsPrivateEndpointResource created on azure
// for more information of creating StreamAnalyticsPrivateEndpointResource, please refer to the document of StreamAnalyticsPrivateEndpointResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "sjrg";
string clusterName = "testcluster";
string privateEndpointName = "testpe";
ResourceIdentifier streamAnalyticsPrivateEndpointResourceId = StreamAnalyticsPrivateEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, privateEndpointName);
StreamAnalyticsPrivateEndpointResource streamAnalyticsPrivateEndpoint = client.GetStreamAnalyticsPrivateEndpointResource(streamAnalyticsPrivateEndpointResourceId);

// invoke the operation
await streamAnalyticsPrivateEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
