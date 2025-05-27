using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/StaticCidrs_Delete.json
// this example is just showing the usage of "StaticCidrs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticCidrResource created on azure
// for more information of creating StaticCidrResource, please refer to the document of StaticCidrResource
string subscriptionId = "11111111-1111-1111-1111-111111111111";
string resourceGroupName = "rg1";
string networkManagerName = "TestNetworkManager";
string poolName = "TestPool";
string staticCidrName = "TestStaticCidr";
ResourceIdentifier staticCidrResourceId = StaticCidrResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, poolName, staticCidrName);
StaticCidrResource staticCidr = client.GetStaticCidrResource(staticCidrResourceId);

// invoke the operation
await staticCidr.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
