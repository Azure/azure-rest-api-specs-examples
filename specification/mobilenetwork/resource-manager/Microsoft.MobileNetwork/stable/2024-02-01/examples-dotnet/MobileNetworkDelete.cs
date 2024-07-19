using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/MobileNetworkDelete.json
// this example is just showing the usage of "MobileNetworks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkResource created on azure
// for more information of creating MobileNetworkResource, please refer to the document of MobileNetworkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string mobileNetworkName = "testMobileNetwork";
ResourceIdentifier mobileNetworkResourceId = MobileNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mobileNetworkName);
MobileNetworkResource mobileNetwork = client.GetMobileNetworkResource(mobileNetworkResourceId);

// invoke the operation
await mobileNetwork.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
