using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Targets_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Target_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EdgeTargetResource
EdgeTargetCollection collection = resourceGroupResource.GetEdgeTargets();

// invoke the operation
string targetName = "testname";
EdgeTargetData data = new EdgeTargetData(new AzureLocation("kckloegmwsjgwtcl"))
{
    Properties = new EdgeTargetProperties(
    "riabrxtvhlmizyhffdpjeyhvw",
    "qjlbshhqzfmwxvvynibkoi",
    new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
    new Dictionary<string, BinaryData>(),
    new string[] { "grjapghdidoao" },
    "octqptfirejhjfavlnfqeiikqx")
    {
        SolutionScope = "testname",
        State = EdgeResourceState.Active,
    },
    ExtendedLocation = new ExtendedLocation
    {
        Name = "szjrwimeqyiue",
    },
    Tags =
    {
    ["key612"] = "vtqzrk"
    },
};
ArmOperation<EdgeTargetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, targetName, data);
EdgeTargetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeTargetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
