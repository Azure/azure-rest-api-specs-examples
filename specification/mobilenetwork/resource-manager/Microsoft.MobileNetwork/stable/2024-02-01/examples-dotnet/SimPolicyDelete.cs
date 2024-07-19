using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimPolicyDelete.json
// this example is just showing the usage of "SimPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkSimPolicyResource created on azure
// for more information of creating MobileNetworkSimPolicyResource, please refer to the document of MobileNetworkSimPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string mobileNetworkName = "testMobileNetwork";
string simPolicyName = "testPolicy";
ResourceIdentifier mobileNetworkSimPolicyResourceId = MobileNetworkSimPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mobileNetworkName, simPolicyName);
MobileNetworkSimPolicyResource mobileNetworkSimPolicy = client.GetMobileNetworkSimPolicyResource(mobileNetworkSimPolicyResourceId);

// invoke the operation
await mobileNetworkSimPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
