using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L3IsolationDomains_clearArpTable_MaximumSet_Gen.json
// this example is just showing the usage of "L3IsolationDomains_clearArpTable" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this L3IsolationDomainResource created on azure
// for more information of creating L3IsolationDomainResource, please refer to the document of L3IsolationDomainResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3IsolationDomainName = "l3IsolationDomainName";
ResourceIdentifier l3IsolationDomainResourceId = L3IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName);
L3IsolationDomainResource l3IsolationDomain = client.GetL3IsolationDomainResource(l3IsolationDomainResourceId);

// invoke the operation
EnableDisableOnResources body = new EnableDisableOnResources()
{
    ResourceIds =
    {
    "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"
    },
};
await l3IsolationDomain.ClearArpTableAsync(WaitUntil.Completed, body);

Console.WriteLine($"Succeeded");
