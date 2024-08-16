using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/DeleteUpdateRuns.json
// this example is just showing the usage of "UpdateRuns_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterUpdateRunResource created on azure
// for more information of creating HciClusterUpdateRunResource, please refer to the document of HciClusterUpdateRunResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
string updateName = "Microsoft4.2203.2.32";
string updateRunName = "23b779ba-0d52-4a80-8571-45ca74664ec3";
ResourceIdentifier hciClusterUpdateRunResourceId = HciClusterUpdateRunResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, updateName, updateRunName);
HciClusterUpdateRunResource hciClusterUpdateRun = client.GetHciClusterUpdateRunResource(hciClusterUpdateRunResourceId);

// invoke the operation
await hciClusterUpdateRun.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
