using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerOrchestratorRuntime.Models;
using Azure.ResourceManager.ContainerOrchestratorRuntime;

// Generated from example definition: 2024-03-01/LoadBalancers_CreateOrUpdate.json
// this example is just showing the usage of "LoadBalancer_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectedClusterLoadBalancerResource created on azure
// for more information of creating ConnectedClusterLoadBalancerResource, please refer to the document of ConnectedClusterLoadBalancerResource
string resourceUri = "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1";
string loadBalancerName = "testlb";
ResourceIdentifier connectedClusterLoadBalancerResourceId = ConnectedClusterLoadBalancerResource.CreateResourceIdentifier(resourceUri, loadBalancerName);
ConnectedClusterLoadBalancerResource connectedClusterLoadBalancer = client.GetConnectedClusterLoadBalancerResource(connectedClusterLoadBalancerResourceId);

// invoke the operation
ConnectedClusterLoadBalancerData data = new ConnectedClusterLoadBalancerData()
{
    Properties = new ConnectedClusterLoadBalancerProperties(new string[]
{
"192.168.50.1/24","192.168.51.2-192.168.51.10"
}, AdvertiseMode.Arp)
    {
        ServiceSelector =
        {
        ["app"] = "frontend",
        },
    },
};
ArmOperation<ConnectedClusterLoadBalancerResource> lro = await connectedClusterLoadBalancer.UpdateAsync(WaitUntil.Completed, data);
ConnectedClusterLoadBalancerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectedClusterLoadBalancerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
