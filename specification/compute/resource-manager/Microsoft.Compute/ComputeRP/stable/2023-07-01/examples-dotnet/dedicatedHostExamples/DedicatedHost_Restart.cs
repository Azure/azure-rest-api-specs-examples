using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/dedicatedHostExamples/DedicatedHost_Restart.json
// this example is just showing the usage of "DedicatedHosts_Restart" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedHostResource created on azure
// for more information of creating DedicatedHostResource, please refer to the document of DedicatedHostResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string hostGroupName = "myDedicatedHostGroup";
string hostName = "myHost";
ResourceIdentifier dedicatedHostResourceId = DedicatedHostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostGroupName, hostName);
DedicatedHostResource dedicatedHost = client.GetDedicatedHostResource(dedicatedHostResourceId);

// invoke the operation
await dedicatedHost.RestartAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
