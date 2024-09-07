using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/FeaturesetVersion/backfill.json
// this example is just showing the usage of "FeaturesetVersions_Backfill" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningFeatureSetVersionResource created on azure
// for more information of creating MachineLearningFeatureSetVersionResource, please refer to the document of MachineLearningFeatureSetVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
string version = "string";
ResourceIdentifier machineLearningFeatureSetVersionResourceId = MachineLearningFeatureSetVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name, version);
MachineLearningFeatureSetVersionResource machineLearningFeatureSetVersion = client.GetMachineLearningFeatureSetVersionResource(machineLearningFeatureSetVersionResourceId);

// invoke the operation
FeatureSetVersionBackfillContent content = new FeatureSetVersionBackfillContent()
{
    DisplayName = "string",
    Description = "string",
    FeatureWindow = new FeatureWindow()
    {
        FeatureWindowStart = DateTimeOffset.Parse("2020-01-01T12:34:56.999+00:51"),
        FeatureWindowEnd = DateTimeOffset.Parse("2020-01-01T12:34:56.999+00:51"),
    },
    DataAvailabilityStatus =
    {
    DataAvailabilityStatus.None
    },
    JobId = "string",
    Tags =
    {
    ["string"] = "string",
    },
    ResourceInstanceType = "string",
    SparkConfiguration =
    {
    ["string"] = "string",
    },
};
ArmOperation<FeaturesetVersionBackfillResponse> lro = await machineLearningFeatureSetVersion.BackfillAsync(WaitUntil.Completed, content);
FeaturesetVersionBackfillResponse result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
