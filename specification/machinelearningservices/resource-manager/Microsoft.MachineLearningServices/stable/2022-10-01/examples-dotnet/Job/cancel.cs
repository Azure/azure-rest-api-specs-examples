using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/cancel.json
// this example is just showing the usage of "Jobs_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningJobResource created on azure
// for more information of creating MachineLearningJobResource, please refer to the document of MachineLearningJobResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string id = "string";
ResourceIdentifier machineLearningJobResourceId = MachineLearningJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, id);
MachineLearningJobResource machineLearningJob = client.GetMachineLearningJobResource(machineLearningJobResourceId);

// invoke the operation
await machineLearningJob.CancelAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
