using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ManagedNetwork;

// Generated from example definition: specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetworkGroup/ManagedNetworkGroupsDelete.json
// this example is just showing the usage of "ManagedNetworkGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedNetworkGroupResource created on azure
// for more information of creating ManagedNetworkGroupResource, please refer to the document of ManagedNetworkGroupResource
string subscriptionId = "subscriptionA";
string resourceGroupName = "myResourceGroup";
string managedNetworkName = "myManagedNetwork";
string managedNetworkGroupName = "myManagedNetworkGroup1";
ResourceIdentifier managedNetworkGroupResourceId = ManagedNetworkGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedNetworkName, managedNetworkGroupName);
ManagedNetworkGroupResource managedNetworkGroup = client.GetManagedNetworkGroupResource(managedNetworkGroupResourceId);

// invoke the operation
await managedNetworkGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
