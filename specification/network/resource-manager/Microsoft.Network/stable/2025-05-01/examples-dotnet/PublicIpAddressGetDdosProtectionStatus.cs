using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PublicIpAddressGetDdosProtectionStatus.json
// this example is just showing the usage of "PublicIPAddresses_DdosProtectionStatus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublicIPAddressResource created on azure
// for more information of creating PublicIPAddressResource, please refer to the document of PublicIPAddressResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string publicIPAddressName = "test-pip";
ResourceIdentifier publicIPAddressResourceId = PublicIPAddressResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publicIPAddressName);
PublicIPAddressResource publicIPAddress = client.GetPublicIPAddressResource(publicIPAddressResourceId);

// invoke the operation
ArmOperation<PublicIPDdosProtectionStatusResult> lro = await publicIPAddress.DdosProtectionStatusAsync(WaitUntil.Completed);
PublicIPDdosProtectionStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
