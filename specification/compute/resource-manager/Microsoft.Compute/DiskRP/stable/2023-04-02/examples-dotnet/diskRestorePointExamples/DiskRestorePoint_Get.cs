using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskRestorePointExamples/DiskRestorePoint_Get.json
// this example is just showing the usage of "DiskRestorePoint_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorePointResource created on azure
// for more information of creating RestorePointResource, please refer to the document of RestorePointResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string restorePointGroupName = "rpc";
string vmRestorePointName = "vmrp";
ResourceIdentifier restorePointResourceId = RestorePointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, restorePointGroupName, vmRestorePointName);
RestorePointResource restorePoint = client.GetRestorePointResource(restorePointResourceId);

// get the collection of this DiskRestorePointResource
DiskRestorePointCollection collection = restorePoint.GetDiskRestorePoints();

// invoke the operation
string diskRestorePointName = "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745";
bool result = await collection.ExistsAsync(diskRestorePointName);

Console.WriteLine($"Succeeded: {result}");
