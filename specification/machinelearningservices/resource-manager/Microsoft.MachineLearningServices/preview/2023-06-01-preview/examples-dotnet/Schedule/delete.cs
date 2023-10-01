using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Schedule/delete.json
// this example is just showing the usage of "Schedules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningScheduleResource created on azure
// for more information of creating MachineLearningScheduleResource, please refer to the document of MachineLearningScheduleResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningScheduleResourceId = MachineLearningScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningScheduleResource machineLearningSchedule = client.GetMachineLearningScheduleResource(machineLearningScheduleResourceId);

// invoke the operation
await machineLearningSchedule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
