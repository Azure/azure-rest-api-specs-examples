using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Advisor;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/DeleteSuppression.json
// this example is just showing the usage of "Suppressions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SuppressionContractResource created on azure
// for more information of creating SuppressionContractResource, please refer to the document of SuppressionContractResource
string resourceUri = "resourceUri";
string recommendationId = "recommendationId";
string name = "suppressionName1";
ResourceIdentifier suppressionContractResourceId = SuppressionContractResource.CreateResourceIdentifier(resourceUri, recommendationId, name);
SuppressionContractResource suppressionContract = client.GetSuppressionContractResource(suppressionContractResourceId);

// invoke the operation
await suppressionContract.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
