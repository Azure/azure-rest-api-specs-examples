using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/getDiagnosticSetting.json
// this example is just showing the usage of "DiagnosticSettings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiagnosticSettingResource created on azure
// for more information of creating DiagnosticSettingResource, please refer to the document of DiagnosticSettingResource
string resourceUri = "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
string name = "mysetting";
ResourceIdentifier diagnosticSettingResourceId = DiagnosticSettingResource.CreateResourceIdentifier(resourceUri, name);
DiagnosticSettingResource diagnosticSetting = client.GetDiagnosticSettingResource(diagnosticSettingResourceId);

// invoke the operation
DiagnosticSettingResource result = await diagnosticSetting.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiagnosticSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
