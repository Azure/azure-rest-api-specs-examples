using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_CreateOrUpdate.json
// this example is just showing the usage of "Disks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabUserResource created on azure
// for more information of creating DevTestLabUserResource, please refer to the document of DevTestLabUserResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string userName = "{userId}";
ResourceIdentifier devTestLabUserResourceId = DevTestLabUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName);
DevTestLabUserResource devTestLabUser = client.GetDevTestLabUserResource(devTestLabUserResourceId);

// get the collection of this DevTestLabDiskResource
DevTestLabDiskCollection collection = devTestLabUser.GetDevTestLabDisks();

// invoke the operation
string name = "{diskName}";
DevTestLabDiskData data = new DevTestLabDiskData(default)
{
    DiskType = DevTestLabStorageType.Standard,
    DiskSizeGiB = 1023,
    LeasedByLabVmId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/vmName"),
};
ArmOperation<DevTestLabDiskResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DevTestLabDiskResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabDiskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
