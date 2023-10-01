using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/OnlineDeployment/getLogs.json
// this example is just showing the usage of "OnlineDeployments_GetLogs" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningOnlineDeploymentResource created on azure
// for more information of creating MachineLearningOnlineDeploymentResource, please refer to the document of MachineLearningOnlineDeploymentResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg123";
string workspaceName = "workspace123";
string endpointName = "testEndpoint";
string deploymentName = "testDeployment";
ResourceIdentifier machineLearningOnlineDeploymentResourceId = MachineLearningOnlineDeploymentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, endpointName, deploymentName);
MachineLearningOnlineDeploymentResource machineLearningOnlineDeployment = client.GetMachineLearningOnlineDeploymentResource(machineLearningOnlineDeploymentResourceId);

// invoke the operation
MachineLearningDeploymentLogsContent content = new MachineLearningDeploymentLogsContent()
{
    ContainerType = MachineLearningContainerType.StorageInitializer,
    Tail = 0,
};
MachineLearningDeploymentLogs result = await machineLearningOnlineDeployment.GetLogsAsync(content);

Console.WriteLine($"Succeeded: {result}");
