using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs.Models;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/WorkloadNetworks_UpdateDhcp.json
// this example is just showing the usage of "WorkloadNetworkDhcp_UpdateDhcp" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadNetworkDhcpResource created on azure
// for more information of creating WorkloadNetworkDhcpResource, please refer to the document of WorkloadNetworkDhcpResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string dhcpId = "dhcp1";
ResourceIdentifier workloadNetworkDhcpResourceId = WorkloadNetworkDhcpResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, dhcpId);
WorkloadNetworkDhcpResource workloadNetworkDhcp = client.GetWorkloadNetworkDhcpResource(workloadNetworkDhcpResourceId);

// invoke the operation
WorkloadNetworkDhcpData data = new WorkloadNetworkDhcpData
{
    Properties = new WorkloadNetworkDhcpServer
    {
        ServerAddress = "40.1.5.1/24",
        LeaseTime = 86400L,
        Revision = 1L,
    },
};
ArmOperation<WorkloadNetworkDhcpResource> lro = await workloadNetworkDhcp.UpdateAsync(WaitUntil.Completed, data);
WorkloadNetworkDhcpResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkloadNetworkDhcpData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
