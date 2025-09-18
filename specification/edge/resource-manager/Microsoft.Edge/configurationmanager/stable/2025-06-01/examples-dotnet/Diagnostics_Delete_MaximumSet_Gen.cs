using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Diagnostics_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "Diagnostic_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeDiagnosticResource created on azure
// for more information of creating EdgeDiagnosticResource, please refer to the document of EdgeDiagnosticResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string diagnosticName = "testname";
ResourceIdentifier edgeDiagnosticResourceId = EdgeDiagnosticResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diagnosticName);
EdgeDiagnosticResource edgeDiagnostic = client.GetEdgeDiagnosticResource(edgeDiagnosticResourceId);

// invoke the operation
await edgeDiagnostic.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
