using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L2IsolationDomains_Update_MaximumSet_Gen.json
// this example is just showing the usage of "L2IsolationDomains_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricL2IsolationDomainResource created on azure
// for more information of creating NetworkFabricL2IsolationDomainResource, please refer to the document of NetworkFabricL2IsolationDomainResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string l2IsolationDomainName = "example-l2Domain";
ResourceIdentifier networkFabricL2IsolationDomainResourceId = NetworkFabricL2IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l2IsolationDomainName);
NetworkFabricL2IsolationDomainResource networkFabricL2IsolationDomain = client.GetNetworkFabricL2IsolationDomainResource(networkFabricL2IsolationDomainResourceId);

// invoke the operation
NetworkFabricL2IsolationDomainPatch patch = new NetworkFabricL2IsolationDomainPatch()
{
    Annotation = "annotation1",
    Mtu = 6000,
    Tags =
    {
    ["keyID"] = "keyValue",
    },
};
ArmOperation<NetworkFabricL2IsolationDomainResource> lro = await networkFabricL2IsolationDomain.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricL2IsolationDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricL2IsolationDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
