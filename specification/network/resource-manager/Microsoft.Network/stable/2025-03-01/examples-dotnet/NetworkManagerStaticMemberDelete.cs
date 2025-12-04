using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerStaticMemberDelete.json
// this example is just showing the usage of "StaticMembers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkGroupStaticMemberResource created on azure
// for more information of creating NetworkGroupStaticMemberResource, please refer to the document of NetworkGroupStaticMemberResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SampleRG";
string networkManagerName = "TestNM";
string networkGroupName = "testNetworkGroup";
string staticMemberName = "testStaticMember";
ResourceIdentifier networkGroupStaticMemberResourceId = NetworkGroupStaticMemberResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, networkGroupName, staticMemberName);
NetworkGroupStaticMemberResource networkGroupStaticMember = client.GetNetworkGroupStaticMemberResource(networkGroupStaticMemberResourceId);

// invoke the operation
await networkGroupStaticMember.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
