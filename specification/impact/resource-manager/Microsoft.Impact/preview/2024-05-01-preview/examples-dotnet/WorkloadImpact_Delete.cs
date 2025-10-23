using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/WorkloadImpact_Delete.json
// this example is just showing the usage of "WorkloadImpact_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadImpactResource created on azure
// for more information of creating WorkloadImpactResource, please refer to the document of WorkloadImpactResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string workloadImpactName = "impact-001";
ResourceIdentifier workloadImpactResourceId = WorkloadImpactResource.CreateResourceIdentifier(subscriptionId, workloadImpactName);
WorkloadImpactResource workloadImpact = client.GetWorkloadImpactResource(workloadImpactResourceId);

// invoke the operation
await workloadImpact.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
