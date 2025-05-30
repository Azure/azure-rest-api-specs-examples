using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L2IsolationDomains_Create_MaximumSet_Gen.json
// this example is just showing the usage of "L2IsolationDomains_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkFabricL2IsolationDomainResource
NetworkFabricL2IsolationDomainCollection collection = resourceGroupResource.GetNetworkFabricL2IsolationDomains();

// invoke the operation
string l2IsolationDomainName = "example-l2domain";
NetworkFabricL2IsolationDomainData data = new NetworkFabricL2IsolationDomainData(new AzureLocation("eastus"), new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"), 501)
{
    Annotation = "annotation",
    Mtu = 1500,
    Tags =
    {
    ["keyID"] = "keyValue"
    },
};
ArmOperation<NetworkFabricL2IsolationDomainResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, l2IsolationDomainName, data);
NetworkFabricL2IsolationDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricL2IsolationDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
