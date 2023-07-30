using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.json
// this example is just showing the usage of "InternalNetworks_ListByL3IsolationDomain" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricL3IsolationDomainResource created on azure
// for more information of creating NetworkFabricL3IsolationDomainResource, please refer to the document of NetworkFabricL3IsolationDomainResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string l3IsolationDomainName = "example-l3domain";
ResourceIdentifier networkFabricL3IsolationDomainResourceId = NetworkFabricL3IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName);
NetworkFabricL3IsolationDomainResource networkFabricL3IsolationDomain = client.GetNetworkFabricL3IsolationDomainResource(networkFabricL3IsolationDomainResourceId);

// get the collection of this NetworkFabricInternalNetworkResource
NetworkFabricInternalNetworkCollection collection = networkFabricL3IsolationDomain.GetNetworkFabricInternalNetworks();

// invoke the operation and iterate over the result
await foreach (NetworkFabricInternalNetworkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkFabricInternalNetworkData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
