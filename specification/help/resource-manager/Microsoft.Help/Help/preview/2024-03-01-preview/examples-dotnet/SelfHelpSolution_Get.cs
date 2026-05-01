using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SelfHelp.Models;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: 2024-03-01-preview/SelfHelpSolution_Get.json
// this example is just showing the usage of "SolutionResourceSelfHelp_GetSelfHelpSolutionById" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string solutionId = "SolutionId1";
SelfHelpSolutionResultResource result = await tenantResource.GetSelfHelpSolutionResultAsync(solutionId);

Console.WriteLine($"Succeeded: {result}");
