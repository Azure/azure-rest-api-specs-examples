using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LabServices;
using Azure.ResourceManager.LabServices.Models;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/stopVirtualMachine.json
// this example is just showing the usage of "VirtualMachines_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LabVirtualMachineResource created on azure
// for more information of creating LabVirtualMachineResource, please refer to the document of LabVirtualMachineResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string labName = "testlab";
string virtualMachineName = "template";
ResourceIdentifier labVirtualMachineResourceId = LabVirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, virtualMachineName);
LabVirtualMachineResource labVirtualMachine = client.GetLabVirtualMachineResource(labVirtualMachineResourceId);

// invoke the operation
await labVirtualMachine.StopAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
