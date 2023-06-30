using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabrics_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkFabrics_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricResource created on azure
// for more information of creating NetworkFabricResource, please refer to the document of NetworkFabricResource
string subscriptionId = "70A68989-24C0-4FA7-B1B2-8A2BB6D10CA8";
string resourceGroupName = "rgNetworkFabrics";
string networkFabricName = "lrhjxlxlhgvufessdcuetcwnto";
ResourceIdentifier networkFabricResourceId = NetworkFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricName);
NetworkFabricResource networkFabric = client.GetNetworkFabricResource(networkFabricResourceId);

// invoke the operation
await networkFabric.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
