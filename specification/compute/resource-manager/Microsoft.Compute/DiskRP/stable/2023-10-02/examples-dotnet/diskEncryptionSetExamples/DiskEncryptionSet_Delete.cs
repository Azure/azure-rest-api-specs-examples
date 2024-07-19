using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Delete.json
// this example is just showing the usage of "DiskEncryptionSets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskEncryptionSetResource created on azure
// for more information of creating DiskEncryptionSetResource, please refer to the document of DiskEncryptionSetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskEncryptionSetName = "myDiskEncryptionSet";
ResourceIdentifier diskEncryptionSetResourceId = DiskEncryptionSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskEncryptionSetName);
DiskEncryptionSetResource diskEncryptionSet = client.GetDiskEncryptionSetResource(diskEncryptionSetResourceId);

// invoke the operation
await diskEncryptionSet.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
