using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachines_Restart.json
// this example is just showing the usage of "BareMetalMachines_Restart" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BareMetalMachineResource created on azure
// for more information of creating BareMetalMachineResource, please refer to the document of BareMetalMachineResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string bareMetalMachineName = "bareMetalMachineName";
ResourceIdentifier bareMetalMachineResourceId = BareMetalMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, bareMetalMachineName);
BareMetalMachineResource bareMetalMachine = client.GetBareMetalMachineResource(bareMetalMachineResourceId);

// invoke the operation
await bareMetalMachine.RestartAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
