using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Consoles_Delete.json
// this example is just showing the usage of "Consoles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConsoleResource created on azure
// for more information of creating ConsoleResource, please refer to the document of ConsoleResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string virtualMachineName = "virtualMachineName";
string consoleName = "default";
ResourceIdentifier consoleResourceId = ConsoleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineName, consoleName);
ConsoleResource console = client.GetConsoleResource(consoleResourceId);

// invoke the operation
await console.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
