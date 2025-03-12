using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteVMMServer.json
// this example is just showing the usage of "VmmServers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmServerResource created on azure
// for more information of creating ScVmmServerResource, please refer to the document of ScVmmServerResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vmmServerName = "ContosoVMMServer";
ResourceIdentifier scVmmServerResourceId = ScVmmServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmmServerName);
ScVmmServerResource scVmmServer = client.GetScVmmServerResource(scVmmServerResourceId);

// invoke the operation
await scVmmServer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
