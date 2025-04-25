using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ArizeAIObservabilityEval.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ArizeAIObservabilityEval;

// Generated from example definition: 2024-10-01-preview/Organizations_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArizeAIObservabilityEvalOrganizationResource created on azure
// for more information of creating ArizeAIObservabilityEvalOrganizationResource, please refer to the document of ArizeAIObservabilityEvalOrganizationResource
string subscriptionId = "4DEBE8B4-8BA4-42F8-AE50-FBEF318751D1";
string resourceGroupName = "rgopenapi";
string organizationname = "test-organization-1";
ResourceIdentifier arizeAIObservabilityEvalOrganizationResourceId = ArizeAIObservabilityEvalOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationname);
ArizeAIObservabilityEvalOrganizationResource arizeAIObservabilityEvalOrganization = client.GetArizeAIObservabilityEvalOrganizationResource(arizeAIObservabilityEvalOrganizationResourceId);

// invoke the operation
await arizeAIObservabilityEvalOrganization.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
