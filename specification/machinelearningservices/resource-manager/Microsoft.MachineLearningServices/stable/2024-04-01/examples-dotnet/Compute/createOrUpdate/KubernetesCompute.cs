using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/createOrUpdate/KubernetesCompute.json
// this example is just showing the usage of "Compute_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string workspaceName = "workspaces123";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningComputeResource
MachineLearningComputeCollection collection = machineLearningWorkspace.GetMachineLearningComputes();

// invoke the operation
string computeName = "compute123";
MachineLearningComputeData data = new MachineLearningComputeData(new AzureLocation("eastus"))
{
    Properties = new MachineLearningKubernetesCompute
    {
        Properties = new MachineLearningKubernetesProperties
        {
            Namespace = "default",
            DefaultInstanceType = "defaultInstanceType",
            InstanceTypes =
            {
            ["defaultInstanceType"] = new MachineLearningInstanceTypeSchema
            {
            NodeSelector = {},
            Resources = new MachineLearningInstanceTypeSchemaResources
            {
            Requests =
            {
            ["cpu"] = "1",
            ["memory"] = "4Gi",
            ["nvidia.com/gpu"] = null
            },
            Limits =
            {
            ["cpu"] = "1",
            ["memory"] = "4Gi",
            ["nvidia.com/gpu"] = null
            },
            },
            }
            },
        },
        Description = "some compute",
        ResourceId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2"),
    },
};
ArmOperation<MachineLearningComputeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, computeName, data);
MachineLearningComputeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningComputeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
