using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateMaintenanceWindows.json
// this example is just showing the usage of "SqlPoolMaintenanceWindows_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseMaintenanceWindowResource created on azure
// for more information of creating SynapseMaintenanceWindowResource, please refer to the document of SynapseMaintenanceWindowResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "samplerg";
string workspaceName = "testworkspace";
string sqlPoolName = "testsp";
ResourceIdentifier synapseMaintenanceWindowResourceId = SynapseMaintenanceWindowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseMaintenanceWindowResource synapseMaintenanceWindow = client.GetSynapseMaintenanceWindowResource(synapseMaintenanceWindowResourceId);

// invoke the operation
string maintenanceWindowName = "current";
SynapseMaintenanceWindowData data = new SynapseMaintenanceWindowData
{
    TimeRanges = {new SynapseMaintenanceWindowTimeRange
    {
    DayOfWeek = SynapseDayOfWeek.Saturday,
    StartOn = XmlConvert.ToTimeSpan("00:00:00"),
    Duration = XmlConvert.ToTimeSpan("PT60M"),
    }},
};
await synapseMaintenanceWindow.CreateOrUpdateAsync(WaitUntil.Completed, maintenanceWindowName, data);

Console.WriteLine("Succeeded");
