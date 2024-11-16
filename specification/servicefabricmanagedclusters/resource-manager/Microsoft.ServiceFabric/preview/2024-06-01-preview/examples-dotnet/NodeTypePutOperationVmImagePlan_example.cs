using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/NodeTypePutOperationVmImagePlan_example.json
// this example is just showing the usage of "NodeTypes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedClusterResource created on azure
// for more information of creating ServiceFabricManagedClusterResource, please refer to the document of ServiceFabricManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
ResourceIdentifier serviceFabricManagedClusterResourceId = ServiceFabricManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ServiceFabricManagedClusterResource serviceFabricManagedCluster = client.GetServiceFabricManagedClusterResource(serviceFabricManagedClusterResourceId);

// get the collection of this ServiceFabricManagedNodeTypeResource
ServiceFabricManagedNodeTypeCollection collection = serviceFabricManagedCluster.GetServiceFabricManagedNodeTypes();

// invoke the operation
string nodeTypeName = "BE";
ServiceFabricManagedNodeTypeData data = new ServiceFabricManagedNodeTypeData()
{
    IsPrimary = false,
    VmInstanceCount = 10,
    DataDiskSizeInGB = 200,
    VmSize = "Standard_D3",
    VmImagePublisher = "testpublisher",
    VmImageOffer = "windows_2022_test",
    VmImageSku = "win_2022_test_20_10_gen2",
    VmImageVersion = "latest",
    VmImagePlan = new VmImagePlan()
    {
        Name = "win_2022_test_20_10_gen2",
        Product = "windows_2022_test",
        Publisher = "testpublisher",
    },
};
ArmOperation<ServiceFabricManagedNodeTypeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, nodeTypeName, data);
ServiceFabricManagedNodeTypeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedNodeTypeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
