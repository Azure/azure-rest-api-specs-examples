using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/getDiagnosticSettingsCategory.json
// this example is just showing the usage of "DiagnosticSettingsCategory_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this DiagnosticSettingsCategoryResource
string resourceUri = "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
DiagnosticSettingsCategoryCollection collection = client.GetDiagnosticSettingsCategories(new ResourceIdentifier(resourceUri));

// invoke the operation
string name = "WorkflowRuntime";
NullableResponse<DiagnosticSettingsCategoryResource> response = await collection.GetIfExistsAsync(name);
DiagnosticSettingsCategoryResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DiagnosticSettingsCategoryData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
