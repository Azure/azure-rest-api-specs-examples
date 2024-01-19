using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/DeleteTarget.json
// this example is just showing the usage of "Targets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ChaosTargetResource created on azure
// for more information of creating ChaosTargetResource, please refer to the document of ChaosTargetResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string parentProviderNamespace = "Microsoft.Compute";
string parentResourceType = "virtualMachines";
string parentResourceName = "exampleVM";
string targetName = "Microsoft-Agent";
ResourceIdentifier chaosTargetResourceId = ChaosTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName);
ChaosTargetResource chaosTarget = client.GetChaosTargetResource(chaosTargetResourceId);

// invoke the operation
await chaosTarget.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
