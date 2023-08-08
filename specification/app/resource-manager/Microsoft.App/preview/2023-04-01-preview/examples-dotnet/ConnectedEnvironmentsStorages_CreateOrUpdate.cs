using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2023-04-01-preview/examples/ConnectedEnvironmentsStorages_CreateOrUpdate.json
// this example is just showing the usage of "ConnectedEnvironmentsStorages_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppConnectedEnvironmentStorageResource created on azure
// for more information of creating ContainerAppConnectedEnvironmentStorageResource, please refer to the document of ContainerAppConnectedEnvironmentStorageResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string connectedEnvironmentName = "env";
string storageName = "jlaw-demo1";
ResourceIdentifier containerAppConnectedEnvironmentStorageResourceId = ContainerAppConnectedEnvironmentStorageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, connectedEnvironmentName, storageName);
ContainerAppConnectedEnvironmentStorageResource containerAppConnectedEnvironmentStorage = client.GetContainerAppConnectedEnvironmentStorageResource(containerAppConnectedEnvironmentStorageResourceId);

// invoke the operation
ContainerAppConnectedEnvironmentStorageData data = new ContainerAppConnectedEnvironmentStorageData()
{
    ConnectedEnvironmentStorageAzureFile = new ContainerAppAzureFileProperties()
    {
        AccountName = "account1",
        AccountKey = "key",
        AccessMode = ContainerAppAccessMode.ReadOnly,
        ShareName = "share1",
    },
};
ArmOperation<ContainerAppConnectedEnvironmentStorageResource> lro = await containerAppConnectedEnvironmentStorage.UpdateAsync(WaitUntil.Completed, data);
ContainerAppConnectedEnvironmentStorageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppConnectedEnvironmentStorageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
