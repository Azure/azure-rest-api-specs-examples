using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Datadog;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Get.json
// this example is just showing the usage of "Monitors_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DatadogMonitorResource
DatadogMonitorResourceCollection collection = resourceGroupResource.GetDatadogMonitorResources();

// invoke the operation
string monitorName = "myMonitor";
bool result = await collection.ExistsAsync(monitorName);

Console.WriteLine($"Succeeded: {result}");
