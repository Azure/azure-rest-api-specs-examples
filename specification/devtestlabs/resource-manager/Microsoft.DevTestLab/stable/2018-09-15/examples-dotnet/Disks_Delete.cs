using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_Delete.json
// this example is just showing the usage of "Disks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabDiskResource created on azure
// for more information of creating DevTestLabDiskResource, please refer to the document of DevTestLabDiskResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string userName = "{userId}";
string name = "{diskName}";
ResourceIdentifier devTestLabDiskResourceId = DevTestLabDiskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName, name);
DevTestLabDiskResource devTestLabDisk = client.GetDevTestLabDiskResource(devTestLabDiskResourceId);

// invoke the operation
await devTestLabDisk.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
