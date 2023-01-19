using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetwork;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetworkGroup/ManagedNetworkGroupsGet.json
// this example is just showing the usage of "ManagedNetworkGroups_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ManagedNetworkGroupResource
ManagedNetworkGroupCollection collection = managedNetwork.GetManagedNetworkGroups();

// invoke the operation
string managedNetworkGroupName = "myManagedNetworkGroup1";
bool result = await collection.ExistsAsync(managedNetworkGroupName);

Console.WriteLine($"Succeeded: {result}");
