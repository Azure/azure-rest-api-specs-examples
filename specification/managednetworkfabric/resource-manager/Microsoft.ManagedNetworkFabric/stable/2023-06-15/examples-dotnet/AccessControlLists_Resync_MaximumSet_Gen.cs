using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_Resync_MaximumSet_Gen.json
// this example is just showing the usage of "AccessControlLists_Resync" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricAccessControlListResource created on azure
// for more information of creating NetworkFabricAccessControlListResource, please refer to the document of NetworkFabricAccessControlListResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string accessControlListName = "example-acl";
ResourceIdentifier networkFabricAccessControlListResourceId = NetworkFabricAccessControlListResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accessControlListName);
NetworkFabricAccessControlListResource networkFabricAccessControlList = client.GetNetworkFabricAccessControlListResource(networkFabricAccessControlListResourceId);

// invoke the operation
ArmOperation<StateUpdateCommonPostActionResult> lro = await networkFabricAccessControlList.ResyncAsync(WaitUntil.Completed);
StateUpdateCommonPostActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
