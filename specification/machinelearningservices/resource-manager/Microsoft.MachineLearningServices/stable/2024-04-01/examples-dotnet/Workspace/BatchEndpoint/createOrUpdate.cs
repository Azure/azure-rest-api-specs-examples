using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/BatchEndpoint/createOrUpdate.json
// this example is just showing the usage of "BatchEndpoints_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningBatchEndpointResource
MachineLearningBatchEndpointCollection collection = machineLearningWorkspace.GetMachineLearningBatchEndpoints();

// invoke the operation
string endpointName = "testEndpointName";
MachineLearningBatchEndpointData data = new MachineLearningBatchEndpointData(new AzureLocation("string"), new MachineLearningBatchEndpointProperties(MachineLearningEndpointAuthMode.AmlToken)
{
    DefaultsDeploymentName = "string",
    Description = "string",
    Properties =
    {
    ["string"] = "string"
    },
})
{
    Kind = "string",
    Identity = new ManagedServiceIdentity("SystemAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("string")] = new UserAssignedIdentity()
        },
    },
    Sku = new MachineLearningSku("string")
    {
        Tier = MachineLearningSkuTier.Free,
        Size = "string",
        Family = "string",
        Capacity = 1,
    },
    Tags = { },
};
ArmOperation<MachineLearningBatchEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, endpointName, data);
MachineLearningBatchEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningBatchEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
