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

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_ListAppServices_MaximumSet_Gen.json
// this example is just showing the usage of "Monitors_ListAppServices" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NewRelicMonitorResource created on azure
// for more information of creating NewRelicMonitorResource, please refer to the document of NewRelicMonitorResource
string subscriptionId = "nqmcgifgaqlf";
string resourceGroupName = "rgNewRelic";
string monitorName = "fhcjxnxumkdlgpwanewtkdnyuz";
ResourceIdentifier newRelicMonitorResourceId = NewRelicMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
NewRelicMonitorResource newRelicMonitorResource = client.GetNewRelicMonitorResource(newRelicMonitorResourceId);

// invoke the operation and iterate over the result
NewRelicAppServicesGetContent content = new NewRelicAppServicesGetContent("ruxvg@xqkmdhrnoo.hlmbpm")
{
    AzureResourceIds =
    {
    new ResourceIdentifier("pvzrksrmzowobuhxpwiotnpcvjbu")
    },
};
await foreach (NewRelicObservabilityAppServiceInfo item in newRelicMonitorResource.GetAppServicesAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
