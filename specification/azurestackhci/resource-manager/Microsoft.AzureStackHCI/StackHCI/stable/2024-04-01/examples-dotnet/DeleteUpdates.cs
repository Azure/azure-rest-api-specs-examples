using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/DeleteUpdates.json
// this example is just showing the usage of "Updates_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterUpdateResource created on azure
// for more information of creating HciClusterUpdateResource, please refer to the document of HciClusterUpdateResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
string updateName = "Microsoft4.2203.2.32";
ResourceIdentifier hciClusterUpdateResourceId = HciClusterUpdateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, updateName);
HciClusterUpdateResource hciClusterUpdate = client.GetHciClusterUpdateResource(hciClusterUpdateResourceId);

// invoke the operation
await hciClusterUpdate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
