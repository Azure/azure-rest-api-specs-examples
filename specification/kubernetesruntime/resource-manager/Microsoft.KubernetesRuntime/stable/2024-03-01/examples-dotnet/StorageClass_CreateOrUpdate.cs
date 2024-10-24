using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerOrchestratorRuntime.Models;
using Azure.ResourceManager.ContainerOrchestratorRuntime;

// Generated from example definition: 2024-03-01/StorageClass_CreateOrUpdate.json
// this example is just showing the usage of "StorageClassResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ConnectedClusterStorageClassResource
string resourceUri = "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceUri));
ConnectedClusterStorageClassCollection collection = client.GetConnectedClusterStorageClasses(scopeId);

// invoke the operation
string storageClassName = "testrwx";
ConnectedClusterStorageClassData data = new ConnectedClusterStorageClassData()
{
    Properties = new ConnectedClusterStorageClassProperties(new RwxStorageClassTypeProperties("default")),
};
ArmOperation<ConnectedClusterStorageClassResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, storageClassName, data);
ConnectedClusterStorageClassResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectedClusterStorageClassData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
