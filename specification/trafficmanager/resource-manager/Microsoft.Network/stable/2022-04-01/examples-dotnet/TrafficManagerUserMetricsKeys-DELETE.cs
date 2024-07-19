using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.TrafficManager;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/TrafficManagerUserMetricsKeys-DELETE.json
// this example is just showing the usage of "TrafficManagerUserMetricsKeys_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficManagerUserMetricsResource created on azure
// for more information of creating TrafficManagerUserMetricsResource, please refer to the document of TrafficManagerUserMetricsResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier trafficManagerUserMetricsResourceId = TrafficManagerUserMetricsResource.CreateResourceIdentifier(subscriptionId);
TrafficManagerUserMetricsResource trafficManagerUserMetrics = client.GetTrafficManagerUserMetricsResource(trafficManagerUserMetricsResourceId);

// invoke the operation
await trafficManagerUserMetrics.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
