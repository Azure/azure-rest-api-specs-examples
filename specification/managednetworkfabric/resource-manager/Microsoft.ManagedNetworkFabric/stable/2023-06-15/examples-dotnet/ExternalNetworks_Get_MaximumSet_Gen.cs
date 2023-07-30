using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_Get_MaximumSet_Gen.json
// this example is just showing the usage of "ExternalNetworks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricExternalNetworkResource created on azure
// for more information of creating NetworkFabricExternalNetworkResource, please refer to the document of NetworkFabricExternalNetworkResource
string subscriptionId = "42EEDB3B-8E17-46E3-B0B4-B1CD9842D90D";
string resourceGroupName = "rgL3IsolationDomains";
string l3IsolationDomainName = "yhtr";
string externalNetworkName = "fltpszzikbalrzaqq";
ResourceIdentifier networkFabricExternalNetworkResourceId = NetworkFabricExternalNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName, externalNetworkName);
NetworkFabricExternalNetworkResource networkFabricExternalNetwork = client.GetNetworkFabricExternalNetworkResource(networkFabricExternalNetworkResourceId);

// invoke the operation
NetworkFabricExternalNetworkResource result = await networkFabricExternalNetwork.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricExternalNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
