using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Workloads.Models;
using Azure.ResourceManager.Workloads;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_Update.json
// this example is just showing the usage of "SapLandscapeMonitor_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapLandscapeMonitorResource created on azure
// for more information of creating SapLandscapeMonitorResource, please refer to the document of SapLandscapeMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "mySapMonitor";
ResourceIdentifier sapLandscapeMonitorResourceId = SapLandscapeMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
SapLandscapeMonitorResource sapLandscapeMonitor = client.GetSapLandscapeMonitorResource(sapLandscapeMonitorResourceId);

// invoke the operation
SapLandscapeMonitorData data = new SapLandscapeMonitorData
{
    Grouping = new SapLandscapeMonitorPropertiesGrouping
    {
        Landscape = {new SapLandscapeMonitorSidMapping
        {
        Name = "Prod",
        TopSid = {"SID1", "SID2"},
        }},
        SapApplication = {new SapLandscapeMonitorSidMapping
        {
        Name = "ERP1",
        TopSid = {"SID1", "SID2"},
        }},
    },
    TopMetricsThresholds = {new SapLandscapeMonitorMetricThresholds
    {
    Name = "Instance Availability",
    Green = 90,
    Yellow = 75,
    Red = 50,
    }},
};
SapLandscapeMonitorResource result = await sapLandscapeMonitor.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapLandscapeMonitorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
