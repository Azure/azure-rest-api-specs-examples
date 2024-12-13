using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/AzureFirewallDelete.json
// this example is just showing the usage of "AzureFirewalls_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureFirewallResource created on azure
// for more information of creating AzureFirewallResource, please refer to the document of AzureFirewallResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string azureFirewallName = "azurefirewall";
ResourceIdentifier azureFirewallResourceId = AzureFirewallResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureFirewallName);
AzureFirewallResource azureFirewall = client.GetAzureFirewallResource(azureFirewallResourceId);

// invoke the operation
await azureFirewall.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
