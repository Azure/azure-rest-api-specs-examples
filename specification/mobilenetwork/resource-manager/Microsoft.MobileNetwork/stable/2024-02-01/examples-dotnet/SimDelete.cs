using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimDelete.json
// this example is just showing the usage of "Sims_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkSimResource created on azure
// for more information of creating MobileNetworkSimResource, please refer to the document of MobileNetworkSimResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testResourceGroupName";
string simGroupName = "testSimGroup";
string simName = "testSim";
ResourceIdentifier mobileNetworkSimResourceId = MobileNetworkSimResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, simGroupName, simName);
MobileNetworkSimResource mobileNetworkSim = client.GetMobileNetworkSimResource(mobileNetworkSimResourceId);

// invoke the operation
await mobileNetworkSim.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
