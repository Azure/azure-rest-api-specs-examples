using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallPacketCaptureOperation.json
// this example is just showing the usage of "AzureFirewalls_PacketCaptureOperation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureFirewallResource created on azure
// for more information of creating AzureFirewallResource, please refer to the document of AzureFirewallResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string azureFirewallName = "azureFirewall1";
ResourceIdentifier azureFirewallResourceId = AzureFirewallResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureFirewallName);
AzureFirewallResource azureFirewall = client.GetAzureFirewallResource(azureFirewallResourceId);

// invoke the operation
FirewallPacketCaptureRequestContent content = new FirewallPacketCaptureRequestContent
{
    DurationInSeconds = 300,
    NumberOfPacketsToCapture = 5000,
    SasUri = new Uri("someSASURL"),
    FileName = "azureFirewallPacketCapture",
    Protocol = AzureFirewallNetworkRuleProtocol.Any,
    Flags = {new AzureFirewallPacketCaptureFlags
    {
    FlagsType = AzureFirewallPacketCaptureFlagsType.Syn,
    }, new AzureFirewallPacketCaptureFlags
    {
    FlagsType = AzureFirewallPacketCaptureFlagsType.Fin,
    }},
    Filters = {new AzureFirewallPacketCaptureRule
    {
    Sources = {"20.1.1.0"},
    Destinations = {"20.1.2.0"},
    DestinationPorts = {"4500"},
    }, new AzureFirewallPacketCaptureRule
    {
    Sources = {"10.1.1.0", "10.1.1.1"},
    Destinations = {"10.1.2.0"},
    DestinationPorts = {"123", "80"},
    }},
    Operation = AzureFirewallPacketCaptureOperationType.Status,
};
ArmOperation<AzureFirewallPacketCaptureResult> lro = await azureFirewall.PacketCaptureOperationAsync(WaitUntil.Completed, content);
AzureFirewallPacketCaptureResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
