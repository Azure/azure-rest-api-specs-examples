using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;
using Azure.ResourceManager.Support.Models;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyProblemClassifications.json
// this example is just showing the usage of "ProblemClassificationsNoSubscription_classifyProblems" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportAzureServiceResource created on azure
// for more information of creating SupportAzureServiceResource, please refer to the document of SupportAzureServiceResource
string problemServiceName = "serviceId1";
ResourceIdentifier supportAzureServiceResourceId = SupportAzureServiceResource.CreateResourceIdentifier(problemServiceName);
SupportAzureServiceResource supportAzureService = client.GetSupportAzureServiceResource(supportAzureServiceResourceId);

// invoke the operation
ServiceProblemClassificationContent content = new ServiceProblemClassificationContent("Can not connect to Windows VM");
ServiceProblemClassificationListResult result = await supportAzureService.ClassifyServiceProblemAsync(content);

Console.WriteLine($"Succeeded: {result}");
