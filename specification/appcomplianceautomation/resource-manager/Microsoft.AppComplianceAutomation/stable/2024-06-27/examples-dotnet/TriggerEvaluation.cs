using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppComplianceAutomation.Models;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/TriggerEvaluation.json
// this example is just showing the usage of "ProviderActions_TriggerEvaluation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
TriggerEvaluationContent content = new TriggerEvaluationContent(new string[]
{
});
ArmOperation<TriggerEvaluationResult> lro = await tenantResource.TriggerEvaluationProviderActionAsync(WaitUntil.Completed, content);
TriggerEvaluationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
