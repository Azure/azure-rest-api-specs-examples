using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StoragePool;
using Azure.ResourceManager.StoragePool.Models;

// Generated from example definition: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Start.json
// this example is just showing the usage of "DiskPools_Start" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskPoolResource created on azure
// for more information of creating DiskPoolResource, please refer to the document of DiskPoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string diskPoolName = "myDiskPool";
ResourceIdentifier diskPoolResourceId = DiskPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskPoolName);
DiskPoolResource diskPool = client.GetDiskPoolResource(diskPoolResourceId);

// invoke the operation
await diskPool.StartAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
