using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteVirtualMachineTemplate.json
// this example is just showing the usage of "VirtualMachineTemplates_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmVirtualMachineTemplateResource created on azure
// for more information of creating ScVmmVirtualMachineTemplateResource, please refer to the document of ScVmmVirtualMachineTemplateResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string virtualMachineTemplateName = "HRVirtualMachineTemplate";
ResourceIdentifier scVmmVirtualMachineTemplateResourceId = ScVmmVirtualMachineTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineTemplateName);
ScVmmVirtualMachineTemplateResource scVmmVirtualMachineTemplate = client.GetScVmmVirtualMachineTemplateResource(scVmmVirtualMachineTemplateResourceId);

// invoke the operation
await scVmmVirtualMachineTemplate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
