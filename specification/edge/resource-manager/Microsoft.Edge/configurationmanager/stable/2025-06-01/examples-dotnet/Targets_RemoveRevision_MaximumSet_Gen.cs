using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Targets_RemoveRevision_MaximumSet_Gen.json
// this example is just showing the usage of "Targets_RemoveRevision" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeTargetResource created on azure
// for more information of creating EdgeTargetResource, please refer to the document of EdgeTargetResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string targetName = "testname";
ResourceIdentifier edgeTargetResourceId = EdgeTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, targetName);
EdgeTargetResource edgeTarget = client.GetEdgeTargetResource(edgeTargetResourceId);

// invoke the operation
RemoveRevisionContent content = new RemoveRevisionContent(new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"), "tomwmqybqomwkfaeukjneva");
await edgeTarget.RemoveRevisionAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
