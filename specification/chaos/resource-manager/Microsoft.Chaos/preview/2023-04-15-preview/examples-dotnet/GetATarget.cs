using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetATarget.json
// this example is just showing the usage of "Targets_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TargetResource created on azure
// for more information of creating TargetResource, please refer to the document of TargetResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string parentProviderNamespace = "Microsoft.Compute";
string parentResourceType = "virtualMachines";
string parentResourceName = "exampleVM";
string targetName = "Microsoft-Agent";
ResourceIdentifier targetResourceId = TargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName);
TargetResource target = client.GetTargetResource(targetResourceId);

// invoke the operation
TargetResource result = await target.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TargetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
