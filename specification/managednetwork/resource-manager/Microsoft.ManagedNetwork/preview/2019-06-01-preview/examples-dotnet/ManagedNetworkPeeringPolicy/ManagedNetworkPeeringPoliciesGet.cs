using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetwork;
using Azure.ResourceManager.ManagedNetwork.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetworkPeeringPolicy/ManagedNetworkPeeringPoliciesGet.json
// this example is just showing the usage of "ManagedNetworkPeeringPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedNetworkResource created on azure
// for more information of creating ManagedNetworkResource, please refer to the document of ManagedNetworkResource
string subscriptionId = "subscriptionA";
string resourceGroupName = "myResourceGroup";
string managedNetworkName = "myManagedNetwork";
ResourceIdentifier managedNetworkResourceId = ManagedNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedNetworkName);
ManagedNetworkResource managedNetwork = client.GetManagedNetworkResource(managedNetworkResourceId);

// get the collection of this ManagedNetworkPeeringPolicyResource
ManagedNetworkPeeringPolicyCollection collection = managedNetwork.GetManagedNetworkPeeringPolicies();

// invoke the operation
string managedNetworkPeeringPolicyName = "myHubAndSpoke";
NullableResponse<ManagedNetworkPeeringPolicyResource> response = await collection.GetIfExistsAsync(managedNetworkPeeringPolicyName);
ManagedNetworkPeeringPolicyResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagedNetworkPeeringPolicyData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
