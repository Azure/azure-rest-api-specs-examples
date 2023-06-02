using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/CreateOrUpdateAvailabilityGroupListener.json
// this example is just showing the usage of "AvailabilityGroupListeners_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlVmGroupResource created on azure
// for more information of creating SqlVmGroupResource, please refer to the document of SqlVmGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlVmGroupName = "testvmgroup";
ResourceIdentifier sqlVmGroupResourceId = SqlVmGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlVmGroupName);
SqlVmGroupResource sqlVmGroup = client.GetSqlVmGroupResource(sqlVmGroupResourceId);

// get the collection of this AvailabilityGroupListenerResource
AvailabilityGroupListenerCollection collection = sqlVmGroup.GetAvailabilityGroupListeners();

// invoke the operation
string availabilityGroupListenerName = "agl-test";
AvailabilityGroupListenerData data = new AvailabilityGroupListenerData()
{
    AvailabilityGroupName = "ag-test",
    LoadBalancerConfigurations =
    {
    new AvailabilityGroupListenerLoadBalancerConfiguration()
    {
    PrivateIPAddress = new AvailabilityGroupListenerPrivateIPAddress()
    {
    IPAddress = IPAddress.Parse("10.1.0.112"),
    SubnetResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
    },
    LoadBalancerResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test"),
    ProbePort = 59983,
    SqlVmInstances =
    {
    new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3")
    },
    }
    },
    Port = 1433,
};
ArmOperation<AvailabilityGroupListenerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, availabilityGroupListenerName, data);
AvailabilityGroupListenerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AvailabilityGroupListenerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
