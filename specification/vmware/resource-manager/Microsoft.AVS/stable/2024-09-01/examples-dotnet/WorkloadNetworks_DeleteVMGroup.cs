using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/WorkloadNetworks_DeleteVMGroup.json
// this example is just showing the usage of "WorkloadNetworkVMGroup_DeleteVmGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadNetworkVmGroupResource created on azure
// for more information of creating WorkloadNetworkVmGroupResource, please refer to the document of WorkloadNetworkVmGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string vmGroupId = "vmGroup1";
ResourceIdentifier workloadNetworkVmGroupResourceId = WorkloadNetworkVmGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, vmGroupId);
WorkloadNetworkVmGroupResource workloadNetworkVmGroup = client.GetWorkloadNetworkVmGroupResource(workloadNetworkVmGroupResourceId);

// invoke the operation
await workloadNetworkVmGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
