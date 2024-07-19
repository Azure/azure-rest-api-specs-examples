using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/ServiceCreate.json
// this example is just showing the usage of "Services_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkResource created on azure
// for more information of creating MobileNetworkResource, please refer to the document of MobileNetworkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string mobileNetworkName = "testMobileNetwork";
ResourceIdentifier mobileNetworkResourceId = MobileNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mobileNetworkName);
MobileNetworkResource mobileNetwork = client.GetMobileNetworkResource(mobileNetworkResourceId);

// get the collection of this MobileNetworkServiceResource
MobileNetworkServiceCollection collection = mobileNetwork.GetMobileNetworkServices();

// invoke the operation
string serviceName = "TestService";
MobileNetworkServiceData data = new MobileNetworkServiceData(new AzureLocation("eastus"), 255, new PccRuleConfiguration[]
{
new PccRuleConfiguration("default-rule",255,new MobileNetworkServiceDataFlowTemplate[]
{
new MobileNetworkServiceDataFlowTemplate("IP-to-server",MobileNetworkSdfDirectionS.Uplink,new string[]
{
"ip"
},new string[]
{
"10.3.4.0/24"
})
{
Ports =
{
},
}
})
{
RuleQosPolicy = new PccRuleQosPolicy(new Ambr("500 Mbps","1 Gbps"))
{
FiveQi = 9,
AllocationAndRetentionPriorityLevel = 9,
PreemptionCapability = MobileNetworkPreemptionCapability.NotPreempt,
PreemptionVulnerability = MobileNetworkPreemptionVulnerability.Preemptable,
},
TrafficControl = MobileNetworkTrafficControlPermission.Enabled,
}
})
{
    ServiceQosPolicy = new MobileNetworkQosPolicy(new Ambr("500 Mbps", "1 Gbps"))
    {
        FiveQi = 9,
        AllocationAndRetentionPriorityLevel = 9,
        PreemptionCapability = MobileNetworkPreemptionCapability.NotPreempt,
        PreemptionVulnerability = MobileNetworkPreemptionVulnerability.Preemptable,
    },
};
ArmOperation<MobileNetworkServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceName, data);
MobileNetworkServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MobileNetworkServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
