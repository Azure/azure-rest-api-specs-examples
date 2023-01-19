using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetProblemClassification.json
// this example is just showing the usage of "ProblemClassifications_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportAzureServiceResource created on azure
// for more information of creating SupportAzureServiceResource, please refer to the document of SupportAzureServiceResource
string serviceName = "service_guid";
ResourceIdentifier supportAzureServiceResourceId = SupportAzureServiceResource.CreateResourceIdentifier(serviceName);
SupportAzureServiceResource supportAzureService = client.GetSupportAzureServiceResource(supportAzureServiceResourceId);

// get the collection of this ProblemClassificationResource
ProblemClassificationCollection collection = supportAzureService.GetProblemClassifications();

// invoke the operation
string problemClassificationName = "problemClassification_guid";
bool result = await collection.ExistsAsync(problemClassificationName);

Console.WriteLine($"Succeeded: {result}");
