using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: 2025-03-01-preview/NodeTypePatchOperationAutoScale_example.json
// this example is just showing the usage of "NodeType_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedNodeTypeResource created on azure
// for more information of creating ServiceFabricManagedNodeTypeResource, please refer to the document of ServiceFabricManagedNodeTypeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string nodeTypeName = "BE";
ResourceIdentifier serviceFabricManagedNodeTypeResourceId = ServiceFabricManagedNodeTypeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, nodeTypeName);
ServiceFabricManagedNodeTypeResource serviceFabricManagedNodeType = client.GetServiceFabricManagedNodeTypeResource(serviceFabricManagedNodeTypeResourceId);

// invoke the operation
ServiceFabricManagedNodeTypePatch patch = new ServiceFabricManagedNodeTypePatch
{
    Tags =
    {
    ["a"] = "b"
    },
    Sku = new NodeTypeSku(10)
    {
        Name = "Standard_S0",
        Tier = "Standard",
    },
};
ArmOperation<ServiceFabricManagedNodeTypeResource> lro = await serviceFabricManagedNodeType.UpdateAsync(WaitUntil.Completed, patch);
ServiceFabricManagedNodeTypeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedNodeTypeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
