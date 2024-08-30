using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Monitor/preview/2023-10-01-preview/examples/PipelineGroupCreateSyslogsWithNetworking.json
// this example is just showing the usage of "PipelineGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this PipelineGroupResource
PipelineGroupCollection collection = resourceGroupResource.GetPipelineGroups();

// invoke the operation
string pipelineGroupName = "plGroup1";
PipelineGroupData data = new PipelineGroupData(new AzureLocation("eastus2"))
{
    ExtendedLocation = new ExtendedLocation()
    {
        Name = "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/myTestCustomLocation",
    },
    Receivers =
    {
    new PipelineGroupReceiver(PipelineGroupReceiverType.Syslog,"syslog-receiver1")
    {
    Syslog = new SyslogReceiver("0.0.0.0:514"),
    }
    },
    Processors =
    {
    },
    Exporters =
    {
    new PipelineGroupExporter(PipelineGroupExporterType.AzureMonitorWorkspaceLogs,"my-workspace-logs-exporter1")
    {
    AzureMonitorWorkspaceLogs = new MonitorWorkspaceLogsExporter(new MonitorWorkspaceLogsApiConfig(new Uri("https://logs-myingestion-eb0s.eastus-1.ingest.monitor.azure.com"),"Custom-MyTableRawData_CL","dcr-00000000000000000000000000000000",new MonitorWorkspaceLogsSchemaMap(new MonitorWorkspaceLogsRecordMap[]
    {
    new MonitorWorkspaceLogsRecordMap("body","Body"),new MonitorWorkspaceLogsRecordMap("severity_text","SeverityText"),new MonitorWorkspaceLogsRecordMap("time_unix_nano","TimeGenerated")
    })))
    {
    Concurrency = new MonitorWorkspaceLogsExporterConcurrencyConfiguration()
    {
    WorkerCount = 4,
    BatchQueueSize = 100,
    },
    },
    }
    },
    Service = new PipelineGroupService(new PipelineGroupServicePipeline[]
{
new PipelineGroupServicePipeline("MyPipelineForLogs1",PipelineGroupServicePipelineType.Logs,new string[]
{
"syslog-receiver1"
},new string[]
{
"my-workspace-logs-exporter1"
})
{
Processors =
{
},
}
}),
    NetworkingConfigurations =
    {
    new PipelineGroupNetworkingConfiguration(PipelineGroupExternalNetworkingMode.LoadBalancerOnly,new PipelineGroupNetworkingRoute[]
    {
    new PipelineGroupNetworkingRoute("syslog-receiver1")
    })
    {
    Host = "azuremonitorpipeline.contoso.com",
    }
    },
    Tags =
    {
    ["tag1"] = "A",
    ["tag2"] = "B",
    },
};
ArmOperation<PipelineGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, pipelineGroupName, data);
PipelineGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PipelineGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
