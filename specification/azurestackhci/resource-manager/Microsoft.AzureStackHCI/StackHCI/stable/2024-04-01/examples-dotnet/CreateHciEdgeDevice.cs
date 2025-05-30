using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/CreateHciEdgeDevice.json
// this example is just showing the usage of "EdgeDevices_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this HciEdgeDeviceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1";
HciEdgeDeviceCollection collection = client.GetHciEdgeDevices(new ResourceIdentifier(resourceUri));

// invoke the operation
string edgeDeviceName = "default";
HciEdgeDeviceData data = new HciArcEnabledEdgeDevice
{
    Properties = new HciArcEnabledEdgeDeviceProperties
    {
        DeviceConfiguration = new HciEdgeDeviceConfiguration
        {
            NicDetails = {new HciEdgeDeviceNicDetail
            {
            AdapterName = "ethernet",
            InterfaceDescription = "NDIS 6.70 ",
            ComponentId = "VMBUS{f8615163-df3e-46c5-913f-f2d2f965ed0g} ",
            DriverVersion = "10.0.20348.1547 ",
            IPv4Address = "10.10.10.10",
            SubnetMask = "255.255.255.0",
            DefaultGateway = "10.10.10.1",
            DnsServers = {"100.10.10.1"},
            DefaultIsolationId = "0",
            }},
            DeviceMetadata = "",
        },
    },
};
ArmOperation<HciEdgeDeviceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, edgeDeviceName, data);
HciEdgeDeviceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciEdgeDeviceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
