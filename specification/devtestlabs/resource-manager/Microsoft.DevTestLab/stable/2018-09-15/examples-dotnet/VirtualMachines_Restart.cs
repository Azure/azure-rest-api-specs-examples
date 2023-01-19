using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Restart.json
// this example is just showing the usage of "VirtualMachines_Restart" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabVmResource created on azure
// for more information of creating DevTestLabVmResource, please refer to the document of DevTestLabVmResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string name = "{vmName}";
ResourceIdentifier devTestLabVmResourceId = DevTestLabVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, name);
DevTestLabVmResource devTestLabVm = client.GetDevTestLabVmResource(devTestLabVmResourceId);

// invoke the operation
await devTestLabVm.RestartAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
