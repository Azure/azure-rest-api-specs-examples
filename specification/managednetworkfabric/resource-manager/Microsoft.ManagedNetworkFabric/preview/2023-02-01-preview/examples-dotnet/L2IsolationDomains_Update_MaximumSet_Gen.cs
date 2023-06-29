using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_Update_MaximumSet_Gen.json
// this example is just showing the usage of "L2IsolationDomains_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this L2IsolationDomainResource created on azure
// for more information of creating L2IsolationDomainResource, please refer to the document of L2IsolationDomainResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l2IsolationDomainName = "l2IsolationDomainName";
ResourceIdentifier l2IsolationDomainResourceId = L2IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l2IsolationDomainName);
L2IsolationDomainResource l2IsolationDomain = client.GetL2IsolationDomainResource(l2IsolationDomainResourceId);

// invoke the operation
L2IsolationDomainPatch patch = new L2IsolationDomainPatch()
{
    Mtu = 9000,
};
ArmOperation<L2IsolationDomainResource> lro = await l2IsolationDomain.UpdateAsync(WaitUntil.Completed, patch);
L2IsolationDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
L2IsolationDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
