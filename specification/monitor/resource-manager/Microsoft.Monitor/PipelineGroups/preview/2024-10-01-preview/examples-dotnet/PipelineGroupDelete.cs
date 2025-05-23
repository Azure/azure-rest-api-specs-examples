using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Monitor/PipelineGroups/preview/2024-10-01-preview/examples/PipelineGroupDelete.json
// this example is just showing the usage of "PipelineGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PipelineGroupResource created on azure
// for more information of creating PipelineGroupResource, please refer to the document of PipelineGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string pipelineGroupName = "plGroup1";
ResourceIdentifier pipelineGroupResourceId = PipelineGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, pipelineGroupName);
PipelineGroupResource pipelineGroup = client.GetPipelineGroupResource(pipelineGroupResourceId);

// invoke the operation
await pipelineGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
