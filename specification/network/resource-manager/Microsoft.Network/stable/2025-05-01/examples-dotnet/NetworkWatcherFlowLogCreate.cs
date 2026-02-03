using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherFlowLogCreate.json
// this example is just showing the usage of "FlowLogs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkWatcherResource created on azure
// for more information of creating NetworkWatcherResource, please refer to the document of NetworkWatcherResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkWatcherName = "nw1";
ResourceIdentifier networkWatcherResourceId = NetworkWatcherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkWatcherName);
NetworkWatcherResource networkWatcher = client.GetNetworkWatcherResource(networkWatcherResourceId);

// get the collection of this FlowLogResource
FlowLogCollection collection = networkWatcher.GetFlowLogs();

// invoke the operation
string flowLogName = "fl";
FlowLogData data = new FlowLogData
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity()
        },
    },
    TargetResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg"),
    StorageId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe"),
    EnabledFilteringCriteria = "srcIP=158.255.7.8 || dstPort=56891",
    RecordTypes = "B,E",
    Enabled = true,
    Format = new FlowLogProperties
    {
        FormatType = FlowLogFormatType.Json,
        Version = 1,
    },
    Location = new AzureLocation("centraluseuap"),
};
ArmOperation<FlowLogResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, flowLogName, data);
FlowLogResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FlowLogData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
