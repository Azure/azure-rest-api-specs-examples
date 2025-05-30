using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/BatchDeployment/update.json
// this example is just showing the usage of "BatchDeployments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningBatchDeploymentResource created on azure
// for more information of creating MachineLearningBatchDeploymentResource, please refer to the document of MachineLearningBatchDeploymentResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string endpointName = "testEndpointName";
string deploymentName = "testDeploymentName";
ResourceIdentifier machineLearningBatchDeploymentResourceId = MachineLearningBatchDeploymentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, endpointName, deploymentName);
MachineLearningBatchDeploymentResource machineLearningBatchDeployment = client.GetMachineLearningBatchDeploymentResource(machineLearningBatchDeploymentResourceId);

// invoke the operation
MachineLearningBatchDeploymentPatch patch = new MachineLearningBatchDeploymentPatch
{
    PartialBatchDeploymentDescription = "string",
    Tags = { },
};
ArmOperation<MachineLearningBatchDeploymentResource> lro = await machineLearningBatchDeployment.UpdateAsync(WaitUntil.Completed, patch);
MachineLearningBatchDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningBatchDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
