using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/listDiagnosticSettings.json
// this example is just showing the usage of "DiagnosticSettings_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this DiagnosticSettingResource
string resourceUri = "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
DiagnosticSettingCollection collection = client.GetDiagnosticSettings(new ResourceIdentifier(resourceUri));

// invoke the operation and iterate over the result
await foreach (DiagnosticSettingResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DiagnosticSettingData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
