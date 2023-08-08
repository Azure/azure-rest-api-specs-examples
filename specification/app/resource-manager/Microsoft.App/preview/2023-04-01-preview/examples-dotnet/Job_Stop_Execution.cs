using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2023-04-01-preview/examples/Job_Stop_Execution.json
// this example is just showing the usage of "Jobs_StopExecution" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppJobExecutionResource created on azure
// for more information of creating ContainerAppJobExecutionResource, please refer to the document of ContainerAppJobExecutionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string jobName = "testcontainerAppsJob0";
string jobExecutionName = "jobExecution1";
ResourceIdentifier containerAppJobExecutionResourceId = ContainerAppJobExecutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName, jobExecutionName);
ContainerAppJobExecutionResource containerAppJobExecution = client.GetContainerAppJobExecutionResource(containerAppJobExecutionResourceId);

// invoke the operation
await containerAppJobExecution.StopExecutionJobAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
