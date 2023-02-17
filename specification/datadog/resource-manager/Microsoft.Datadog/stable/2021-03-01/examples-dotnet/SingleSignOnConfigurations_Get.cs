using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/SingleSignOnConfigurations_Get.json
// this example is just showing the usage of "SingleSignOnConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatadogMonitorResource created on azure
// for more information of creating DatadogMonitorResource, please refer to the document of DatadogMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
ResourceIdentifier datadogMonitorResourceId = DatadogMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
DatadogMonitorResource datadogMonitorResource = client.GetDatadogMonitorResource(datadogMonitorResourceId);

// get the collection of this DatadogSingleSignOnResource
DatadogSingleSignOnResourceCollection collection = datadogMonitorResource.GetDatadogSingleSignOnResources();

// invoke the operation
string configurationName = "default";
bool result = await collection.ExistsAsync(configurationName);

Console.WriteLine($"Succeeded: {result}");
