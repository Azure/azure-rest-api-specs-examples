using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_DeletePublicIPs.json
// this example is just showing the usage of "WorkloadNetworks_DeletePublicIP" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadNetworkPublicIPResource created on azure
// for more information of creating WorkloadNetworkPublicIPResource, please refer to the document of WorkloadNetworkPublicIPResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string publicIPId = "publicIP1";
ResourceIdentifier workloadNetworkPublicIPResourceId = WorkloadNetworkPublicIPResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, publicIPId);
WorkloadNetworkPublicIPResource workloadNetworkPublicIP = client.GetWorkloadNetworkPublicIPResource(workloadNetworkPublicIPResourceId);

// invoke the operation
await workloadNetworkPublicIP.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
