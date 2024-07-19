using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GetVMInstanceGuestAgent.json
// this example is just showing the usage of "VMInstanceGuestAgents_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmGuestAgentResource created on azure
// for more information of creating ScVmmGuestAgentResource, please refer to the document of ScVmmGuestAgentResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier scVmmGuestAgentResourceId = ScVmmGuestAgentResource.CreateResourceIdentifier(resourceUri);
ScVmmGuestAgentResource scVmmGuestAgent = client.GetScVmmGuestAgentResource(scVmmGuestAgentResourceId);

// invoke the operation
ScVmmGuestAgentResource result = await scVmmGuestAgent.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScVmmGuestAgentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
