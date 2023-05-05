using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2023-01-01-preview/examples/GetDiagnosticForKeyVaultResource.json
// this example is just showing the usage of "Diagnostics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SelfHelpDiagnosticResource created on azure
// for more information of creating SelfHelpDiagnosticResource, please refer to the document of SelfHelpDiagnosticResource
string scope = "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read";
string diagnosticsResourceName = "VMNotWorkingInsight";
ResourceIdentifier selfHelpDiagnosticResourceId = SelfHelpDiagnosticResource.CreateResourceIdentifier(scope, diagnosticsResourceName);
SelfHelpDiagnosticResource selfHelpDiagnosticResource = client.GetSelfHelpDiagnosticResource(selfHelpDiagnosticResourceId);

// invoke the operation
SelfHelpDiagnosticResource result = await selfHelpDiagnosticResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SelfHelpDiagnosticResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
