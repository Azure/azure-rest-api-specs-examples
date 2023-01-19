using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetwork;
using Azure.ResourceManager.ManagedNetwork.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetworkPeeringPolicy/ManagedNetworkPeeringPoliciesDelete.json
// this example is just showing the usage of "ManagedNetworkPeeringPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedNetworkPeeringPolicyResource created on azure
// for more information of creating ManagedNetworkPeeringPolicyResource, please refer to the document of ManagedNetworkPeeringPolicyResource
string subscriptionId = "subscriptionA";
string resourceGroupName = "myResourceGroup";
string managedNetworkName = "myManagedNetwork";
string managedNetworkPeeringPolicyName = "myHubAndSpoke";
ResourceIdentifier managedNetworkPeeringPolicyResourceId = ManagedNetworkPeeringPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName);
ManagedNetworkPeeringPolicyResource managedNetworkPeeringPolicy = client.GetManagedNetworkPeeringPolicyResource(managedNetworkPeeringPolicyResourceId);

// invoke the operation
await managedNetworkPeeringPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
