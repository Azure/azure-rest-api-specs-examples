using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ConnectedVMwarevSphere;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-12-01/examples/DeleteCluster.json
// this example is just showing the usage of "Clusters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VMwareClusterResource created on azure
// for more information of creating VMwareClusterResource, please refer to the document of VMwareClusterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string clusterName = "HRCluster";
ResourceIdentifier vMwareClusterResourceId = VMwareClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
VMwareClusterResource vMwareCluster = client.GetVMwareClusterResource(vMwareClusterResourceId);

// invoke the operation
await vMwareCluster.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
