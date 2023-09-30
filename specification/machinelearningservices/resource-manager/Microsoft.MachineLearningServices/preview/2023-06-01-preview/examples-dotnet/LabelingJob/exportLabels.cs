using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/LabelingJob/exportLabels.json
// this example is just showing the usage of "LabelingJobs_ExportLabels" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningLabelingJobResource created on azure
// for more information of creating MachineLearningLabelingJobResource, please refer to the document of MachineLearningLabelingJobResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "workspace-1234";
string workspaceName = "testworkspace";
string id = "testLabelingJob";
ResourceIdentifier machineLearningLabelingJobResourceId = MachineLearningLabelingJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, id);
MachineLearningLabelingJobResource machineLearningLabelingJob = client.GetMachineLearningLabelingJobResource(machineLearningLabelingJobResourceId);

// invoke the operation
ExportSummary body = new DatasetExportSummary();
ArmOperation<ExportSummary> lro = await machineLearningLabelingJob.ExportLabelsAsync(WaitUntil.Completed, body);
ExportSummary result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
