using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Get_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkToNetworkInterconnects_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricResource created on azure
// for more information of creating NetworkFabricResource, please refer to the document of NetworkFabricResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkFabricName = "example-fabric";
ResourceIdentifier networkFabricResourceId = NetworkFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricName);
NetworkFabricResource networkFabric = client.GetNetworkFabricResource(networkFabricResourceId);

// get the collection of this NetworkToNetworkInterconnectResource
NetworkToNetworkInterconnectCollection collection = networkFabric.GetNetworkToNetworkInterconnects();

// invoke the operation
string networkToNetworkInterconnectName = "example-nni";
bool result = await collection.ExistsAsync(networkToNetworkInterconnectName);

Console.WriteLine($"Succeeded: {result}");
