using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/NetworkSecurityGroups_Delete.json
// this example is just showing the usage of "NetworkSecurityGroup_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciVmNetworkSecurityGroupResource created on azure
// for more information of creating HciVmNetworkSecurityGroupResource, please refer to the document of HciVmNetworkSecurityGroupResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string networkSecurityGroupName = "testnsg";
ResourceIdentifier hciVmNetworkSecurityGroupResourceId = HciVmNetworkSecurityGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityGroupName);
HciVmNetworkSecurityGroupResource hciVmNetworkSecurityGroup = client.GetHciVmNetworkSecurityGroupResource(hciVmNetworkSecurityGroupResourceId);

// invoke the operation
await hciVmNetworkSecurityGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
