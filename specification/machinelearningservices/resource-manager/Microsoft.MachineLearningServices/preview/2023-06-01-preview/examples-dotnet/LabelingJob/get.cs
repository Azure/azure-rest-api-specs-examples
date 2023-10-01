using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/LabelingJob/get.json
// this example is just showing the usage of "LabelingJobs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "workspace-1234";
string workspaceName = "testworkspace";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningLabelingJobResource
MachineLearningLabelingJobCollection collection = machineLearningWorkspace.GetMachineLearningLabelingJobs();

// invoke the operation
string id = "testLabelingJob";
bool? includeJobInstructions = true;
bool? includeLabelCategories = true;
NullableResponse<MachineLearningLabelingJobResource> response = await collection.GetIfExistsAsync(id, includeJobInstructions: includeJobInstructions, includeLabelCategories: includeLabelCategories);
MachineLearningLabelingJobResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningLabelingJobData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
