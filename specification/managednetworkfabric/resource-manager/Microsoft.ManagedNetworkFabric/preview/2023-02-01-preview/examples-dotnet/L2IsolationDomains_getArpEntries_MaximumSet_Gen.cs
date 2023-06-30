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

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_getArpEntries_MaximumSet_Gen.json
// this example is just showing the usage of "L2IsolationDomains_getArpEntries" operation, for the dependent resources, they will have to be created separately.

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
ArmOperation<IDictionary<string, ARPProperties>> lro = await l2IsolationDomain.GetArpEntriesAsync(WaitUntil.Completed);
IDictionary<string, ARPProperties> result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
