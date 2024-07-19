using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/deleteAutoscaleSetting.json
// this example is just showing the usage of "AutoscaleSettings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutoscaleSettingResource created on azure
// for more information of creating AutoscaleSettingResource, please refer to the document of AutoscaleSettingResource
string subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
string resourceGroupName = "TestingMetricsScaleSet";
string autoscaleSettingName = "MySetting";
ResourceIdentifier autoscaleSettingResourceId = AutoscaleSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, autoscaleSettingName);
AutoscaleSettingResource autoscaleSetting = client.GetAutoscaleSettingResource(autoscaleSettingResourceId);

// invoke the operation
await autoscaleSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
