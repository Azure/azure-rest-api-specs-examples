using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/preview/2022-06-15-preview/examples/OperationResultsGet.json
// this example is just showing the usage of "OperationResults_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string operationResultId = "exampleid";
ArmOperation<OperationResultsDescription> lro = await subscriptionResource.GetOperationResultAsync(WaitUntil.Completed, operationResultId);
OperationResultsDescription result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
