using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.NewRelicObservability;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-03-01/examples/Monitors_SwitchBilling_MaximumSet_Gen.json
// this example is just showing the usage of "Monitors_SwitchBilling" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NewRelicMonitorResource created on azure
// for more information of creating NewRelicMonitorResource, please refer to the document of NewRelicMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgNewRelic";
string monitorName = "fhcjxnxumkdlgpwanewtkdnyuz";
ResourceIdentifier newRelicMonitorResourceId = NewRelicMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
NewRelicMonitorResource newRelicMonitorResource = client.GetNewRelicMonitorResource(newRelicMonitorResourceId);

// invoke the operation
NewRelicSwitchBillingContent content = new NewRelicSwitchBillingContent("ruxvg@xqkmdhrnoo.hlmbpm")
{
    AzureResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgNewRelic/providers/NewRelic.Observability/monitors/fhcjxnxumkdlgpwanewtkdnyuz"),
    OrganizationId = "k",
    PlanData = new NewRelicPlanDetails
    {
        UsageType = NewRelicObservabilityUsageType.Payg,
        NewRelicPlanBillingCycle = "Yearly",
        PlanDetails = "tbbiaga",
        EffectiveOn = DateTimeOffset.Parse("2022-12-05T14:11:37.786Z"),
    },
};
NewRelicMonitorResource result = await newRelicMonitorResource.SwitchBillingAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NewRelicMonitorResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
