using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NewRelicObservability;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Monitors_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Monitors_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "hfmjmpyqgezxkp";
string resourceGroupName = "rgNewRelic";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NewRelicMonitorResource
NewRelicMonitorResourceCollection collection = resourceGroupResource.GetNewRelicMonitorResources();

// invoke the operation
string monitorName = "cdlymktqw";
bool result = await collection.ExistsAsync(monitorName);

Console.WriteLine($"Succeeded: {result}");
