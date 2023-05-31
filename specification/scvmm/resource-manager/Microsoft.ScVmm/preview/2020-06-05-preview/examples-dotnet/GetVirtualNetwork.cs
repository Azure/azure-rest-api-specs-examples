using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ArcScVmm;
using Azure.ResourceManager.ArcScVmm.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetVirtualNetwork.json
// this example is just showing the usage of "VirtualNetworks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmVirtualNetworkResource created on azure
// for more information of creating ScVmmVirtualNetworkResource, please refer to the document of ScVmmVirtualNetworkResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string virtualNetworkName = "HRVirtualNetwork";
ResourceIdentifier scVmmVirtualNetworkResourceId = ScVmmVirtualNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName);
ScVmmVirtualNetworkResource scVmmVirtualNetwork = client.GetScVmmVirtualNetworkResource(scVmmVirtualNetworkResourceId);

// invoke the operation
ScVmmVirtualNetworkResource result = await scVmmVirtualNetwork.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScVmmVirtualNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
