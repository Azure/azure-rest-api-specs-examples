using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ConfigurationPolicyGroupDelete.json
// this example is just showing the usage of "ConfigurationPolicyGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnServerConfigurationPolicyGroupResource created on azure
// for more information of creating VpnServerConfigurationPolicyGroupResource, please refer to the document of VpnServerConfigurationPolicyGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string vpnServerConfigurationName = "vpnServerConfiguration1";
string configurationPolicyGroupName = "policyGroup1";
ResourceIdentifier vpnServerConfigurationPolicyGroupResourceId = VpnServerConfigurationPolicyGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName);
VpnServerConfigurationPolicyGroupResource vpnServerConfigurationPolicyGroup = client.GetVpnServerConfigurationPolicyGroupResource(vpnServerConfigurationPolicyGroupResourceId);

// invoke the operation
await vpnServerConfigurationPolicyGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
