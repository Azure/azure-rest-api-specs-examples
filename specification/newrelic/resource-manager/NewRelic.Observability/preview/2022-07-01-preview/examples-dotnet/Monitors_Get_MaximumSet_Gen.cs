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

// this example assumes you already have this NewRelicMonitorResource created on azure
// for more information of creating NewRelicMonitorResource, please refer to the document of NewRelicMonitorResource
string subscriptionId = "hfmjmpyqgezxkp";
string resourceGroupName = "rgNewRelic";
string monitorName = "cdlymktqw";
ResourceIdentifier newRelicMonitorResourceId = NewRelicMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
NewRelicMonitorResource newRelicMonitorResource = client.GetNewRelicMonitorResource(newRelicMonitorResourceId);

// invoke the operation
NewRelicMonitorResource result = await newRelicMonitorResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NewRelicMonitorResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
