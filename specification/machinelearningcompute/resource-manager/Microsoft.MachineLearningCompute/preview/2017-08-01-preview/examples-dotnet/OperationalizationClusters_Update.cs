using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearningCompute;
using Azure.ResourceManager.MachineLearningCompute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/OperationalizationClusters_Update.json
// this example is just showing the usage of "OperationalizationClusters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalizationClusterResource created on azure
// for more information of creating OperationalizationClusterResource, please refer to the document of OperationalizationClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string clusterName = "myCluster";
ResourceIdentifier operationalizationClusterResourceId = OperationalizationClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
OperationalizationClusterResource operationalizationCluster = client.GetOperationalizationClusterResource(operationalizationClusterResourceId);

// invoke the operation
OperationalizationClusterPatch patch = new OperationalizationClusterPatch()
{
    Tags =
    {
    ["key1"] = "value1",
    },
};
OperationalizationClusterResource result = await operationalizationCluster.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OperationalizationClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
