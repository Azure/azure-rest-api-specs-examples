using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerOrchestratorRuntime.Models;
using Azure.ResourceManager.ContainerOrchestratorRuntime;

// Generated from example definition: 2024-03-01/StorageClass_Update.json
// this example is just showing the usage of "StorageClassResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectedClusterStorageClassResource created on azure
// for more information of creating ConnectedClusterStorageClassResource, please refer to the document of ConnectedClusterStorageClassResource
string resourceUri = "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1";
string storageClassName = "testrwx";
ResourceIdentifier connectedClusterStorageClassResourceId = ConnectedClusterStorageClassResource.CreateResourceIdentifier(resourceUri, storageClassName);
ConnectedClusterStorageClassResource connectedClusterStorageClass = client.GetConnectedClusterStorageClassResource(connectedClusterStorageClassResourceId);

// invoke the operation
ConnectedClusterStorageClassPatch patch = new ConnectedClusterStorageClassPatch()
{
    Properties = new StorageClassPropertiesUpdate()
    {
        TypeProperties = new StorageClassTypePropertiesUpdate()
        {
            BackingStorageClassName = "default",
        },
    },
};
ArmOperation<ConnectedClusterStorageClassResource> lro = await connectedClusterStorageClass.UpdateAsync(WaitUntil.Completed, patch);
ConnectedClusterStorageClassResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectedClusterStorageClassData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
